package core

import (
	"amapdistricts/convertor"
	"amapdistricts/fetcher"
	"amapdistricts/parser"
	"amapdistricts/saver"
	"amapdistricts/types"
	"log"
	"time"
)

//TODO 其实不用维护一个conditionList,业务暂时根据adCode来查找城市，但是有个问题，
// 有的district的adCode和street的adCode一样，如果不加以控制，程序一直会陷入死循环

//单机版执行核心，采用队列方式实现

type SimpleCore struct {
	//队列
	Out []types.QueryCondition
	//持久化接口
	Saver saver.SaveIface
}

func (s *SimpleCore) Run(condition ...types.QueryCondition) {
	startTime := time.Now()
	//初始化
	for _, value := range condition {
		s.Out = append(s.Out, value)
	}
	for len(s.Out) > 0 {
		queryCondition := s.Out[0]
		s.Out = s.Out[1:]
		result := worker(queryCondition)
		//处理结果
		s.dealWithResult(result)
		//发起下次请求
		districts := result.Districts
		s.generateQueryRequest(queryCondition, districts)
	}
	log.Printf("程序运行结束,运行时间为%v秒,总共获取到数据%v条", time.Now().Sub(startTime).Seconds(), 1)
}

//递归解析结果为最小单位
var count int

func (s *SimpleCore) dealWithResult(district types.District) {
	if len(district.Districts) > 0 {
		for _, value := range district.Districts {
			//为地区赋值parentCode
			value.ParentAdCode = district.AdCode
			count++
			s.Saver.Save(value)
			s.dealWithResult(value)
		}
	}

}
func (s *SimpleCore) generateQueryRequest(condition types.QueryCondition, districts [] types.District) {
	for _, district := range districts {
		if district.Level == types.CITY {
			condition.SubDistrict = "2"
			condition.KeyWords = district.AdCode
			condition.Filter = district.AdCode
			s.Out = append(s.Out, condition)
		} else if district.Level == types.PROVINCE ||
			district.Level == types.COUNTRY {
			condition.SubDistrict = "1"
			condition.KeyWords = district.AdCode
			condition.Filter = district.AdCode
			s.Out = append(s.Out, condition)
		}
	}
}

func worker(condition types.QueryCondition) types.District {
	//1.转换请求
	url := convertor.ParseUrl(condition)
	//2.从服务器获取数据
	contents, err := fetcher.Fetcher(url)
	if err != nil {
		panic(err)
	}
	//3.校验并转换数据
	parseResult, err := parser.ParseDistrict(contents)
	result := parseResult.Districts[0]
	if err != nil {
		panic(err)
	}
	log.Printf("开始请求 %v 的数据url为：%s", result.Name, url)
	//4.返回结果
	return result
}
