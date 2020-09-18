package vector3

import "math"

type Box struct {
	Min Vector
	Max Vector
}

func NewBox(min, max Vector) *Box {
	return &Box{min, max}
}

// 点包含
func (t *Box) ContainsPoint(pt *Vector) bool {
	return pt[0] >= t.Min[0] && pt[0] <= t.Max[0] &&
		pt[1] >= t.Min[1] && pt[1] <= t.Max[1] &&
		pt[2] >= t.Min[2] && pt[2] <= t.Max[2]
}

// 中心点
func (t *Box) Center() Vector {
	c := Add(&t.Min, &t.Max)
	c.Scale(0.5)
	return c
}

// 相交
func (t *Box) Intersects(o *Box) bool {
	if t.Min[0] > o.Max[0] || t.Max[0] < o.Min[0] {
		return false
	}
	if t.Min[1] > o.Max[1] || t.Max[1] < o.Min[1] {
		return false
	}
	if t.Min[2] > o.Max[2] || t.Max[2] < o.Min[2] {
		return false
	}
	return true
}

// 相交 并返回相交的那个box
func (t *Box) Intersects2(o *Box) *Box {
	if t.Min[0] > o.Max[0] || t.Max[0] < o.Min[0] {
		return nil
	}
	if t.Min[1] > o.Max[1] || t.Max[1] < o.Min[1] {
		return nil
	}
	if t.Min[2] > o.Max[2] || t.Max[2] < o.Min[2] {
		return nil
	}
	return &Box{
		Min: Vector{float32(math.Max(float64(t.Min[0]), float64(o.Min[0]))), float32(math.Max(float64(t.Min[1]), float64(o.Min[1]))), float32(math.Max(float64(t.Min[2]), float64(o.Min[2])))},
		Max: Vector{float32(math.Min(float64(t.Max[0]), float64(o.Max[0]))), float32(math.Min(float64(t.Max[1]), float64(o.Max[1]))), float32(math.Min(float64(t.Max[2]), float64(o.Max[2])))},
	}
}

// 合并放大box
func (t *Box) Join(o *Box) {
	t.Min = Min(&t.Min, &o.Min)
	t.Max = Max(&t.Max, &o.Max)
}

// 合并放大box
func Joined(a, o *Box) *Box {
	var joinbox Box
	joinbox.Min = Min(&a.Min, &o.Min)
	joinbox.Max = Max(&a.Max, &o.Max)
	return &joinbox
}
