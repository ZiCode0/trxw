package main

import (
	"bufio"
	"fmt"
	"github.com/LukeEuler/trx-go/key"
	"log"
	_ "log"
	"os"
	"trxw/mods"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func newWallet() mods.MyWallet {
	/*
		Generate Tron wallet using local code
	*/
	k, err := key.NewKey()
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}
	hexAddr := hexutil.Encode([]byte(k.Address()))
	wallet := mods.MyWallet{PrivateKey: k.PrivateKey(), Address: k.Address(), HexAddress: hexAddr}
	fmt.Printf("Address: %s, Private Key: %s\n", k.Address(), k.PrivateKey())
	return wallet
}

func loadWallet(keyString string) mods.MyWallet {
	/*
		Load Tron wallet locally using string>bytes
	*/
	k, err := key.NewKeyFromHex(keyString)
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}
	hexAddr := hexutil.Encode([]byte(k.Address()))
	wallet := mods.MyWallet{PrivateKey: k.PrivateKey(), Address: k.Address(), HexAddress: hexAddr}
	fmt.Printf("[+] Address: %s loaded..\n", k.Address())
	return wallet
}

func main() {
	// generate new wallet
	//wallet := newWallet()

	// load TestNet wallet
	//wallet := loadWallet(os.Getenv("KEY_TEST"))

	// load MainNet wallet
	wallet := loadWallet(os.Getenv("KEY_MAIN"))

	// wait input trx
	fmt.Print("  Waiting input of trc20 tokens. Press <Enter> to continue..")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	// TODO: add swap tokens
	//mods.SwapToken(sunSwapPairContractUSDT, 1000000, urlNodeGrpc)

	receiveAddress := os.Getenv("REC_ADDRESS")
	// transfer trc20 tokens
	txHash := mods.TransferTrc20(wallet, receiveAddress, mods.TokenUSDTMainNet, 1000000)
	fmt.Println("tx hash: " + txHash)
	fmt.Printf("url: %s/#/transaction/%s\n", mods.Network.UrlNetScanner, txHash)
}
