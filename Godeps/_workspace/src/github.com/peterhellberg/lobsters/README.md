# lobsters

Go library for the Lobsters API

[![GoDoc](https://godoc.org/github.com/peterhellberg/lobsters?status.svg)](https://godoc.org/github.com/peterhellberg/lobsters)
[![Build Status](https://travis-ci.org/peterhellberg/lobsters.svg?branch=master)](https://travis-ci.org/peterhellberg/lobsters)

## Installation

```bash
go get -u github.com/peterhellberg/lobsters
```

## Service

The client delegates to an implementation of interface:
[StoriesService](https://godoc.org/github.com/peterhellberg/lobsters#StoriesService)

## Example usage

Showing the hottest stories

```go
package main

import (
  "fmt"

  "github.com/peterhellberg/lobsters"
)

func main() {
  l := lobsters.NewClient(nil)

  stories, err := l.Hottest()
  if err != nil {
    panic(err)
  }

  for _, story := range stories {
    fmt.Println(story.Title)
    fmt.Println(story.URL, "\n")
  }
}
```

## License

> *The MIT License (MIT)*
>
> Copyright (c) 2014 [Peter Hellberg](http://c7.se/)
>
> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included in all
> copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
> SOFTWARE.
