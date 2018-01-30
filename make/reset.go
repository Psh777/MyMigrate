package make


import (

"../postgres"

)

func Reset() error{
	err := postgres.Reset()
	return err
}
