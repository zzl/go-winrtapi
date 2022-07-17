package main

import (
	"github.com/zzl/go-winrtapi/winrt"
)

func main() {

	winrt.Initialize()

	sUri := "https://github.com/zzl/go-winrtapi/README.md"
	uri := winrt.NewUri_CreateUri(sUri)
	println(uri.Get_Domain())
	println(uri.Get_Path())

}
