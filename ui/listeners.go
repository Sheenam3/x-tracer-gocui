package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/Sheenam3/x-tracer-gocui/pkg"
	"github.com/Sheenam3/x-tracer-gocui/events"
	//"log"
)

func refreshLogs(e events.Event) {


	if e, ok := e.(events.EmptyMessage); ok {

        	g.Update(func(g *gocui.Gui) error {
		view, err := g.View("logs")
		if err != nil {
			return err
		}
		view.Clear()
		pn := e.Pn
		_, _ = fmt.Fprint(view, pkg.GetActiveLogs(pn))

		g.SetViewOnTop("logs")
		g.SetCurrentView("logs")
//		ox, oy := view.Origin()
		view.Autoscroll = true
//		view.SetOrigin(ox, oy+1)
/*		view.SetCursor(0,2)
		if l, err := getViewLine(g, view); err != nil || l == "" {
				view.SetCursor(0, 2)
		}*/
		
		return nil
		})
	}
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
