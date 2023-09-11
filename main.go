package main

// NOTE: There should be NO space between the comments and the `import "C"` line.
// The -ldl is sometimes necessary to fix linker errors about `dlsym`.

/*
#cgo LDFLAGS: ./libhello.a -ldl
#include "./hello.h"
*/
import "C"

type Element [6]uint64

type G1Affine struct {
	X, Y Element
}

func call_rust_thing(points []G1Affine, scalars []Element) {

}

func main() {
	C.hello(C.CString("world"))
	C.whisper(C.CString("this is code from the static library"))
	g := C.struct_Dummy{
		thing1: C.CString("cheese"),
		thing2: C.CString("breeze"),
	}
	C.do_thing(&g)
}
