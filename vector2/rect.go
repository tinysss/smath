/*
 * @Author: sealon
 * @Date: 2020-09-18 15:30:28
 * @Last Modified by: sealon
 * @Last Modified time: 2020-09-18 16:24:01
 * @Desc: rect封装
 */
package vector2

type Rect struct {
	Min Vector
	Max Vector
}

func NewRect(min, max Vector) *Rect {
	return &Rect{min, max}
}

// 点包含
func (t *Rect) ContainsPoint(pt *Vector) bool {
	return pt[0] >= t.Min[0] && pt[0] <= t.Max[0] &&
		pt[1] >= t.Min[1] && pt[1] <= t.Max[1]
}

// rect包含
func (t *Rect) Contains(o *Rect) bool {
	return o.Min[0] >= t.Min[0] && o.Max[0] <= t.Max[0] &&
		o.Min[1] >= t.Min[1] && o.Max[1] <= t.Max[1]
}

// rect相交
func (t *Rect) Intersects(o *Rect) bool {
	return (o.Min[0] <= t.Max[0] && o.Max[0] >= t.Min[0] && o.Min[1] <= t.Max[1] && o.Max[1] >= t.Min[1]) ||
		(t.Min[0] <= o.Max[0] && t.Max[0] >= o.Min[0] && t.Min[1] <= o.Max[1] && t.Max[1] >= o.Min[1])
}
