package student

import (
	"fmt"

	"github.com/asadlive84/school/api-gateway/pkg/config"
	pb "github.com/asadlive84/school/pb/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.StudentServiceClient
}

func InitServiceClient(c *config.Config) pb.StudentServiceClient {
	cc, err := grpc.Dial(c.STUDENT_SVC_URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewStudentServiceClient(cc)

}
