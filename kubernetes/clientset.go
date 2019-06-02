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

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Clientset struct {
	*kubernetes.Clientset
	*rest.Config
}

func NewForConfig(c *rest.Config) (*Clientset, error) {
	k, err := kubernetes.NewForConfig(c)
	if err != nil {
		return nil, err
	}
	cs := &Clientset{Clientset: k, Config: c}
	instrument(cs, nil)
	return cs, nil
}

// WithContext creates a new instance of *Clientset with `ctx` propagated to
// it's components' RESTClient instances
func (c *Clientset) WithContext(ctx context.Context) (*Clientset, error) {
	cs, err := NewForConfig(c.Config)
	if err != nil {
		return nil, err
	}
	if err := instrument(cs, ctx); err != nil {
		return nil, err
	}
	return cs, nil
}

// WithContext tries to cast the kubernetes.Interface sent to *Clientset
// and wrap it with `ctx`
func WithContext(c kubernetes.Interface, ctx context.Context) (kubernetes.Interface, error) {
	if v, ok := c.(*Clientset); ok {
		return v.WithContext(ctx)
	}
	return nil, &NotClientsetError{}
}

// TryWithContext will return either a *Clientset wrapping `ctx` or the original
// kubernetes.Interface if an error occurs
func TryWithContext(c kubernetes.Interface, ctx context.Context) kubernetes.Interface {
	k, err := WithContext(c, ctx)
	if err != nil {
		return c
	}
	return k
}
