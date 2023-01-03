package psns

import (
	"fmt"
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

func TestGenerate(t *testing.T) {
	generate, err := Generate(CMP, CP101, "02", "01", "01", 9999)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(generate.String())
	fmt.Println(generate.ZhString())
}
