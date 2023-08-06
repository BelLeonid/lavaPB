package main

import (
	"context"
	"log"
	"reflect"
	"time"

	pb "github.com/BelLeonid/lavaPB/osmosisTest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addrTest = "localhost:3333"
	addrCheck = "grpc.osmosis.zone:9090"
)

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
	clientTest := createClient(addrTest)
	clientCheck := createClient(addrCheck)

	testLatestBlock := GetLatestBlock(clientTest)
	checkLatestBlock := GetLatestBlock(clientCheck)
	log.Println("Comparing GetLatestBlock")
	log.Println(reflect.DeepEqual(testLatestBlock, checkLatestBlock))

	testSync := GetSyncing(clientTest)
	checkSynck := GetSyncing(clientCheck)
	log.Println("Comparing GetSyncing")
	log.Println(reflect.DeepEqual(testSync, checkSynck))
}

func GetLatestBlock (client pb.ServiceClient) (result *pb.GetLatestBlockResponse){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	log.Println("Sending grpc")
	result, err := client.GetLatestBlock(ctx, &pb.GetRequest{}) // works
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Printing result: %s", result)
	return
}

func GetLatestValidatorSet (client pb.ServiceClient) (result *pb.GetLatestValidatorSetResponse){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	log.Println("Sending grpc")
	result, err := client.GetLatestValidatorSet(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Printing result: %s", result)
	return
}

func GetSyncing (client pb.ServiceClient) (result *pb.GetSyncingResponse){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	log.Println("Sending grpc")
	result, err := client.GetSyncing(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Printing result: %s", result)
	return
}

func GetNodeInfo (client pb.ServiceClient) (result *pb.GetNodeInfoResponse){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	log.Println("Sending grpc")
	result, err := client.GetNodeInfo(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Printing result: %s", result)
	return
}

func GetValidatorSetByHeight (client pb.ServiceClient, height int64) (result *pb.GetValidatorSetByHeightResponse){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	log.Println("Sending grpc")
	result, err := client.GetValidatorSetByHeight(ctx, &pb.GetHeightRequest{Height: height})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Printing result: %s", result)
	return
}

func GetBlockByHeight (client pb.ServiceClient, height int64) (result *pb.GetBlockByHeightResponse){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	log.Println("Sending grpc")
	result, err := client.GetBlockByHeight(ctx, &pb.GetHeightRequest{Height: height})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Printing result: %s", result)
	return
}