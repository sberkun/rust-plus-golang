package main

// NOTE: There should be NO space between the comments and the `import "C"` line.
// The -ldl is sometimes necessary to fix linker errors about `dlsym`.

/*
#cgo LDFLAGS: ./libhello.a -ldl
#include "./hello.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Element [6]uint64

type G1Affine struct {
	X, Y Element
}

type G1Jac struct {
	X, Y, Z Element
}

func (p *G1Jac) call_rust_thing(points []G1Affine, scalars []Element) {
	C.multi_scalar_mult_wrapper(
		unsafe.Pointer(p),
		unsafe.Pointer(&points[0]),
		unsafe.Pointer(&scalars[0]),
		C.ulong(len(points)),
	)
}

func main() {
	points := []G1Affine{
		{X: Element{1, 2, 3, 4, 5, 6}, Y: Element{7, 8, 9, 10, 11, 12}},
		{},
		{},
		{},
		{X: Element{1, 2, 3, 4, 5, 34234}, Y: Element{7, 8, 9, 10, 11, 234234}},
	}

	scalars := []Element{
		{4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5},
		{},
		{},
		{10, 20, 30, 40, 50, 60},
	}

	p := G1Jac{}
	p.call_rust_thing(points, scalars)
	fmt.Println(p)
}
