/*
 * Copyright (c) 2019 TFG Co <backend@tfgco.com>
 * Author: TFG Co <backend@tfgco.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package rest

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

type Client struct {
	rest.Interface
	ctx context.Context
}

// New creates a wrapper over a rest.Interface instance
// This wrapper is useful to force usage of a context.Context
// by the wrapped requests
func New(wrapped rest.Interface) *Client {
	return &Client{Interface: wrapped}
}

// WithContext returns a clone of the *Client with an specific `ctx`
// that will be used by any *rest.Request as a result of calling any *Client
// methods
func (r *Client) WithContext(ctx context.Context) *Client {
	return &Client{
		Interface: r.Interface,
		ctx:       ctx,
	}
}

func (r *Client) Verb(verb string) *rest.Request {
	return r.Interface.Verb(verb).Context(r.ctx)
}

func (r *Client) Post() *rest.Request {
	return r.Interface.Post().Context(r.ctx)
}

func (r *Client) Put() *rest.Request {
	return r.Interface.Put().Context(r.ctx)
}

func (r *Client) Patch(pt types.PatchType) *rest.Request {
	return r.Interface.Patch(pt).Context(r.ctx)
}

func (r *Client) Get() *rest.Request {
	return r.Interface.Get().Context(r.ctx)
}

func (r *Client) Delete() *rest.Request {
	return r.Interface.Delete().Context(r.ctx)
}
