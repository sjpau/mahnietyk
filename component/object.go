package component

import "github.com/theonlymoby/mahnietyk/assets"

type Object struct {
	X      float64
	Y      float64
	VX     float64
	VY     float64
	Alive  bool
	Width  int
	Height int
}

func (o *Object) CollideWith(n *Object) bool {
	if o.X+float64(o.Width) >= n.X &&
		o.X <= n.X+float64(n.Width) &&
		o.Y+float64(o.Height) >= n.Y &&
		o.Y <= n.Y+float64(n.Height) {
		return true
	}
	return false
}

func (o *Object) Die() {
	o.Alive = false
	assets.BubbleHit.Rewind()
	assets.BubbleHit.Play()
}
