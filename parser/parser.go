package parser

import (
	"amapdistricts/config"
	"amapdistricts/types"
	"encoding/json"
	"fmt"
)

func ParseDistrict(originJson []byte) (result types.ParseResult, err error) {
	parseResult := types.ParseResult{}
	valid := json.Valid(originJson)
	if !valid {
		return parseResult, fmt.Errorf("invalid json input:%s", originJson)
	}
	err = json.Unmarshal(originJson, &parseResult)
	if err != nil {
		return parseResult, err
	}
	err = validResult(parseResult)
	if err != nil {
		return parseResult, err
	}
	return parseResult, nil
}

func validResult(result types.ParseResult) (err error) {
	if result.Status != config.GetSuccess {
		return fmt.Errorf("api response status is fail:%s", result.Status)
	}
	if result.InfoCode != config.InfoCodeSuccess {
		return fmt.Errorf("infoCode is not expected, code is :%s,error info is %s", result.Status, result.Info)
	}
	return nil
}
