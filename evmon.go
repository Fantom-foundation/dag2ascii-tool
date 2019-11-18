package main

import (
	//"encoding/hex"
	//"context"

	//"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

	//"github.com/Fantom-foundation/go-lachesis/ethapi"

	//"github.com/btcsuite/btcd/rpcclient"
	//"github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/core/types"
	//"math/big"
	//"log"

	//"github.com/ethereum/go-ethereum/accounts"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/state"
	//"github.com/ethereum/go-ethereum/core/vm"
	//"github.com/ethereum/go-ethereum/ethdb"
	//"github.com/ethereum/go-ethereum/params"
	//"github.com/ethereum/go-ethereum/ethclient"

	//"github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	//"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"

	//"github.com/Fantom-foundation/go-lachesis/evm_core"
	"github.com/Fantom-foundation/go-lachesis/hash"

	"github.com/Fantom-foundation/go-lachesis/inter"
)

type DAG_event_t struct {
	child   string
	parents hash.Events
	parnum  int
	idx     uint32
	isread  bool
}

func hextostr(a []hexutil.Bytes) string {
	s := fmt.Sprint(a)
	s = strings.TrimSuffix(s, "]")
	s = strings.TrimPrefix(s, "[")
	return s
}

func main() {

	// host := string("http://18.222.120.223:4003")
	host := string("http://127.0.0.1:18545")
	// host := string("enode://869daf0a0967b52029031af42860fc778feeb2f6e19eb701894953f347b37829c2391b9232ab4cdf1e2f4e6da48dc8bcb83392f45070c91894a9892883bbd5fc@127.0.0.1:5050")
	fmt.Println("====== DAG Monitor Ev ===============================================================")
	Node1, err := rpc.DialHTTP(host)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("=== Connected to lachesis Node (%s)\n", host)
	}
	fmt.Println("====== Start GetHeads ======")
	fmt.Println("=== Getting Event Heads...")
	var Heads []hexutil.Bytes
	err = Node1.Call(&Heads, "debug_getHeads", 1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("=== Right...got events heads:\n")
	}
	EvHash := hextostr(Heads)
	fmt.Println(EvHash)
	var (
		ev inter.Event
		//ev hash.Event
		//dag       []DAG_event_t
		//dp DAG_event_t
		//cur_index uint32 = 0
		//i         uint32
		//k         uint32
	)

	fmt.Println("=== Ask for first event - head of our DAG ===") //===============================

	//err = Node1.Call(&ev, "debug_getEv", 1)
	err = Node1.Call(&ev, "debug_getEvent", EvHash, false)

	fmt.Println(ev.String())
	fmt.Print("Creator - ")
	fmt.Println(ev.Creator.String())
	fmt.Print("Parents - ")
	mm := ev.Parents
	fmt.Println(mm)

	/*
		if err != nil {
			log.Fatal(err)
		} else {
			dp.child = ev.String()
			fmt.Printf("child - ")
			fmt.Println(dp.child)
			dp.parents = ev.Parents
			fmt.Printf("parents - ")
			fmt.Println(dp.parents)
			dp.parnum = len(dp.parents)
			fmt.Printf("parnum - %d \n", dp.parnum)
			dp.isread = false
			dp.idx = cur_index

			dag = append(dag, dp)
			cur_index++

			fmt.Printf("DAG first point added, idx = %d\n", cur_index)
		}*/
	//=========================================================================
	//fmt.Print(dp.parents)

	/*
		fmt.Println("=== Ask for Event by Event")
		for i<cur_index {  // we must check all "parents-unread" events and
		                   // get their parents till stub events

			for k = range dag[cur_index-1] {
			err = Node1.Call(ev, "debug_getEvent", EvHash, true)
			if err != nil { continue } else {

				dp.child = EvHash
				dp.parents = ev.Parents
				dp.parnum = len(dp.parents)
				dp.isread = false
				dp.idx = cur_index
				dag = append(dag, dp)
				cur_index++
				fmt.Printf("DAG point added, idx = %d\n", cur_index)
			  }
		    }
				//fmt.Println("Right...got Event")
				//fmt.Printf("Header: ")
				//fmt.Println(ev.EventHeader.String())
				//fmt.Printf("Epoch:  ")
				//fmt.Println(ev.Epoch)
				//fmt.Printf("Version:  ")
				//fmt.Println(ev.Version)
				//fmt.Printf("IsSelf:  ")
				//fmt.Println(ev.IsSelfParent(ev.Hash()))
				//fmt.Printf("Frame:  ")
				//fmt.Println(ev.Frame)
				//fmt.Printf("Creator:  ")
				//fmt.Println(ev.Creator)
				//fmt.Printf("%d Parents:  ", len(ev.Parents))
				//fmt.Print(ev.Parents)

				fmt.Println(ev.Parents.String())
				fmt.Printf("Seq:  ")
				fmt.Print(ev.Seq)
			}
		}
	*/

	//fmt.Printf("\n:::::::::::   %d Events in this Epoch \n", len(ev))

	fmt.Println("=====================================================================================")
}
