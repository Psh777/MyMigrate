package make

import (

	"../postgres"

)

func Create(s, path string) error{
	err := postgres.Create(s, path)
	return err
}
