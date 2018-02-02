package make


import (

	"../postgres"

)

func CurrentVersion() (int, error){
	i, err := postgres.GetCurrentVersion()
	if err != nil {
		return 0, err
	}
	return i, nil
}

func SetCurrent(version int) {
	postgres.SetCurrent(version)
}