package core

import (
	"amapdistricts/types"
	"log"
	"time"
)

//并发版是核心实现
//采用channel
type ConcurrentCore struct {
	Scheduler Scheduler
	WorkCount int
	SaverChan chan types.District
}

type Scheduler interface {
	Submit(condition types.QueryCondition)
	WorkChan() chan types.QueryCondition
	Run()
}

func (s *ConcurrentCore) Run(seeds ...types.QueryCondition) {
	startTime := time.Now()
	defer
		log.Printf("程序运行结束,运行时间为%v秒", time.Now().Sub(startTime).Seconds())
	out := make(chan types.District)
	//新建channel
	s.Scheduler.Run()
	for i := 0; i < s.WorkCount; i++ {
		s.CreateWorker(s.Scheduler.WorkChan(), out)
	}

	for _, value := range seeds {
		s.Scheduler.Submit(value)
	}
	for {
		select {
		case district := <-out:
			//处理结果
			go func() { s.SaverChan <- district }()
			//获取新请求
			for _, dist := range district.Districts {
				s.submitSubTask(dist)
			}
			//5秒之后默认关闭
		case <-time.After(time.Second * 5):
			break
		}
	}

}

func (s *ConcurrentCore) submitSubTask(dist types.District) {
	condition := types.QueryCondition{}
	if dist.Level == types.CITY {
		condition.SubDistrict = "2"
		s.component(dist, condition)
	} else if dist.Level == types.PROVINCE ||
		dist.Level == types.COUNTRY {
		condition.SubDistrict = "1"
		s.component(dist, condition)
	}
}
func (s *ConcurrentCore) component(dist types.District, condition types.QueryCondition) {
	condition.KeyWords = dist.AdCode
	condition.Filter = dist.AdCode
	s.Scheduler.Submit(condition)
}

//创建worker
func (s *ConcurrentCore) CreateWorker(in chan types.QueryCondition, out chan types.District) {
	go func() {
		for {
			condition := <-in
			district := worker(condition)
			out <- district
		}
	}()
}
