/*
 * @Author: sealon
 * @Date: 2020-11-10 14:55:56
 * @Last Modified by: sealon
 * @Last Modified time: 2020-11-11 17:21:30
 * @Desc:
 */
package mat4

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"unsafe"

	math "github.com/barnex/fmath"

	"github.com/tinysss/smath/sutil"

	"github.com/tinysss/smath/generic"
	"github.com/tinysss/smath/mat2"
	"github.com/tinysss/smath/mat3"
	"github.com/tinysss/smath/vector3"
	"github.com/tinysss/smath/vector4"
)

// 列存储 每个vec代表一列
type Mat4 [4]vector4.Vector

var (
	Zero = Mat4{}
	// 单位阵
	Ident = Mat4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
)

func New(v1, v2, v3, v4 vector4.Vector) *Mat4 {
	return &Mat4{v1, v2, v3, v4}
}
func NewEmpty() *Mat4 {
	l_ret := Ident
	return &l_ret
}

func FromNew(other generic.T) *Mat4 {
	r := Ident
	cols := other.Cols()
	rows := other.Rows()

	if cols != rows || cols < 4 || cols > 5 {
		panic(fmt.Sprintf("unsupported type. cols=%d rows=%d ", cols, rows))
	}

	cols = 4
	rows = 4

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			r[col][row] = other.Get(col, row)
		}
	}
	return &r
}

func (t *Mat4) Array() *[16]float32 {
	return (*[16]float32)(unsafe.Pointer(t))
}

//-------------------------------------------- 实现generic.T begin-------------------------------------
func (t *Mat4) Cols() int {
	return 4
}

func (t *Mat4) Rows() int {
	return 4
}

func (t *Mat4) Size() int {
	return 16
}

func (t *Mat4) Slice() []float32 {
	return t.Array()[:]
}

func (t *Mat4) Get(col, row int) float32 {
	return t[col][row]
}

func (t *Mat4) IsZero() bool {
	return *t == Zero
}

func (t Mat4) String() string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 4, 4, 1, ' ', tabwriter.AlignRight)
	for i := range t {
		fmt.Fprintf(w, "%.3f\t%.3f\t%.3f\t%.3f\t\n", t[i][0], t[i][1], t[i][2], t[i][3])
	}
	w.Flush()

	return buf.String()
}

//-------------------------------------------- 实现generic.T end -------------------------------------

func (t *Mat4) Scale(f float32) *Mat4 {
	t[0][0] *= f
	t[1][1] *= f
	t[2][2] *= f
	return t
}

func (t *Mat4) Scaled(f float32) Mat4 {
	result := *t
	result.Scale(f)
	return result
}

func (t *Mat4) Mul(f float32) *Mat4 {

	for i := range t {
		t[i][0] *= f
		t[i][1] *= f
		t[i][2] *= f
		t[i][3] *= f
	}
	return t
}

func (t *Mat4) Muled(f float32) Mat4 {
	result := *t
	result.Mul(f)
	return result
}

func (t *Mat4) MultMatrix(m *Mat4) *Mat4 {
	for i := range t {
		col := vector4.Vector{t[0][i], t[1][i], t[2][i], t[3][i]}
		t[0][i] = vector4.Dot(&m[0], &col)
		t[1][i] = vector4.Dot(&m[1], &col)
		t[2][i] = vector4.Dot(&m[2], &col)
		t[3][i] = vector4.Dot(&m[3], &col)
	}
	return t
}

func (t *Mat4) Trace() float32 {
	return t[0][0] + t[1][1] + t[2][2] + t[3][3]
}

func (t *Mat4) Trace3() float32 {
	return t[0][0] + t[1][1] + t[2][2]
}

func (t *Mat4) AssignMat2x2(m *mat2.Mat2) *Mat4 {
	*t = Mat4{
		vector4.Vector{t[0][0], t[0][1], 0, 0},
		vector4.Vector{t[1][0], t[1][1], 0, 0},
		vector4.Vector{0, 0, 1, 0},
		vector4.Vector{0, 0, 0, 1},
	}
	return t
}

