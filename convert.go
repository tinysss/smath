/*
 * @Author: sealon
 * @Date: 2020-11-05 18:05:17
 * @Last Modified by: sealon
 * @Last Modified time: 2020-11-05 21:11:53
 * @Desc: 转换
 */
package smath

import (
	"github.com/tinysss/smath/mat3"
	"github.com/tinysss/smath/quat"
)

// quat ->mat3
func QuatToMat3(quat *quat.Quaternion) mat3.Mat3 {
	return mat3.Mat3{
		{1 - 2*quat[1]*quat[1] - 2*quat[2]*quat[2], 2*quat[0]*quat[1] + 2*quat[3]*quat[2], 2*quat[0]*quat[3] - 2*quat[3]*quat[1]},
		{2*quat[0]*quat[1] - 2*quat[3]*quat[2], 1 - 2*quat[0]*quat[0] - 2*quat[2]*quat[2], 2*quat[1]*quat[2] + 2*quat[3]*quat[0]},
		{2*quat[0]*quat[2] + 2*quat[3]*quat[1], 2*quat[1]*quat[2] - 2*quat[3]*quat[0], 1 - 2*quat[0]*quat[0] - 2*quat[1]*quat[1]},
	}
}

// mat3 ->　quat
func Mat3ToQuat(mat3 *mat3.Mat3) quat.Quaternion {
	return quat.Ident
}
