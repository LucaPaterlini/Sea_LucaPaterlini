package main

import (
	pb "../pipe"
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestUniformNumbers(t *testing.T) {
	testInput := "(0888) 1151"
	expectedOutput := "+448881151"
	response := UniformNumbers(testInput)
	if response != expectedOutput {
		t.Errorf("Wrong output, expected %s got %s instead", expectedOutput, response)
	}
}

func TestSendgRPC(t *testing.T) {
	// preparing the client
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		t.Errorf("Unable to connect to %s: %s", ADDRESS, err.Error())
	}
	c := pb.NewTransferClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUTCONNECTION)
	defer cancel()

	err = SendgRPC([]string{"1", "a", "b", "c"}, c, ctx)
	if err != nil {
		t.Error("Unable to send a record:", err.Error())
	}
}
