package homework

import (
	"fmt"
	"testing"
)

func Test3GetMessage(t *testing.T) {

	wanted_message := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	value := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	getMessage := sortMapValues(value)

	fmt.Println(wanted_message)
	fmt.Println(getMessage)
}
