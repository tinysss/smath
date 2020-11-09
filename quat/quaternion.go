/*
 * @Author: sealon
 * @Date: 2020-10-16 14:32:31
 * @Last Modified by: sealon
 * @Last Modified time: 2020-11-09 15:15:58
 * @Desc:
 */
package quat

import (
	math "github.com/barnex/fmath"
	"github.com/tinysss/smath/sutil"

	"github.com/tinysss/smath/vector3"
	"github.com/tinysss/smath/vector4"
)

type Quaternion [4]float32 // [vw]

var (
	Zero  = Quaternion{}
	Ident = Quaternion{0, 0, 0, 1} // 单位四元数
)

// 模
func (t *Quaternion) Norm() float32 {
	return t[0]*t[0] + t[1]*t[1] + t[2]*t[2] + t[3]*t[3]
}

// 模
func (t *Quaternion) NormSqrt() float32 {
	return math.Sqrt(t.Norm())
}

// 模
func (t *Quaternion) Len() float32 {
	return t.NormSqrt()
}

// 归一化
func (t *Quaternion) Normalize() *Quaternion {
	norm := t.Norm()
	if norm != 1 && norm != 0 {
		ool := 1 / math.Sqrt(norm)
		t[0] *= ool
		t[1] *= ool
		t[2] *= ool
		t[3] *= ool
	}
	return t
}

// 归一化
func (t *Quaternion) Normalized() Quaternion {
	norm := t.Norm()
	if norm != 1 && norm != 0 {
		ool := 1 / math.Sqrt(norm)
		return Quaternion{
			t[0] * ool,
			t[1] * ool,
			t[2] * ool,
			t[3] * ool,
		}
	} else {
		return *t
	}
}

// 标准数判断
func (t *Quaternion) IsNormalQuat() bool {
	return math.Abs(t.Norm()-1) <= 0.0001
}

// 返回 绕axis旋转angle的四元数
func FromAxisAngle(axis *vector3.Vector, angle float32) Quaternion {
	axisnor := axis.Normalized()
	angle *= 0.5
	sina, cosa := math.Sincos(angle)
	l_quat := Quaternion{
		axisnor[0] * sina,
		axisnor[1] * sina,
		axisnor[2] * sina,
		cosa,
	}
	// return l_quat
	return l_quat.Normalized()
}

// 返回 绕axis旋转angle的四元数
func NewFromAxisAngle(axis *vector3.Vector, angle float32) *Quaternion {
	l_quat := FromAxisAngle(axis, angle)
	return &l_quat
}

// 返回 绕x轴选装angle的四元数
func FromXAxisAngle(angle float32) Quaternion {
	angle *= 0.5
	sina, cosa := math.Sincos(angle)

	// 不需要归一化(sina^2+cosa^2=1)
	// l_quat := Quaternion{
	// 	float32(sina), 0, 0, float32(cosa),
	// }
	// return l_quat.Normalized()
	return Quaternion{sina, 0, 0, cosa}
}

// 返回 绕y轴选装angle的四元数
func FromYAxisAngle(angle float32) Quaternion {
	angle *= 0.5
	sina, cosa := math.Sincos(angle)
	return Quaternion{0, sina, 0, cosa}
	// return l_quat.Normalized()
}

// 返回 绕z轴选装angle的四元数
func FromZAxisAngle(angle float32) Quaternion {
	angle *= 0.5
	sina, cosa := math.Sincos(angle)
	return Quaternion{0, 0, float32(sina), cosa}
	// return l_quat.Normalized()
}

// 返回 hpb(I2O)欧拉角构造的四元数  (使用限制角)
func FromEulerAngles(yHead, xPitch, zBank float32) Quaternion {
	xPitch, yHead, zBank = sutil.CanonizeEuler(xPitch, yHead, zBank)
	yHead /= 2.0
	xPitch /= 2.0
	zBank /= 2.0

	sh, ch := math.Sincos(yHead)
	sp, cp := math.Sincos(xPitch)
	sb, cb := math.Sincos(zBank)

	// qy := FromYAxisAngle(yHead)
	// qx := FromXAxisAngle(xPitch)
	// qz := FromZAxisAngle(zBank)
	// return Mul3(&qy, &qx, &qz)

	return Quaternion{
		ch*sp*cb + sh*cp*sb,
		sh*cp*cb - ch*sp*sb,
		ch*cp*sb - sh*sp*cb,
		ch*cp*cb + sh*sp*sb,
	}

}

