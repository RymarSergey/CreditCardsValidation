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
		"pan" : "5375414112236376",
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

	for i, j := len(cardPan)-1, 1; i >= 0; i-- {
		temp, r := strconv.ParseInt(string(cardPan[i]), 10, 8)
		if r != nil {
			sum = 0
			break
		}
		if j%2 == 0 {
			temp = temp * 2
			if temp > 9 {
				temp -= 9
			}
		}
		j++
		sum += temp
	}
	return (sum != 0) && (sum%10 == 0)
}
