package events



type SendLogEvent struct {
	Pid string
        ProbeName string
        Log string
        TimeStamp string

}


type ReceiveLogEvent struct {

	ProbeName string
	Sys_Time string
	T string
	Pid string
	Pname string
	Ip string
	Saddr string
	Daddr string
	Dport string

}
