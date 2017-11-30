package hongbao

import (
	"fmt"
	"testing"
)

func TestSendHongbao(t *testing.T) {
	fmt.Println("Test: Generate Hong Bao ...\n")
	var sum float32
	result := GenerateMoneyVector(float32(10), 8)
	for _, v := range result {
		sum += v
	}
	fmt.Println(result)
	if sum == float32(10) {
		fmt.Println("  ... Passed\n")
	} else {
		t.Error("  ... Fail\n")
	}
}
