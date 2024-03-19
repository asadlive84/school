package services

import (
	"context"
	"os"
	"time"

	// "log"
	// "net/http"

	pb "github.com/asadlive84/school/pb/student"

	log "github.com/sirupsen/logrus"
)

func (s *Server) GetDistrictList(ctx context.Context, req *pb.GetDistrictListRequest) (*pb.GetDistrictListResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req}).Info("Service: GetDistrictList method")

	std, err := s.Q.GetDistrictList()
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("get list an error in query GetDistrictList %+v", err)
		return &pb.GetDistrictListResponse{
			District: nil,
		}, nil
	}

	districtList := make([]*pb.District, 0, 1)

	for _, k := range std {
		districtList = append(districtList, &pb.District{
			Id:   k.Id,
			Name: k.Name,
		})
	}

	logger.WithFields(log.Fields{"Name": "GetDistrictList"}).Info("Successfully find GetDistrictList")

	return &pb.GetDistrictListResponse{
		District: districtList,
	}, nil

}
