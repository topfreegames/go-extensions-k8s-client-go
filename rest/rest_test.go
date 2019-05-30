package rest_test

import (
	"context"
	"testing"

	restExtensions "github.com/topfreegames/go-extensions-k8s-client-go/rest"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

type clientMock struct {
	calls map[string]int
	ctx   context.Context
}

func newRestMock() *clientMock {
	return &clientMock{calls: map[string]int{}}
}

func (r *clientMock) GetRateLimiter() flowcontrol.RateLimiter {
	r.calls["GetRateLimiter"]++
	return flowcontrol.NewFakeNeverRateLimiter()
}

func (r *clientMock) Verb(verb string) *rest.Request {
	r.calls["Verb"]++
	return &rest.Request{}
}

func (r *clientMock) Post() *rest.Request {
	r.calls["Post"]++
	return &rest.Request{}
}

func (r *clientMock) Put() *rest.Request {
	r.calls["Put"]++
	return &rest.Request{}
}

func (r *clientMock) Patch(pt types.PatchType) *rest.Request {
	r.calls["Patch"]++
	return &rest.Request{}
}

func (r *clientMock) Get() *rest.Request {
	r.calls["Get"]++
	return &rest.Request{}
}

func (r *clientMock) Delete() *rest.Request {
	r.calls["Delete"]++
	return &rest.Request{}
}

func (r *clientMock) APIVersion() schema.GroupVersion {
	r.calls["APIVersion"]++
	return schema.GroupVersion{}
}

func TestWithContext(t *testing.T) {
	m := newRestMock()
	r := restExtensions.New(m)
	ctx := context.Background()
	if m.calls["Verb"] != 0 {
		t.Fatalf("Expected calls[Verb] to be 0. Got %d", m.calls["Verb"])
	}
	rr := r.WithContext(ctx)
	if rr == r {
		t.Fatal("Expected WithContext to return a new instance of *restExtensions.Client")
	}
	rr.Verb("POST")
	if m.calls["Verb"] != 1 {
		t.Fatalf("Expected calls[Verb] to be 1. Got %d", m.calls["Verb"])
	}
}

func TestVerb(t *testing.T) {
	m := newRestMock()
	r := restExtensions.New(m)
	if m.calls["Verb"] != 0 {
		t.Fatalf("Expected calls[Verb] to be 0. Got %d", m.calls["Verb"])
	}
	r.Verb("POST")
	if m.calls["Verb"] != 1 {
		t.Fatalf("Expected calls[Verb] to be 1. Got %d", m.calls["Verb"])
	}
}

func TestPost(t *testing.T) {
	m := newRestMock()
	r := restExtensions.New(m)
	if m.calls["Post"] != 0 {
		t.Fatalf("Expected calls[Post] to be 0. Got %d", m.calls["Post"])
	}
	r.Post()
	if m.calls["Post"] != 1 {
		t.Fatalf("Expected calls[Post] to be 1. Got %d", m.calls["Post"])
	}
}

func TestPut(t *testing.T) {
	m := newRestMock()
	r := restExtensions.New(m)
	if m.calls["Put"] != 0 {
		t.Fatalf("Expected calls[Put] to be 0. Got %d", m.calls["Put"])
	}
	r.Put()
	if m.calls["Put"] != 1 {
		t.Fatalf("Expected calls[Put] to be 1. Got %d", m.calls["Put"])
	}
}

func TestPatch(t *testing.T) {
	m := newRestMock()
	r := restExtensions.New(m)
	if m.calls["Patch"] != 0 {
		t.Fatalf("Expected calls[Patch] to be 0. Got %d", m.calls["Patch"])
	}
	r.Patch(types.JSONPatchType)
	if m.calls["Patch"] != 1 {
		t.Fatalf("Expected calls[Patch] to be 1. Got %d", m.calls["Patch"])
	}
}

func TestGet(t *testing.T) {
	m := newRestMock()
	r := restExtensions.New(m)
	if m.calls["Get"] != 0 {
		t.Fatalf("Expected calls[Get] to be 0. Got %d", m.calls["Get"])
	}
	r.Get()
	if m.calls["Get"] != 1 {
		t.Fatalf("Expected calls[Get] to be 1. Got %d", m.calls["Get"])
	}
}

func TestDelete(t *testing.T) {
	m := newRestMock()
	r := restExtensions.New(m)
	if m.calls["Delete"] != 0 {
		t.Fatalf("Expected calls[Delete] to be 0. Got %d", m.calls["Delete"])
	}
	r.Delete()
	if m.calls["Delete"] != 1 {
		t.Fatalf("Expected calls[Delete] to be 1. Got %d", m.calls["Delete"])
	}
}
