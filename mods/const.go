package mods

type MyWallet struct {
	PrivateKey string `json:"privateKey"`
	Address    string `json:"address"`
	HexAddress string `json:"hexAddress"`
}

type NetworkStruct struct {
	UrlNetScanner   string `json:"urlNetScanner"`
	UrlNodeGrpc     string `json:"urlNodeGrpc"`
	UrlSolidityGrpc string `json:"urlSolidityGrpc"`
}

var Network = NetworkStruct{
	// MainNet scanner
	UrlNetScanner: "https://tronscan.org/",
	// TestNet scanner
	// 	urlNetScanner = "https://shasta.tronscan.org/"

	// Public MainNet nodes. More: https://github.com/tronprotocol/Documentation/blob/master/TRX/Official_Public_Node.md
	UrlNodeGrpc:     "35.180.51.163:50051",
	UrlSolidityGrpc: "35.180.51.163:50061",
	// Public TestNet nodes. More: https://developers.tron.network/docs/networks#testnet
	//	urlNodeGrpc = "grpc.shasta.trongrid.io:50051"
	//	urlSolidityGrpc = "grpc.shasta.trongrid.io:50061"
}

const (
	// Trc20BalanceOf Trc20TransferMethodSignature SunSwapETHForExactTokens : Function signs
	Trc20BalanceOf               = "70a08231"
	Trc20TransferMethodSignature = "a9059cbb"
	SunSwapETHForExactTokens     = "723e6d2d"
	// FeeLimit : Transfer params
	FeeLimit = 100000000
	// TokenUSDTMainNet TokenUSDTTestNet : Token addresses
	TokenUSDTMainNet = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	TokenUSDTTestNet = "TB5NSkyzxkzi3eHW87NwFE6TmtTmnZw61y"
	// SunSwapPairContractUSDT : Trade pair contracts
	SunSwapPairContractUSDT = "TQn9Y2khEsLJW1ChVWFMSMeRDow5KcbLSE"
)
