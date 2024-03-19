package services

import (
	"context"
	"os"
	"strconv"
	"time"

	// "log"
	// "net/http"

	pb "github.com/asadlive84/school/pb/student"

	log "github.com/sirupsen/logrus"
)

func (s *Server) GetUpazillaListByDistrictId(ctx context.Context, req *pb.GetUpazillaListByDistrictIdRequest) (*pb.GetUpazillaListByDistrictIdResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req}).Info("Service: GetUpazillaListByDistrictId method")

	dictrictid, err := strconv.Atoi(req.DictrictId)
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("value convert issue %+v", err)
		return &pb.GetUpazillaListByDistrictIdResponse{
			Upazilla: []*pb.Upazilla{},
		}, nil
	}

	getUpazillaList, err := s.Q.GetUpazillasByDistrictID(int32(dictrictid))
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("get list an error in query GetUpazillaListByDistrictId %+v", err)
		return &pb.GetUpazillaListByDistrictIdResponse{
			Upazilla: []*pb.Upazilla{},
		}, nil
	}

	upazillaList := make([]*pb.Upazilla, 0, 1)

	for _, k := range getUpazillaList {
		upazillaList = append(upazillaList, &pb.Upazilla{
			Id:   k.Id,
			Name: k.Name,
		})
	}

	logger.WithFields(log.Fields{"Name": "GetUpazillaListByDistrictId"}).Info("Successfully find GetUpazillaListByDistrictId")

	return &pb.GetUpazillaListByDistrictIdResponse{
		Upazilla: upazillaList,
	}, nil

}
