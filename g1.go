package main

type fpElement [6]uint64

type frElement [4]uint64

type G1Affine struct {
	X, Y fpElement
}

type G1Jac struct {
	X, Y, Z fpElement
}

// IsInfinity checks if the point is infinity
// in affine, it's encoded as (0,0)
// (0,0) is never on the curve for j=0 curves
func (p *G1Affine) IsInfinity() bool {
	xZero := (p.X[0] | p.X[1] | p.X[2] | p.X[3] | p.X[4] | p.X[5]) == 0
	yZero := (p.Y[0] | p.Y[1] | p.Y[2] | p.Y[3] | p.Y[4] | p.Y[5]) == 0
	return xZero && yZero
}
