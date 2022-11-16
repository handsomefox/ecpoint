package ecpoint

import (
	"math/big"
	"reflect"
	"testing"
)

func TestBasePointGGet(t *testing.T) {
	tests := []struct {
		name      string
		wantPoint ECPoint
	}{
		{
			name: "empty point",
			wantPoint: ECPoint{
				X: &big.Int{},
				Y: &big.Int{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPoint := BasePointGGet(); !reflect.DeepEqual(gotPoint, tt.wantPoint) {
				t.Errorf("BasePointGGet() = %v, want %v", gotPoint, tt.wantPoint)
			}
		})
	}
}

func TestECPointGen(t *testing.T) {
	type args struct {
		x *big.Int
		y *big.Int
	}

	tests := []struct {
		name      string
		args      args
		wantPoint ECPoint
	}{
		{
			name: "point with args",
			args: args{
				x: big.NewInt(10),
				y: big.NewInt(-30),
			},
			wantPoint: ECPoint{
				X: big.NewInt(10),
				Y: big.NewInt(-30),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPoint := ECPointGen(tt.args.x, tt.args.y); !reflect.DeepEqual(gotPoint, tt.wantPoint) {
				t.Errorf("ECPointGen() = %v, want %v", gotPoint, tt.wantPoint)
			}
		})
	}
}

func TestIsOnCurveCheck(t *testing.T) {
	type args struct {
		a ECPoint
	}

	tests := []struct {
		name  string
		args  args
		wantC bool
	}{
		{
			name: "on curve",
			args: args{
				a: ECPointGen(curve.Params().Gx, curve.Params().Gy),
			},
			wantC: true,
		},
		{
			name: "off curve",
			args: args{
				a: ECPointGen(new(big.Int).SetInt64(1), new(big.Int).SetInt64(1)),
			},
			wantC: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := IsOnCurveCheck(tt.args.a); gotC != tt.wantC {
				t.Errorf("IsOnCurveCheck() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestAddECPoints(t *testing.T) {
	type args struct {
		a ECPoint
		b ECPoint
	}

	tests := []struct {
		name  string
		args  args
		wantC ECPoint
	}{{
		name: "add",
		args: args{
			a: ECPointGen(curve.Params().Gx, curve.Params().Gy),
			b: ECPointGen(curve.Params().Gx, curve.Params().Gy),
		},
		wantC: ECPointGen(curve.Add(curve.Params().Gx, curve.Params().Gy, curve.Params().Gx, curve.Params().Gy)),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := AddECPoints(tt.args.a, tt.args.b); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("AddECPoints() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestDoubleECPoints(t *testing.T) {
	type args struct {
		a ECPoint
	}

	tests := []struct {
		name  string
		args  args
		wantC ECPoint
	}{
		{
			name: "double",
			args: args{
				a: ECPointGen(curve.Params().Gx, curve.Params().Gy),
			},
			wantC: ECPointGen(curve.Double(curve.Params().Gx, curve.Params().Gy)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := DoubleECPoints(tt.args.a); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("DoubleECPoints() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestScalarMult(t *testing.T) {
	type args struct {
		a ECPoint
		k big.Int
	}

	wantx, wanty := curve.ScalarMult(curve.Params().Gx, curve.Params().Gy, curve.Params().Gx.Bytes())

	tests := []struct {
		name  string
		args  args
		wantC ECPoint
	}{{
		name: "",
		args: args{
			a: ECPointGen(curve.Params().Gx, curve.Params().Gy),
			k: *curve.Params().Gx,
		},
		wantC: ECPointGen(wantx, wanty),
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := ScalarMult(tt.args.a, tt.args.k); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("ScalarMult() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
