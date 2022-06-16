package mods

import (
	"log"
	"sync"

	"github.com/LukeEuler/trx-go/trx/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
https://developers.tron.network/docs/networks
*/

var (
	once           sync.Once
	Wallet         api.WalletClient
	WalletSolidity api.WalletSolidityClient
)

func InitClient(urlNode string, urlNodeSolidity string) {
	f := func() {
		conn, err := grpc.Dial(urlNode, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		Wallet = api.NewWalletClient(conn)

		conn, err = grpc.Dial(urlNodeSolidity, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal(err)
		}
		WalletSolidity = api.NewWalletSolidityClient(conn)
	}
	once.Do(f)
}
