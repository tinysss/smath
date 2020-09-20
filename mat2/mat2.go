/*
 * @Author: sealon
 * @Date: 2020-09-20 10:01:43
 * @Last Modified by: sealon
 * @Last Modified time: 2020-09-21 01:20:54
 * @Desc:  使用列存储
 */
package mat2

import (
	"fmt"
	"unsafe"

	"github.com/tinysss/smath/generic"
	"github.com/tinysss/smath/vector2"
)

type Mat2 [2]vector2.Vector // 列存储

var (
	Zero = Mat2{}
	// 单位阵
	Ident = Mat2{
		{1, 0},
		{0, 1},
	}
)

func New(v1, v2 vector2.Vector) *Mat2 {
	return &Mat2{v1, v2}
}

func FromNew(other generic.T) *Mat2 {
	r := Ident
	cols := other.Cols()
	rows := other.Rows()

	if cols != rows || cols < 2 || cols > 4 {
		panic(fmt.Sprintf("unsupported type. cols=%d rows=%d ", cols, rows))
	}

	cols = 2
	rows = 2

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			r[col][row] = other.Get(col, row)
		}
	}
	return &r
}

func (t *Mat2) Array() *[4]float32 {
	return (*[4]float32)(unsafe.Pointer(t))
}

//-------------------------------------------- 实现generic.T begin-------------------------------------
func (t *Mat2) Cols() int {
	return 2
}

func (t *Mat2) Rows() int {
	return 2
}

func (t *Mat2) Size() int {
	return 4
}

func (t *Mat2) Slice() []float32 {
	return t.Array()[:]
}

func (t *Mat2) Get(col, row int) float32 {
	return t[col][row]
}

func (t *Mat2) IsZero() bool {
	return *t == Zero
}

//-------------------------------------------- 实现generic.T end -------------------------------------

func (t *Mat2) Scale(f float32) *Mat2 {
	t[0][0] *= f
	t[1][1] *= f
	return t
}

func (t *Mat2) Scaled(f float32) Mat2 {
	r := *t
	return *r.Scale(f)
}

// 返回缩放比例的向量
func (t *Mat2) Scaling() vector2.Vector {
	return vector2.Vector{t[0][0], t[1][1]}
}

func (t *Mat2) SetScaling(s *vector2.Vector) *Mat2 {
	t[0][0] = s[0]
	t[1][1] = s[1]
	return t
}

func (t *Mat2) Trace() float32 {
	return t[0][0] + t[1][1]
}

func (t *Mat2) MulVec2(v *vector2.Vector) vector2.Vector {
	return vector2.Vector{
		t[0][0]*v[0] + t[1][0]*v[1],
		t[0][1]*v[0] + t[1][1]*v[1],
	}
}

func (t *Mat2) AssignMul(a, b *Mat2) *Mat2 {
	t[0] = a.MulVec2(&b[0])
	t[1] = a.MulVec2(&b[1])
	return t
}

// 转置
func (t *Mat2) Transpose() *Mat2 {
	t[0][1], t[1][0] = t[1][0], t[0][1]
	return t
}
