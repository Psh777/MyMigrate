package make

import (
	"../files"
	"../postgres"
)

func Init(path string) error {

	err := postgres.InitialTable()

	//restore
	sql, err := files.Read(path + "/grab.sql")
	if err == nil {
		postgres.RunSql(sql)
	}

	return err
}
