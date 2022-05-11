package homework

import (
	"fmt"
	"testing"
)

func Test1GetMessage(t *testing.T) {

	wanted_message := []int64{7, 6, 5, 4, 3, 2}
	value := []int64{1, 2, 5, 15}
	getMessage := reverse(value)

	fmt.Println(wanted_message)
	fmt.Println(getMessage)

}
