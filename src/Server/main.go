// Package main this package contain the implentation of the receiver
// service , it writes on the database the new record or return
// an ack of the failure with a fail message and the serial of the request
package main

import (
	pipe "../pipe"
	"fmt"
	"github.com/globalsign/mgo"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
)

var accountTable *mgo.Collection

type transferServer struct{}

// AddUpdateRecord implement the functionality of receiving a Rercord and Upsert it in the mongodb database
func (s *transferServer) AddUpdateRecord( stream  pipe.Transfer_AddUpdateRecordServer)(err error){
	fmt.Println("hello")
	for {
		var item *pipe.Record
		item, err = stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Receiver: %v\n", item)
		_, err = accountTable.UpsertId(item.Id, &item)
		// checking the upInsert and returning a Ack object over the error Stream
		var msg = ""
		if err != nil {
			msg = fmt.Sprint("Error on Update: ", item.Id, " err: ", err.Error())
		}
		ack := &pipe.Ack{Err: err != nil, Message: msg}
		// halting the read loop if the connections is not working
		if err = stream.Send(ack); err != nil {
				log.Printf("Stream connection failed: %v", err.Error())
				return err
			}
		}
}


// signalHandling the reception of a signal and halt the service gracefully
func signalHandling(){
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt,os.Kill)
	go func() {
		select {
		case <-c:
			fmt.Print("Server: I got stopipeed!")
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
		pipe.RegisterTransferServer(s, &transferServer{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}