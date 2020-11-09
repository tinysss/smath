/*
 * @Author: sealon
 * @Date: 2020-11-05 18:05:17
 * @Last Modified by: sealon
 * @Last Modified time: 2020-11-09 15:13:24
 * @Desc: 转换
 */
package smath

import (
	math "github.com/barnex/fmath"
	"github.com/tinysss/smath/mat3"
	"github.com/tinysss/smath/quat"
)

// quat ->mat3
func QuatToMat3(quat *quat.Quaternion) mat3.Mat3 {
	return mat3.Mat3{
		{1 - 2*quat[1]*quat[1] - 2*quat[2]*quat[2], 2*quat[0]*quat[1] + 2*quat[3]*quat[2], 2*quat[0]*quat[2] - 2*quat[3]*quat[1]},
		{2*quat[0]*quat[1] - 2*quat[3]*quat[2], 1 - 2*quat[0]*quat[0] - 2*quat[2]*quat[2], 2*quat[1]*quat[2] + 2*quat[3]*quat[0]},
		{2*quat[0]*quat[2] + 2*quat[3]*quat[1], 2*quat[1]*quat[2] - 2*quat[3]*quat[0], 1 - 2*quat[0]*quat[0] - 2*quat[1]*quat[1]},
	}
}

// mat3 ->　quat
func Mat3ToQuat(mat3 *mat3.Mat3) quat.Quaternion {

	l_quat := quat.Ident
	l_tr := mat3.Trace()
	l_x := mat3.Get(0, 0) - mat3.Get(1, 1) - mat3.Get(2, 2)
	l_y := mat3.Get(1, 1) - mat3.Get(0, 0) - mat3.Get(2, 2)
	l_z := mat3.Get(2, 2) - mat3.Get(0, 0) - mat3.Get(1, 1)
	l_w := l_tr

	l_bigidx := 0
	l_bigval := l_x
	if l_y > l_bigval {
		l_bigidx = 1
		l_bigval = l_y
	}
	if l_z > l_bigval {
		l_bigidx = 2
		l_bigval = l_z
	}
	if l_w > l_bigval {
		l_bigidx = 3
		l_bigval = l_w
	}

	l_bigval = math.Sqrt(l_bigval+1.0) * 0.5 // s
	l_scale := 0.25 / l_bigval               //1/4s
	switch l_bigidx {
	case 3: // w
		l_quat[3] = l_bigval
		l_quat[0] = (mat3.Get(1, 2) - mat3.Get(2, 1)) * l_scale
		l_quat[1] = (mat3.Get(2, 0) - mat3.Get(0, 2)) * l_scale
		l_quat[2] = (mat3.Get(0, 1) - mat3.Get(1, 0)) * l_scale
	case 0: // x
		l_quat[0] = l_bigval
		l_quat[3] = (mat3.Get(1, 2) - mat3.Get(2, 1)) * l_scale
		l_quat[1] = (mat3.Get(0, 1) + mat3.Get(1, 0)) * l_scale
		l_quat[2] = (mat3.Get(2, 0) + mat3.Get(0, 2)) * l_scale

	case 1: // y
		l_quat[1] = l_bigval
		l_quat[3] = (mat3.Get(2, 0) - mat3.Get(0, 2)) * l_scale
		l_quat[0] = (mat3.Get(0, 1) + mat3.Get(1, 0)) * l_scale
		l_quat[2] = (mat3.Get(1, 2) + mat3.Get(2, 1)) * l_scale

	case 2: // z
		l_quat[2] = l_bigval
		l_quat[3] = (mat3.Get(0, 1) - mat3.Get(1, 0)) * l_scale
		l_quat[0] = (mat3.Get(2, 0) + mat3.Get(0, 2)) * l_scale
		l_quat[1] = (mat3.Get(1, 2) + mat3.Get(2, 1)) * l_scale
	}

	return l_quat.Normalized()
}

// euler -> mat3
// func QuatToMat3(quat *quat.Quaternion) mat3.Mat3 {
// }

// mat3 -> euler
// func QuatToMat3(quat *quat.Quaternion) mat3.Mat3 {
// }
