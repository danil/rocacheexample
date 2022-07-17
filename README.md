# Struct cache example for Go

[![Build Status](https://cloud.drone.io/api/badges/danil/structcacheexample/status.svg)](https://cloud.drone.io/danil/structcacheexample)
[![Go Reference](https://pkg.go.dev/badge/github.com/danil/structcacheexample.svg)](https://pkg.go.dev/github.com/danil/structcacheexample)

Source files are distributed under the BSD-style license.

## About

The software is considered to be at a alpha level of readiness,
its extremely slow and allocates a lots of memory.

## Benchmark

```sh
$ go test -count=1 -race -bench ./...
goos: linux
goarch: amd64
pkg: github.com/danil/structcacheexample
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkCache/cache-8             13396         89845 ns/op
PASS
ok      github.com/danil/structcacheexample 2.143s
```
