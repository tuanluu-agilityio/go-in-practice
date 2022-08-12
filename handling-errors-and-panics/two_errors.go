package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("The request timed out")
var ErrReject = errors.New("The request was rejected")

func main() {
	response, err := SendRequest("Hello")
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying.")
		response, err = SendRequest("Hello")
	}
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func SendRequest(req string) (string, error) {
	switch rand.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrReject
	default:
		return "", ErrTimeout
	}
}