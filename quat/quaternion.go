/*
 * @Author: sealon
 * @Date: 2020-10-16 14:32:31
 * @Last Modified by: sealon
 * @Last Modified time: 2020-10-17 12:21:10
 * @Desc:
 */
package quat

import (
	math "github.com/barnex/fmath"

	"github.com/tinysss/smath/vector3"
	"github.com/tinysss/smath/vector4"
)

type Quaternion [4]float32 // [vw]

var (
	Zero  = Quaternion{}
	Ident = Quaternion{0, 0, 0, 1}
)

// 模
func (t *Quaternion) Norm() float32 {
	return t[0]*t[0] + t[1]*t[1] + t[2]*t[2] + t[3]*t[3]
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

// 返回 绕axis旋转angle的四元数
func FromAxisAngle(axis *vector3.Vector, angle float32) Quaternion {
	angle *= 0.5
	sina, cosa := math.Sincos(angle)
	l_quat := Quaternion{
		axis[0] * sina,
		axis[1] * sina,
		axis[2] * sina,
		cosa,
	}
	// return l_quat
	return l_quat.Normalized()
}

func FromAxisAngle2(axis *vector3.Vector, angle float32) Quaternion {
	angle *= 0.5
	sina, cosa := math.Sincos(angle)
	l_quat := Quaternion{
		axis[0] * sina,
		axis[1] * sina,
		axis[2] * sina,
		cosa,
	}
	return l_quat
	// return l_quat.Normalized()
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

// 返回 hpb(I2O unity是这种)欧拉角代表的四元数
func FromEulerAnglesI2O(yHead, xPitch, zBank float32) Quaternion {

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

	// 20201017 效率更高
	return Quaternion{
		ch*sp*cb + sh*cp*sb,
		sh*cp*cb - ch*sp*sb,
		ch*cp*sb - sh*sp*cb,
		ch*cp*cb + sh*sp*sb,
	}

}

// 返回 rpy(O2I)欧拉角代表的四元数
// func FromEulerAnglesO2I(zRoll, xPitch, yYaw float32) Quaternion {

// 	yYaw /= 2.0
// 	xPitch /= 2.0
// 	zRoll /= 2.0

// 	sr, cr := math.Sincos(zRoll)
// 	sp, cp := math.Sincos(xPitch)
// 	sh, ch := math.Sincos(yYaw)

// 	return Quaternion{
// 		cr*sp*ch - sr*cp*sh,
// 		cr*cp*sh + sr*sp*ch,
// 		cr*sp*sh + sr*cp*ch,
// 		cr*cp*ch - sr*sp*sh,
// 	}

// }

func FromVec4(v *vector4.Vector) Quaternion {
	return Quaternion(*v)
}

func (t *Quaternion) Vec4() vector4.Vector {
	return vector4.Vector(*t)
}

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
func (t *Quaternion) RotateVec3(v *vector3.Vector) {
	p := Quaternion{v[0], v[1], v[2], 0}
	qinv := t.Inversed()

	q := Mul3(t, &p, &qinv)
	v[0] = q[0]
	v[1] = q[2]
	v[2] = q[3]

}

func (t *Quaternion) RotatedVec3(v *vector3.Vector) vector3.Vector {
	p := Quaternion{v[0], v[1], v[2], 0}
	qinv := t.Inversed()
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

func (t *Quaternion) Scaled(c float32) Quaternion {
	return Quaternion{t[0] * c, t[1] * c, t[2] * c, t[3] * c}
}

// 点积
func (t *Quaternion) Dot(o *Quaternion) float32 {
	return t[3]*o[3] + t[0]*o[0] + t[1]*o[1] + t[2]*o[2]
}

// 点积
func Dot(a, b *Quaternion) float32 {
	return a.Dot(b)
}

// 2个乘 ps:(乘积的的模 = 模的乘积)单位四元数乘单位四元数结果仍是单位四元数
func Mul(a, b *Quaternion) Quaternion {
	q := Quaternion{
		a[3]*b[0] + a[0]*b[3] + a[1]*b[2] - a[2]*b[1],
		a[3]*b[1] + a[1]*b[3] + a[2]*b[0] - a[0]*b[2],
		a[3]*b[2] + a[2]*b[3] + a[0]*b[1] - a[1]*b[0],
		a[3]*b[3] - a[0]*b[0] - a[1]*b[1] - a[2]*b[2],
	}
	return q.Normalized()
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
