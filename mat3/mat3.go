/*
 * @Author: sealon
 * @Date: 2020-09-29 14:50:05
 * @Last Modified by: sealon
 * @Last Modified time: 2020-09-30 18:01:50
 * @Desc:
 */
package mat3

import (
	"fmt"
	"unsafe"

	"github.com/tinysss/smath/generic"
	"github.com/tinysss/smath/mat2"
	"github.com/tinysss/smath/vector3"
	"github.com/ungerik/go3d/vec2"
)

type Mat3 [3]vector3.Vector // 列存储

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

// v' = M * v
func (t *Mat3) MulVec3(v *vector3.Vector) vector3.Vector {
	return vector3.Vector{
		t[0][0]*v[0] + t[1][0]*v[1] + t[2][0]*v[2],
		t[0][1]*v[0] + t[1][1]*v[1] + t[2][1]*v[2],
		t[0][2]*v[0] + t[1][2]*v[1] + t[2][2]*v[2],
	}
}

func (t *Mat3) AssignMul(a, b *Mat3) *Mat3 {
	t[0] = a.MulVec3(&b[0])
	t[1] = a.MulVec3(&b[1])
	t[2] = a.MulVec3(&b[2])
	return t
}

func (t *Mat3) AssignMat2x2(m *mat2.Mat2) *Mat3 {
	*t = Mat3{
		vector3.Vector{m[0][0], m[1][0], 0},
		vector3.Vector{m[0][1], m[1][1], 0},
		vector3.Vector{0, 0, 1},
	}
	return t
}
