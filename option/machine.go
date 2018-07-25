package option

var devMachine = &machine{
	MongoURL:    "localhost:27017",
	MongoDBName: "postaldep",
}

var prodMachine = &machine{
	MongoURL:    "<onlyfillwhenneeded>:<your_port>",
	MongoDBName: "postaldep",
}
