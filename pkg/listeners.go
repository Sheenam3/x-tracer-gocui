package pkg

import (

	"github.com/Sheenam3/x-tracer-gocui/database"
	"github.com/Sheenam3/x-tracer-gocui/events"
	"os"
//	"fmt"
)


func receiveLog(e events.Event) {
/*	f, err := os.Create("test.txt")
	if err != nil {
     	   fmt.Println(err)
       		 return
    	}
 l, err := f.WriteString("Tamahsa shuru")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }*/
   	if e, ok := e.(events.ReceiveLogEvent); ok {

		tcp := events.ReceiveLogEvent{ProbeName: e.ProbeName,
                                     Sys_Time:  e.Sys_Time,
                                     T:         e.T,
                                     Pid:       e.Pid,
                                     Pname:     e.Pname,
                                     Ip:        e.Ip,
                                     Saddr:     e.Saddr,
                                     Daddr:     e.Daddr,
                                     Dport:     e.Dport,
                                     Sport:     e.Sport,
	//			     Tx_kb:	"0",
	//			     Rx_kb:	"0",
	//			     Ms:	"0",
		}
		tcplogs := database.TcpLog(tcp)

/* fmt.Println(l, "after copying")
		   
		
		if e.ProbeName == "tcplife"{

			fmt.Println(l,"inside tcplife")
			err := database.UpdateTcpLifeLogs(tcplogs)
			fmt.Println(l,"after tcplife")
			if err != nil {
			fmt.Println(l,"error aa gayi hawwwwwwwwwwww")	
				 err = f.Close()
		    if err != nil {
	        	fmt.Println(err)
        		return
    		}

				fmt.Println(err)
				os.Exit(1)
			}

		
		}else{	*/
			err := database.UpdateLogs(tcplogs)
			if err != nil {
	
				os.Exit(1)
			}
		

//		fmt.Println(l,"before calling next event")
		
		events.PublishEvent("logs:refresh", events.EmptyMessage{Pn: e.ProbeName})
//				fmt.Println(l,"after calling next event")
	

	}
}

func tcplifeLog(e events.Event) {
	if e, ok := e.(events.TcpLifeLogEvent); ok {

		tcp := events.TcpLifeLogEvent{TimeStamp: e.TimeStamp,
				     ProbeName: e.ProbeName,
                                     Sys_Time:  e.Sys_Time,
		                     Pid:       e.Pid,
                                     Pname:     e.Pname,
		                     Laddr:     e.Laddr,
                                     Lport:     e.Lport,
                                     Raddr:     e.Raddr,
                                     Rport:     e.Rport,
				     Tx_kb:	e.Tx_kb,
				     Rx_kb:	e.Rx_kb,
				     Ms:	e.Ms,
		}
		tcplogs := database.TcpLifeLog(tcp)
		err := database.UpdateTcpLifeLogs(tcplogs)
		if err != nil {

			os.Exit(1)
		}
		events.PublishEvent("logs:refresh", events.EmptyMessage{Pn: e.ProbeName})

	}
}


func execsnoopLog(e events.Event) {
	if e, ok := e.(events.ExecSnoopLogEvent); ok {

		tcp := events.ExecSnoopLogEvent{TimeStamp: e.TimeStamp,
				     ProbeName: e.ProbeName,
                                     Sys_Time:  e.Sys_Time,
		                     T:         e.T,
                                     Pname:     e.Pname,
		                     Pid:       e.Pid,
                                     Ppid:      e.Ppid,
                                     Ret:     	e.Ret,
                                     Args:     	e.Args,
		}
		eslogs := database.ExecSnoopLog(tcp)
		err := database.UpdateEsLogs(eslogs)
		if err != nil {

			os.Exit(1)
		}
		events.PublishEvent("logs:refresh", events.EmptyMessage{Pn: e.ProbeName})

	}
}


func biosnoopLog(e events.Event) {
	if e, ok := e.(events.BioSnoopLogEvent); ok {

		tcp := events.BioSnoopLogEvent{TimeStamp: e.TimeStamp,
				     ProbeName: e.ProbeName,
                                     Sys_Time:  e.Sys_Time,
		                     T:       	e.T,
                                     Pname:     e.Pname,
				     Pid:	e.Pid,
		                     Disk:      e.Disk,
                                     Rw:        e.Rw,
                                     Sector:    e.Sector,
                                     Bytes:     e.Bytes,
				     Lat:	e.Lat,
		}
		bslogs := database.BioSnoopLog(tcp)
		err := database.UpdateBsLogs(bslogs)
		if err != nil {

			os.Exit(1)
		}
		events.PublishEvent("logs:refresh", events.EmptyMessage{Pn: e.ProbeName})

	}
}


func cachestatLog(e events.Event) {
	if e, ok := e.(events.CacheStatLogEvent); ok {

		tcp := events.CacheStatLogEvent{TimeStamp: e.TimeStamp,
				     ProbeName: e.ProbeName,
                                     Sys_Time:  e.Sys_Time,
		                     Pid:       e.Pid,
                                     Uid:     	e.Uid,
		                     Cmd:     	e.Cmd,
                                     Hits:      e.Hits,
                                     Miss:     	e.Miss,
                                     Dirties:   e.Dirties,
				     Read_hit:	e.Read_hit,
				     Write_hit:	e.Write_hit,
		}
		cslogs := database.CacheStatLog(tcp)
		err := database.UpdateCsLogs(cslogs)
		if err != nil {

			os.Exit(1)
		}
		events.PublishEvent("logs:refresh", events.EmptyMessage{Pn: e.ProbeName})

	}
}


func SubscribeListeners() {
	events.Subscribe(tcplifeLog, 	"log:tcplife")
	events.Subscribe(receiveLog, 	"log:receive")
	events.Subscribe(execsnoopLog, 	"log:execsnoop")
	events.Subscribe(biosnoopLog, 	"log:biosnoop")
	events.Subscribe(cachestatLog, 	"log:cachestat")
}