func NewFromEulerAngles(yHead, xPitch, zBank float32) *Quaternion {
	l_quat := FromEulerAngles(yHead, xPitch, zBank)
	return &l_quat
}

func FromVec4(v *vector4.Vector) Quaternion {
	return Quaternion(*v)
}

func (t *Quaternion) Vec4() vector4.Vector {
	return vector4.Vector(*t)
}

// 提取欧拉角
// func (t *Quaternion) ToEulerAngles() (yHead, xPitch, zBank float32) {

// }

// 提取轴角
func (t *Quaternion) AxisAngle() (axis vector3.Vector, angle float32) {
	angle = math.Acos(t[3]) * 2
	// sina := math.Sin(angle / 2)
	sina := math.Sqrt(1 - t[3]*t[3])
	// 防止sina趋于0，这里改成乘法
	ooSin := float32(1)
	if math.Abs(sina) > 0.0001 {
		ooSin = 1 / sina
	}
	axis[0] = t[0] * ooSin
	axis[1] = t[1] * ooSin
	axis[2] = t[2] * ooSin
	return
}

// 共轭
func (t *Quaternion) Conjugate() *Quaternion {
	t[0] = -t[0]
	t[1] = -t[1]
	t[2] = -t[2]
	return t
}

func (t *Quaternion) Conjugated() Quaternion {
	return Quaternion{-t[0], -t[1], -t[2], t[3]}
}

// 逆
func (t *Quaternion) Inverse() *Quaternion {
	l_cjgated := t.Conjugated()
	l_inv := l_cjgated.Scaled(1 / t.Dot(t))
	t[0] = l_inv[0]
	t[1] = l_inv[1]
	t[2] = l_inv[2]
	t[3] = l_inv[3]
	return t
}

func (t *Quaternion) Inversed() Quaternion {
	l_cjgated := t.Conjugated()
	return l_cjgated.Scaled(1 / t.Dot(t))
}

// p` = qpq-1 旋转
// t必须为标准数
func (t *Quaternion) RotateVec3(v *vector3.Vector) {
	p := Quaternion{v[0], v[1], v[2], 0}
	qinv := t.Conjugated() // 标准数的共轭=逆
	q := Mul3(t, &p, &qinv)
	v[0] = q[0]
	v[1] = q[1]
	v[2] = q[2]
}

func (t *Quaternion) RotatedVec3(v *vector3.Vector) vector3.Vector {
	p := Quaternion{v[0], v[1], v[2], 0}
	qinv := t.Conjugated()
	q := Mul3(t, &p, &qinv)
	return vector3.Vector{q[0], q[1], q[2]}
}

// scale
func (t *Quaternion) Scale(c float32) *Quaternion {
	t[0] *= c
	t[1] *= c
	t[2] *= c
	t[3] *= c
	return t
}

func (t Quaternion) Scaled(c float32) Quaternion {
	return Quaternion{t[0] * c, t[1] * c, t[2] * c, t[3] * c}
}

func (t *Quaternion) Sub(q Quaternion) *Quaternion {
	t[0] -= q[0]
	t[1] -= q[1]
	t[2] -= q[2]
	t[3] -= q[3]
	return t
}

func (t *Quaternion) Subed(q Quaternion) Quaternion {
	res := *t
	res[0] -= q[0]
	res[1] -= q[1]
	res[2] -= q[2]
	res[3] -= q[3]
	return res
}

func (t *Quaternion) Add(q Quaternion) *Quaternion {
	t[0] += q[0]
	t[1] += q[1]
	t[2] += q[2]
	t[3] += q[3]
	return t
}

func (t Quaternion) Added(q Quaternion) Quaternion {
	res := t
	res[0] += q[0]
	res[1] += q[1]
	res[2] += q[2]
	res[3] += q[3]
	return res
}

// 点积[1,-1]  越大表明两个角位移越接近
// PS: t,o 为标准数且方向相同才有意义
func (t *Quaternion) Dot(o *Quaternion) float32 {
	return t[3]*o[3] + t[0]*o[0] + t[1]*o[1] + t[2]*o[2]
}

// a -> b 的角位移是否是最短的(因为有两个方向)
func (t *Quaternion) IsShortestRotation(b *Quaternion) bool {
	return Dot(t, b) >= 0
}

