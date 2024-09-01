package main

import (
	"doom-launcher/internal/view"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			// Manage the rendering state.
			gtx := app.NewContext(&ops, e)

			view.DrawMainView(gtx, theme)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}
	}
}