func (t *Mat4) AssignMat3x3(m *mat3.Mat3) *Mat4 {
	*t = Mat4{
		vector4.Vector{t[0][0], t[0][1], t[0][2], 0},
		vector4.Vector{t[1][0], t[1][1], t[1][2], 0},
		vector4.Vector{t[2][0], t[2][1], t[2][2], 0},
		vector4.Vector{0, 0, 0, 1},
	}
	return t
}

// v` = v * M
func (t *Mat4) MulVec4(v *vector4.Vector) vector4.Vector {
	return vector4.Vector{
		t[0][0]*v[0] + t[1][0]*v[1] + t[2][0]*v[2] + t[3][0]*v[3],
		t[0][1]*v[0] + t[1][1]*v[1] + t[2][1]*v[2] + t[3][1]*v[3],
		t[0][2]*v[0] + t[1][2]*v[1] + t[2][2]*v[2] + t[3][2]*v[3],
		t[0][3]*v[0] + t[1][3]*v[1] + t[2][3]*v[2] + t[3][3]*v[3],
	}
}

func (t *Mat4) AssignMul(a, b *Mat4) *Mat4 {
	t[0] = a.MulVec4(&b[0])
	t[1] = a.MulVec4(&b[1])
	t[2] = a.MulVec4(&b[2])
	t[3] = a.MulVec4(&b[3])
	return t
}

func (t *Mat4) TransformVec4(v *vector4.Vector) {
	x := t[0][0]*v[0] + t[1][0]*v[1] + t[2][0]*v[2] + t[3][0]*v[3]
	y := t[0][1]*v[0] + t[1][1]*v[1] + t[2][1]*v[2] + t[3][1]*v[3]
	z := t[0][2]*v[0] + t[1][2]*v[1] + t[2][2]*v[2] + t[3][2]*v[3]
	v[3] = t[0][3]*v[0] + t[1][3]*v[1] + t[2][3]*v[2] + t[3][3]*v[3]
	v[0] = x
	v[1] = y
	v[2] = z
}

func (t *Mat4) MulVec3(v *vector3.Vector) vector3.Vector {
	v4 := vector4.Vector{v[0], v[1], v[2], 1}
	v4 = t.MulVec4(&v4)
	return v4.Vec3DividedByW()
}

func (t *Mat4) TransformVec3(v *vector3.Vector) {
	x := t[0][0]*v[0] + t[1][0]*v[1] + t[2][0]*v[2] + t[3][0]
	y := t[0][1]*v[0] + t[1][1]*v[1] + t[2][1]*v[2] + t[3][1]
	z := t[0][2]*v[0] + t[1][2]*v[1] + t[2][2]*v[2] + t[3][2]
	w := t[0][3]*v[0] + t[1][3]*v[1] + t[2][3]*v[2] + t[3][3]
	if sutil.FloatEqual(w, 0) {
		w = 1
	}
	oow := 1 / w
	v[0] = x * oow
	v[1] = y * oow
	v[2] = z * oow
}

func (t *Mat4) MulVec3W(v *vector3.Vector, w float32) vector3.Vector {
	result := *v
	t.TransformVec3W(&result, w)
	return result
}

func (t *Mat4) TransformVec3W(v *vector3.Vector, w float32) {
	x := t[0][0]*v[0] + t[1][0]*v[1] + t[2][0]*v[2] + t[3][0]*w
	y := t[0][1]*v[0] + t[1][1]*v[1] + t[2][1]*v[2] + t[3][1]*w
	v[2] = t[0][2]*v[0] + t[1][2]*v[1] + t[2][2]*v[2] + t[3][2]*w
	v[0] = x
	v[1] = y
}

func (t *Mat4) SetTranslation(v *vector3.Vector) *Mat4 {
	t[3][0] = v[0]
	t[3][1] = v[1]
	t[3][2] = v[2]
	return t
}

func (t *Mat4) Translate(v *vector3.Vector) *Mat4 {
	t[3][0] += v[0]
	t[3][1] += v[1]
	t[3][2] += v[2]
	return t
}

func (t *Mat4) TranslateX(dx float32) *Mat4 {
	t[3][0] += dx
	return t
}

func (t *Mat4) TranslateY(dy float32) *Mat4 {
	t[3][1] += dy
	return t
}

func (t *Mat4) TranslateZ(dz float32) *Mat4 {
	t[3][2] += dz
	return t
}

