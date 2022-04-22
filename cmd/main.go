package main

import (
	"github.com/BAN1ce/Ack/boot"
	"github.com/BAN1ce/Ack/pkg/api"
)

func main() {
	app := boot.NewApp()
	app.RegisterProvider(api.NewServer())
	if err := app.Run(); err != nil {
		panic(err)
	}

}
