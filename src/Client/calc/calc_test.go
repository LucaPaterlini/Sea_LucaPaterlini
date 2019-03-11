package calc

import (
	"reflect"
	"testing"
)
import "../../pipe"

func TestUniformNumbers(t *testing.T) {
	testInput := "(0888) 1151"
	expectedOutput := "+448881151"
	response := UniformNumbers(testInput)
	if response != expectedOutput {
		t.Errorf("Wrong output, expected %s got %s instead", expectedOutput, response)
	}
}

func TestRecordInit(t *testing.T) {
	// positive test
	expected := &pipe.Record{Id:111,Name:"luca",Email:"paterlini.luca@gmail.com",Phone:"+447523615177"}
	received,err :=RecordInit([]string{"111","luca","paterlini.luca@gmail.com","07523hello(=6)15177"})
	if err!=nil {
		t.Error(err.Error())
	}else if !reflect.DeepEqual(expected,received){
		t.Errorf("Err, no error received but, expected %v and received %v",expected,received)
	}
	// negative test
	_,err =RecordInit([]string{"hello","luca","paterlini.luca@gmail.com","07523hello(=6)15177"})
	if err==nil {t.Error("expected error due to the wrong input in (str) in Id")}
}
