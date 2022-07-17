// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package structcacheexample

import (
	"errors"
	"sync"

	"go.uber.org/atomic"
)

type T struct {
	foo      []byte
	bar      []byte
	caching  chan struct{}
	uncached *atomic.Bool
}

func New() *T {
	v := &T{
		caching:  make(chan struct{}, 1),
		uncached: atomic.NewBool(false),
	}
	v.caching <- struct{}{}
	return v
}

func (v *T) String() string {
	if err := v.cache(); err != nil {
		return err.Error()
	}
	return string(v.foo)
}

func (v *T) Bar() string {
	if err := v.cacheBar(); err != nil {
		return err.Error()
	}
	return string(v.bar)
}

var ErrUncached = errors.New("closed")

func (v *T) cacheBar() error {
	if err := v.cache(); err != nil {
		return err
	}

	v.bar = (*pool.Get().(*[]byte))[:0]
	v.bar = append(v.bar, []byte("Hello, Bar!")...)
	return nil
}

func (v *T) cache() error {
	if v.caching == nil {
		return ErrUncached
	}

	if v.uncached.Load() {
		return ErrUncached
	}

	if _, caching := <-v.caching; caching {
		defer close(v.caching)
	} else {
		return nil
	}

	v.foo = (*pool.Get().(*[]byte))[:0]
	v.foo = append(v.foo, []byte("Hello, Foo!")...)
	return nil
}

func (v *T) Close() {
	if v.uncached.Load() {
		return
	}
	v.uncache()
}

func (v *T) uncache() {
	defer v.uncached.Store(true)

	if v.foo != nil {
		pool.Put(&v.foo)
	}

	if v.bar != nil {
		pool.Put(&v.bar)
	}
}

var pool = sync.Pool{New: func() interface{} { return &[]byte{} }}
