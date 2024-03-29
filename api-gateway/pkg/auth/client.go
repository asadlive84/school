package auth

import (
	"fmt"

	// "github.com/asadlive84/school/api-gateway/pkg/auth/pb"
	pb "github.com/asadlive84/school/pb/auth"
	"github.com/asadlive84/school/api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)

}
