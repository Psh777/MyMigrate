package flags

import (
	"flag"
	"../make"
	"../types"
	"net/url"
	"fmt"
	"net"
	"../postgres"
	"os"
	"strings"
)

func Flags() {

	path, url, up, down, initial, reset, grab, help, version, create := ReadFlags()

	if help == true {
		Help()
		return
	}

	if len(url) > 0 {
		Configuration(url)
	}

	if initial == true {
		err := make.Init(path)
		if err != nil {
			fmt.Println("Init error", err)
			os.Exit(0)
		} else {
			fmt.Println("Init done!")
		}
		return
	}

	if reset == true {
		err := make.Reset()
		if err != nil {
			fmt.Println("Reset error", err)
			os.Exit(0)
		} else {
			fmt.Println("Reset done!")
		}
		return
	}

	if grab == true {
		err := make.Grab(path)
		if err != nil {
			fmt.Println("Grab error", err)
			os.Exit(0)
		} else {
			fmt.Println("Grab done!")
		}
		return
	}

	if version == true {
		i, err := make.CurrentVersion()
		if err != nil {
			fmt.Println("Current version error", err)
			os.Exit(0)
		} else {
			fmt.Println("Current version:", i)
		}
		return
	}

	if len(create) > 0 {
		err := make.Create(create, path)
		if err != nil {
			fmt.Println("Create error", err)
			os.Exit(0)
		} else {
			fmt.Println("Create done!")
		}
		return
	}

	if up > 0 {
		make.MakeUp(up, path)
		return
	}

	if down > 0 {
		make.MakeDown(down, path)
		return
	}

	Help()

}

func ReadFlags() (string, string, int, int, bool, bool, bool, bool, bool, string) {
	var url = flag.String("url", "", "placeholder")
	var path = flag.String("path", "", "placeholder")
	var up = flag.Int("up", 0, "placeholder")
	var down = flag.Int("down", 0, "placeholder")
	var create = flag.String("create", "", "placeholder")

	flag.Parse()

	fmt.Println("path:", *path)
	fmt.Println("url:", *url)

	if len(*path) < 1 {
		*path = "."
	}

	var initial = false
	if len(flag.Args()) > 0 {
		if flag.Args()[0] == "init" {
			initial = true
		}
	}

	var reset = false
	if len(flag.Args()) > 0 {
		if flag.Args()[0] == "reset" {
			reset = true
		}
	}

	var grab = false
	if len(flag.Args()) > 0 {
		if flag.Args()[0] == "grab" {
			grab = true
		}
	}

	var help = false
	if len(flag.Args()) > 0 {
		if flag.Args()[0] == "help" {
			help = true
		}
	}

	var current = false
	if len(flag.Args()) > 0 {
		if flag.Args()[0] == "current" {
			current = true
		}
	}

	return *path, *url, *up, *down, initial, reset, grab, help, current, strings.Replace(*create, " ", "_", -1)
}

func UrlParse(s string) types.MyConfig {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	host, port, _ := net.SplitHostPort(u.Host)

	var config types.MyConfig
	config.PostgresScheme = u.Scheme
	config.PostgresBase = strings.Replace(u.Path, "/", "", -1)
	config.PostgresHost = host
	config.PostgresPort = port
	config.PostgresUser = u.User.Username()
	config.PostgresPassword, _ = u.User.Password()

	return config
}

func Configuration(url string) types.MyConfig {
	var config types.MyConfig
	if len(url) > 0 {
		config = UrlParse(url)
	}

	err := postgres.Init(config)
	if err != nil {
		fmt.Println("Postgres Url or Config error", err)
		os.Exit(0)
		//panic(err)
	}
	return config
}

func Help() {
	var help string = `
Arg:
	init		Initializing migration\n
	reset		Reset migration (for reinstall)\n
	grab		Saving data for migration\n
	current		Current version

Flags:
	-create $subject	Create version
	-up $version		Up version (int version id)
	-down $version		Down version (int version id)

Sample usage:
	//dev env
	go run main.go -url postgresql://user:password@localhost:5432/base -path init
	go run main.go -url postgresql://user:password@localhost:5432/base -path ./sql create initial_version
	go run main.go -url postgresql://user:password@localhost:5432/base -path ./sql grab

	//stage env
	go run main.go -url postgresql://user:password@localhost:5432/base -path ./sql init
	go run main.go -url postgresql://user:password@localhost:5432/base -path ./sql up 1
`

	fmt.Println(help)
}
