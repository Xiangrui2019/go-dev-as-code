package lxd

import (
	"log"
	"os"

	lxc "github.com/lxc/lxd/client"
)

var LXDClient lxc.InstanceServer

func init() {
	var err error

	LXDClient, err = lxc.ConnectLXD(os.Getenv("LXD_SERVER_ADDR"), &lxc.ConnectionArgs{
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Fatal(err)
	}
}
