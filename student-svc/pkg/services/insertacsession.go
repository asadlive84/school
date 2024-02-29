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

func (s *Server) InsertAcademicSession(ctx context.Context, req *pb.InsertAcademicSessionRequest) (*pb.InsertAcademicSessionResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": ""}).Info("Service: InsertAcademicSession method")

	err := s.Q.InsertAcademicSession(q.AcademicSession{
		StudentId: req.StudentId,
		SessionId: req.SessionId,
	})

	if err != nil {
		logger.WithFields(log.Fields{"Name": ""}).Errorf("an error in query InsertAcademicSession %+v", err)
		return nil, err
	}

	logger.WithFields(log.Fields{"Name": ""}).Info("Successfully created InsertAcademicSession")

	return &pb.InsertAcademicSessionResponse{}, nil
}
