package db

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB
var err error

type User struct {
    gorm.Model
    ID uint `gorm:"primary_key"`
    Name string
}

var structList []interface{}

func init() {
    structList = append(structList, &User{})
}

func SetupDB() *gorm.DB {
    db, err = gorm.Open("mysql", "root:@/goprac?charset=utf8&loc=Local")
    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    for i := range structList {
        db.AutoMigrate(structList[i])
    }
    return db
}