// 幂运算 (t是标准四元数才有意义)
func (t *Quaternion) Pow(exponent float32) *Quaternion {
	// 单位四元数的任意次方仍然是单位四元数
	if math.Abs(t[3]) >= 0.9999 {
		return t
	}

	angle := math.Acos(t[3]) // 半角
	newAngle := angle * exponent
	t[3] = math.Cos(newAngle)
	l_mult := angle / newAngle
	t[0] *= l_mult
	t[1] *= l_mult
	t[2] *= l_mult
	return t
}

func (t *Quaternion) Powed(exponent float32) Quaternion {
	// 单位四元数的任意次方仍然是单位四元数
	if math.Abs(t[3]) >= 0.9999 {
		return *t
	}

	angle := math.Acos(t[3]) // 半角
	newAngle := angle * exponent
	l_quat := *t
	l_quat[3] = math.Cos(newAngle)
	l_mult := angle / newAngle
	l_quat[0] *= l_mult
	l_quat[1] *= l_mult
	l_quat[2] *= l_mult
	return l_quat
}

// a•b=|a||b|cosθ  点积[1,-1]
func Dot(a, b *Quaternion) float32 {
	return a.Dot(b)
}

// a -> b 的角位移是否是最短的(因为有两个方向)
func IsShortestRotation(a, b *Quaternion) bool {
	return Dot(a, b) >= 0
}

// 乘
func Mul(a, b *Quaternion) Quaternion {
	q := Quaternion{
		a[3]*b[0] + a[0]*b[3] + a[1]*b[2] - a[2]*b[1],
		a[3]*b[1] + a[1]*b[3] + a[2]*b[0] - a[0]*b[2],
		a[3]*b[2] + a[2]*b[3] + a[0]*b[1] - a[1]*b[0],
		a[3]*b[3] - a[0]*b[0] - a[1]*b[1] - a[2]*b[2],
	}
	return q
}

// 3个乘
func Mul3(a, b, c *Quaternion) Quaternion {
	q := Mul(a, b)
	return Mul(&q, c)
}

// 4个乘
func Mul4(a, b, c, d *Quaternion) Quaternion {
	q := Mul(a, b)
	q = Mul(&q, c)
	return Mul(&q, d)
}

// 差四元数　（ad=b  求d = a-1 * b ）
func DiffQuat(a, b *Quaternion) Quaternion {
	ainv := a.Inversed()
	d := Mul(&ainv, b)
	return d.Normalized()
}

func FromToQuat(from, to vector3.Vector) Quaternion {
	from.Normalize()
	to.Normalize()
	cr := vector3.Cross(&from, &to)
	sr := math.Sqrt(2 * (1 + vector3.Dot(&from, &to)))
	oosr := 1 / sr

	q := Quaternion{cr[0] * oosr, cr[1] * oosr, cr[2] * oosr, sr * 0.5}
	return q.Normalized()
}

func Clamp(a, low, high float32) float32 {
	if a < low {
		return low
	} else if a > high {
		return high
	}

	return a
}

func Lerp(a, b *Quaternion, t float32) *Quaternion {
	l_res := a.Added(b.Subed((*a).Scaled(t)))
	return &l_res
}

func NLerp(a, b *Quaternion, t float32) *Quaternion {
	return Lerp(a, b, t).Normalize()
}

// [0-2pi] 顺时针
func Slerp(a, b *Quaternion, t float32) Quaternion {
	if t <= 0.0 {
		return *a
	} else if t >= 1.0 {
		return *b
	}
	dot := Dot(a, b)
	if dot > 0.9995 {
		return *NLerp(a, b, t)
	}

	dot = Clamp(dot, -1, 1) // cosalpha
	theta := math.Acos(dot) * t

	s, c := math.Sincos(theta)

	resl := b.Subed(a.Scaled(dot))
	resl.Normalize()

	return a.Scaled(c).Added(resl.Scaled(s))
}

// [-pi,pi] 选择近角方向
func SmartSlerp(a, b *Quaternion, t float32) Quaternion {
	if t <= 0.0 {
		return *a
	} else if t >= 1.0 {
		return *b
	}
	dot := Dot(a, b)
	if dot > 0.9995 {
		return *NLerp(a, b, t)
	} else if dot <= 0.0 {
		temp := b.Scaled(-1)
		b = &temp
		dot = -dot
	}

	dot = Clamp(dot, -1, 1) // cosalpha
	theta := math.Acos(dot) * t

	s, c := math.Sincos(theta)

	resl := b.Subed(a.Scaled(dot))
	resl.Normalize()

	return a.Scaled(c).Added(resl.Scaled(s))
}
