package services

import (
	"context"
	"os"
	"time"

	// "log"
	"net/http"

	"github.com/asadlive84/school/auth-svc/pkg/pb"
	q "github.com/asadlive84/school/auth-svc/pkg/query"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"email": req.Email}).Info("Service: Register method")

	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

	if err != nil {
		logger.WithFields(log.Fields{"email": req.Email}).Errorf("an error in hashed password %+v", err)
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "account not created",
		}, nil
	}

	user, err := s.Q.GetUserByEmail(req.Email)
	if err != nil && err != q.NotFound {
		logger.WithFields(log.Fields{"email": req.Email}).Errorf("an error in get user query %+v", err)
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	if user != nil {
		logger.WithFields(log.Fields{"email": req.Email}).Info("email is exists")
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	if req.Phone != "" {
		user, err = s.Q.GetUserByPhone(req.Phone)
		if err != nil && err != q.NotFound {
			logger.WithFields(log.Fields{"phone": req.Email}).Errorf("an error in get user query %+v", err)
			return &pb.RegisterResponse{
				Status: http.StatusBadRequest,
			}, nil
		}

		if user != nil {
			logger.WithFields(log.Fields{"phone": req.Email}).Info("phone is exists")
			return &pb.RegisterResponse{
				Status: http.StatusBadRequest,
			}, nil
		}
	}

	err = s.Q.InsertUser(q.User{
		Email:    req.Email,
		Password: string(hashedPasswd),
		Phone:    req.Phone,
	})

	if err != nil {
		logger.WithFields(log.Fields{"email": req.Email}).Errorf("an error in query %+v", err)
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	logger.WithFields(log.Fields{"email": req.Email}).Info("Successfully created user")

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}
