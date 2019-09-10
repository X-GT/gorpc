package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Listener int

type Pay struct {
	Name string
	Amount int64
}

func (l *Listener) TopUp(line []byte, ark *bool) error {
	var payuser Pay
	err := json.Unmarshal(line, &payuser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("berhasil ditambahkan %s dengan nominal %v", payuser.Name, payuser.Amount)
	return nil
}

func (l *Listener) GetAmount(line []byte, ark *bool) error {
	var payuser Pay
	err := json.Unmarshal(line, &payuser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("saldo %s dengan nominal %v", payuser.Name, payuser.Amount)
	return nil
}

func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	fmt.Println(listener)
	rpc.Register(listener)
	rpc.Accept(inbound)
}