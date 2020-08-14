
package database

import (
	memdb "github.com/hashicorp/go-memdb"
	"time"
//	"fmt"
	"os"
)



var(

	db *memdb.MemDB
	tldb *memdb.MemDB
	es *memdb.MemDB
	bs *memdb.MemDB
	cs *memdb.MemDB

)


func Init(){
	var err error
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
		"tcpconnect": &memdb.TableSchema{
			Name: "tcpconnect",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "Timestamp"},
				},
				"pn": &memdb.IndexSchema{
					Name:    "pn",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
				},
				"sys_time": &memdb.IndexSchema{
					Name:    "sys_time",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
				},

				"t": &memdb.IndexSchema{
					Name:    "t",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "T"},
				},
				"pid": &memdb.IndexSchema{
					Name:    "pid",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pid"},
				},

				"pname": &memdb.IndexSchema{
					Name:    "pname",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pname"},
				},
				"ip": &memdb.IndexSchema{
					Name:    "ip",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Ip"},
				},

				"saddr": &memdb.IndexSchema{
					Name:    "saddr",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Saddr"},
				},
				"daddr": &memdb.IndexSchema{
					Name:    "daddr",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Daddr"},
				},
				"dport": &memdb.IndexSchema{
					Name:    "dport",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Dport"},
				},
				"sport": &memdb.IndexSchema{
					Name:    "sport",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Sport"},
				},
				
			},
		},
	},
}




	schematcplife := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
		"tcplife": &memdb.TableSchema{
			Name: "tcplife",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "TimeStamp"},
				},
				"pn": &memdb.IndexSchema{
					Name:    "pn",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
				},
				"sys_time": &memdb.IndexSchema{
					Name:    "sys_time",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
				},

				"pid": &memdb.IndexSchema{
					Name:    "pid",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pid"},
				},

				"pname": &memdb.IndexSchema{
					Name:    "pname",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pname"},
				},
				

				"laddr": &memdb.IndexSchema{
					Name:    "laddr",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Laddr"},
				},
				"lport": &memdb.IndexSchema{
					Name:    "lport",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Lport"},
				},
				"raddr": &memdb.IndexSchema{
					Name:    "raddr",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Raddr"},
				},
				"rport": &memdb.IndexSchema{
					Name:    "rport",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Rport"},
				},

				"tx_kb": &memdb.IndexSchema{
					Name:    "tx_kb",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Tx_kb"},
				},

				"rx_kb": &memdb.IndexSchema{
					Name:    "rx_kb",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Rx_kb"},
				},				

				"ms": &memdb.IndexSchema{
					Name:    "ms",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Ms"},
				},
			},
		},
	},
}

	//Schema for Execsnoop
	schemaes := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
		"execsnoop": &memdb.TableSchema{
			Name: "execsnoop",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "TimeStamp"},
				},
				"pn": &memdb.IndexSchema{
					Name:    "pn",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
				},
				"sys_time": &memdb.IndexSchema{
					Name:    "sys_time",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
				},


				"pname": &memdb.IndexSchema{
					Name:    "pname",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pname"},
				},

				"pid": &memdb.IndexSchema{
					Name:    "pid",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pid"},
				},

				"ppid": &memdb.IndexSchema{
					Name:    "ppid",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Ppid"},
				},
				"ret": &memdb.IndexSchema{
					Name:    "ret",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Ret"},
				},
				"args": &memdb.IndexSchema{
					Name:    "args",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Args"},
				},
			},
		},
	},
}


	//Schema for Biosnoop

	schemabs := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
		"biosnoop": &memdb.TableSchema{
			Name: "biosnoop",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "TimeStamp"},
				},
				"pn": &memdb.IndexSchema{
					Name:    "pn",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
				},
				"sys_time": &memdb.IndexSchema{
					Name:    "sys_time",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
				},

				"t": &memdb.IndexSchema{
					Name:    "t",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "T"},
				},
				"pname": &memdb.IndexSchema{
					Name:    "pname",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pname"},
				},

				"pid": &memdb.IndexSchema{
					Name:    "pid",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pid"},
				},

				"disk": &memdb.IndexSchema{
					Name:    "disk",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Disk"},
				},
				"rw": &memdb.IndexSchema{
					Name:    "rw",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Rw"},
				},

				"sector": &memdb.IndexSchema{
					Name:    "sector",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Sector"},
				},
				"bytes": &memdb.IndexSchema{
					Name:    "bytes",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Bytes"},
				},
				"lat": &memdb.IndexSchema{
					Name:    "lat",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Lat"},
				},
			},
		},
	},
}




	//Schema for Cachestat

	schemacs := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
		"cachestat": &memdb.TableSchema{
			Name: "cachestat",
			Indexes: map[string]*memdb.IndexSchema{
				"id": &memdb.IndexSchema{
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.IntFieldIndex{Field: "TimeStamp"},
				},
				"pn": &memdb.IndexSchema{
					Name:    "pn",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
				},
				"sys_time": &memdb.IndexSchema{
					Name:    "sys_time",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Sys_Time"},
				},

				"pid": &memdb.IndexSchema{
					Name:    "pid",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pid"},
				},

				"uid": &memdb.IndexSchema{
					Name:    "uid",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Uid"},
				},
				"cmd": &memdb.IndexSchema{
					Name:    "cmd",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Cmd"},
				},

				"hits": &memdb.IndexSchema{
					Name:    "hits",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Hits"},
				},
				"miss": &memdb.IndexSchema{
					Name:    "miss",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Miss"},
				},
				"dirties": &memdb.IndexSchema{
					Name:    "dirties",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Dirties"},
				},
				"rh": &memdb.IndexSchema{
					Name:    "rh",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Read_hit"},
				},
				"wh": &memdb.IndexSchema{
					Name:    "wh",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Write_hit"},
				},

			},
		},
	},
}



