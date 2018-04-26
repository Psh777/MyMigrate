package make

import (
	"../postgres"
	"os"
	"io/ioutil"
	"fmt"
)

func Grab(path string) error {
	list, err := postgres.Grab()

	if err != nil {
		return err
	}

	var aaa string
	for i := 0; i < len(list); i++ {
		//file_base, _ := ioutil.ReadFile(path + "/grab.sql")
		file := fmt.Sprintf("INSERT INTO my_migrate_versions (id_version, current_version, subject) VALUES (%v, %v, '%s');", list[i].IdVersion, false, list[i].Subject)
		//fmt.Println("write", file)
		aaa = aaa + file + "\n"
	}
	os.Remove(path + "/grab.sql")
	ioutil.WriteFile(path + "/grab.sql", []byte(aaa), 777)
	return err
}