func (t *Mat4) Scaling() vector4.Vector {
	return vector4.Vector{t[0][0], t[1][1], t[2][2], t[3][3]}
}

func (t *Mat4) SetScaling(s *vector4.Vector) *Mat4 {
	t[0][0] = s[0]
	t[1][1] = s[1]
	t[2][2] = s[2]
	t[3][3] = s[3]
	return t
}

func (t *Mat4) ScaleVec3(s *vector3.Vector) *Mat4 {
	t[0][0] *= s[0]
	t[1][1] *= s[1]
	t[2][2] *= s[2]
	return t
}

func (t *Mat4) AssignXRotation(angle float32) *Mat4 {
	sina, cosa := math.Sincos(angle)

	t[0][0] = 1
	t[0][1] = 0
	t[0][2] = 0
	t[0][3] = 0

	t[1][0] = 0
	t[1][1] = cosa
	t[1][2] = sina
	t[1][3] = 0

	t[2][0] = 0
	t[2][1] = -sina
	t[2][2] = cosa
	t[2][3] = 0

	t[3][0] = 0
	t[3][1] = 0
	t[3][2] = 0
	t[3][3] = 1

	return t
}

func (t *Mat4) AssignYRotation(angle float32) *Mat4 {
	sina, cosa := math.Sincos(angle)

	t[0][0] = cosa
	t[0][1] = 0
	t[0][2] = -sina
	t[0][3] = 0

	t[1][0] = 0
	t[1][1] = 1
	t[1][2] = 0
	t[1][3] = 0

	t[2][0] = sina
	t[2][1] = 0
	t[2][2] = cosa
	t[2][3] = 0

	t[3][0] = 0
	t[3][1] = 0
	t[3][2] = 0
	t[3][3] = 1

	return t
}

func (t *Mat4) AssignZRotation(angle float32) *Mat4 {
	sina, cosa := math.Sincos(angle)

	t[0][0] = cosa
	t[0][1] = sina
	t[0][2] = 0
	t[0][3] = 0

	t[1][0] = -sina
	t[1][1] = cosa
	t[1][2] = 0
	t[1][3] = 0

	t[2][0] = 0
	t[2][1] = 0
	t[2][2] = 1
	t[2][3] = 0

	t[3][0] = 0
	t[3][1] = 0
	t[3][2] = 0
	t[3][3] = 1

	return t
}

func (t *Mat4) AssignCoordinateSystem(x, y, z *vector3.Vector) *Mat4 {
	t[0][0] = x[0]
	t[0][1] = x[1]
	t[0][2] = x[2]
	t[0][3] = 0

	t[1][0] = y[0]
	t[1][1] = y[1]
	t[1][2] = y[2]
	t[1][3] = 0

	t[2][0] = z[0]
	t[2][1] = z[1]
	t[2][2] = z[2]
	t[2][3] = 0

	t[3][0] = 0
	t[3][1] = 0
	t[3][2] = 0
	t[3][3] = 1

	return t
}

// 通过euler构建mat3
func (t *Mat4) AssignEulerRotation(yHead, xPitch, zBank float32) *Mat4 {
	xPitch, yHead, zBank = sutil.CanonizeEuler(xPitch, yHead, zBank)

	sh, ch := math.Sincos(yHead)
	sp, cp := math.Sincos(xPitch)
	sb, cb := math.Sincos(zBank)

	t[0][0] = ch*cb + sh*sp*sb
	t[0][1] = sb * cp
	t[0][2] = -sh*cb + ch*sp*sb
	t[0][3] = 0

	t[1][0] = -ch*sb + sh*sp*cb
	t[1][1] = cb * cp
	t[1][2] = sb*sh + ch*sp*cb
	t[1][3] = 0

	t[2][0] = sh * cp
	t[2][1] = -sp
	t[2][2] = ch * cp
	t[2][3] = 0

	t[3][0] = 0
	t[3][1] = 0
	t[3][2] = 0
	t[3][3] = 1

	return t
}

