/*
 * @Author: sealon
 * @Date: 2020-10-26 17:38:17
 * @Last Modified by: sealon
 * @Last Modified time: 2020-10-27 16:12:03
 * @Desc:
 */
package sutil

import (
	math "github.com/barnex/fmath"
)

var Epsilon float32 = 1e-4
var MinNormal = float32(1.1754943508222875e-38)
var MinValue = float32(math.SmallestNonzeroFloat32)
var MaxValue = float32(math.MaxFloat32)

const KPi = math.Pi
const K2Pi = KPi * 2.0
const KPiOver2 = KPi / 2.0
const K1OverPi = 1.0 / KPi
const K1Over2Pi = 1.0 / K2Pi

const Rad2Deg = 180.0 / math.Pi
const Deg2Rad = math.Pi / 180
