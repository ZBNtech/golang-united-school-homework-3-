package homework

import "testing"

func TestGetMessage(t *testing.T) {

	wanted_message := float32(3.5)
	value := [15]float32{1, 2, 3, 4, 5, 6}
	getMessage := average(value)

	if wanted_message != getMessage {
		t.Errorf("Get Message %v, want %v", getMessage, wanted_message)
	}

}
