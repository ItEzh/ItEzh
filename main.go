package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("ПОДПИШИСЬ НА КАНАЛ!")
	w.Resize(fyne.NewSize(500, 500))

	circle1 := canvas.NewCircle(color.NRGBA{70, 8, 163, 240})
	circle1.StrokeColor = color.NRGBA{255, 0, 0, 255}
	circle1.StrokeWidth = 3

	rectangle1 := canvas.NewRectangle(color.NRGBA{17, 94, 6, 255})
	rectangle1.StrokeColor = color.NRGBA{11, 23, 9, 255}
	rectangle1.StrokeWidth = 3

	line1 := canvas.NewLine(color.NRGBA{195, 0, 0, 255})
	line1.StrokeWidth = 5

	icon1 := widget.NewIcon(theme.ComputerIcon())

	btn := widget.NewButton("Нажми", func() {
		rectangle1.Hide()

		circle1.FillColor = color.NRGBA{45, 36, 36, 255}
		circle1.StrokeWidth = 10

		line1.StrokeWidth = 10
		line1.StrokeColor = color.NRGBA{44, 24, 52, 255}
	})

	content1 := container.NewGridWithColumns(2, circle1, rectangle1)
	content2 := container.NewGridWithColumns(2, line1, icon1)

	content := container.NewGridWithRows(3, content1, content2, btn)

	w.SetContent(content)
	w.ShowAndRun()
}
