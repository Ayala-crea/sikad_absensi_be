package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func BimbinganDB() *gorm.DB {
	Connectionstrings := os.Getenv("POSTGRES")
	DB, err := gorm.Open(postgres.Open(Connectionstrings), &gorm.Config{ // Ubah ke postgres.Open
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
		fmt.Println("Cant Connect to Database : " + err.Error())
	} else {
		fmt.Println("Successfully connected to the database!")
	}
	return DB
}
