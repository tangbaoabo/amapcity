package convertor

import (
	"amapdistricts/config"
	"amapdistricts/types"
	"reflect"
	"strings"
)

func ParseUrl(condition types.QueryCondition) (url string) {
	return config.UrlPrefix + config.AppKey + convertObjToStrUrl(condition)
}

//拼装请求参数
func convertObjToStrUrl(src interface{}) (strUrl string) {
	var str []string
	typeOf := reflect.TypeOf(src)
	valueOf := reflect.ValueOf(src)
	fieldNum := typeOf.NumField()
	for i := 0; i < fieldNum; i++ {
		if strings.TrimSpace(valueOf.Field(i).String()) != "" {
			field := typeOf.Field(i)
			value := valueOf.Field(i)
			str = append(str, "&", field.Name, "=", value.String())
		}
	}
	return strings.ToLower(strings.Join(str, ""))
}
