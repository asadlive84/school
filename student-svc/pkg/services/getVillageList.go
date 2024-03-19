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

func (s *Server) GetVillageOrRoadListByUnionId(ctx context.Context, req *pb.GetVillageOrRoadListByUnionIdRequest) (*pb.GetVillageOrRoadListByUnionIdResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req}).Info("Service: GetVillagesOrRoadsByUnionID method")

	unionId, err := strconv.Atoi(req.UiononId)
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("value convert issue %+v", err)
		return &pb.GetVillageOrRoadListByUnionIdResponse{
			VillageOrRoad: []*pb.VillageOrRoad{},
		}, nil
	}

	getUpazillaList, err := s.Q.GetVillagesOrRoadsByUnionID(int32(unionId))
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("get list an error in query GetVillageOrRoadListByUnionId %+v", err)
		return &pb.GetVillageOrRoadListByUnionIdResponse{
			VillageOrRoad: []*pb.VillageOrRoad{},
		}, nil
	}

	villageOrRoadList := make([]*pb.VillageOrRoad, 0, 1)

	for _, k := range getUpazillaList {
		villageOrRoadList = append(villageOrRoadList, &pb.VillageOrRoad{
			Id:   k.Id,
			Name: k.Name,
		})
	}

	logger.WithFields(log.Fields{"Name": "GetVillagesOrRoadsByUnionID"}).Info("Successfully find GetVillageOrRoadListByUnionId")

	return &pb.GetVillageOrRoadListByUnionIdResponse{
		VillageOrRoad: villageOrRoadList,
	}, nil

}
