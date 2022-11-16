package ecpoint

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
)

var curve = elliptic.P256()

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

// G-generator receiving.
func BasePointGGet() (point ECPoint) {
	return ECPoint{
		X: new(big.Int),
		Y: new(big.Int),
	}
}

// ECPoint creation with pre-defined parameters.
func ECPointGen(x, y *big.Int) (point ECPoint) {
	return ECPoint{
		X: new(big.Int).Set(x),
		Y: new(big.Int).Set(y),
	}
}

// P ∈ CURVE?
func IsOnCurveCheck(a ECPoint) (c bool) {
	return curve.IsOnCurve(a.X, a.Y)
}

// P + Q.
func AddECPoints(a, b ECPoint) (c ECPoint) {
	x, y := curve.Add(a.X, a.Y, b.X, b.Y)
	return ECPoint{
		X: x,
		Y: y,
	}
}

// 2P.
func DoubleECPoints(a ECPoint) (c ECPoint) {
	x, y := curve.Double(a.X, a.Y)
	return ECPoint{
		X: x,
		Y: y,
	}
}

// k * P.
func ScalarMult(a ECPoint, k big.Int) (c ECPoint) {
	x, y := curve.ScalarMult(a.X, a.Y, k.Bytes())
	return ECPoint{
		X: x,
		Y: y,
	}
}

// Convert point to string.
func ECPointToString(point ECPoint) (s string) {
	return fmt.Sprintf("x=%s, y=%s", point.X, point.Y)
}

// Print point.
func PrintECPoint(point ECPoint) {
	fmt.Println(ECPointToString(point))
}
