package pkg

import (
//	"context"
	"github.com/Sheenam3/x-tracer-gocui/database"
	"github.com/Sheenam3/x-tracer-gocui/events"
//	pb "github.com/Sheenam3/x-tracer-gocui/api"
//	"log"
	"os"
//	"strconv"
)
/*
func sendLog(e events.Event) {
	if e, ok := e.(events.SendLogEvent); ok {
		stream, err := client.RouteLog(context.Background())
		
		if err != nil {
			log.Panic(err)
		}
		n, err := strconv.ParseInt(e.Pid, 10, 64)
		err = stream.Send(&pb.Log{
			Pid:       n,
			ProbeName: e.ProbeName,
			Log:       e.Log,
			TimeStamp: e.TimeStamp,
		})

		if err != nil {
			log.Panic(err)
		}

		resp, err := stream.CloseAndRecv()
	
		log.Printf("Response from the Server ;) : %v", resp.Res)
		if err != nil {
			log.Panic(err)	
			//events.PublishEvent("modal:display", events.DisplayModalEvent{Message: "Connection to peer lost. Retrying"})
			//time.Sleep(1 * time.Millisecond)
			//sendMsg(e)
		} /*else {
			err = database.SaveConversation(activeContact.Address, 0, e.Message)

			if err != nil {
				// @TODO figure out how to handle error
				log.Panic(err)
			}

			events.PublishEvent("modal:hide", events.EmptyMessage{})
			events.PublishEvent("log:refresh", events.EmptyMessage{})
		}

	} else {
		// @TODO if e is not of SendMessageEvent type
		// ignore for the time being
	}
}
*/
func receiveLog(e events.Event) {
	if e, ok := e.(events.ReceiveLogEvent); ok {

		
		err := database.UpdateLogs(e.ProbeName, e.Sys_Time, e.T, e.Pid, e.Pname, e.Ip, e.Saddr, e.Daddr, e.Dport)
		if err != nil {
			// @TODO figure out how to handle error
			os.Exit(1)
		}
		events.PublishEvent("logs:refresh", events.EmptyMessage{})

	}
}

func SubscribeListeners() {
//	events.Subscribe(sendLog, "log:send")
	events.Subscribe(receiveLog, "log:receive")
}
