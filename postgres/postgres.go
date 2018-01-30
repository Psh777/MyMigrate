package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"../types"
	//"fmt"
)

var DB *sql.DB

func Init(myconfig types.MyConfig)(error) {

	var dataSource string = myconfig.PostgresScheme + "://" + myconfig.PostgresUser + ":" + myconfig.PostgresPassword + "@" + myconfig.PostgresHost + ":" + myconfig.PostgresPort + "/" + myconfig.PostgresBase + "?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", dataSource)
	if err != nil {
		return err
	}

	if err = DB.Ping();
		err != nil {
		return err
	}

	//fmt.Println("Postgress connect OK")

	return nil
}

func Check() bool {

	err := DB.Ping()

	if err != nil {
		return false
	}

	return true
}


