/*
 * @Author: sealon
 * @Date: 2020-09-10 17:29:11
 * @Last Modified by: sealon
 * @Last Modified time: 2020-09-20 14:35:18
 * @Desc: 向量或矩阵的公共接口
 */
package generic

type T interface {
	Cols() int

	Rows() int

	Size() int

	Slice() []float32

	Get(col, row int) float32

	IsZero() bool
}
