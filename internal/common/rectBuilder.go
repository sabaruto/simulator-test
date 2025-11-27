package common

import "github.com/tfriedel6/canvas"

type RectBuilder struct {
	cv           *canvas.Canvas
	x            float64
	y            float64
	width        float64
	height       float64
	fillColour   string
	strokeColour string
}

func NewRectBuilder() *RectBuilder {
	return &RectBuilder{}
}

func (r *RectBuilder) X(x float64) *RectBuilder {
	r.x = x
	return r
}

func (r *RectBuilder) Y(y float64) *RectBuilder {
	r.y = y
	return r
}

func (r *RectBuilder) Width(width float64) *RectBuilder {
	r.width = width
	return r
}

func (r *RectBuilder) Height(height float64) *RectBuilder {
	r.height = height
	return r
}

func (r *RectBuilder) FillColour(fillColour string) *RectBuilder {
	r.fillColour = fillColour
	return r
}

func (r *RectBuilder) StrokeColour(strokeColour string) *RectBuilder {
	r.strokeColour = strokeColour
	return r
}

func (r *RectBuilder) Canvas(cv *canvas.Canvas) *RectBuilder {
	r.cv = cv
	return r
}

func (r *RectBuilder) Build() {
	if len(r.fillColour) > 0 {
		r.cv.SetFillStyle(r.fillColour)
	}

	if len(r.strokeColour) > 0 {
		r.cv.SetStrokeStyle(r.strokeColour)
	}

	r.cv.BeginPath()
	r.cv.FillRect(r.x, r.y, r.width, r.height)
	r.cv.StrokeRect(r.x, r.y, r.width, r.height)
}
