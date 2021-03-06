/*
 * @Author: sealon
 * @Date: 2020-09-29 14:50:05
 * @Last Modified by: sealon
 * @Last Modified time: 2020-11-11 10:14:23
 * @Desc:
 */
package mat3

import (
	"fmt"
	"unsafe"

	math "github.com/barnex/fmath"
	"github.com/tinysss/smath/generic"
	"github.com/tinysss/smath/mat2"
	"github.com/tinysss/smath/sutil"
	"github.com/tinysss/smath/vector3"
	"github.com/ungerik/go3d/vec2"
)

// 列存储 每个vec代表一列
type Mat3 [3]vector3.Vector

var (
	Zero = Mat3{}
	// 单位阵
	Ident = Mat3{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
)

func New(v1, v2, v3 vector3.Vector) *Mat3 {
	return &Mat3{v1, v2, v3}
}
func NewEmpty() *Mat3 {
	l_ret := Ident
	return &l_ret
}

func FromNew(other generic.T) *Mat3 {
	r := Ident
	cols := other.Cols()
	rows := other.Rows()

	if cols != rows || cols < 3 || cols > 4 {
		panic(fmt.Sprintf("unsupported type. cols=%d rows=%d ", cols, rows))
	}

	cols = 3
	rows = 3

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			r[col][row] = other.Get(col, row)
		}
	}
	return &r
}

func (t *Mat3) Array() *[9]float32 {
	return (*[9]float32)(unsafe.Pointer(t))
}

//-------------------------------------------- 实现generic.T begin-------------------------------------
func (t *Mat3) Cols() int {
	return 3
}

func (t *Mat3) Rows() int {
	return 3
}

func (t *Mat3) Size() int {
	return 9
}

func (t *Mat3) Slice() []float32 {
	return t.Array()[:]
}

func (t *Mat3) Get(col, row int) float32 {
	return t[col][row]
}

func (t *Mat3) IsZero() bool {
	return *t == Zero
}

//-------------------------------------------- 实现generic.T end -------------------------------------

func (t *Mat3) Scale(f float32) *Mat3 {
	t[0][0] *= f
	t[1][1] *= f
	t[2][2] *= f
	return t
}

func (t *Mat3) Scaled(f float32) Mat3 {
	r := *t
	return *r.Scale(f)
}

func (t *Mat3) Scaling() vector3.Vector {
	return vector3.Vector{t[0][0], t[1][1], t[2][2]}
}

func (t *Mat3) SetScaling(s *vector3.Vector) *Mat3 {
	t[0][0] = s[0]
	t[1][1] = s[1]
	t[2][2] = s[2]
	return t
}

func (t *Mat3) ScaleVec2(s *vec2.T) *Mat3 {
	t[0][0] *= s[0]
	t[1][1] *= s[1]
	return t
}

func (t *Mat3) SetTranslation(s *vec2.T) *Mat3 {
	t[2][0] = s[0]
	t[2][1] = s[1]
	return t
}

func (t *Mat3) Translate(s *vec2.T) *Mat3 {
	t[2][0] += s[0]
	t[2][1] += s[1]
	return t
}

func (t *Mat3) TranslateX(dx float32) *Mat3 {
	t[2][0] += dx
	return t
}

func (t *Mat3) TranslateY(dy float32) *Mat3 {
	t[2][1] += dy
	return t
}

// 迹
func (t *Mat3) Trace() float32 {
	return t[0][0] + t[1][1] + t[2][2]
}

func (t *Mat3) Mul(s float32) *Mat3 {
	t[0].Scale(s)
	t[1].Scale(s)
	t[2].Scale(s)
	return t
}

// v' = v * M
func (t *Mat3) MulVec3(v *vector3.Vector) vector3.Vector {
	return vector3.Vector{
		t[0][0]*v[0] + t[1][0]*v[1] + t[2][0]*v[2],
		t[0][1]*v[0] + t[1][1]*v[1] + t[2][1]*v[2],
		t[0][2]*v[0] + t[1][2]*v[1] + t[2][2]*v[2],
	}
}

func (t *Mat3) Mul3x3(o *Mat3) *Mat3 {
	l_temp := *t
	t.AssignMul(&l_temp, o)
	return t
}

func (t *Mat3) AssignMul(a, b *Mat3) *Mat3 {
	t[0] = a.MulVec3(&b[0])
	t[1] = a.MulVec3(&b[1])
	t[2] = a.MulVec3(&b[2])
	return t
}

func (t *Mat3) AssignMat2x2(m *mat2.Mat2) *Mat3 {
	*t = Mat3{
		vector3.Vector{m[0][0], m[0][1], 0},
		vector3.Vector{m[1][0], m[1][1], 0},
		vector3.Vector{0, 0, 1},
	}
	return t
}

// 变换v，直接将结果给v
func (t *Mat3) TransformVec3(v *vector3.Vector) {
	vx := t[0][0]*v[0] + t[1][0]*v[1] + t[2][0]*v[2]
	vy := t[0][1]*v[0] + t[1][1]*v[1] + t[2][1]*v[2]
	v[2] = t[0][2]*v[0] + t[1][2]*v[1] + t[2][2]*v[2]
	v[0] = vx
	v[1] = vy
}

