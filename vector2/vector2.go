/*
 * @Author: sealon
 * @Date: 2020-09-10 16:58:56
 * @Last Modified by: sealon
 * @Last Modified time: 2020-09-17 19:06:50
 * @Desc: float32 2D向量封装
 */
package vector2

import (
	"math"

	"github.com/tinysss/smath/generic"
)

type Vector [2]float32

var (
	Zero   = Vector{}
	UnitX  = Vector{1, 0}
	UnitY  = Vector{0, 1}
	UnitXY = Vector{1, 1}
	MinVal = Vector{-math.MaxFloat32, -math.MaxFloat32}
	MaxVal = Vector{math.MaxFloat32, math.MaxFloat32}
)

func New(f1, f2 float32) *Vector {
	return &Vector{f1, f2}
}

func FromNew(other generic.T) *Vector {
	return &Vector{other.Get(0, 0), other.Get(0, 1)}
}

//-------------------------------------------- 实现generic.T begin-------------------------------------
func (t *Vector) Cols() int {
	return 2
}

func (t *Vector) Rows() int {
	return 1
}

func (t *Vector) Size() int {
	return 2
}

func (t *Vector) Slice() []float32 {
	return t[:]
}

func (t *Vector) Get(row, col int) float32 {
	return t[col]
}

func (t *Vector) IsZero() bool {
	return t[0] == 0 && t[1] == 0
}

//-------------------------------------------- 实现generic.T end -------------------------------------

func (t *Vector) X() float32 {
	return t[0]
}
func (t *Vector) Y() float32 {
	return t[1]
}

func (t *Vector) Length() float32 {
	return float32(math.Hypot(float64(t[0]), float64(t[1])))
}

func (t *Vector) LengthSqr() float32 {
	return t[0]*t[0] + t[1]*t[1]
}

// 缩放自身
func (t *Vector) Scale(ratio float32) *Vector {
	t[0] *= ratio
	t[1] *= ratio
	return t
}

// 返回缩放自身的拷贝，自身不受影响
func (t *Vector) Scaled(ratio float32) Vector {
	return Vector{t[0] * ratio, t[1] * ratio}
}

// 逆暂且求相反向量
func (t *Vector) Invert(ratio float32) *Vector {
	t[0] = -t[0]
	t[1] = -t[1]
	return t
}

// 返回逆自身的拷贝，自身不受影响
func (t *Vector) Inverted() Vector {
	return Vector{-t[0], -t[1]}
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

func (t *Vector) Add(v *Vector) *Vector {
	t[0] += v[0]
	t[1] += v[1]
	return t
}

func (t *Vector) Sub(v *Vector) *Vector {
	t[0] -= v[0]
	t[1] -= v[1]
	return t
}

func (t *Vector) Mul(v *Vector) *Vector {
	t[0] *= v[0]
	t[1] *= v[1]
	return t
}

// >0逆时针
func (t *Vector) Rotate(angle float32) *Vector {
	*t = t.Rotated(angle)
	return t
}

//
// x1 = |R| * （x0 * cosB / |R| - y0 * sinB / |R|） =>  x1 = x0 * cosB - y0 * sinB
// y1 = |R| * （y0 * cosB / |R| + x0 * sinB / |R|） =>  y1 = x0 * sinB + y0 * cosB
func (t *Vector) Rotated(angle float32) Vector {
	sinA := float32(math.Sin(float64(angle)))
	cosA := float32(math.Cos(float64(angle)))

	return Vector{
		t[0]*cosA - t[1]*sinA,
		t[0]*sinA + t[1]*cosA,
	}
}

// >0逆时针 绕任意点旋转
func (t *Vector) RotateAroundPoint(point *Vector, angle float32) *Vector {
	return t.Sub(point).Rotate(angle).Add(point)
}

// 逆时针旋转90度，不用Rotate方法是为了加速运算
func (t *Vector) Rotate90degLeft() *Vector {
	l_temp := t[0]
	t[0] = -t[1]
	t[1] = l_temp
	return t
}

// 顺时针旋转90度，不用Rotate方法是为了加速运算
func (t *Vector) Rotate90degRight() *Vector {
	l_temp := t[0]
	t[0] = t[1]
	t[1] = -l_temp
	return t
}

// 相对于x轴的弧度, 返回[-PI,PI]
func (t *Vector) Angle() float32 {
	return float32(math.Atan2(float64(t[1]), float64(t[0])))
}

// 限定在　min max之间
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

// 限定在　min max之间. 返回拷贝
func (t *Vector) Clamped(min, max *Vector) Vector {
	res := *t
	res.Clamp(min, max)
	return res
}

// 限定在0 - 1
func (t *Vector) Clamp01() *Vector {
	return t.Clamp(&Zero, &UnitXY)
}

//  限定在0 - 1. 返回拷贝
func (t *Vector) Clamped01() Vector {
	res := *t
	res.Clamp01()
	return res
}

func Add(a, b *Vector) Vector {
	return Vector{a[0] + b[0], a[1] + b[1]}
}

func Sub(a, b *Vector) Vector {
	return Vector{a[0] - b[0], a[1] - b[1]}
}

func Mul(a, b *Vector) Vector {
	return Vector{a[0] * b[0], a[1] * b[1]}
}

func Dot(a, b *Vector) float32 {
	return a[0]*b[0] + a[1]*b[1]
}

/*
	a0  b0
	a1	b1
*/
func Cross(a, b *Vector) Vector {
	return Vector{
		a[1]*b[0] - a[0]*b[1],
		a[0]*b[1] - a[1]*b[0], // >0逆时针  <0顺时针
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
func Angle2(a, b *Vector) float32 {
	l_angle := Angle(a, b)
	if a[0]*b[1] > a[1]*b[0] {
		return l_angle
	} else {
		return -l_angle
	}
}

// a 到 b是否是向左旋转
func IsLeftWinding(a, b *Vector) bool {
	// l_ab := b.Rotated(-a.Angle())
	// return l_ab.Angle() > 0
	return a[0]*b[1] > a[1]*b[0]
}

// a 到 b是否是向右旋转
func IsRightWinding(a, b *Vector) bool {
	// l_ab := b.Rotated(-a.Angle())
	// return l_ab.Angle() < 0
	return a[0]*b[1] < a[1]*b[0]
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
	}
}
