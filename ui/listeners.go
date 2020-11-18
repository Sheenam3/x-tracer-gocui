package ui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/Sheenam3/x-tracer-gocui/pkg"
	"github.com/Sheenam3/x-tracer-gocui/events"
	//"log"

)

func refreshIntegratedLogs(e events.Event) {


	if e, ok := e.(events.EmptyMessage); ok {

        	g.Update(func(g *gocui.Gui) error {

			pn := e.Pn
			if pn == "tcplife"{
				viewtl, err := g.View("tcplife")
				if err != nil {
					return err
				}
				viewtl.Clear()

				_, _ = fmt.Fprint(viewtl, pkg.GetActiveLogs(pn))

				g.SetViewOnTop("tcplife")
//				g.SetCurrentView("tcplife")

				viewtl.Autoscroll = true

				return nil
			/*}else if pn == "biosnoop"{
                                view, err := g.View("logs")
                                if err != nil {
                                        return err
                                }
                                view.Clear()

                                _, _ = fmt.Fprint(view, pkg.GetActiveLogs(pn))

                                g.SetViewOnTop("logs")
                                g.SetCurrentView("logs")

                                view.Autoscroll = true

                                return nil
			*/}else if pn == "cachestat"{
                                viewcs, err := g.View("cachestat")
                                if err != nil {
                                        return err
                                }
                                viewcs.Clear()

                                _, _ = fmt.Fprint(viewcs, pkg.GetActiveLogs(pn))

                                g.SetViewOnTop("cachestat")
//                                g.SetCurrentView("cachestat")

                                viewcs.Autoscroll = true

                                return nil
			}else if pn == "execsnoop"{
                                viewes, err := g.View("execsnoop")
                                if err != nil {
                                        return err
                                }
                                viewes.Clear()

                                _, _ = fmt.Fprint(viewes, pkg.GetActiveLogs(pn))

                                g.SetViewOnTop("execsnoop")
//                                g.SetCurrentView("execsnoop")

                                viewes.Autoscroll = true

                                return nil
			}else{
				view, err := g.View("tcplogs")
                                if err != nil {
                                        return err
                                }
                                view.Clear()

                                _, _ = fmt.Fprint(view, pkg.GetActiveLogs(pn))
                                g.SetViewOnTop("tcplogs")
//                                g.SetCurrentView("tcplogs")

                                view.Autoscroll = true

                                return nil
			}

		})
	}
}


func refreshSingleLogs(e events.Event) {


	if e, ok := e.(events.EmptyMessage); ok {

        	g.Update(func(g *gocui.Gui) error {

			pn := e.Pn
			view, err := g.View("logs")
			if err != nil {
				return err
			}
			view.Clear()

			_, _ = fmt.Fprint(view, pkg.GetActiveLogs(pn))

			g.SetViewOnTop("logs")
			g.SetCurrentView("logs")

			view.Autoscroll = true

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
	events.Subscribe(refreshIntegratedLogs, "logs:refreshinteg")
	events.Subscribe(refreshSingleLogs, "logs:refreshsingle")

//	events.Subscribe(displayModal, "modal:display")
	
}
