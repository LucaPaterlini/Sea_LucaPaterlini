package calc

import (
	"../../pipe"
	"../config"
	"errors"
	"strconv"
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
	phoneNumberFiltered = config.PHONEPREFIX + phoneNumberFiltered
	return
}

// RecordInit initializing the item to send
func RecordInit(recordS []string)(item *pipe.Record,err error){
	recordS[3] = UniformNumbers(recordS[3])
	id, err := strconv.ParseInt(recordS[0], 0, 64)
	if err != nil {
		err = errors.New("invalid id cannot send a not int value:  err: " + err.Error())
		return
	}
	item = &pipe.Record{Id: id, Name: recordS[1], Email: recordS[2], Phone: recordS[3]}
	return
}
