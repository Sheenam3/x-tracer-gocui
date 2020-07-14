package database

import (
	godb "github.com/hashicorp/go-memdb"
	"fmt"
)



var(

	db *godb.MemDB

)


func Init(){

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
		"tcpconnect": &memdb.TableSchema{
			Name: "tcpconnect",
			Indexes: map[string]*memdb.IndexSchema{
				"pn": &memdb.IndexSchema{
					Name:    "pn",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "ProbeName"},
				},
				"sys_time": &memdb.IndexSchema{
					Name:    "sys_time",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "Sys_Time"},
				},

				"t": &memdb.IndexSchema{
					Name:    "t",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "T"},
				},
				"pid": &memdb.IndexSchema{
					Name:    "pid",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "Pid"},
				},

				"pname": &memdb.IndexSchema{
					Name:    "pname",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Pname"},
				},
				"ip": &memdb.IndexSchema{
					Name:    "ip",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "Ip"},
				},

				"saddr": &memdb.IndexSchema{
					Name:    "saddr",
					Unique:  false,
					Indexer: &memdb.StringFieldIndex{Field: "Saddr"},
				},
				"daddr": &memdb.IndexSchema{
					Name:    "daddr",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "Daddr"},
				},
				"dport": &memdb.IndexSchema{
					Name:    "dport",
					Unique:  false,
					Indexer: &memdb.IntFieldIndex{Field: "Dport"},
				},
			},
		},
	},
}


//Create a new data base
db, err := memdb.NewMemDB(schema)
if err != nil {
	panic(err)
}

}



func UpdateLogs(pn string, st string, t string, pid string, pname string, ip string, saddr string, daddr string, dport string) error{

txn := db.Txn(true)

logs := *Log{
	&Log{pn, st, t, pid, pname, ip, saddr, daddr, dport},
	}

if err := txn.Insert("tcpconnect", logs); err!= nil{
	return err
}	

txn.Commit()


return nil

}


func GetLogs() ([]*Log){

txn = db.Txn(false)
defer txn.Abort()


it, err := txn.Get("tcpconnect", "sys_time")
if err != nil {
	panic(err)
}

fmt.Println("All the people:")
var logs []*Log

for  obj := it.Next(); obj != nil; obj = it.Next() {
	p := obj.(*Log)
	logs = append(logs, p)
}

return logs
}
