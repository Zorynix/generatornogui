package main

import (
	"utils/utils"

	_ "golang.org/x/sys/windows"
)

func main() {

	utils.DBParse()
	utils.Startapp()
}
