package main

import (
	"fmt"
	"log"
	"github.com/onrik/ethrpc"
)

func main() {

	client := ethrpc.New("http://127.0.0.1:8545")

	version, err := client.Web3ClientVersion()

	if err != nil { log.Fatal(err) }
	fmt.Println(version)

	// Send 1 eth
	txid, err := client.EthSendTransaction(ethrpc.T{
		From:  "0x6247cf0412c6462da2a51d05139e2a3c6c630f0a",
		To:    "0xcfa202c4268749fbb5136f2b68f7402984ed444b",
		Value: ethrpc.Eth1(),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(txid)
}
