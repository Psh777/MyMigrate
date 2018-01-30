package make

import (
	"../files"
	"strconv"
	"fmt"
	"../postgres"
	"os"
)

func MakeDown(version int, path string) {

	current, err := postgres.GetCurrentVersion()
	if err != nil {
		return
	}
	list, err := postgres.GetVersionsDown(current, version)

	if err != nil {
		return
	}

	for i := 0; i < len(list); i++ {
		var filename = path + "/" + strconv.Itoa(list[i].IdVersion) + "_" + list[i].Subject + "_down." + "sql"
		fmt.Println("Read file:", filename)
		sql, err := files.Read(filename)
		if err !=nil {
			os.Exit(1)
		}
		postgres.RunSql(sql)
	}

	postgres.ClearCurrent()
	postgres.SetCurrent(version)

}
