package conf

import (
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"strings"
)

// default config if config file not found
var config = map[string]string{
	"Name": "ShareX",
	"Host": "localhost",
	"Port": "8080",
}

var dbConn *gorm.DB
var dbURL = "root:Qwer1234@tcp(just:3306)/justlike?charset=utf8&parseTime=True"

func InitConfig(path string) map[string]string {
	//config = make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return config
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

func GetDB() *gorm.DB {
	if dbConn != nil {
		return dbConn
	}
	//db, err := gorm.Open("mysql", dbURL)
	//if err != nil {
	//	log.Println(err)
	//	return nil
	//}
	//
	//db.AutoMigrate(&Paste{})
	//dbConn = db
	//return db
	return nil
}
