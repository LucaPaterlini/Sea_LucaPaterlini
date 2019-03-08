package main

import "time"

const (
	MONGODBHOSTS      = "127.0.0.1"
	MONGODBDATABASE   = "Sea"
	NAMECOLLECTION    = "addressBook"
	SERVICEPORT       = ":11235"
	TIMEOUTDATABASE   = 10 * time.Second
	TIMEOUTCONNECTION = 2 * time.Second
)
