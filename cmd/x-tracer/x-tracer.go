package main

import (

        "github.com/Sheenam3/x-tracer-gocui/ui"
        "github.com/Sheenam3/x-tracer-gocui/pkg/streamserver"
        "log"
        "time"

)



func main() {


        server := streamserver.New("6666")
        go server.StartServer()


        ui.InitGui()

        for {
                log.Println("From x-tracer- Sleeping")
                time.Sleep(10 * time.Second)
        }


}

