package services

import (
	"context"
	"os"
	"time"

	// "log"
	"net/http"

	pb "github.com/asadlive84/school/pb/student"
	q "github.com/asadlive84/school/student-svc/pkg/query"

	log "github.com/sirupsen/logrus"
)

func (s *Server) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req.Student.Name}).Info("Service: Register method")

	stdId, err := s.Q.InsertStudent(q.Student{
		StdId:        req.Student.StdId,
		Name:         req.Student.Name,
		Email:        req.Student.Email,
		NameBn:       req.Student.NameBn,
		FathersName:  req.Student.FathersName,
		MothersName:  req.Student.MothersName,
		DOB:          req.Student.Dob.AsTime(),
		Gender:       req.Student.Gender,
		BloodGroup:   req.Student.BloodGroup,
		MobileNumber: req.Student.MobileNumber,
		Session:      req.Student.Session,
		ClassName:    req.Student.ClassName,
		ClassSection: req.Student.ClassName,
		Address:      req.Student.Address,
		Religion:     req.Student.Religion,
		IsActive:     true,
		// CreatedAt:    time.Time{},
		// UpdatedAt:    time.Time{},
		// CreatedBy:    "",
		// UpdatedBy:    "",
	})

	if err != nil {
		logger.WithFields(log.Fields{"Name": req.Student.Name}).Errorf("insert an error in query Student %+v", err)
		return &pb.InsertResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	logger.WithFields(log.Fields{"Name": req.Student.Name}).Info("Successfully created Student")

	return &pb.InsertResponse{
		Status: http.StatusCreated,
		StdId:  stdId,
	}, nil
}
