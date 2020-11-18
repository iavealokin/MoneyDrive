package store_test

import (
	"os"
	"testing")



var (
	databaseURL string
)

func TestMain(m *testing.M){
 databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL== ""{
		databaseURL = "host=localhost user=Alexandr password=Cfyz11005310 dbname=moneydrive_dev sslmode=disable"
	}


	os.Exit(m.Run())
}