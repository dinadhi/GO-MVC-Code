package option

import "os"


type machine struct {
	MongoURL    string
	MongoDBName string
}

var Machine *machine

func init() {
	if os.Getenv("GO_ENV") == "dev" {
		Machine = devMachine
		return
	}
	Machine = prodMachine
}
