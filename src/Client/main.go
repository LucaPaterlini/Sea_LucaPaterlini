// Package main this package read a csv and send each record over to the port
// mentioned in the configuration in cost.go

package main

import (
	"../pipe"
	"./calc"
	"./config"
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

// FetchCsvSendgRPC read from the csv, invoke the checks and send it over the stream
func FetchCsvSendgRPC(r *csv.Reader,stream pipe.Transfer_AddUpdateRecordClient, chanTerm <-chan os.Signal, wg *sync.WaitGroup){
	defer wg.Done()
	// loop over each record of the csv
	var recordS []string
	var errRead error
	serial := int64(0)
	errRead=nil
	L:for errRead != io.EOF {
			serial +=1
			select {
			case <-chanTerm:
				fmt.Print("Client: I got stopipeed!")
				break L
			default:
				recordS, errRead = r.Read()
				// avoid first line and fields with not digit id
				if len(recordS) == 0 {
					continue
				}
				if _, errRead = strconv.Atoi(recordS[0]); errRead != nil {
					continue
				}
				// prepare the record and send it
				item,err := calc.RecordInit(recordS)
				item.Serial = serial
				if err!=nil{log.Println(err.Error())}
				log.Println("Sending",item)
				err = stream.Send(item)

				// returning the ack object, potential improvement
				// resend object that have not been sent correctly
				reply, err := stream.Recv()
				if err != nil {
					log.Println("Local Error:",err.Error())
					break L
				}
				if reply.Err {
					log.Println("Server error:",reply.Message," Seq Num.",reply.Serial)
					break L
				}
				// added for debug to test the effect of signals
				// time.Sleep(time.Second * 2)
			}
		}
		if errs := stream.CloseSend();errs!=nil{
			fmt.Println("Error while closing the channel",errRead.Error())
			return
		}
		if errRead == io.EOF {fmt.Println("GOOD JOB!!Each record of the csv has been sent to the server")}
}

func main() {
	// channel that received the signals form the term
	chanTerm := make(chan os.Signal, 1)
	signal.Notify(chanTerm, os.Interrupt,os.Kill)
	// initialize the read of the csv
	f, err := os.Open(config.READERCSVPATH)
	if err !=nil{log.Println(err.Error());os.Exit(1)}
	reader := csv.NewReader(bufio.NewReader(f))
	// prepare the connection on the client side
	conn, err := grpc.Dial(config.ADDRESS, grpc.WithInsecure())
	if err !=nil{log.Println(err.Error());os.Exit(2)}
	client := pipe.NewTransferClient(conn)
	if err != nil {
		log.Printf("Unable to connect to %s: %s", config.ADDRESS, err.Error())
		os.Exit(1)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.AddUpdateRecord(ctx)
	if err!=nil{
		log.Println("Error while initializing the stream channel",err.Error())
		os.Exit(3)
	}
	//creating a waiting group to wait the routing FetchCsvSendgRPC
	var wg sync.WaitGroup
	wg.Add(1)
	go FetchCsvSendgRPC(reader,stream,chanTerm,&wg)
	wg.Wait()
}
