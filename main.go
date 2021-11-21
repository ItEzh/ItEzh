package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Подпишись на канал!")
	w.Resize(fyne.NewSize(300, 400))

	url, _ := url.Parse("https://www.youtube.com/channel/UCXX0w2eLads0EJeRrAkl_Xg/videos")

	link := widget.NewHyperlink("Подпишись!", url)

	w.SetContent(container.NewVBox(link))

	w.ShowAndRun()
}
