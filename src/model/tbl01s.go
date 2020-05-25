package model

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// Users ユーザー情報のテーブル情報
type Tbl01 struct {
	ID           int `json:"id" gorm:"column:id;"`
	Name         string `json:"name" gorm:"column:name;"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:created_at;"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"column:updated_at;"`
}

func Create(name string) string{
	db, err := sql.Open("mysql", "ichiro:ichiro@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	fmt.Println(name)

	//query := `INSERT INTO "tbl01" ("name") VALUES ("`+name+`")`
	//query := `INSERT INTO tbl01 (name) VALUES ("aaa")`

	//query := "INSERT INTO `tbl01` (`name`) VALUES ('"+name+"')"
	//query := "INSERT INTO `tbl01` (`name`) VALUES ('えええ'); UPDATE `tbl01` SET `name`='インジェクション' WHERE `id` = 3; "
	query := "UPDATE `tbl01` SET `name`='インジェクション' WHERE `id` = 3; "

	fmt.Println(query)
	defer db.Close()
	_, err = db.Exec(query)
	if err != nil {
		return "作成失敗"
		panic(err.Error())
	}
	return "作成完了"

}

func FindByID(id string, pass string) bool {
	db, err := sql.Open("mysql", "ichiro:ichiro@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	fmt.Println(id, pass)

	query := "Select name FROM `tbl01` WHERE `id` = '"+id+"' AND `name` = '"+pass+"'"

	fmt.Println(query)

	defer db.Close()
	var name string
	err = db.QueryRow(query).Scan(&name)
	if err != nil {
		return false
	}
	fmt.Println(name)

	if name != "" {
		return true
	}

	return false
}

func GetAll() []Tbl01 {
	db, err := sqlConnect()
	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var tbl01 []Tbl01
	db.Order("id").Find(&tbl01)

	db.Close()

	return tbl01
}

func FindUser(name string) string {
	db, err := sql.Open("mysql", "ichiro:ichiro@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	fmt.Println(name)

	query := "Select name FROM `tbl01` WHERE `name` = '"+name+"'"

	fmt.Println(query)

	defer db.Close()

	//rows, err := db.Query(query)
	//if err != nil {
	//	return false
	//}
	//fmt.Println(name)
	//
	//if name != "" {
	//	return true
	//}
	//
	//return false
	return ""
}


// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "ichiro"
	PASS := "ichiro"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "mydb"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}
