package main

import (
	//"encoding/hex"

	//"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"io/ioutil"
	"net/http"
	"strings"
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
	//"github.com/ethereum/go-ethereum/rpc"
	//"github.com/Fantom-foundation/go-lachesis/evm_core"
	//"github.com/Fantom-foundation/go-lachesis/hash"
)

/*
type DAG_event_t struct {
	child   string
	parents hash.Events
	parnum  int
	idx     uint32
	isread  bool
}*/

type EventHeader struct {
	ClaimedTime      int64    `json:"claimedTime"`
	Creator          string   `json:"creator"`
	Epoch            int64    `json:"epoch"`
	ExtraData        string   `json:"extraData"`
	Frame            int64    `json:"frame"`
	GasPowerLeft     int64    `json:"gasPowerLeft"`
	GasPowerUsed     int64    `json:"gasPowerUsed"`
	Hash             string   `json:"hash"`
	IsRoot           bool     `json:"isRoot"`
	Lamport          int64    `json:"lamport"`
	MedianTime       int64    `json:"medianTime"`
	Parents          []string `json:"parents"`
	PrevEpochHash    string   `json:"prevEpochHash"`
	Seq              int64    `json:"seq"`
	TransactionsRoot string   `json:"transactionsRoot"`
	Version          int      `json:"version"`
}
type Event struct {
	EventHeader
	Transactions []string `json:"transactions"`
}

type GetHeadResponse struct {
	JsonRPC string   `json:"jsonrpc"`
	Id      int64    `json:"id"`
	Heads   []string `json:"result"`
}
type RPC struct {
	url string
}

func (rpc *RPC) call(reqBody string) ([]byte, error) {
	reqIO := strings.NewReader(reqBody)
	resp, err := http.Post(rpc.url, "application/json", reqIO)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// 	log.Printf("REQUEST: %s\nRESPONSE: %s\n", reqBody, body)
	return body, nil
}
func NewRPC(host string, port int) *RPC {
	return &RPC{
		url: "http://" + host + ":" + strconv.FormatInt(int64(port), 10) + "/",
	}
}

type EventResponse struct {
	JsonRPC string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Result  Event  `json:"result"`
}

func GetA(a string) string {
	s := fmt.Sprint(a)
	s = strings.TrimSuffix(s, "]")
	s = strings.TrimPrefix(s, "[")
	return s
}
func (rpc *RPC) GetEvent(hash string) (*Event, error) {
	req := `{"jsonrpc":"2.0","method":"debug_getEvent","params":["` + hash + `", true],"id":1}`
	body, err := rpc.call(req)
	if err != nil {
		log.Printf("Call RPC error: %s\n", err)
		return nil, err
	}
	head := EventResponse{}
	err = json.Unmarshal(body, &head)
	if err != nil {
		log.Printf("Json parse response debug_getHeads body error: %s\n", err)
		return nil, err
	}
	return &head.Result, nil
}

type DAG_vertex struct {
	Name    string
	Vertex  string
	Parents []string
	Index   int
	PIndex  []int
}

func PrintVertex(vtx DAG_vertex) {
	par_len := len(vtx.Parents)
	fmt.Printf("idx: %4d, name=%s %d Parents - ", vtx.Index, vtx.Name, par_len)
	fmt.Println(vtx.Parents)
}

func main() {

	fmt.Println("====== DAG Monitor Ev ===============================================================")
	fmt.Println("=== Ask for first event - head of our DAG ===")

	req := (`{"jsonrpc":"2.0","method":"debug_getHeads","params":[-1],"id":1}`)
	myrpc := NewRPC("127.0.0.1", 18545)

	var current_index, added_index, i, par_len int
	var eve *Event
	var erre error
	var dag []DAG_vertex
	var vtx DAG_vertex

	body, err := myrpc.call(req)
	var tip GetHeadResponse
	if err == nil {
		current_index = 0
		added_index = 0
		//==== Here we have the Head of DAG to start collect vertexes ===
		fmt.Printf(">>> Heads: ")
		err = json.Unmarshal(body, &tip)
		fmt.Println(tip.Heads)
		for i = 0; i < len(tip.Heads); i++ {
			eve, erre = myrpc.GetEvent(tip.Heads[i])
			if erre == nil {
				vtx.Name = "0."
				vtx.Vertex = eve.EventHeader.Hash
				vtx.Parents = eve.EventHeader.Parents
				vtx.Index = i
				par_len = len(vtx.Parents)
				added_index += par_len
				dag = append(dag, vtx)
				PrintVertex(vtx)
			}
		}
		fmt.Println("======================================================")
		for current_index < added_index {
			par_len = len(dag[current_index].Parents)
			// go across all parens and add parent vertexes
			for i = 0; i < par_len; i++ {
				eve, erre = myrpc.GetEvent(dag[current_index].Parents[i])
				if erre == nil {
					vtx.Name = "."
					vtx.Vertex = eve.EventHeader.Hash
					vtx.Parents = eve.EventHeader.Parents
					vtx.Index = current_index
					added_index++
					vtx.PIndex = append(vtx.PIndex, added_index)
					dag = append(dag, vtx)
					PrintVertex(vtx)
				}
			}
			current_index++
		}
	}

	/*
		body, err := myrpc.call(req)
		head := EventResponse{}
		err = json.Unmarshal(body, &head)
		fmt.Print("========================== Parents - ")
		fmt.Println(head.Result.EventHeader.Parents)

		//err = Node1.Call(&ev, "debug_getEv", 1)
		//err = Node1.Call(&ev, "debug_getEvent", EvHash, true)

		   	fmt.Println(ev.String())
		   	fmt.Print("Creator - ")
		   	fmt.Println(ev.Creator.String())
		   	fmt.Print("Parents - ")
		   	fmt.Println(ev.Parents)

		   	fmt.Println("=== Ask for first parent - head of our DAG ===") //===============================
		       EvHash = ev.Parents[0].String()

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
