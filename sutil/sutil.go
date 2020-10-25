/*
 * @Author: sealon
 * @Date: 2020-10-26 17:38:02
 * @Last Modified by: sealon
 * @Last Modified time: 2020-10-26 21:03:58
 * @Desc:
 */
package sutil

import math "github.com/barnex/fmath"

// 限制在[-pi,pi]
func WrapPi(theta float32) float32 {
	// if math.Abs(math.Abs(theta)-KPi) <= 0.0001 {
	// 	return theta
	// }
	theta += math.Pi
	theta -= math.Floor(theta*K1Over2Pi) * K2Pi
	theta -= math.Pi
	return theta
}

// 			   90,-50,200
// func Canonize(pitch, heading, bank float32) {
// 	fmt.Println("input   : ", pitch*sutil.Rad2Deg, heading*sutil.Rad2Deg, bank*sutil.Rad2Deg)
// 	// 先将pitch限制在[-pi,pi]
// 	pitch = sutil.WrapPi(pitch)

// 	90
// 	// 将pitch限制在[-90,90]
// 	if pitch < -sutil.KPiOver2 {
// 		fmt.Println("111")
// 		// 若pitch(-90,-180]
// 		pitch = -sutil.KPi - pitch
// 		heading += pitch
// 		bank += pitch
// 	} else if pitch > sutil.KPiOver2 {
// 		// 若pitch(90,180]
// 		fmt.Println("222")
// 		pitch = sutil.KPi - pitch    90
// 		heading += sutil.KPi         130
// 		bank += sutil.KPi            430
// 	}
// 	if math.Abs(pitch) > (sutil.KPiOver2 - 0.0001) {
// 		fmt.Println("333:", math.Abs(pitch), sutil.KPiOver2-0.0001)
// 		heading += bank       560
// 		bank = 0.0
// 	} else {
// 		fmt.Println("444")
// 		bank = sutil.WrapPi(bank)
// 	}
// 	heading = sutil.WrapPi(heading)
// 	fmt.Println("output: ", pitch*sutil.Rad2Deg, heading*sutil.Rad2Deg, bank*sutil.Rad2Deg)
// }
