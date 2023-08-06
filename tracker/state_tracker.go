package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"time"

	"golang.org/x/exp/slices"

	pb "github.com/BelLeonid/lavaPB/osmosisTest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = "localhost:3333"
)

type BlockInfo struct {
	Height int64 `json:"height"`
	Hash string `json:"hash"`
}

type TestResult struct {
	TestResult []BlockInfo `json:"test_result"`
}

func createClient (address string) (client pb.ServiceClient) {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("Established connection to server: %v", address)
	client = pb.NewServiceClient(connection)
	log.Println("Established new client")
	return
}

func main()  {
	// Creating two clients - to my server and directly to endpoint
	client := createClient(address)
	var uniqBlockIds []string;
	testResult := TestResult{TestResult: []BlockInfo{}}
	// We count until we have met five different blocks during the runtime
	for len(uniqBlockIds) < 5 {
		latestBlock := GetLatestBlock(client)
		log.Println("GetLatestBlock Info")
		hash := latestBlock.GetBlockId().GetHash()
		// Here I use base64 encoding, since that is the format that's returned from direct grpcurl call
		encodedHash := base64.StdEncoding.EncodeToString([]byte(hash))
		log.Printf("Hash: %v", encodedHash)
		height := latestBlock.GetBlock().GetHeader().GetHeight()
		log.Printf("Height: %v", height)

		// if the hash is new - we found another block to add
		if !slices.Contains(uniqBlockIds, encodedHash) {
			uniqBlockIds = append(uniqBlockIds, encodedHash)
			blockData := BlockInfo{Height: height, Hash: encodedHash}
			testResult.TestResult = append(testResult.TestResult, blockData)
			file, _ := json.MarshalIndent(testResult, "", " ")
			_ = os.WriteFile("test.json", file, 0644)
		}
	}
}

func GetLatestBlock (client pb.ServiceClient) (result *pb.GetLatestBlockResponse){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	log.Println("Sending grpc")
	result, err := client.GetLatestBlock(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return
}
