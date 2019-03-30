package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"stackoverflow/models"
	"strings"
)

func main() {
	db, err := gorm.Open("mysql", "root:example@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Language{})

	//admin := models.User{
	//	Name: "Admin",
	//}
	//db.Save(&admin)

	jsonText := `{
	"Name": "EN",
	"Users": [
		{
			"ID": 1
		}
	]
}`

	language := models.Language{}

	decoder := json.NewDecoder(strings.NewReader(jsonText))
	if err := decoder.Decode(&language); err != nil {
		panic(err)
	}

	if err := db.Save(&language).Error; err != nil {
		panic(err)
	}

	fmt.Println(language)
}
