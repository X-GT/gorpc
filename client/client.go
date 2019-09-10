package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/rpc"
	"os"
)

type Pay struct {
	Name string
	Amount int64
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		payuser := &Pay{Name: string(line), Amount:1000}
		b, err := json.Marshal(payuser)
		var reply bool
		err = client.Call("Listener.TopUp", b, &reply)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Call("Listener.GetAmount", b, &reply)
		if err != nil {
			log.Fatal(err)
		}
	}
}