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

func (s *Server) CheckStudent(ctx context.Context, req *pb.CheckStudentRequest) (*pb.CheckStudentResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"UserID": req.UserID}).Info("Service: checking user method with UserID")

	if req.UserID == "" {
		logger.WithFields(log.Fields{"UserID": req.UserID}).Errorln("user id is empty")
		return &pb.CheckStudentResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	user, err := s.Q.GetUserByID(req.UserID)
	if err != nil && err != q.NotFound {
		logger.WithFields(log.Fields{"UserID": req.UserID}).Errorf("an error in get user query %+v", err)
		return &pb.CheckStudentResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	if (user == q.User{}) {
		logger.WithFields(log.Fields{"UserID": req.UserID}).Errorln("user is empty")
		return &pb.CheckStudentResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	res := &pb.CheckStudentResponse{
		Status:  http.StatusCreated,
		Message: "success",
	}

	logger.WithFields(log.Fields{"UserID": req.UserID}).Info("login successfully")
	return res, nil
}
