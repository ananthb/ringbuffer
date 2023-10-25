// Copyright 2019 smallnest, 2023 Ananth Bhaskararaman. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package ringbuffer

import "fmt"

func ExampleRingBuffer() {
	rb := New(1024)
	_, _ = rb.Write([]byte("abcd"))
	fmt.Println(rb.Length())
	fmt.Println(rb.Free())
	buf := make([]byte, 4)

	_, _ = rb.Read(buf)
	fmt.Println(string(buf))
	// Output: 4
	// 1020
	// abcd
}
