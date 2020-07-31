package pkg

import (
	"strings"
	"fmt"
	pb "github.com/Sheenam3/x-tracer-gocui/api"
	"github.com/Sheenam3/x-tracer-gocui/events"
	"github.com/Sheenam3/x-tracer-gocui/database"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	)


type StreamServer struct {
	//port string
}

var (

	port    string
)

func (s *StreamServer)RouteLog(stream pb.SentLog_RouteLogServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Response{
				Res:                  "Stream closed",
			})
		}
		if err != nil {
			return err
		}
//		fmt.Println("\n", r.Log)

	        parse := strings.Fields(string(r.Log))

//		fmt.Println("PID:",r.Pid)

		if r.ProbeName == "tcpconnect"{
			events.PublishEvent("log:receive", events.ReceiveLogEvent{ProbeName: r.ProbeName,
									  Sys_Time: parse[0],
									  T:	    parse[1],
									  Pid:      parse[3],
  									  Pname:    parse[4],
									  Ip:	    parse[5],
									  Saddr:    parse[6],
									  Daddr:    parse[7],
									  Dport:    parse[8],
									  Sport:    "0",
									 })
		}else if r.ProbeName == "tcptracer"{
			events.PublishEvent("log:receive", events.ReceiveLogEvent{ProbeName: r.ProbeName,
									  Sys_Time: parse[0],
									  T:	    parse[1],
									  Pid:      parse[3],
  									  Pname:    parse[4],
									  Ip:	    parse[5],
									  Saddr:    parse[6],
									  Daddr:    parse[7],
									  Dport:    parse[9],
									  Sport:    parse[8],
									 })
		}else if r.ProbeName == "tcpaccept"{
			events.PublishEvent("log:receive", events.ReceiveLogEvent{ProbeName: r.ProbeName,
									  Sys_Time: parse[0],
									  T:	    parse[1],
									  Pid:      parse[3],
  									  Pname:    parse[4],
									  Ip:	    parse[5],
									  Saddr:    parse[8],
									  Daddr:    parse[6],
									  Dport:    parse[7],
									  Sport:    parse[9],
									 })
		}



		
/*
		if r.ProbeName == "tcptracer"{

		//fmt.Println("ProbeName:",r.ProbeName)
                //fmt.Printf("{%s}\n", r.Log)
                fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s |IP->%s | SADDR:%s | DADDR:%s | SPORT:%s | DPORT:%s \n",r.ProbeName,parse[0],parse[1],parse[3],parse[4],parse[5],parse[6],parse[7],parse[8],parse[9])

                }else if r.ProbeName == "tcpaccept"{

                //fmt.Println("ProbeName:",r.ProbeName)
		//fmt.Printf("{%s}\n", r.Log)
                fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | RADDR:%s | RPORT:%s | LADDR:%s | LPORT:%s \n",r.ProbeName,parse[0],parse[1],parse[3],parse[4],parse[5],parse[6],parse[7],parse[8],parse[9])

                }else if r.ProbeName == "tcplife"{

		fmt.Printf("{Probe:%s |Sys_Time: %s |PID:%s | PNAME:%s | LADDRR:%s | LPORT:%s | RADDR:%s | RPORT:%s | TX_KB:%s | RX_KB:%s | MS: %s \n",r.ProbeName,parse[0],parse[2],parse[3],parse[4],parse[5],parse[6],parse[7],parse[8],parse[9],parse[10])

		}else if r.ProbeName == "execsnoop"{
		fmt.Printf("{Probe:%s |Sys_Time: %s | T:%s | PNAME: %s | PID:%s | PPID:%s | RET:%s | ARGS:%s \n",r.ProbeName,parse[0],parse[1],parse[3],parse[4],parse[5],parse[6],parse[7])

		}else if r.ProbeName == "biosnoop"{

		fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s |PNAME: %s | PID:%s | DISK:%s | R/W:%s | SECTOR:%s |BYTES: %s | Lat(ms): %s | \n",r.ProbeName,parse[0],parse[1],parse[2],parse[3],parse[4],parse[5],parse[6],parse[7],parse[9])

		}else if r.ProbeName == "cachetop"{

		fmt.Printf("{Probe:%s |Sys_Time: %s | PID:%s | UID:%s | CMD:%s | HITS:%s | MISS:%s | DIRTIES: %s| READ_HIT%:%s | W_HIT%: %s | \n",r.ProbeName,parse[0],parse[1],parse[2],parse[3],parse[5],parse[6],parse[7],parse[8], parse[9])

		}else{

                fmt.Printf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s \n",r.ProbeName,parse[0],parse[1],parse[3],parse[4],parse[5],parse[6],parse[7],parse[8])
                }*/
		//fmt.Println(r.TimeStamp, "\n")
	}
}

/*func New(servicePort string) *StreamServer{
	return &StreamServer{
		servicePort}
}*/

func SetPort(sport string) {
	port = sport
}

func StartServer(){
	server := grpc.NewServer()
	pb.RegisterSentLogServer(server, &StreamServer{})

	lis, err := net.Listen("tcp", ":"+ port)
	if err != nil {
		log.Fatalln("net.Listen error:", err)
	}

	_ = server.Serve(lis)
}



func GetActiveLogs() string {
	var err error
	logs := database.GetLogs()

	if err != nil {
		log.Panic(err)
	}

	var displayLogs []string


	for _, val := range logs {
		if val.ProbeName == "tcpconnect"{
                displayLogs = append(displayLogs, fmt.Sprintf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s \n", val.ProbeName,val.Sys_Time,val.T, val.Pid,val.Pname, val.Ip, val.Saddr, val.Daddr, val.Dport))
                }else if val.ProbeName == "tcptracer"{
                displayLogs = append(displayLogs, fmt.Sprintf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s | SPORT:%s \n", val.ProbeName,val.Sys_Time,val.T, val.Pid,val.Pname, val.Ip, val.Saddr, val.Daddr, val.Dport, val.Sport))
                }else if val.ProbeName == "tcpaccept"{
                displayLogs = append(displayLogs, fmt.Sprintf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | LADDR:%s | RADDR:%s | LPORT:%s |RPORT: %s \n", val.ProbeName,val.Sys_Time,val.T, val.Pid,val.Pname, val.Ip, val.Saddr, val.Daddr, val.Sport, val.Dport))
		}
		//displayLogs = append(displayLogs, fmt.Sprintf("{Probe:%s |Sys_Time: %s |T: %s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s \n", val.ProbeName,val.Sys_Time,val.T, val.Pid,val.Pname, val.Ip, val.Saddr, val.Daddr, val.Dport))
	}

	return strings.Join(displayLogs, "\n")
}
