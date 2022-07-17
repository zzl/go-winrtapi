package main

import (
	"github.com/zzl/go-winrtapi/winrt"
)

func main() {
	winrt.Initialize()

	sm := winrt.NewStringMap()
	sm.Insert("A", "AA")

	println(sm.HasKey("B"))
	println(sm.HasKey("A"))

}
