# SurflineR

Go Client for the Surfline Regions API.

## Installation

Simple run `go get github.com/mhelmetag/surfliner` and start using it in your own apps!

## Usage

The full example for fetching Areas, Regions and Sub Regions can be found in `examples/main.go` and can be run with `go run examples/main.go`.

To list all Surfline Areas, first instantiate a new `Client` and then call `ListAreas()`:

```go
package main

import "github.com/mhelmetag/surfliner"

client, err := surfliner.DefaultClient()
areas, err := client.ListAreas()
```
