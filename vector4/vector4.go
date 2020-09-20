/*
 * @Author: sealon
 * @Date: 2020-09-17 17:54:30
 * @Last Modified by: sealon
 * @Last Modified time: 2020-09-20 14:37:04
 * @Desc:
 */

package vector4

import (
	"math"

	"github.com/tinysss/smath/generic"
	"github.com/tinysss/smath/vector3"
)

type Vector [4]float32

var (
	Zero     = Vector{}
	UnitXW   = Vector{1, 0, 0, 1}
	UnitYW   = Vector{0, 1, 0, 1}
	UnitZW   = Vector{0, 0, 1, 1}
	UnitXYZW = Vector{1, 1, 1, 1}
	MinVal   = Vector{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32, 1}
	MaxVal   = Vector{math.MaxFloat32, math.MaxFloat32, math.MaxFloat32, 1}
)

func New(f1, f2, f3, f4 float32) *Vector {
	return &Vector{f1, f2, f3, f4}
}

func FromNew(other generic.T) *Vector {
	switch other.Size() {
	case 2:
		return &Vector{other.Get(0, 0), other.Get(0, 1), 0}
	case 3:
		return &Vector{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2)}
	case 4:
		return &Vector{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2), other.Get(0, 3)}
	default:
		panic("unsupported type.")
	}
}

//-------------------------------------------- 实现generic.T begin-------------------------------------
func (t *Vector) Cols() int {
	return 1
}

func (t *Vector) Rows() int {
	return 4
}

func (t *Vector) Size() int {
	return 4
}

func (t *Vector) Slice() []float32 {
	return t[:]
}

func (t *Vector) Get(col, row int) float32 {
	return t[row]
}

func (t *Vector) IsZero() bool {
	return t[0] == 0 && t[1] == 0 && t[2] == 0 && t[3] == 0
}

//-------------------------------------------- 实现generic.T end -------------------------------------

func (t *Vector) X() float32 {
	return t[0]
}
func (t *Vector) Y() float32 {
	return t[1]
}
func (t *Vector) Z() float32 {
	return t[2]
}
func (t *Vector) W() float32 {
	return t[3]
}

func (t *Vector) Length() float32 {
	v4 := t.DividedByW()
	return v4.Length()
}

func (t *Vector) LengthSqr() float32 {
	v4 := t.DividedByW()
	return v4.LengthSqr()
}

// 缩放自身
func (t *Vector) Scale(ratio float32) *Vector {
	t[0] *= ratio
	t[1] *= ratio
	t[2] *= ratio
	return t
}

// 返回缩放自身的拷贝，自身不受影响
func (t *Vector) Scaled(ratio float32) Vector {
	return Vector{t[0] * ratio, t[1] * ratio, t[2] * ratio}
}

// 逆暂且求相反向量
func (t *Vector) Invert(ratio float32) *Vector {
	t[0] = -t[0]
	t[1] = -t[1]
	t[2] = -t[2]
	return t
}

// 返回逆自身的拷贝，自身不受影响
func (t *Vector) Inverted() Vector {
	return Vector{-t[0], -t[1], -t[2]}
}

// 使用vector3 归一化
func (t *Vector) Normalize() *Vector {
	v3 := t.Vec3DividedByW()
	v3.Normalize()
	t[0] = v3[0]
	t[1] = v3[1]
	t[2] = v3[2]
	t[3] = 1
	return t
}

//
func (t *Vector) Normalized() Vector {
	l_temp := *t
	l_temp.Normalize()
	return l_temp
}

// 标准化正交向量
func (t *Vector) Normal() Vector {
	v3 := t.Vector3()
	n3 := v3.Normal()
	return Vector{n3[0], n3[1], n3[2], 1}
}

// 根据W分量取值, 自身
func (t *Vector) DivideByW() *Vector {
	if t[3] == 1 {
		return t
	}
	s := 1 / t[3]
	t[0] *= s
	t[1] *= s
	t[2] *= s
	t[3] = 1
	return t
}

// 根据W分量取值， 拷贝
func (t *Vector) DividedByW() Vector {
	if t[3] == 1 {
		return *t
	}
	s := 1 / t[3]
	return Vector{t[0] * s, t[1] * s, t[2] * s, 1}
}

