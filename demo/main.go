package main

import einfo "github.com/euskadi31/go-einfo"

func main() {
	einfo.Info("test info")
	einfo.Warn("test warn")
	einfo.Error("test error")

	einfo.Begin("Fetch data")
	einfo.End(true, "Cannot fetch data")

	einfo.Begin("Fetch data")
	einfo.End(false, "Cannot fetch data")
}
