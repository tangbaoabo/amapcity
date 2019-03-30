package convertor

import (
	"amapdistricts/types"
	"fmt"
	"testing"
)

func TestParseUrl(t *testing.T) {
	condition := types.QueryCondition{
		KeyWords:    "640300",
		SubDistrict: "1",
		Filter:      "640300",
	}
	url := ParseUrl(condition)
	fmt.Println(url)
}
