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

package kubernetes

import (
	"context"

	http "github.com/topfreegames/go-extensions-http"
	restWrapper "github.com/topfreegames/go-extensions-k8s-client-go/rest"
	"k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
)

type Clientset struct {
	kubernetes.Interface
	restClient *restWrapper.Client
}

// NewForConfig creates a *Clientset that acts as a wrapper over an instance of
// *k8s.io/client-go/kubernetes.Clientset
// OpenTracing instrumentation is set for the underlying *http.Client
// of the rest.Interface sent to all instances created by the wrapped
// *kubernetes.Clientset
func NewForConfig(c *rest.Config) (*Clientset, error) {
	client, err := rest.RESTClientFor(c)
	if err != nil {
		return nil, err
	}
	client.Client = http.New()
	cc := restWrapper.New(client)
	return newWithContext(cc, nil), nil
}

func newWithContext(c *restWrapper.Client, ctx context.Context) *Clientset {
	cc := c.WithContext(ctx)
	return &Clientset{
		Interface:  kubernetes.New(cc),
		restClient: cc,
	}
}

// WithContext creates a new instance of *Clientset with `ctx` propagated to
// it's restClient, and used by all of restClient's methods
func (c *Clientset) WithContext(ctx context.Context) *Clientset {
	return newWithContext(c.restClient, ctx)
}
