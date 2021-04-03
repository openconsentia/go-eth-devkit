package geth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	InfuraMainNetURL = "https://mainnet.infura.io/v3"
)

func GetConnection(infuraMainNetURL string, infuraPID string) (*ethclient.Client, error) {

	netURL := fmt.Sprintf("%s/%s", infuraMainNetURL, infuraPID)
	conn, err := ethclient.Dial(netURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
