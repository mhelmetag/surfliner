# SurflineR

[![Build Status](https://travis-ci.org/mhelmetag/surfliner.svg?branch=master)](https://travis-ci.org/mhelmetag/surfliner)

**NOTE:** I would use Surfline's taxonomy API (https://services.surfline.com/taxonomy) instead of my API (https://github.com/mhelmetag/surfline_regions) and this client. My API and this client were meant to bridge the gap when Surfline didn't have an appropriate JSON API for regions they support. Now they do.

Go Client for the Surfline Regions API.

## Installation

Simply run `go get github.com/mhelmetag/surfliner` and start using it in your own apps!

## Usage

The full example for fetching Areas, Regions and Sub Regions can be found in `examples/main.go` and can be run with `go run examples/main.go`.

To list all Surfline Areas, first instantiate a new `Client` and then call `Areas()`:

```go
package main

import "github.com/mhelmetag/surfliner"

client, err := surfliner.DefaultClient()
areas, err := client.Areas()
```