// 提取euler
func (t *Mat4) ExtractEulerAngles() (yHead, xPitch, zBank float32) {
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

//
func (t *Mat4) Det3x3() float32 {
	return t[0][0]*t[1][1]*t[2][2] +
		t[0][1]*t[1][2]*t[2][0] +
		t[0][2]*t[1][0]*t[2][1] -
		t[0][2]*t[1][1]*t[2][0] -
		t[0][1]*t[1][0]*t[2][2] -
		t[0][0]*t[1][2]*t[2][1]
}

func (t *Mat4) Det() float32 {
	return t[3][0]*t[2][1]*t[1][2]*t[0][3] - t[2][0]*t[3][1]*t[1][2]*t[0][3] - t[3][0]*t[1][1]*t[2][2]*t[0][3] + t[1][0]*t[3][1]*t[2][2]*t[0][3] +
		t[2][0]*t[1][1]*t[3][2]*t[0][3] - t[1][0]*t[2][1]*t[3][2]*t[0][3] - t[3][0]*t[2][1]*t[0][2]*t[1][3] + t[2][0]*t[3][1]*t[0][2]*t[1][3] +
		t[3][0]*t[0][1]*t[2][2]*t[1][3] - t[0][0]*t[3][1]*t[2][2]*t[1][3] - t[2][0]*t[0][1]*t[3][2]*t[1][3] + t[0][0]*t[2][1]*t[3][2]*t[1][3] +
		t[3][0]*t[1][1]*t[0][2]*t[2][3] - t[1][0]*t[3][1]*t[0][2]*t[2][3] - t[3][0]*t[0][1]*t[1][2]*t[2][3] + t[0][0]*t[3][1]*t[1][2]*t[2][3] +
		t[1][0]*t[0][1]*t[3][2]*t[2][3] - t[0][0]*t[1][1]*t[3][2]*t[2][3] - t[2][0]*t[1][1]*t[0][2]*t[3][3] + t[1][0]*t[2][1]*t[0][2]*t[3][3] +
		t[2][0]*t[0][1]*t[1][2]*t[3][3] - t[0][0]*t[2][1]*t[1][2]*t[3][3] - t[1][0]*t[0][1]*t[2][2]*t[3][3] + t[0][0]*t[1][1]*t[2][2]*t[3][3]
}

//
func (t *Mat4) Transpose() *Mat4 {
	t[0][1], t[1][0] = t[1][0], t[0][1]
	t[0][2], t[2][0] = t[2][0], t[0][2]
	t[0][3], t[3][0] = t[3][0], t[0][3]

	t[1][2], t[2][1] = t[2][1], t[1][2]
	t[1][3], t[3][1] = t[3][1], t[1][3]
	t[2][3], t[3][2] = t[3][2], t[2][3]

	return t
}

func (t *Mat4) Transposed() Mat4 {
	l_temp := *t
	l_temp[0][1], l_temp[1][0] = l_temp[1][0], l_temp[0][1]
	l_temp[0][2], l_temp[2][0] = l_temp[2][0], l_temp[0][2]
	l_temp[0][3], l_temp[3][0] = l_temp[3][0], l_temp[0][3]

	l_temp[1][2], l_temp[2][1] = l_temp[2][1], l_temp[1][2]
	l_temp[1][3], l_temp[3][1] = l_temp[3][1], l_temp[1][3]
	l_temp[2][3], l_temp[3][2] = l_temp[3][2], l_temp[2][3]

	return l_temp
}

func (t *Mat4) maskedBlock(blockI, blockJ int) *mat3.Mat3 {
	var m mat3.Mat3
	m_i := 0
	for i := 0; i < 4; i++ {
		if i == blockI {
			continue
		}
		m_j := 0
		for j := 0; j < 4; j++ {
			if j == blockJ {
				continue
			}
			m[m_i][m_j] = t[i][j]
			m_j++
		}
		m_i++
	}
	return &m
}

// adj
func (t *Mat4) adjugate() *Mat4 {
	matOri := *t
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			t[i][j] = matOri.maskedBlock(i, j).Det() * float32(((i+j)%2)*-2+1)
		}
	}

	return t.Transpose()
}

func (t *Mat4) adjugated() Mat4 {
	result := *t
	result.adjugate()
	return result
}

func (t *Mat4) Inv() *Mat4 {
	initialDet := t.Det()
	t.adjugate()
	t.Mul(1 / initialDet)
	return t
}

func (t *Mat4) Inverted() Mat4 {
	result := *t
	result.Inv()
	return result
}
