package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Подпишись на канал!")
	w.Resize(fyne.NewSize(300, 500))

	title := widget.NewLabel("Оформление заказа")

	name_label := widget.NewLabel("Ваше имя:")
	name := widget.NewEntry()

	food_label := widget.NewLabel("Выберите еду для заказа:")
	food := widget.NewCheckGroup([]string{"Пицца", "Торт", "Наггетсы", "Бургер", "Кола"}, func(s []string) {})

	male_label := widget.NewLabel("Ваш пол:")
	male := widget.NewRadioGroup([]string{"Мужской", "Женский"}, func(s string) {})

	result := widget.NewLabel("")

	btn := widget.NewButton("ОФОРМИТЬ", func() {
		username := name.Text
		food_arr := food.Selected
		male_val := male.Selected

		result.SetText(username + "\n" + male_val + "\n")

		for _, i := range food_arr {
			result.SetText(result.Text + i + "\n")
		}
	})

	w.SetContent(container.NewVBox(
		title,
		name_label,
		name,
		food_label,
		food,
		male_label,
		male,
		btn,
		result,
	))

	w.ShowAndRun()
}
