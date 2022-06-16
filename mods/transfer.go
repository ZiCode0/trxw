package mods

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/LukeEuler/trx-go/common"
	"github.com/LukeEuler/trx-go/trx"
	"github.com/LukeEuler/trx-go/trx/api"
	"github.com/LukeEuler/trx-go/trx/core"
	"google.golang.org/protobuf/proto"
	"log"
	"math/big"
	"strconv"
)

//const feeLimit = 100000000

func TransferTrc20(wallet MyWallet, toAddress string, assetAddress string, amount int) string {

	amountSt := strconv.Itoa(amount)
	// key == privKey
	InitClient(Network.UrlNodeGrpc, Network.UrlSolidityGrpc)

	amountS, ok := big.NewInt(0).SetString(amountSt, 10)
	if !ok {
		log.Fatal(fmt.Errorf("invalid amount: [%s]", amountS))
		// log.Entry.Errorf("invalid amount: [%s]", amountS)
	}

	fromBytes, err := common.DecodeCheck(wallet.Address)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	toBytes, err := common.DecodeCheck(toAddress)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	contractBytes, err := common.DecodeCheck(assetAddress)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	toHex := hex.EncodeToString(toBytes)

	ab := common.LeftPadBytes(amountS.Bytes(), 32)

	data := Trc20TransferMethodSignature + "0000000000000000000000000000000000000000000000000000000000000000"[len(toHex)-2:] + toHex[2:]
	data += hex.EncodeToString(ab)
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	ct := &core.TriggerSmartContract{
		OwnerAddress:    fromBytes,
		ContractAddress: contractBytes,
		Data:            dataBytes,
	}

	tx, err := Wallet.TriggerContract(context.Background(), ct)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	tx.Transaction.RawData.FeeLimit = FeeLimit
	rawData, err := proto.Marshal(tx.Transaction.GetRawData())
	if err != nil {
		//panic(err)
		log.Fatal(err)
	}

	h256h := sha256.New()
	h256h.Write(rawData)
	hash := h256h.Sum(nil)
	tx.Txid = hash
	return signAndBroadcast(tx, wallet.PrivateKey)
}

func signAndBroadcast(tx *api.TransactionExtention, key string) string {
	/*
		Sign and broadcast formed api.TransactionExtention object
		Returns tx hash
	*/
	if proto.Size(tx) == 0 {
		// log.Entry.Fatal("bad transaction")
		log.Fatal("bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		//log.Entry.Fatal(string(tx.GetResult().GetMessage()))
		log.Fatal(fmt.Errorf(string(tx.GetResult().GetMessage())))
	}

	signTx, err := trx.SignTransaction(tx.Transaction, key)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	result, err := Wallet.BroadcastTransaction(context.Background(), signTx)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	if !result.GetResult() {
		//log.Entry.Fatal(string(result.GetMessage()))
		log.Fatal(string(result.GetMessage()))
	}
	if result.GetCode() != api.Return_SUCCESS {
		log.Fatalf("%d %s", result.Code, string(result.GetMessage()))
		//panic(fmt.Errorf("%d %s", result.Code, string(result.GetMessage())))
	}

	// get transaction hash
	hash := hex.EncodeToString(tx.GetTxid())
	return hash
}
