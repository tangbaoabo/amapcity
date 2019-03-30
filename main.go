package main

import (
	"amapdistricts/config"
	"amapdistricts/core"
	"amapdistricts/saver"
	"amapdistricts/scheduler"
	"amapdistricts/types"
)

func main() {
	//查询初始条件
	condition := types.QueryCondition{
		KeyWords:    "100000",
		SubDistrict: "1",
	}
	current(condition)

}

func current(condition types.QueryCondition) {
	itemSave := saver.Saver{
		Saver: &saver.MySQLSaver{
			DriverName:     config.MYSQL,
			DataSourceName: config.DataSourceName,
		},
	}
	//获取channel
	districtSaver, err := itemSave.DistrictSaver()
	if err != nil {
		panic(err)
	}
	//配置核心启动
	concurrentCore := core.ConcurrentCore{
		Scheduler: &scheduler.SimpleScheduler{},
		//配置worker数量
		WorkCount: 10,
		SaverChan: districtSaver,
	}
	concurrentCore.Run(condition)

}

func single(condition types.QueryCondition) {
	simpleCore := core.SimpleCore{
		Out:   make([]types.QueryCondition, 0),
		Saver: saver.ConsoleDisplay{},
	}
	simpleCore.Run(condition)
}
