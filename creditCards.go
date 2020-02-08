package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type creditCard struct {
	Pan       string `json:"pan"`
	ValidThru string `json:"valid_thru"`
	Owner     string `json:"owner,omitempty"`
	Cvc2      string `json:"cvc_2"`
}

func main() {

	jSonString := `{
		"pan" : "5375414874236376",
		"valid_thru" : "31/23",
		"owner" : "Rymar Sergey",
		"cvc_2" : "123"
	}`
	cardParse := creditCard{}
	owner := json.Unmarshal([]byte(jSonString), &cardParse)
	if owner != nil {
	}
	fmt.Println(IsCreditCardValid(cardParse.Pan))

}

func IsCreditCardValid(cardPan string) bool {
	cardPan = strings.ReplaceAll(cardPan, " ", "")
	cardPan = strings.ReplaceAll(cardPan, "-", "")

	if len(cardPan) < 13 && len(cardPan) > 19 {
		return false
	}
	var sum int64 = 0

	for i, value := range cardPan {
		temp, r := strconv.ParseInt(string(value), 10, 8)
		if r != nil {
			sum = 0
			break
		}
		if len(cardPan)%2 == 0 {
			if i%2 == 0 {
				temp = chengeDigit(temp)
			}
		} else {
			if i%2 == 1 {
				temp = chengeDigit(temp)
			}
		}
		sum += temp
	}
	return (sum != 0) && (sum%10 == 0)
}

func chengeDigit(temp int64) int64 {

	temp = temp * 2
	if temp > 9 {
		temp -= 9

	}
	return temp
}
