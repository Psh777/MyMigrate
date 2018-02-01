package files

import (
	"io/ioutil"
	//"fmt"
)

func Read(filename string) ([]byte, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		//fmt.Print("error: ", err)
		return file, err
	}
	return file, nil
}