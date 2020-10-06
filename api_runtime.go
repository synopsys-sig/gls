// Copyright 2018 Massimiliano Ghilardi. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !gc,!gccgo

// gls/runtime/id.go is deceivingly simple...
// but requires a lot of trickery to compile and install.
// for gc, prefer assembler implementations id_*.s
// for gccgo, prefer the C implementation wrapped by api_gccgo.go

package gls

import (
	"github.com/synopsys-sig/gls/runtime"
)

// return the current goroutine ID.
//
// note that the returned value is DIFFERENT from most other goroutine libraries:
// this GoID() returns the address, converted to uintptr, of the runtime.g struct.
// NOT the runtime.g.goid field returned by most other libraries.
func GoID() uintptr {
	return runtime.GoID()
}
