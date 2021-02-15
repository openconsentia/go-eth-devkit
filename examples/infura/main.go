package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func parseFlag() map[string]string {

	flagValues := make(map[string]string)

	projectID := flag.String("id", "", "project ID")
	txnHash := flag.String("txn", "", "txn hash")

	flag.Parse()

	if *projectID != "" {
		flagValues["project-id"] = *projectID
	}

	if *txnHash != "" {
		flagValues["txn-hash"] = *txnHash
	}

	return flagValues
}

func queryTransaction(url string, id string, txnHash string) error {
	nettURL := fmt.Sprintf("%s/%s", url, id)
	conn, err := ethclient.Dial(nettURL)
	if err != nil {
		return err
	}

	ctx := context.Background()
	tx, pending, err := conn.TransactionByHash(ctx, common.HexToHash(txnHash))
	if err != nil {
		return err
	}

	if !pending {
		fmt.Printf("Data: %v\n", string(tx.Data()))
		fmt.Printf("Gas: %v", tx.Gas())
	}

	return nil
}

func queryMainNet(id string, txnHash string) error {
	return queryTransaction("https://mainnet.infura.io/v3", id, txnHash)
}

func main() {

	flagValues := parseFlag()

	projectID, ok := flagValues["project-id"]
	log.Println(projectID)
	if !ok {
		log.Fatal("Project ID missing")
	}

	txnHash, ok := flagValues["txn-hash"]
	log.Println(txnHash)
	if !ok {
		log.Fatal("Txn hash missing")
	}

	err := queryMainNet(projectID, txnHash)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
