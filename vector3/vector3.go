/*
 * @Author: sealon
 * @Date: 2020-09-14 10:59:46
 * @Last Modified by: sealon
 * @Last Modified time: 2020-09-14 11:03:06
 * @Desc:
 */
package vector3

import (
	"math"

	"github.com/tinysss/smath/generic"
)

type Vector [3]float32

var (
	Zero    = Vector{}
	UnitX   = Vector{1, 0, 0}
	UnitY   = Vector{0, 1, 0}
	UnitZ   = Vector{0, 0, 1}
	UnitXYZ = Vector{1, 1, 1}
	MinVal  = Vector{-math.MaxFloat32, -math.MaxFloat32, -math.MaxFloat32}
	MaxVal  = Vector{math.MaxFloat32, math.MaxFloat32, math.MaxFloat32}
)

func New(f1, f2, f3 float32) *Vector {
	return &Vector{f1, f2, f3}
}

func FromNew(other generic.T) *Vector {
	switch other.Size() {
	case 2:
		return &Vector{other.Get(0, 0), other.Get(0, 1), 0}
	case 3, 4:
		return &Vector{other.Get(0, 0), other.Get(0, 1), other.Get(0, 2)}
	default:
		panic("unsupported type.")
	}
}

//-------------------------------------------- 实现generic.T begin-------------------------------------
func (t *Vector) Cols() int {
	return 3
}

func (t *Vector) Rows() int {
	return 1
}

func (t *Vector) Size() int {
	return 3
}

func (t *Vector) Slice() []float32 {
	return t[:]
}

func (t *Vector) Get(row, col int) float32 {
	return t[col]
}

func (t *Vector) IsZero() bool {
	return t[0] == 0 && t[1] == 0 && t[2] == 0
}

//-------------------------------------------- 实现generic.T end -------------------------------------

func (t *Vector) Length() float32 {
	return float32(math.Sqrt(float64(t[0]*t[0] + t[1]*t[1] + t[2]*t[2])))
}

func (t *Vector) LengthSqr() float32 {
	return t[0]*t[0] + t[1]*t[1] + t[2]*t[2]
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

func (t *Vector) Abs() *Vector {
	t[0] = float32(math.Abs(float64(t[0])))
	t[1] = float32(math.Abs(float64(t[1])))
	t[2] = float32(math.Abs(float64(t[2])))
	return t
}

func (t *Vector) Absed() Vector {
	return Vector{float32(math.Abs(float64(t[0]))), float32(math.Abs(float64(t[1]))), float32(math.Abs(float64(t[2])))}
}

// 归一化  v norm = (1/|v|)*v
func (t *Vector) Normalize() *Vector {
	l := t.LengthSqr()
	if l == 0 || l == 1 {
		return t
	}
	t.Scale(float32(1 / math.Sqrt(float64(l))))
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
	n := Cross(t, &UnitZ)
	if n.IsZero() {
		return UnitX
	}
	return n.Normalized()
}

func (t *Vector) Add(v *Vector) *Vector {
	t[0] += v[0]
	t[1] += v[1]
	t[2] += v[2]
	return t
}

func (t *Vector) Sub(v *Vector) *Vector {
	t[0] -= v[0]
	t[1] -= v[1]
	t[2] -= v[2]
	return t
}

func (t *Vector) Mul(v *Vector) *Vector {
	t[0] *= v[0]
	t[1] *= v[1]
	t[2] *= v[2]
	return t
}

// Clamp clamps the vector's components to be in the range of min to max.
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
	return t.Clamp(&Zero, &UnitXYZ)
}

func (t *Vector) Clamped01() Vector {
	result := *t
	result.Clamp01()
	return result
}

func Add(a, b *Vector) Vector {
	return Vector{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func SquareDistance(a, b *Vector) float32 {
	d := Sub(a, b)
	return d.LengthSqr()
}

func Distance(a, b *Vector) float32 {
	d := Sub(a, b)
	return d.Length()
}

func Sub(a, b *Vector) Vector {
	return Vector{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func Mul(a, b *Vector) Vector {
	return Vector{a[0] * b[0], a[1] * b[1], a[2] * b[2]}
}

func Dot(a, b *Vector) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

/*
	a0  b0
	a1	b1
	a2  b2
*/
func Cross(a, b *Vector) Vector {
	return Vector{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
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

// 两个分量最小值构成的新向量
func Min(a, b *Vector) Vector {
	l_min := *a
	if l_min[0] > b[0] {
		l_min[0] = b[0]
	}
	if l_min[1] > b[1] {
		l_min[1] = b[1]
	}
	if l_min[2] > b[2] {
		l_min[2] = b[2]
	}
	return l_min
}

// 两个分量最大值构成的新向量
func Max(a, b *Vector) Vector {
	l_max := *a
	if l_max[0] < b[0] {
		l_max[0] = b[0]
	}
	if l_max[1] < b[1] {
		l_max[1] = b[1]
	}
	if l_max[2] < b[2] {
		l_max[2] = b[2]
	}
	return l_max
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
	}
}