// 变换v，不该v，返回变换结果
func (t *Mat3) TransformVec3Ret(v *vector3.Vector) *vector3.Vector {
	l_nv := *v
	t.TransformVec3(&l_nv)
	return &l_nv
}

func (t *Mat3) AssignXRotation(angle float32) *Mat3 {
	sina, cosa := math.Sincos(angle)

	t[0][0] = 1
	t[0][1] = 0
	t[0][2] = 0

	t[1][0] = 0
	t[1][1] = cosa
	t[1][2] = sina

	t[2][0] = 0
	t[2][1] = -sina
	t[2][2] = cosa

	return t
}

func (t *Mat3) AssignYRotation(angle float32) *Mat3 {
	sina, cosa := math.Sincos(angle)

	t[0][0] = cosa
	t[0][1] = 0
	t[0][2] = -sina

	t[1][0] = 0
	t[1][1] = 1
	t[1][2] = 0

	t[2][0] = sina
	t[2][1] = 0
	t[2][2] = cosa

	return t
}

func (t *Mat3) AssignZRotation(angle float32) *Mat3 {
	sina, cosa := math.Sincos(angle)

	t[0][0] = cosa
	t[0][1] = sina
	t[0][2] = 0

	t[1][0] = -sina
	t[1][1] = cosa
	t[1][2] = 0

	t[2][0] = 0
	t[2][1] = 0
	t[2][2] = 1

	return t
}

// 通过euler构建mat3
func (t *Mat3) AssignEulerRotation(yHead, xPitch, zBank float32) *Mat3 {
	xPitch, yHead, zBank = sutil.CanonizeEuler(xPitch, yHead, zBank)

	sh, ch := math.Sincos(yHead)
	sp, cp := math.Sincos(xPitch)
	sb, cb := math.Sincos(zBank)

	t[0][0] = ch*cb + sh*sp*sb
	t[0][1] = sb * cp
	t[0][2] = -sh*cb + ch*sp*sb

	t[1][0] = -ch*sb + sh*sp*cb
	t[1][1] = cb * cp
	t[1][2] = sb*sh + ch*sp*cb

	t[2][0] = sh * cp
	t[2][1] = -sp
	t[2][2] = ch * cp

	return t
}

// 提取euler
func (t *Mat3) ExtractEulerAngles() (yHead, xPitch, zBank float32) {
	sp := -t[2][1]
	if sp >= -0.999 {
		if sp <= 0.999 { // 有效区间 sp(-1,1)
			xPitch = math.Asin(sp)
			yHead = math.Atan2(t[2][0], t[2][2])
			zBank = math.Atan2(t[0][1], t[1][1])
		} else { // sp  >= 0.999  按sinp = 1处理
			xPitch = sutil.KPiOver2
			yHead = math.Atan2(t[1][0], t[0][0])
			zBank = 0
		}
	} else { //sinp <= -0.999  按sinp = -1处理
		xPitch = -sutil.KPiOver2
		yHead = math.Atan2(-t[1][0], t[0][0])
		zBank = 0
	}
	// xPitch, yHead, zBank = sutil.CanonizeEuler(xPitch, yHead, zBank)

	return
}

func (t *Mat3) AssignCoordinateSystem(x, y, z *vector3.Vector) *Mat3 {
	t[0] = *x
	t[1] = *y
	t[2] = *z

	return t
}

// |Mat|
// a11a22a33 + a12a23a31 + a13a21a32- a13a22a31 - a12a21a33 - a11a23a32
func (t *Mat3) Det() float32 {
	return t[0][0]*t[1][1]*t[2][2] +
		t[0][1]*t[1][2]*t[2][0] +
		t[0][2]*t[1][0]*t[2][1] -
		t[0][2]*t[1][1]*t[2][0] -
		t[0][1]*t[1][0]*t[2][2] -
		t[0][0]*t[1][2]*t[2][1]

}

// 逆
func (t *Mat3) Inv() *Mat3 {
	det := t.Det()
	if sutil.FloatEqual(det, 0) {
		return NewEmpty()
	}

	retMat := New(
		vector3.Vector{
			t[1][1]*t[2][2] - t[2][1]*t[1][2],
			t[2][1]*t[0][2] - t[0][1]*t[2][2],
			t[0][1]*t[1][2] - t[1][1]*t[0][2]},

		vector3.Vector{
			t[2][0]*t[1][2] - t[1][0]*t[2][2],
			t[0][0]*t[2][2] - t[2][0]*t[0][2],
			t[1][0]*t[0][2] - t[0][0]*t[1][2]},
		vector3.Vector{
			t[1][0]*t[2][1] - t[2][0]*t[1][1],
			t[2][0]*t[0][1] - t[0][0]*t[2][1],
			t[0][0]*t[1][1] - t[1][0]*t[0][1]})

	return retMat.Mul(1 / det)
}

// 转置
func (t *Mat3) Transpose() *Mat3 {
	t[0][1], t[1][0] = t[1][0], t[0][1]
	t[0][2], t[2][0] = t[2][0], t[0][2]
	t[1][2], t[2][1] = t[2][1], t[1][2]
	return t
}

func Mul(a, b *Mat3) *Mat3 {
	l_matres := Mat3{
		a.MulVec3(&b[0]),
		a.MulVec3(&b[1]),
		a.MulVec3(&b[2]),
	}
	return &l_matres
}
