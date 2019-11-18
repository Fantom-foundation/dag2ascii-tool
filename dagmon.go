package main

import (
	"context"
	"fmt"

	"github.com/Fantom-foundation/go-lachesis/ethapi"

	//"github.com/ethereum/go-ethereum/common/hexutil"

	//"github.com/ethereum/go-ethereum"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	//"github.com/ethereum/go-ethereum/accounts"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/state"
	//"github.com/ethereum/go-ethereum/core/vm"
	//"github.com/ethereum/go-ethereum/ethdb"
	//"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/ethclient"

	//"github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	//"github.com/ethereum/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	//"github.com/Fantom-foundation/go-lachesis/evm_core"
	//"github.com/Fantom-foundation/go-lachesis/hash"
	//"github.com/Fantom-foundation/go-lachesis/inter"
)

func main() {

	fmt.Println("\n\n\n\n====== DAG Monitor 10 ===============================================================")

	Node_Eth, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to mainnet.infura.io")
	}

	Node1, err := ethclient.Dial("http://18.189.195.64:4001")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to lachesis Node1")
	}

	Node2, err := ethclient.Dial("http://18.191.96.173:4002")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to lachesis Node2")
	}

	Node3, err := ethclient.Dial("http://18.222.120.223:4003")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to lachesis Node3")
	}

	Node4, err := ethclient.Dial("http://127.0.0.1:5050")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to lachesis Node4 (local)")
	}

	blk := new(types.Block)

	fmt.Println("=== Get Block by number ===")
	BNum := big.NewInt(200)

	//===========================================================
	blk, err = Node_Eth.BlockByNumber(context.Background(), BNum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Node Ethereum block %d:", BNum)
	fmt.Println(blk)
	//--------------
	blk, err = Node1.BlockByNumber(context.Background(), BNum)
	//    if err != nil { log.Fatal(err) }
	fmt.Printf("Node1 block %d:", BNum)
	fmt.Println(blk)
	//--------------
	blk, err = Node2.BlockByNumber(context.Background(), BNum)
	//    if err != nil { log.Fatal(err) }
	fmt.Printf("Node2 block %d:", BNum)
	fmt.Println(blk)
	//--------------
	blk, err = Node3.BlockByNumber(context.Background(), BNum)
	//    if err != nil { log.Fatal(err) }
	fmt.Printf("Node3 block %d:", BNum)
	fmt.Println(blk)
	//--------------
	blk, err = Node4.BlockByNumber(context.Background(), BNum)
	//    if err != nil { log.Fatal(err) }
	fmt.Printf("Node4 block %d:", BNum)
	fmt.Println(blk)
	//===========================================================

	//func (ec *Client) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {

	//func (ec *Client) getBlock(ctx context.Context, method string, ) (*types.Block, error) {
	//var raw json.RawMessage
	//err2 := Node1.c.CallContext( context.Background(), &raw, "debug_getHeads" )
	//if err != nil {  log.Fatal(err) } else {fmt.Print(raw)}

	// debug_getHeads - returns IDs of all the events with no descendants in current epoch
	// Parameters: none  Returns: `Array` - Array of event IDs
	// curl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"debug_getHeads","params":[],"id":1}' localhost:18545

	//fmt.Println("Extract Header")
	//h := blk.Header()
	//fmt.Println(h)

	// GetHeads returns IDs of all the events with no descendants in current epoch.
	//func (api ) GetHeads(ctx context.Context) ([]hexutil.Bytes, error) {
	//	return eventIDsToHex(api.b.GetHeads(ctx)), nil

	// ap := new(*PublicDebugAPI)

	ev := ethapi.GetHeads(context.Background())
	//fmt.Println(ev)

	// GetHeads returns IDs of all the events with no descendants in current epoch.
	//([]hexutil.Bytes, error)

	rpcl, _ := rpc.Dial("http://18.189.195.64:4001")
	//Node1Client := ethclient.NewClient(rpcl)

	ev = ethapi.Backend().GetHeads(context.Background())
	fmt.Println(ev)

	fmt.Println("=====================================================================")
}
