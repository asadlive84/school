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

func (s *Server) GetUnionListByUpazillaId(ctx context.Context, req *pb.GetUnionListByUpazillaIdRequest) (*pb.GetUnionListByUpazillaIdResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req}).Info("Service: GetUnionsByUpazillaID method")

	upazillaid, err := strconv.Atoi(req.UpazillaId)
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("value convert issue %+v", err)
		return &pb.GetUnionListByUpazillaIdResponse{
			UnionInfo: []*pb.UnionInfo{},
		}, nil
	}

	getUpazillaList, err := s.Q.GetUnionsByUpazillaID(int32(upazillaid))
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("get list an error in query GetUnionsByUpazillaID %+v", err)
		return &pb.GetUnionListByUpazillaIdResponse{
			UnionInfo: []*pb.UnionInfo{},
		}, nil
	}

	unionInfoList := make([]*pb.UnionInfo, 0, 1)

	for _, k := range getUpazillaList {
		unionInfoList = append(unionInfoList, &pb.UnionInfo{
			Id:   k.Id,
			Name: k.Name,
		})
	}

	logger.WithFields(log.Fields{"Name": "GetUpazillaListByDistrictId"}).Info("Successfully find GetUnionsByUpazillaID")

	return &pb.GetUnionListByUpazillaIdResponse{
		UnionInfo: unionInfoList,
	}, nil

}
