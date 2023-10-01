package main

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-winrtapi/winrt"
)

func main() {

	//It's important to call winrt.Initialize first.
	//A default com scope is created in this call.
	winrt.Initialize()

	hc := winrt.NewHttpClient()
	sUrl := "https://www.microsoft.com/robots.txt"
	uri := winrt.NewUri_CreateUri(sUrl)
	gsa := hc.GetStringAsync(uri.IUriRuntimeClass)

	gsa.Put_Progress(func(asyncInfo *winrt.IAsyncOperationWithProgress[string, winrt.HttpProgress],
		progressInfo winrt.HttpProgress) com.Error {
		println("Bytes received:", progressInfo.BytesReceived)
		return com.OK
	})

	gsa.Put_Completed(func(asyncInfo *winrt.IAsyncOperationWithProgress[string, winrt.HttpProgress],
		asyncStatus winrt.AsyncStatus) com.Error {
		content := asyncInfo.GetResults()
		println(content)
		//callbacks are called from runtime thread pool, so PostQuitMessage will not work.
		win32.PostThreadMessage(com.GetContext().TID, win32.WM_QUIT, 0, 0)
		return com.OK
	})

	com.MessageLoop()
}
