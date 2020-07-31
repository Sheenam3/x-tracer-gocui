
package database

import (
	memdb "github.com/hashicorp/go-memdb"
	"time"
)



var(

	db *memdb.MemDB

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


//Create a new data base
db, err = memdb.NewMemDB(schema)
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
