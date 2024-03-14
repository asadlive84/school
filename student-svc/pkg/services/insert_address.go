package services

import (
	"context"
	"os"
	"time"

	// "log"

	pb "github.com/asadlive84/school/pb/student"
	q "github.com/asadlive84/school/student-svc/pkg/query"

	log "github.com/sirupsen/logrus"
)

func (s *Server) InsertAddress(ctx context.Context, req *pb.InsertAddressRequest) (*pb.InsertAddressResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req.Address.StudentId}).Info("Service: InsertAddress method")

	// s.Q.GetDistrictByName()

	_, err := s.Q.InsertAddress(q.Address{
		Info:            req.Address.Info,
		StudentId:       req.Address.StudentId,
		VillageOrRoadId: req.Address.VillageOrRoadId,
		UnionID:         req.Address.UnionId,
		UpazillaId:      req.Address.UpazillaId,
		DistrictId:      req.Address.DistrictId,
		IsPresent:       req.Address.IsPresent,
		IsPermanent:     req.Address.IsPermanent,
	})

	if err != nil {
		logger.WithFields(log.Fields{"Name": req.Address.StudentId}).Errorf("insert an error in query Address %+v", err)
		return &pb.InsertAddressResponse{
			// Status: http.StatusBadRequest,
		}, nil
	}

	logger.WithFields(log.Fields{"Name": req.Address.StudentId}).Info("Successfully created Address")

	return &pb.InsertAddressResponse{
		// Status: http.StatusCreated,
		// StdId:  stdId,
	}, nil
}
