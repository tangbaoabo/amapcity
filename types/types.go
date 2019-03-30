package types

//查询条件

type QueryCondition struct {
	KeyWords    string
	SubDistrict string
	Page        string
	Offset      string
	Extensions  string
	Filter      string
	Callback    string
	Output      string
}

//返回状态码：
//1表示成功，0表示失败
//状态信息：
//when status = 1 ,info = "ok"
//else show the error of response
//返回状态说明:
//10000代表正确，详情参阅info状态表

type ParseResult struct {
	Status    string     `json:"status"`
	Info      string     `json:"info"`
	InfoCode  string     `json:"infocode"`
	Count     string     `json:"count"`
	Districts []District `json:"districts"`
}

//实体结果

type District struct {
	Id           int
	AdCode       string      `json:"adcode"`
	CityCode     interface{} `json:"citycode"`
	Name         string      `json:"name"`
	Center       string      `json:"center"`
	Level        string      `json:"level"`
	ParentAdCode string
	Districts    []District `json:"districts"`
}

//请求枚举
type DistrictLevel int32

const (
	COUNTRY = "country"

	PROVINCE = "province"

	CITY = "city"

	DISTRICT = "district"

	STREET = "street"
)
