package utils

import "math"

type Vector struct {
    V1 float64
    V2 float64
    V3 float64
}

func (v *Vector) X() float64 {
    return v.V1
}
func (v *Vector) Y() float64 {
    return v.V2
}
func (v *Vector) Z() float64 {
    return v.V3
}
func (v *Vector) R() float64 {
    return v.V1
}
func (v *Vector) G() float64 {
    return v.V2
}
func (v *Vector) B() float64 {
    return v.V3
}

func Add(v_1, v_2 Vector) Vector {
    return Vector{v_1.V1 + v_2.V1, v_1.V2 + v_2.V2, v_1.V3 + v_2.V3}
}

func Subtract(v_1, v_2 Vector) Vector {
    return Vector{v_1.V1 - v_2.V1, v_1.V2 - v_2.V2, v_1.V3 - v_2.V3}
}

func Multiply(v_1, v_2 Vector) Vector {
    return Vector{v_1.V1 * v_2.V1, v_1.V2 * v_2.V2, v_1.V3 * v_2.V3}
}

func Divide(v_1, v_2 Vector) Vector {
    return Vector{v_1.V1 / v_2.V1, v_1.V2 / v_2.V2, v_1.V3 / v_2.V3}
}

func AddScalar(v Vector, t float64) Vector {
    return Vector{v.V1 + t, v.V2 + t, v.V3 + t}
}
func SubtractScalar(v Vector, t float64) Vector {
    return Vector{v.V1 - t, v.V2 - t, v.V3 - t}
}
func MultiplyScalar(v Vector, t float64) Vector {
    return Vector{v.V1 * t, v.V2 * t, v.V3 * t}
}

func DivideScalar(v Vector, t float64) Vector {
    return Vector{v.V1 / t, v.V2 / t, v.V3 / t}
}

func (v *Vector) Add(v_other Vector) {
    v.V1 = v.V1 + v_other.V1
    v.V2 = v.V2 + v_other.V2
    v.V3 = v.V3 + v_other.V3
}

func (v *Vector) Subtract(v_other Vector) {
    v.V1 = v.V1 - v_other.V1
    v.V2 = v.V2 - v_other.V2
    v.V3 = v.V3 - v_other.V3
}

func (v *Vector) Multiply(v_other Vector) {
    v.V1 = v.V1 * v_other.V1
    v.V2 = v.V2 * v_other.V2
    v.V3 = v.V3 * v_other.V3
}

func (v *Vector) Divide(v_other Vector) {
    v.V1 = v.V1 / v_other.V1
    v.V2 = v.V2 / v_other.V2
    v.V3 = v.V3 / v_other.V3
}

func Dot(v_1, v_2 Vector) float64 {
    return (v_1.V1 * v_2.V1 + v_1.V2 * v_2.V2 + v_1.V3 * v_2.V3)
}

func (v *Vector) Length() float64 {
    return math.Sqrt(v.V1*v.V1 + v.V2*v.V2 + v.V3*v.V3)
}

func (v *Vector) SquaredLength() float64 {
    return (v.V1*v.V1 + v.V2*v.V2 + v.V3*v.V3)
}

func MakeUnitVector(v Vector) Vector {
    return DivideScalar(v, v.Length())
}
