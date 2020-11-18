package pkg

import (
	"strings"
	"fmt"
	pb "github.com/Sheenam3/x-tracer-gocui/api"
	"github.com/gogo/protobuf/sortkeys"
	"github.com/Sheenam3/x-tracer-gocui/events"
	"github.com/Sheenam3/x-tracer-gocui/database"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
//	"os"
	)


type StreamServer struct {
	//port string
}

var (

	port    string
)

var bufLogs []string
var wbLogs []string
var csbufLogs []string
var cswbLogs []string
//var bsbufLogs []string
//var bswbLogs []string
var esbufLogs []string
var eswbLogs []string
var tlbufLogs []string
var tlwbLogs []string

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


	        parse := strings.Fields(string(r.Log))



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
//									  Tx_kb:    "0",
  //                                                                        Rx_kb:    "0",
   //                                                                       Ms:       "0",
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
//									  Tx_kb:    "0",
//                                                                          Rx_kb:    "0",
//                                                                          Ms:       "0",
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
//									  Tx_kb:    "0",
//                                                                          Rx_kb:    "0",
//                                                                          Ms:       "0",
									 })
		}else if r.ProbeName == "tcplife"{

			events.PublishEvent("log:tcplife", events.TcpLifeLogEvent{TimeStamp: 0,
									  ProbeName:r.ProbeName,
									  Sys_Time: parse[0],
									  Pid:      parse[2],
  									  Pname:    parse[3],
									  Laddr:    parse[4],
									  Lport:    parse[5],
									  Raddr:    parse[6],
									  Rport:    parse[7],
									  Tx_kb:    parse[8],
									  Rx_kb:    parse[9],
									  Ms:	    parse[10],
									 })
		}else if r.ProbeName == "execsnoop"{
			if len(parse) < 8 {
				events.PublishEvent("log:execsnoop", events.ExecSnoopLogEvent{TimeStamp: 0,
                                                                          ProbeName:r.ProbeName,
                                                                          Sys_Time: parse[0],
                                                                          T:        parse[1],
                                                                          Pname:    parse[3],
                                                                          Pid:      parse[4],
                                                                          Ppid:     parse[5],
                                                                          Ret:      parse[6],
                                                                          Args:     parse[3],
                                                                         })

			}else{
				events.PublishEvent("log:execsnoop", events.ExecSnoopLogEvent{TimeStamp: 0,
									  ProbeName:r.ProbeName,
									  Sys_Time: parse[0],
									  T:	    parse[1],
  	                         		                          Pname:    parse[3],
									  Pid:      parse[4],
									  Ppid:	    parse[5],
									  Ret:      parse[6],
									  Args:     parse[7],
									 })
			}
		}else if r.ProbeName == "biosnoop"{

			events.PublishEvent("log:biosnoop", events.BioSnoopLogEvent{TimeStamp: 0,
									  ProbeName:	r.ProbeName,
									  Sys_Time: 	parse[0],
									  T:	    	parse[1],
									  Pname:        parse[2],
  									  Pid:    	parse[3],
									  Disk:    	parse[4],
									  Rw:    	parse[5],
									  Sector:    	parse[6],
									  Bytes:    	parse[7],
									  Lat:    	parse[9],
									 })
		}else if r.ProbeName == "cachestat"{

			events.PublishEvent("log:cachestat", events.CacheStatLogEvent{TimeStamp: 0,
									  ProbeName:  r.ProbeName,
									  Sys_Time: 	parse[0],
									  Pid:      	parse[1],
  									  Uid:    	parse[2],
									  Cmd:    	parse[3],
									  Hits:    	parse[5],
									  Miss:    	parse[6],
									  Dirties:    	parse[7],
									  Read_hit:    	parse[8],
									  Write_hit:    parse[9],
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



func GetActiveLogs(pn string) string {
	var err error


	var keys []int64


	if pn == "tcplife"{
		var tlLogs []string
		logs := database.GetTcpLifeLogs()

		if err != nil {
			log.Panic(err)
		}



		for k := range logs {
			keys = append(keys, k)

		}


		sortkeys.Int64s(keys)


		for _, log := range keys {
			val := logs[log]
			//fmt.Println(val.ProbeName)
                	tlLogs = append(tlLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s | PID:%s | PNAME:%s | LADDR:%s | LPORT:%s | RADDR:%s | RPORT:%s | Tx_kb:%s | Rx_kb:%s | Ms: %s \n", val.ProbeName,val.Sys_Time,val.Pid,val.Pname, val.Laddr, val.Lport, val.Raddr, val.Rport, val.Tx_kb, val.Rx_kb, val.Ms))

		}


			for i := range tlLogs {
				tlbufLogs = append(tlbufLogs,tlLogs[i])
			}
			if len(tlbufLogs) >= 9{

				tlwbLogs = tlbufLogs
				tlbufLogs = nil
				del := database.DeleteTlLogs()
				return strings.Join(tlwbLogs, "\n")
				fmt.Println(del)
			}else{

				return strings.Join(tlwbLogs, "\n")

			}

		//return strings.Join(tlLogs, "\n")

	}else if pn == "execsnoop"{
		var esLogs []string
		logs := database.GetExecSnoopLogs()

		if err != nil {
			log.Panic(err)
		}


		for k := range logs {

			keys = append(keys, k)

		}


		sortkeys.Int64s(keys)

		for _, log := range keys {
			val := logs[log] 
			esLogs = append(esLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s | T:%s | PNAME:%s | PID:%s | PPID:%s | RET:%s | ARGS:%s \n", val.ProbeName,val.Sys_Time,val.T,val.Pname,val.Pid,val.Ppid, val.Ret, val.Args))

		}

		for i := range esLogs {
			esbufLogs = append(esbufLogs,esLogs[i])
		}
		if len(esbufLogs) >= 9{

			eswbLogs = esbufLogs
			esbufLogs = nil
			del := database.DeleteESLogs()
			return strings.Join(eswbLogs, "\n")
			fmt.Println(del)
		}else{

				return strings.Join(eswbLogs, "\n")

		}

	//	return strings.Join(esLogs, "\n")

	}else if pn == "biosnoop"{
		var bsLogs []string
		logs := database.GetBioSnoopLogs()

		if err != nil {
			log.Panic(err)
		}

		for k := range logs {

			keys = append(keys, k)

		}


		sortkeys.Int64s(keys)


		for _, log := range keys {
			val := logs[log]
                	bsLogs = append(bsLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s | T:%s | PNAME:%s | PID:%s | DISK:%s | R/W:%s | SECTOR:%s | BYTES:%s | LAT:%s \n", val.ProbeName,val.Sys_Time,val.T,val.Pname,val.Pid,val.Disk, val.Rw, val.Sector, val.Bytes, val.Lat))

		}
		return strings.Join(bsLogs, "\n")

	}else if pn == "cachestat"{
		var csLogs []string
		logs := database.GetCacheStatLogs()

		if err != nil {
			log.Panic(err)
		}

		for k := range logs {

			keys = append(keys, k)

		}


		sortkeys.Int64s(keys)
/*		f, err := os.OpenFile("cache_db_del.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			    if err != nil {
        			fmt.Println(err)
        			return "file open error"
    			}
		defer f.Close()*/

		for _, log := range keys {
			val := logs[log]
                	csLogs = append(csLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s | PID:%s | UID:%s | CMD:%s | HITS:%s | MISS:%s | DIRTIES:%s | READ_HIT%:%s | WRITE_HIT%:%s \n", val.ProbeName,val.Sys_Time,val.Pid,val.Uid, val.Cmd, val.Hits, val.Miss, val.Dirties, val.Read_hit, val.Write_hit))
			/* _, err := f.WriteString(fmt.Sprintf("{Probe:%s |Sys_Time:%s | PID:%s | UID:%s | CMD:%s | HITS:%s | MISS:%s | DIRTIES:%s | READ_HIT%:%s | WRITE_HIT%:%s \n", val.ProbeName,val.Sys_Time,val.Pid,val.Uid, val.Cmd, val.Hits, val.Miss, val.Dirties, val.Read_hit, val.Write_hit))
			    if err != nil {
			        fmt.Println(err)
			        f.Close()
		        	return "file writing error"
    		    	    }*/

		}


		for i := range csLogs {
                                csbufLogs = append(csbufLogs,csLogs[i])
                        }
                        if len(csbufLogs) >= 9{

                                cswbLogs = csbufLogs
                                csbufLogs = nil
                                del := database.DeleteCSLogs()
                                return strings.Join(cswbLogs, "\n")
                                fmt.Println(del)
                        }else{

                                return strings.Join(cswbLogs, "\n")

                        }

		//return strings.Join(csLogs, "\n")

	}else{
		var tcpLogs []string
		//queue := list.New()
		logs := database.GetLogs()

		if err != nil {
			log.Panic(err)
		}

		for k := range logs {

			keys = append(keys, k)

		}


		sortkeys.Int64s(keys)


		for _, log := range keys {
			val := logs[log]
			if val.ProbeName == "tcpconnect"{
		                tcpLogs = append(tcpLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s |T:%s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s \n", val.ProbeName,val.Sys_Time,val.T, val.Pid,val.Pname, val.Ip, val.Saddr, val.Daddr, val.Dport))

                	}else if val.ProbeName == "tcptracer"{
		                tcpLogs = append(tcpLogs,fmt.Sprintf("{Probe:%s |Sys_Time:%s |T:%s | PID:%s | PNAME:%s | IP:%s | SADDR:%s | DADDR:%s | DPORT:%s | SPORT:%s \n", val.ProbeName,val.Sys_Time,val.T, val.Pid,val.Pname, val.Ip, val.Saddr, val.Daddr, val.Dport, val.Sport))


                	}else if val.ProbeName == "tcpaccept"{
		                tcpLogs = append(tcpLogs, fmt.Sprintf("{Probe:%s |Sys_Time:%s |T:%s | PID:%s | PNAME:%s | IP:%s | LADDR:%s | RADDR:%s | LPORT:%s |RPORT: %s \n", val.ProbeName,val.Sys_Time,val.T, val.Pid,val.Pname, val.Ip, val.Saddr, val.Daddr, val.Sport, val.Dport))
			}
		}



/*		if len(tcpLogs) > 10{

		//fmt.Println(tcpLogs)

			del := database.DeleteTcpLogs()
			return strings.Join(tcpLogs, "\n")
			fmt.Println(del)

		}else{*/
			for i := range tcpLogs {
				bufLogs = append(bufLogs,tcpLogs[i])
			}
			if len(bufLogs) >= 9{

				wbLogs = bufLogs
				bufLogs = nil
				del := database.DeleteTcpLogs()
				return strings.Join(wbLogs, "\n")
				fmt.Println(del)
			}else{

				return strings.Join(wbLogs, "\n")

			}
	//	}

	}

	return "Nothing yet"

}
