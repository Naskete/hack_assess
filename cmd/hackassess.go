package main

import (
	"HackAssess/internal/dao"
	"HackAssess/internal/http"
)

func main() {
	dao.DatabaseInit()
	http.Init()
}