//Create a new data base for tcplogs
db, err = memdb.NewMemDB(schema)
if err != nil {
	panic(err)
}



//Create a new data base for tcplife
tldb, err = memdb.NewMemDB(schematcplife)
if err != nil {
	panic(err)
}


//Create a new data base for execsnoop
es, err = memdb.NewMemDB(schemaes)
if err != nil {
	panic(err)
}

//Create a new data base for biosnoop
bs, err = memdb.NewMemDB(schemabs)
if err != nil {
	panic(err)
}

//Create a new data base for cacahestat
cs, err = memdb.NewMemDB(schemacs)
if err != nil {
	panic(err)
}

}


//func UpdateLogs(pn string, st string, t string, pid string, pname string, ip string, saddr string, daddr string, dport string) error{
func UpdateLogs(log TcpLog) error{

txn := db.Txn(true)
timestamp := time.Now().UnixNano()
logs := []*Log{
	//&Log{timestamp,pn, st, t, pid, pname, ip, saddr, daddr, dport},
	&Log{timestamp, log.ProbeName, log.Sys_Time, log.T, log.Pid, log.Pname, log.Ip, log.Saddr, log.Daddr, log.Dport, log.Sport},
	}

for _, p := range logs {
	if err := txn.Insert("tcpconnect", p); err!= nil{
		return err
	}	
}

txn.Commit()


return nil

}


func UpdateTcpLifeLogs(log TcpLifeLog) error{

txn := tldb.Txn(true)
timestamp := time.Now().UnixNano()
logs := []*TcpLifeLog{
	//&Log{timestamp,pn, st, t, pid, pname, ip, saddr, daddr, dport},
	&TcpLifeLog{timestamp, log.ProbeName, log.Sys_Time,log.Pid, log.Pname,log.Laddr, log.Lport, log.Raddr, log.Rport, log.Tx_kb, log.Rx_kb, log.Ms},
	}

for _, p := range logs {
	if err := txn.Insert("tcplife", p); err!= nil{

		return err
	}	
}

txn.Commit()


return nil

}


