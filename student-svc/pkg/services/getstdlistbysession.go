package services

import (
	"context"
	"fmt"
	"os"
	"time"

	// "log"
	// "net/http"

	pb "github.com/asadlive84/school/pb/student"
	"google.golang.org/protobuf/types/known/timestamppb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) GetStudentListBySession(ctx context.Context, req *pb.GetStudentListBySessionRequest) (*pb.GetStudentListBySessionResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req}).Info("Service: GetStudentListBySession method")

	fmt.Println("#####################req.SessionId#################################")
	fmt.Printf("req.SessionId111: %+v\n", req.SessionId)
	fmt.Println("#####################req.SessionId#################################")

	std, err := s.Q.GetStudentListBySession(req.SessionId)
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("get list an error in query Student %+v", err)
		return &pb.GetStudentListBySessionResponse{
			Student: nil,
		}, nil
	}

	stdb := make([]*pb.Student, 0, 1)

	for _, k := range std {
		stdb = append(stdb, &pb.Student{
			Id:           k.ID,
			StdId:        k.StdId,
			Name:         k.Name,
			NameBn:       k.NameBn,
			FathersName:  k.FathersName,
			MothersName:  k.MothersName,
			Dob:          timestamppb.New(k.DOB),
			Gender:       k.Gender,
			BloodGroup:   k.BloodGroup,
			MobileNumber: k.MobileNumber,
			Session:      k.Session,
			ClassName:    k.ClassName,
			ClassSection: k.ClassSection,
			ClassGroup:   "",
			ClassRoll:    "",
			Address:      k.Address,
			Religion:     "",
			// CreatedAt:    &timestamppb.Timestamp{},
			// UpdatedAt:    &timestamppb.Timestamp{},
			// CreatedBy:    "",
			// UpdatedBy:    "",
			// Email:        "",
		})
	}

	logger.WithFields(log.Fields{"Name": "req.Student.Name"}).Info("Successfully find Student")

	return &pb.GetStudentListBySessionResponse{
		Student: stdb,
	}, nil

}
