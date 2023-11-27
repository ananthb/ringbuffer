# ringbuffer

[![Go Reference](https://pkg.go.dev/badge/github.com/ananthb/ringbuffer.svg)](https://pkg.go.dev/github.com/ananthb/ringbuffer) [![CI](https://github.com/ananthb/ringbuffer/actions/workflows/go.yml/badge.svg)](https://github.com/ananthb/ringbuffer/actions/workflows/go.yml)

A circular buffer (ring buffer) in Go, implementing standard Go interfaces
for reads and writes.

[![Ring Buffer](circular_buffer_animation.gif)](https://github.com/smallnest/ringbuffer)

```go
  rb := New(1024)

  // write
  rb.Write([]byte("abcd"))
  fmt.Println(rb.Length())
  fmt.Println(rb.Free())

  // read
  buf := make([]byte, 4)
  rb.Read(buf)
  fmt.Println(string(buf))
```

## [LICENSE](LICENSE)

Copyright (c) 2019 smallnest, 2023 Ananth Bhaskararaman

ringbuffer is available under the terms of the MIT license.
