package make

import (
	"../files"
	"../postgres"
	"fmt"
)

func Init(path string) error {

	err := postgres.InitialTable()
	if err == nil {
		fmt.Println("Init migrate DB table created.")
	}

	//restore
	sql, err := files.Read(path + "/grab.sql")
	if err == nil {
		postgres.RunSql(sql)
	} else {
		fmt.Println("Grab file not faund. Backup not upload.")
	}

	return err
}
