package psns

import (
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	for _, v := range []string{
		"P02010201202210210005",
		"V04080201202210210002",
		"T02040201202210210003",
	} {
		serialNumber, err := Parse(v)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(serialNumber, serialNumber.ZhString(), serialNumber.CM.K.IsV())
	}
}
