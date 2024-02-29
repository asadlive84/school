package services

import (
	"context"
	"os"
	"time"

	// "net/http"

	pb "github.com/asadlive84/school/pb/student"

	log "github.com/sirupsen/logrus"
)

func (s *Server) GetClassList(ctx context.Context, req *pb.GetClassListRequest) (*pb.GetClassListResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"GetClassList": "req.UserID"}).Info("Service: checking user method with GetClassList")

	classList, err := s.Q.GetClassList()
	if err != nil {
		return nil, err
	}

	cls := make([]*pb.SessionClass, 0, 1)

	for _, k := range classList {
		cls = append(cls, &pb.SessionClass{
			SessionId: k.SessionId,
			Class:     k.Class,
			Year:      k.Year,
		})
	}

	res := &pb.GetClassListResponse{
		Session: cls,
	}

	logger.WithFields(log.Fields{"GetClassList": ""}).Info("GetClassList successfully")
	return res, nil
}
