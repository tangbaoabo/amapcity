package saver

import (
	"amapdistricts/types"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type MySQLSaver struct {
	DriverName     string
	DataSourceName string
}


func (s *MySQLSaver) Save(district types.District) {
	count++
	log.Printf("chan got new Data: #%d %v", count, district)
	switch district.CityCode.(type) {
	case []interface{}:
		district.CityCode = ""
	default:
	}
	db, err := sql.Open(s.DriverName, s.DataSourceName)
	defer db.Close()
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO tb_city(`id`,`name`,`ad_code`,`city_code`,`center`,`level`,`parent_ad_code`)VALUES(null,?,?,?,?,?,?)")
	defer stmt.Close()
	checkErr(err)
	_, err = stmt.Exec(district.Name, district.AdCode, district.CityCode, district.Center, district.Level, district.ParentAdCode)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
