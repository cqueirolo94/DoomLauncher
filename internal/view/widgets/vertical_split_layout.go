package widgets

import (
	"image"

	l "gioui.org/layout"
	"gioui.org/op"
)

type VerticalSplitLayout struct {
	// 0 is center, -1 is completely to the left, 1 is completely to the right.
	Ratio float32
}

func (s VerticalSplitLayout) Layout(left, right l.Widget) l.Widget {
	return func(gtx l.Context) l.Dimensions {
		proportion := (s.Ratio + 1) / 2
		leftsize := int(proportion * float32(gtx.Constraints.Max.X))

		rightoffset := leftsize
		rightsize := gtx.Constraints.Min.X - rightoffset

		{
			gtx := gtx
			gtx.Constraints = l.Exact(image.Pt(leftsize, gtx.Constraints.Max.Y))
			left(gtx)
		}

		{
			gtx := gtx
			gtx.Constraints = l.Exact(image.Pt(rightsize, gtx.Constraints.Max.Y))
			trans := op.Offset(image.Pt(rightoffset, 0)).Push(gtx.Ops)
			right(gtx)
			trans.Pop()
		}

		return l.Dimensions{Size: gtx.Constraints.Max}
	}
}

func NewVerticalSplitLayout(ratio float32, leftWidget, rightWidget l.Widget) l.Widget {
	return VerticalSplitLayout{Ratio: ratio}.Layout(
		leftWidget,
		rightWidget,
	)
}
