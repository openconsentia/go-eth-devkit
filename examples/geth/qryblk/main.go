package main

import (
	"context"
	"goethdevkit/examples/geth"
	"log"
	"math/big"
	"os"
)

func main() {
	projectID, ok := os.LookupEnv("PID")
	if !ok {
		log.Fatal("Project ID enviroment variable not set")
	}

	client, err := geth.GetConnection(geth.InfuraMainNetURL, projectID)
	if err != nil {
		log.Fatalf("Failed to connect. Reason: %v", err)
	}

	blkNum, err := client.BlockNumber(context.TODO())
	if err != nil {
		log.Fatalf("Unable to get last block number. Reason: %v", err)
	}

	log.Printf("Last block number: %v", blkNum)
	header, err := client.HeaderByNumber(context.TODO(), big.NewInt(int64(blkNum)))
	if err != nil {
		log.Fatalf("Failed to get header. Reason: %v", err)
	}

	log.Printf("Last block number from header: %v", header.Number)

	blk, err := client.BlockByNumber(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to get block. Reason: %v", err)
	}

	log.Printf("Number of transactions %v", len(blk.Transactions()))
}
