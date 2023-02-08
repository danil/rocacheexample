# Struct cache example for Go

[![Build Status](https://cloud.drone.io/api/badges/danil/rocacheexample/status.svg)](https://cloud.drone.io/danil/rocacheexample)
[![Go Reference](https://pkg.go.dev/badge/github.com/danil/rocacheexample.svg)](https://pkg.go.dev/github.com/danil/rocacheexample)

Source files are distributed under the BSD-style license.

## About

The software is considered to be at an alpha level of readiness,
its extremely slow and allocates a lots of memory.

## Benchmark

```sh
$ go test -count=1 -race -bench ./...
goos: linux
goarch: amd64
pkg: github.com/danil/rocacheexample
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkCache/cache-8             13396         89845 ns/op
PASS
ok      github.com/danil/rocacheexample 2.143s
```
