package saver

import (
	"amapdistricts/types"
	"log"
)

type ConsoleDisplay struct {
}

var count int

func (ConsoleDisplay) Save(district types.District) {
	count++
	log.Printf("chan got new Data: #%d %v", count, district)
}
