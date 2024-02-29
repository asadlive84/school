package services

import (
	"context"
	"os"
	"time"

	// "log"
	// "net/http"

	pb "github.com/asadlive84/school/pb/student"
	"google.golang.org/protobuf/types/known/timestamppb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) StudentList(ctx context.Context, req *pb.StudentListRequest) (*pb.StudentListResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req}).Info("Service: Register method")

	std, err := s.Q.StudentList()

	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("list an error in query Student %+v", err)
		return &pb.StudentListResponse{
			Student: nil,
		}, nil
	}

	stdb := make([]*pb.Student, 0, 1)

	for _, k := range std {
		stdb = append(stdb, &pb.Student{
			Id:           k.ID,
			Name:         k.Name,
			NameBn:       k.NameBn,
			FathersName:  k.FathersName,
			MothersName:  k.MothersName,
			Session:      k.Session,
			Gender:       k.Gender,
			MobileNumber: k.MobileNumber,
			ClassName:    k.ClassName,
			StdId:        k.StdId,
			Dob:          timestamppb.New(k.DOB),
		})
	}

	logger.WithFields(log.Fields{"Name": "req.Student.Name"}).Info("Successfully find Student")

	return &pb.StudentListResponse{
		Student: stdb,
	}, nil

}
