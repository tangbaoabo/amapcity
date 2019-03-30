package saver

import (
	"amapdistricts/types"
)

type Saver struct {
	Saver SaveIface
}

type SaveIface interface {
	Save(district types.District)
}

func (s *Saver) DistrictSaver() (district chan types.District, err error) {
	out := make(chan types.District)
	go func() {
		for {
			district := <-out
			s.dealWithResult(district)
		}
	}()
	return out, nil
}

//递归调用处理结果
func (s *Saver) dealWithResult(district types.District) {
	if len(district.Districts) > 0 {
		for _, value := range district.Districts {
			//为地区赋值parentCode
			value.ParentAdCode = district.AdCode
			//持久化
			s.Saver.Save(value)
			s.dealWithResult(value)
		}
	}
}
