package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Недокалькулятор")

	label1 := widget.NewLabel("Введите первое число:")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Введите второе число:")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("")

	btn := widget.NewButton("ПОСЧИТАЙ", func() {
		n1, err := strconv.ParseFloat(entry1.Text, 64)
		n2, er := strconv.ParseFloat(entry2.Text, 64)

		if err != nil || er != nil {
			answer.SetText("ОШИБКА ВВОДА!")
		} else {
			sum := n1 + n2
			sub := n1 - n2
			mul := n1 * n2
			div := n1 / n2

			answer.SetText(fmt.Sprintf("(+) %f\n(-) %f\n(*) %f\n(/) %f", sum, sub, mul, div))
		}
	})

	w.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		btn,
		answer,
	))

	w.ShowAndRun()
}