// 根据W分量取值， vector3拷贝
func (t *Vector) Vec3DividedByW() vector3.Vector {
	if t[3] == 1 {
		return vector3.Vector{t[0], t[1], t[2]}
	}
	s := 1 / t[3]
	return vector3.Vector{t[0] * s, t[1] * s, t[2] * s}
}

// 转vector3
func (t *Vector) Vector3() vector3.Vector {
	return vector3.Vector{t[0], t[1], t[2]}
}

func (t *Vector) AssignVec3(v *vector3.Vector) *Vector {
	t[0] = v[0]
	t[1] = v[1]
	t[2] = v[2]
	t[3] = 1
	return t
}

// ps:w不同时，统一转1
func (t *Vector) Add(v *Vector) *Vector {
	if t[3] == v[3] {
		t[0] += v[0]
		t[1] += v[1]
		t[2] += v[2]
		return t
	}
	t.DivideByW()
	v3 := v.Vec3DividedByW()
	t[0] += v3[0]
	t[1] += v3[1]
	t[2] += v3[2]
	return t
}

func (t *Vector) Sub(v *Vector) *Vector {
	if t[3] == v[3] {
		t[0] -= v[0]
		t[1] -= v[1]
		t[2] -= v[2]
		return t
	}

	t.DivideByW()
	v3 := v.Vec3DividedByW()
	t[0] -= v3[0]
	t[1] -= v3[1]
	t[2] -= v3[2]
	return t
}

func (t *Vector) Clamp(min, max *Vector) *Vector {
	for i := range t {
		if t[i] < min[i] {
			t[i] = min[i]
		} else if t[i] > max[i] {
			t[i] = max[i]
		}
	}
	return t
}

func (t *Vector) Clamped(min, max *Vector) Vector {
	result := *t
	result.Clamp(min, max)
	return result
}

func (t *Vector) Clamp01() *Vector {
	return t.Clamp(&Zero, &UnitXYZW)
}

func (t *Vector) Clamped01() Vector {
	result := *t
	result.Clamp01()
	return result
}

func Add(a, b *Vector) Vector {
	if a[3] == b[3] {
		return Vector{a[0] + b[0], a[1] + b[1], a[2] + b[2], a[3]}
	}
	a1 := a.Vec3DividedByW()
	b1 := b.Vec3DividedByW()
	return Vector{a1[0] + b1[0], a1[1] + b1[1], a1[2] + b1[2], 1}
}

func Sub(a, b *Vector) Vector {
	if a[3] == b[3] {
		return Vector{a[0] - b[0], a[1] - b[1], a[2] - b[2], a[3]}
	}
	a1 := a.Vec3DividedByW()
	b1 := b.Vec3DividedByW()
	return Vector{a1[0] - b1[0], a1[1] - b1[1], a1[2] - b1[2], 1}
}

func Dot(a, b *Vector) float32 {
	a3 := a.Vec3DividedByW()
	b3 := b.Vec3DividedByW()
	return vector3.Dot(&a3, &b3)
}

func Dot4(a, b *Vector) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

func Cross(a, b *Vector) Vector {
	a3 := a.Vec3DividedByW()
	b3 := b.Vec3DividedByW()
	c3 := vector3.Cross(&a3, &b3)
	return Vector{c3[0], c3[1], c3[2], 1}
}

// a,b夹角  [0,pi]
// a·b=|a|·|b|·cosθ
func Angle(a, b *Vector) float32 {
	v := Dot(a, b) / (a.Length() * b.Length())
	// 避免NaN
	if v > 1. {
		v = v - 2
	} else if v < -1. {
		v = v + 2
	}
	return float32(math.Acos(float64(v)))
}

// a,b夹角  [-pi,pi]
func Angle2(a, b, up *Vector) float32 {
	l_angle := Angle(a, b)
	if l_angle == 0 {
		return l_angle
	}
	l_normal := Cross(a, b)
	if Dot(&l_normal, up) > 0 {
		return l_angle
	} else {
		return -l_angle
	}
}

// a - b的插值  t[0,1]
func Interpolate(a, b *Vector, t float32) Vector {
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	t1 := 1 - t
	return Vector{
		a[0]*t1 + b[0]*t,
		a[1]*t1 + b[1]*t,
		a[2]*t1 + b[2]*t,
		a[3]*t1 + b[3]*t,
	}
}
