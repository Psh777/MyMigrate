package make

import (
	"../files"
	"../postgres"
	"fmt"
)

func Init(path string) error {

	last, _ := postgres.GetCurrentVersion()

	postgres.DropMainTable()
	err := postgres.InitialTable()
	if err == nil {
		fmt.Println("Init migrate DB table created.")
	}

	//restore
	sql, err := files.Read(path + "/grab.sql")
	if err == nil {
		_ = postgres.RunSql(sql)
	} else {
		fmt.Println("Grab file not found. Backup not upload.")
	}

	if last > 0 {
		postgres.SetCurrent(last)
	}
	return err
}
