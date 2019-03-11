package main

import (
	pp "../pipe"
	"fmt"
	"github.com/globalsign/mgo"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

var accountTable *mgo.Collection

type transferServer struct{
	itemC chan pp.Record
}


//func RegisterTransferServer() pp.TransferServer {
//	return &transferServer{itemC: make(chan pp.Record, NSTREAM)}
//}


func (s *transferServer) AddUpdateRecord( errS  pp.Transfer_AddUpdateRecordServer)(err error){
	for item := range s.itemC {
		log.Printf("Receiver: %v\n", item)
		_, err = accountTable.UpsertId(item.Id, &item)
		// checking the upsert and returning a Ack object over the error Stream
		if err != nil {
			msg:=fmt.Sprint("Error on Update: ", item.Id, " err: ",err.Error())
			ack := &pp.Ack{Err:err==nil,Message:msg}
			// halting the read loop if the connections is not working
			if err = errS.Send(ack); err != nil {
				s.itemC <- item
				log.Printf("Stream connection failed: %v", err)
				return
			}
			log.Println("Error on Update: ", item.Id, " err: ",err)
		}
	}
	return
}

func signalHandling(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,os.Kill)
	go func() {
		select {
		case <-c:
			fmt.Print("Server: I got stopped!")
			os.Exit(1)
		}
	}()
}

func main() {
	// handling the signals
	signalHandling()
	// connecting to the database
	mongoSession, err := mgo.DialWithTimeout(MONGODBHOSTS+"/"+MONGODBDATABASE, TIMEOUTDATABASE)
	if err != nil {
		log.Println("error While connecting to MongoDb", err.Error())
	} else {
		accountTable = mongoSession.DB(MONGODBDATABASE).C(NAMECOLLECTION)

		lis, err := net.Listen("tcp", SERVICEPORT)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pp.RegisterTransferServer(s, &transferServer{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}