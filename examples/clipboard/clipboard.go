package main

import (
	"fmt"
	"github.com/zzl/go-winrtapi/winrt"
)

func main() {

	winrt.Initialize()

	htmlDescription := "<h1>Hello WinRT clipboard</h1><hr/>" +
		"<div><em>Powered by go-winrtapi</em></div>"

	textDescription := "Hello WinRT clipboard\n---------------------\n" +
		"Powered by go-winrtapi"

	//
	dataPackage := winrt.NewDataPackage()

	htmlFormatHelperStatics := winrt.NewIHtmlFormatHelperStatics()
	htmlFormat := htmlFormatHelperStatics.CreateHtmlFormat(htmlDescription)
	dataPackage.SetHtmlFormat(htmlFormat)
	dataPackage.SetText(textDescription)

	//
	print("Press ENTER to set clipboard content..")
	_, _ = fmt.Scanln()
	clipboard := winrt.NewIClipboardStatics()
	clipboard.SetContent(dataPackage.IDataPackage)
	clipboard.Flush()
	println("clipboard content set.\n")

	//
	print("Press ENTER to clear the clipboard..")
	_, _ = fmt.Scanln()
	clipboard.Clear()
	println("clipboard cleared.\n")

	//
	print("Press ENTER to exit..")
	_, _ = fmt.Scanln()

}
