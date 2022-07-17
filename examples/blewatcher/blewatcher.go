package main

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-winrtapi/winrt"
)

func main() {

	winrt.Initialize()

	watcher := winrt.NewBluetoothLEAdvertisementWatcher()
	watcher.Put_ScanningMode(winrt.BluetoothLEScanningMode_Active)

	watcher.Add_Received(func(sender *winrt.IBluetoothLEAdvertisementWatcher,
		args *winrt.IBluetoothLEAdvertisementReceivedEventArgs) com.Error {

		println(args.Get_AdvertisementType())
		println(args.Get_RawSignalStrengthInDBm())
		ad := args.Get_Advertisement()
		uuids := ad.Get_ServiceUuids()
		count := uuids.Get_Size()
		for n := uint32(0); n < count; n++ {
			uuid := uuids.GetAt(n)
			sUuid, _ := win32.GuidToStr(&uuid)
			println(sUuid)
		}
		return com.OK
	})
	watcher.Start()

	com.MessageLoop()
	winrt.Uninitialize()
}
