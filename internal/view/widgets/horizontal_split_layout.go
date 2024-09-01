package widgets

import (
	"image"

	l "gioui.org/layout"
	"gioui.org/op"
)

type HorizontalSplitLayout struct {
	// 0 is center, -1 is completely to the top, 1 is completely to the bottom.
	Ratio float32
}

func (s HorizontalSplitLayout) Layout(top, bottom l.Widget) l.Widget {
	return func(gtx l.Context) l.Dimensions {
		proportion := (s.Ratio + 1) / 2
		topsize := int(proportion * float32(gtx.Constraints.Max.Y))

		bottomoffset := topsize
		bottomsize := gtx.Constraints.Max.Y - bottomoffset

		{
			gtx := gtx
			gtx.Constraints = l.Exact(image.Pt(gtx.Constraints.Max.X, topsize))
			top(gtx)
		}

		{
			gtx := gtx
			gtx.Constraints = l.Exact(image.Pt(gtx.Constraints.Max.X, bottomsize))
			trans := op.Offset(image.Pt(0, bottomoffset)).Push(gtx.Ops)
			bottom(gtx)
			trans.Pop()
		}

		return l.Dimensions{Size: gtx.Constraints.Max}
	}
}

func NewHorizontalSplitLayout(ratio float32, topWidget, bottomWidget l.Widget) l.Widget {
	return HorizontalSplitLayout{Ratio: ratio}.Layout(
		topWidget,
		bottomWidget,
	)
}
