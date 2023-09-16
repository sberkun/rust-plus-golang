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

func call_multi_scalar_init(points []G1Affine) unsafe.Pointer {
	type RustG1Affine struct {
		X, Y     [6]uint64
		infinity bool
	}

	rust_points := make([]RustG1Affine, len(points))
	for i := 0; i < len(points); i++ {
		rust_points[i] = RustG1Affine{
			points[i].X,
			points[i].Y,
			points[i].IsInfinity(),
		}
	}

	ctx := C.multi_scalar_init_wrapper(
		unsafe.Pointer(&rust_points[0]),
		C.ulong(len(rust_points)),
	)
	return ctx
}

func (p *G1Jac) call_multi_scalar_mult(ctx unsafe.Pointer, scalars []frElement) {
	C.multi_scalar_mult_wrapper(
		unsafe.Pointer(p),
		ctx,
		unsafe.Pointer(&scalars[0]),
		C.ulong(len(scalars)),
	)
}

func main() {
	points := []G1Affine{
		{X: fpElement{1, 2, 3, 4, 5, 6}, Y: fpElement{7, 8, 9, 10, 11, 12}},
		{},
		{},
		{},
		{X: fpElement{1, 2, 3, 4, 5, 34234}, Y: fpElement{7, 8, 9, 10, 11, 234234}},
	}

	scalars := []frElement{
		{4, 4, 4, 4},
		{5, 5, 5, 5},
		{},
		{},
		{10, 20, 30, 40},
	}

	p := G1Jac{}
	ctx := call_multi_scalar_init(points)
	p.call_multi_scalar_mult(ctx, scalars)
	fmt.Println(p)
}
