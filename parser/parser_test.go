package parser

import (
	"amapdistricts/convertor"
	"amapdistricts/fetcher"
	"amapdistricts/types"
	"fmt"
	"testing"
)

func TestParseDistrict(t *testing.T) {
	condition := types.QueryCondition{
		KeyWords:    "640300",
		SubDistrict: "1",
		Filter:      "640300",
	}
	url := convertor.ParseUrl(condition)
	fmt.Println(url)
	contents, err := fetcher.Fetcher(url)
	if err != nil {
		panic(err)
	}
	result, err := ParseDistrict(contents)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v",result.Districts)
}
