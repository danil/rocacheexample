// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rocacheexample_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/danil/rocacheexample"
)

func BenchmarkCache(b *testing.B) {
	b.ReportAllocs()

	b.Run("cache", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			test()
		}
	})
}

func test() {
	foo := rocacheexample.New()
	defer foo.Close()

	bar := rocacheexample.New()
	defer bar.Close()

	xyz := rocacheexample.New()
	defer xyz.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = foo.String()
		foo.Close()
		_ = foo.String()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = bar.Bar()
		bar.Close()
		_ = bar.Bar()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ = fmt.Sprint(xyz.String(), "\n", xyz.Bar(), "\n")
		xyz.Close()
		_ = fmt.Sprint(xyz.String(), "\n", xyz.Bar(), "\n")
	}()

	wg.Wait()
}
