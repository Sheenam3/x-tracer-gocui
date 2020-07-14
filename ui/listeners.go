package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/Sheenam3/x-tracer-gocui/pkg"
	"github.com/Sheenam3/x-tracer-gocui/events"
	//"log"
)

func refreshLogs(e events.Event) {
	g.Update(func(g *gocui.Gui) error {
		view, err := g.View("logs")
		if err != nil {
			return err
		}
		view.Clear()
		_, _ = fmt.Fprint(view, pkg.GetActiveLogs())

		return nil
	})
}
/*
var isModalDisplayed = false

func displayModal(e events.Event) {
	if !isModalDisplayed {
		maxX, maxY := g.Size()

		var modalMsg string
		if e, ok := e.(events.DisplayModalEvent); ok {
			modalMsg = e.Message

			g.Update(func(g *gocui.Gui) error {
				v, err := g.SetView("modal",
					maxX/2-50,
					maxY/2-2,
					maxX/2+50,
					maxY/2+2,
				)

				if err != nil {
					// @TODO handle error
					if err != gocui.ErrUnknownView {
						return err
					}
				}

				v.Wrap = true
				v.Frame = true

				// explicitly ignore error
				_, _ = fmt.Fprint(v, modalMsg)

				return nil
			})
		}

		isModalDisplayed = true
	}
}
*/
func SubscribeListeners() {
	events.Subscribe(refreshLogs, "logs:refresh")

//	events.Subscribe(displayModal, "modal:display")
	
}
