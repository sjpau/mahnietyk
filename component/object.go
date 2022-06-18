package component

type Object struct {
	X      float64
	Y      float64
	VX     float64
	VY     float64
	Alive  bool
	CanDie bool
	Width  int
	Height int
}

func (o *Object) CollideWith(n *Object) bool {
	if o.X+float64(o.Width) >= n.X && // r1 right edge past r2 left
		o.X <= n.X+float64(n.Width) && // r1 left edge past r2 right
		o.Y+float64(o.Height) >= n.Y && // r1 top edge past r2 bottom
		o.Y <= n.Y+float64(n.Height) { // r1 bottom edge past r2 top
		return true
	}
	return false
}
