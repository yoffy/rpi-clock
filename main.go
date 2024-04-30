package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"time"
)

const (
	// Raspberry Piのディスプレイサイズに合わせること
	width = 1920
	height = 1080
)

func main() {
	a := app.New()
	w := a.NewWindow("clock")
	w.Resize(fyne.NewSize(width, height))

	t := canvas.NewText("Loading...", color.White)
	t.TextSize = 100
	t.TextStyle.Bold = true
	t.Alignment = fyne.TextAlignCenter
	w.SetContent(container.NewCenter(t))
	w.SetFullScreen(true)

	go func() {
		tick := time.NewTicker(time.Second)
		for {
			<-tick.C
			t.Text = time.Now().Format("15:04:05")
			t.Refresh()
		}
	}()

	w.ShowAndRun()
}
