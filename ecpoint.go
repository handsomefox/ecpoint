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
	return ECPointGen(curve.Add(a.X, a.Y, b.X, b.Y))
}

// 2P.
func DoubleECPoints(a ECPoint) (c ECPoint) {
	return ECPointGen(curve.Double(a.X, a.Y))
}

// k * P.
func ScalarMult(a ECPoint, k big.Int) (c ECPoint) {
	return ECPointGen(curve.ScalarMult(a.X, a.Y, k.Bytes()))
}

// Convert point to string.
func ECPointToString(point ECPoint) (s string) {
	return fmt.Sprintf("x=%s, y=%s", point.X, point.Y)
}

// Print point.
func PrintECPoint(point ECPoint) {
	fmt.Println(ECPointToString(point))
}
