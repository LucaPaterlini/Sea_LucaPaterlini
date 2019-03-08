// main this package read a csv and send each record over to the port
// mentioned in the configuration in cost.go

package main

import (
	pb "../pipeProto"
	"bufio"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
)

// UniformNumbers standardize all the phone numbers in the formate [PREFIX]PHONENUMBER
func UniformNumbers(phoneNumber string) (phoneNumberFiltered string) {
	for _, c := range phoneNumber {
		if unicode.IsDigit(c) {
			phoneNumberFiltered += string(c)
		}
	}
	if len(phoneNumberFiltered) == 0 {
		return
	}
	if phoneNumberFiltered[0] == '0' {
		phoneNumberFiltered = phoneNumberFiltered[1:]
	}
	phoneNumberFiltered = PHONEPREFIX + phoneNumberFiltered
	return
}

// SendgRPC handle the connection , intialize and send a new record over the connection
func SendgRPC(record []string, c pb.TransferClient, ctx context.Context) (err error) {
	id, err := strconv.ParseInt(record[0], 0, 64)
	if err != nil {
		err = errors.New("invalid id cannot send a not int value: " + record[0] + "err: " + err.Error())
		return
	}
	item := &pb.Record{Id: id, Name: record[1], Email: record[1], Phone: record[2]}
	_, err = c.AddUpdateRecord(ctx, item)

	for err != nil {
		log.Println("Error on server response ID:", item.Id, " Error: ", err.Error())
		time.Sleep(WAITATTEMPT)
		var conn *grpc.ClientConn
		conn, err = grpc.Dial(ADDRESS, grpc.WithInsecure())
		c := pb.NewTransferClient(conn)
		_, err = c.AddUpdateRecord(ctx, item)
	}
	return
}

func main() {
	// initialize the read of the csv
	f, err := os.Open(READERCSVPATH)
	r := csv.NewReader(bufio.NewReader(f))
	// prepare the connection on the client side
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Printf("Unable to connect to %s: %s", ADDRESS, err.Error())
	} else {
		c := pb.NewTransferClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), TIMEOUTCONNECTION)
		defer cancel()
		// loop over each record of the csv
		if err == nil {
			var record []string
			for err != io.EOF {
				record, err = r.Read()
				// avoid first line and fields with not digit id
				if len(record) == 0 {
					continue
				}
				if _, err = strconv.Atoi(record[0]); err != nil {
					continue
				}
				record[3] = UniformNumbers(record[3])
				err = SendgRPC(record, c, ctx)
				if err != nil {
					break
				}
			}
		}
		fmt.Println("GOOD JOB!!Each record of the csv has been sent to the server ")
	}
}
