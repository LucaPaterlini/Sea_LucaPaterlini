package mock_pipe

import (
	"../pipeProto"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/empty"
	"testing"
	"time"
)

func TestAddUpdateRecord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransferClient := NewMockTransferClient(ctrl)

	mockTransferClient.EXPECT().AddUpdateRecord(
		gomock.Any(),
		gomock.Any(),
	).Return(&empty.Empty{}, nil)
	testAddUpdateRecord(t,mockTransferClient)
}

func testAddUpdateRecord(t *testing.T, client pipe.TransferClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := client.AddUpdateRecord(ctx, &pipe.Record{Name: "unit_test"})
	if err != nil {
		t.Errorf("mocking failed")
	}
}