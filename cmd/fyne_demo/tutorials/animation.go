package tutorials

import (
	"image/color"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func makeAnimationScreen(_ fyne.Window) fyne.CanvasObject {
	curves := makeAnimationCurves()
	curves.Move(fyne.NewPos(0, 140+theme.Padding()))
	return fyne.NewContainerWithoutLayout(makeAnimationCanvas(), curves)
}

func makeAnimationCanvas() fyne.CanvasObject {
	rect := canvas.NewRectangle(color.Black)
	rect.Resize(fyne.NewSize(410, 140))

	a := canvas.NewColorRGBAAnimation(theme.PrimaryColorNamed(theme.ColorBlue), theme.PrimaryColorNamed(theme.ColorGreen),
		time.Second*3, func(c color.Color) {
			rect.FillColor = c
			canvas.Refresh(rect)
		})
	a.Repeat = true
	a.Start()

	var a2 *fyne.Animation
	i := widget.NewIcon(theme.CheckButtonCheckedIcon())
	a2 = canvas.NewPositionAnimation(fyne.NewPos(0, 0), fyne.NewPos(350, 80), time.Second*3, func(p fyne.Position) {
		i.Move(p)

		width := int(10 + (float64(p.X) / 7))
		i.Resize(fyne.NewSize(width, width))
	})
	a2.Repeat = true
	a2.Curve = fyne.AnimationLinear
	a2.Start()

	return fyne.NewContainerWithoutLayout(rect, i)
}

func makeAnimationCurves() fyne.CanvasObject {
	label1, box1, a1 := makeAnimationCurveItem("EaseInOut", fyne.AnimationEaseInOut, 0)
	label2, box2, a2 := makeAnimationCurveItem("EaseIn", fyne.AnimationEaseIn, 30+theme.Padding())
	label3, box3, a3 := makeAnimationCurveItem("EaseOut", fyne.AnimationEaseOut, 60+theme.Padding()*2)
	label4, box4, a4 := makeAnimationCurveItem("Linear", fyne.AnimationLinear, 90+theme.Padding()*3)

	start := widget.NewButton("Compare", func() {
		a1.Start()
		a2.Start()
		a3.Start()
		a4.Start()
	})
	start.Resize(start.MinSize())
	start.Move(fyne.NewPos(0, 120+theme.Padding()*4))
	return fyne.NewContainerWithoutLayout(label1, label2, label3, label4, box1, box2, box3, box4, start)
}

func makeAnimationCurveItem(label string, curve fyne.AnimationCurve, yOff int) (
	text *widget.Label, box fyne.CanvasObject, anim *fyne.Animation) {
	text = widget.NewLabel(label)
	text.Alignment = fyne.TextAlignCenter
	text.Resize(fyne.NewSize(380, 30))
	text.Move(fyne.NewPos(0, yOff))
	box = canvas.NewRectangle(theme.TextColor())
	box.Resize(fyne.NewSize(30, 30))
	box.Move(fyne.NewPos(0, yOff))

	anim = canvas.NewPositionAnimation(
		fyne.NewPos(0, yOff), fyne.NewPos(380, yOff), time.Second, func(p fyne.Position) {
			box.Move(p)
		})
	anim.Curve = curve
	return
}