//update execsnoop table
func UpdateEsLogs(log ExecSnoopLog) error{

txn := es.Txn(true)
timestamp := time.Now().UnixNano()
logs := []*ExecSnoopLog{

	&ExecSnoopLog{timestamp, log.ProbeName, log.Sys_Time, log.T, log.Pname, log.Pid, log.Ppid, log.Ret, log.Args },
	}

for _, p := range logs {
	if err := txn.Insert("execsnoop", p); err!= nil{
		return err
	}	
}

txn.Commit()


return nil

}



//update biosnoop  table
func UpdateBsLogs(log BioSnoopLog) error{

txn := bs.Txn(true)
timestamp := time.Now().UnixNano()
logs := []*BioSnoopLog{
	//&Log{timestamp,pn, st, t, pid, pname, ip, saddr, daddr, dport},
	&BioSnoopLog{timestamp, log.ProbeName, log.Sys_Time, log.T, log.Pname, log.Pid, log.Disk, log.Rw, log.Sector, log.Bytes, log.Lat},
	}

for _, p := range logs {
	if err := txn.Insert("biosnoop", p); err!= nil{
		return err
	}	
}

txn.Commit()


return nil

}


//update cachestat table
func UpdateCsLogs(log CacheStatLog) error{

txn := cs.Txn(true)
timestamp := time.Now().UnixNano()
logs := []*CacheStatLog{

	&CacheStatLog{timestamp, log.ProbeName, log.Sys_Time, log.Pid, log.Uid, log.Cmd, log.Hits, log.Miss, log.Dirties, log.Read_hit, log.Write_hit},
	}

for _, p := range logs {
	if err := txn.Insert("cachestat", p); err!= nil{
		return err
	}	
}

txn.Commit()


return nil

}


func GetLogs() ([]*Log){

txn := db.Txn(false)
defer txn.Abort()



it, err := txn.Get("tcpconnect", "id")
if err != nil {
	panic(err)
}


var logs []*Log

for  obj := it.Next(); obj != nil; obj = it.Next() {
	p := obj.(*Log)
	logs = append(logs, p)
}

return logs
}



func GetTcpLifeLogs() (map[int64]*TcpLifeLog){

txn := tldb.Txn(false)
defer txn.Abort()


logs := make(map[int64]*TcpLifeLog)


it, err := txn.Get("tcplife", "id")
if err != nil {
	panic(err)
}


//var logs []*TcpLifeLog


for  obj := it.Next(); obj != nil; obj = it.Next() {
	p := obj.(*TcpLifeLog)
	timestamp := p.TimeStamp
	logs[timestamp] = p
//	logs = append(logs, p)
}

return logs
}

//Get execsnoop logs

func GetExecSnoopLogs() ([]*ExecSnoopLog){

txn := es.Txn(false)
defer txn.Abort()



it, err := txn.Get("execsnoop", "id")
if err != nil {
	panic(err)
	os.Exit(1)
}


var logs []*ExecSnoopLog

for  obj := it.Next(); obj != nil; obj = it.Next() {
	p := obj.(*ExecSnoopLog)
	logs = append(logs, p)
}

return logs
}


//Get Biosnoop logs

func GetBioSnoopLogs() ([]*BioSnoopLog){

txn := bs.Txn(false)
defer txn.Abort()



it, err := txn.Get("biosnoop", "id")
if err != nil {
	panic(err)
}


var logs []*BioSnoopLog

for  obj := it.Next(); obj != nil; obj = it.Next() {
	p := obj.(*BioSnoopLog)
	logs = append(logs, p)
}

return logs
}



//Get Cachestat logs

func GetCacheStatLogs() ([]*CacheStatLog){

txn := cs.Txn(false)
defer txn.Abort()



it, err := txn.Get("cachestat", "id")
if err != nil {
	panic(err)
}


var logs []*CacheStatLog

for  obj := it.Next(); obj != nil; obj = it.Next() {
	p := obj.(*CacheStatLog)
	logs = append(logs, p)
}

return logs
}

