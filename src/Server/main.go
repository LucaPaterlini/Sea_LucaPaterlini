package main

import (
	pb "../pipeProto"
	"context"
	"github.com/globalsign/mgo"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"log"
	"net"
)

var accountTable *mgo.Collection

type server struct{}

func (s *server) AddUpdateRecord(ctx context.Context, in *pb.Record) (_ *empty.Empty, err error) {
	log.Printf("Receiver: %v\n", in)
	_, err = accountTable.UpsertId(in.Id, &in)
	if err != nil {
		log.Println("Updated", in.Id)
	}
	return &empty.Empty{}, nil
}

func main() {
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
		pb.RegisterTransferServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
