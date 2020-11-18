package main

import (
	"flag"
        "github.com/Sheenam3/x-tracer-gocui/ui"
        "github.com/Sheenam3/x-tracer-gocui/pkg"
        "github.com/Sheenam3/x-tracer-gocui/database"
        "github.com/Sheenam3/x-tracer-gocui/events"
//        "log"
 //       "time"

)



func main() {

	database.Init()


	ui.SubscribeListeners()
	pkg.SubscribeListeners()

	go events.Run()


	port := flag.String("port", "6666", "")
	pkg.SetPort(*port)
        go pkg.StartServer()

        ui.InitGui()

/*        for {
                log.Println("From x-tracer- Sleeping")
                time.Sleep(10 * time.Second)
        }*/


}

