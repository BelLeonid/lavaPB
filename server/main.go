package main

import (
	"context"
	"log"
	"net"
	"time"

	pbOsmosis "github.com/BelLeonid/lavaPB/osmosisTest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

var (
	osmosisAddress = "grpc.osmosis.zone:9090"
	osmosisClient = CreateClient(osmosisAddress)
)

func CreateClient(address string) (client pbOsmosis.ServiceClient)  {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("Established connection")
	client = pbOsmosis.NewServiceClient(conn)
	log.Println("Established new client")
	return
}

type myGRPCServer struct {
	pbOsmosis.UnimplementedServiceServer
}

func (m *myGRPCServer) GetLatestBlock(ctx context.Context, request *pbOsmosis.GetRequest) (*pbOsmosis.GetLatestBlockResponse, error) {
	// Cancel request if we have no response for 15 seconds - I tested, it should be enough for server to send all the data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	osmosisResponse, err := osmosisClient.GetLatestBlock(ctx, request) // works
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return osmosisResponse, err
}

func (m *myGRPCServer) GetLatestValidatorSet(ctx context.Context, request *pbOsmosis.GetRequest) (*pbOsmosis.GetLatestValidatorSetResponse, error) {
	// Cancel request if we have no response for 15 seconds - I tested, it should be enough for server to send all the data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	osmosisResponse, err := osmosisClient.GetLatestValidatorSet(ctx, request) // works
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return osmosisResponse, err
}

func (m *myGRPCServer) GetNodeInfo(ctx context.Context, request *pbOsmosis.GetRequest) (*pbOsmosis.GetNodeInfoResponse, error) {
	// Cancel request if we have no response for 15 seconds - I tested, it should be enough for server to send all the data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	osmosisResponse, err := osmosisClient.GetNodeInfo(ctx, request) // works
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return osmosisResponse, err
}

func (m *myGRPCServer) GetSyncing(ctx context.Context, request *pbOsmosis.GetRequest) (*pbOsmosis.GetSyncingResponse, error) {
	// Cancel request if we have no response for 15 seconds - I tested, it should be enough for server to send all the data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	osmosisResponse, err := osmosisClient.GetSyncing(ctx, request) // works
	if err != nil {
		log.Printf("Error: %v", err)
	}
	return osmosisResponse, err
}

func (m *myGRPCServer) GetBlockByHeight(ctx context.Context, request *pbOsmosis.GetHeightRequest) (*pbOsmosis.GetBlockByHeightResponse, error) {
	// Cancel request if we have no response for 15 seconds - I tested, it should be enough for server to send all the data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	osmosisResponse, err := osmosisClient.GetBlockByHeight(ctx, request) // works
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return osmosisResponse, err
}

func (m *myGRPCServer) GetValidatorSetByHeight(ctx context.Context, request *pbOsmosis.GetHeightRequest) (*pbOsmosis.GetValidatorSetByHeightResponse, error) {
	// Cancel request if we have no response for 15 seconds - I tested, it should be enough for server to send all the data
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15 )
	defer cancel()
	osmosisResponse, err := osmosisClient.GetValidatorSetByHeight(ctx, request) // works
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return osmosisResponse, err
}

func main() {
	lis, err := net.Listen("tcp", ":3333")
	if err != nil {
	   log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	// I added reflection so it's possible to get a list of the available services
	reflection.Register(server)
	myServiceServer := &myGRPCServer{}
	pbOsmosis.RegisterServiceServer(server, myServiceServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
	   log.Fatalf("failed to serve: %v", err)
	}
 }