package main

import (
	"flag"
	"fmt"
	"goethdevkit/internal/gethwrap"
	"goethdevkit/internal/infurawrap"
	"log"
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

	mainNetURL := infurawrap.InfuraMainNetURL
	result, err := gethwrap.InfuraQueryTxnByHash(mainNetURL, projectID, txnHash, 1)
	if err != nil {
		log.Fatalf("Query transaction failed. Reason: %s", err.Error())
	}

	fmt.Println("ChainID: ", result.ChainID)
	fmt.Println("Nounce:  ", result.Nounce)
	fmt.Println("Gas limit", result.GasLimit)
	fmt.Println("Value", result.Value)
	fmt.Println("Payload", string(result.Payload))
}
