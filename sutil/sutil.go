/*
 * @Author: sealon
 * @Date: 2020-10-26 17:38:02
 * @Last Modified by: sealon
 * @Last Modified time: 2020-10-27 19:29:57
 * @Desc:
 */
package sutil

import (
	math "github.com/barnex/fmath"
)

func Abs(a float32) float32 {
	if a < 0 {
		return -a
	} else if a == 0 {
		return 0
	}

	return a
}

func FloatEqual(a, b float32) bool {
	return FloatEqualThreshold(a, b, Epsilon)
}

func FloatEqualThreshold(a, b, epsilon float32) bool {
	if a == b {
		return true
	}

	// diff := math.Abs(a - b)
	// if a*b == 0 || diff < MinNormal {
	// 	return diff < epsilon*epsilon
	// }

	// return diff/(Abs(a)+Abs(b)) < epsilon

	if a > b {
		return a-b < epsilon
	} else {
		return b-a < epsilon
	}
}

func Clamp(a, low, high float32) float32 {
	if a < low {
		return low
	} else if a > high {
		return high
	}

	return a
}

func ClampFunc(low, high float32) func(float32) float32 {
	return func(a float32) float32 {
		return Clamp(a, low, high)
	}
}

func IsClamped(a, low, high float32) bool {
	return a >= low && a <= high
}

// min,max
func SetMin(a, b *float32) {
	if *b < *a {
		*a = *b
	}
}

// max,min
func SetMax(a, b *float32) {
	if *a < *b {
		*a = *b
	}
}

func Round(v float32, precision int) float32 {
	p := float32(precision)
	t := v * math.Pow(10, p)
	if t > 0 {
		return math.Floor(t+0.5) / math.Pow(10, p)
	}
	return math.Ceil(t-0.5) / math.Pow(10, p)
}

// [-pi,pi]
func WrapPi(theta float32) float32 {
	// for theta > math.Pi {
	// 	theta -= K2Pi
	// }
	// for theta < -math.Pi {
	// 	theta += K2Pi
	// }
	// return theta

	theta += math.Pi
	theta -= math.Floor(theta*K1Over2Pi) * K2Pi
	theta -= math.Pi

	return theta

}

// [-180,180]
func WrapAngle(angle float32) float32 {
	for angle > 180 {
		angle -= 360
	}
	for angle < -180 {
		angle += 360
	}
	return angle
}

// [0,360]
func WrapAngle360(angle float32) float32 {
	angle = WrapAngle(angle)
	if angle < 0 {
		angle += 360
	}
	return angle
}

// 限制欧拉 pitch[-pi/2,pi/2] heading[-pi,pi] bank[-pi,pi]
func CanonizeEuler(pitch, heading, bank float32) (rp, rh, rb float32) {
	pitch = WrapPi(pitch)
	if pitch < -KPiOver2 {
		pitch = -math.Pi - pitch
		if pitch > 0 {
			heading += pitch
			bank += pitch
		} else {
			heading += math.Pi
			bank += math.Pi
		}

	} else if pitch > KPiOver2 {
		pitch = math.Pi - pitch
		if pitch >= 0 {
			heading += math.Pi
			bank += math.Pi
		} else {
			heading += pitch
			bank += pitch
		}
	}

	if math.Abs(pitch) > (KPiOver2 - 0.001) {
		if pitch > 0 {
			heading -= bank
		} else {
			heading += bank
		}

		bank = 0.0
	} else {
		bank = WrapPi(bank)
	}
	heading = WrapPi(heading)
	return pitch, heading, bank
}

// 限制欧拉 pitch[-90,90] heading[-180,180] bank[-180,180]
func CanonizeEulerAngle(pitch, heading, bank float32) (rp, rh, rb float32) {
	pitch = WrapAngle(pitch)
	if pitch < -90 {
		pitch = -180 - pitch
		if pitch > 0 {
			heading += pitch
			bank += pitch
		} else {
			heading += 180
			bank += 180
		}

	} else if pitch > 90 {
		pitch = 180 - pitch
		if pitch >= 0 {
			heading += 180
			bank += 180
		} else {
			heading += pitch
			bank += pitch
		}
	}

	if math.Abs(pitch) > (90 - 0.001) {
		if pitch > 0 {
			heading -= bank
		} else {
			heading += bank
		}

		bank = 0.0
	} else {
		bank = WrapAngle(bank)
	}
	heading = WrapAngle(heading)
	return pitch, heading, bank
}
