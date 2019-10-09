package main

import (
	"log"
	"os"

	"github.com/wilriker/goduetapiclient/connection"
	"github.com/wilriker/goduetapiclient/connection/initmessages"
	"github.com/wilriker/goduetapiclient/types"
)

const (
	LocalSock = "/home/manuel/tmp/duet.sock"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	switch os.Args[1] {
	case "subscribe":
		subscribe()
	case "intercept":
		intercept()
	case "command":
		if len(os.Args) > 2 {
			for _, c := range os.Args[2:] {
				command(c)
			}
		} else {
			command("")
		}
	}
}

func command(code string) {
	cc := connection.CommandConnection{}
	err := cc.Connect(LocalSock)
	if err != nil {
		panic(err)
	}
	defer cc.Close()
	if code != "" {
		r, err := cc.PerformSimpleCode(code, types.SPI)
		if err != nil {
			panic(err)
		}
		log.Println(r)
	} else {
		mm, err := cc.GetSerializedMachineModel()
		if err != nil {
			panic(err)
		}
		log.Println(string(mm))
	}
}

func subscribe() {
	sc := connection.SubscribeConnection{}
	err := sc.Connect(initmessages.SubscriptionModeFull, "", LocalSock)
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	m, err := sc.GetMachineModel()
	if err != nil {
		panic(err)
	}
	log.Println(m)
}

func intercept() {
	bc := connection.InterceptConnection{}
	err := bc.Connect(initmessages.InterceptionModePre, LocalSock)
	if err != nil {
		panic(err)
	}
	defer bc.Close()
	for {
		c, err := bc.ReceiveCode()
		if err != nil {
			panic(err)
		}
		log.Println(c)
		err = bc.IgnoreCode()
		if err != nil {
			panic(err)
		}
	}
}
