package make

import (

	"../postgres"

)

func Create(s, path string) error{
	id, err := postgres.GetLastVersion()
	id = id + 1
	if err != nil {
		return err
	}
	err = postgres.Create(id, s, path)
	return err
}
