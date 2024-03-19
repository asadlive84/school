package services

import (
	"context"
	"fmt"
	"os"

	// "strconv"
	"time"

	// "log"
	// "net/http"

	pb "github.com/asadlive84/school/pb/student"
	"github.com/asadlive84/school/student-svc/pkg/query"
	"google.golang.org/protobuf/types/known/timestamppb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) SearchStudent(ctx context.Context, req *pb.SearchStudentRequest) (*pb.SearchStudentResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req}).Info("Service: SearchStudent method")

	fmt.Println("========================================")
	fmt.Printf("DistrictId %+v\n", req.DistrictId)
	fmt.Printf("UpazillaId %+v\n", req.UpazillaId)
	fmt.Printf("UnonId %+v\n", req.UnonId)
	fmt.Printf("VillageOrRoadId %+v\n", req.VillageOrRoadId)
	fmt.Println("========================================")

	// upazillaid, err := strconv.Atoi(req.UpazillaId)
	// if err != nil {
	// 	logger.WithFields(log.Fields{"Name": err}).Errorf("value convert issue %+v", err)
	// 	return &pb.SearchStudentResponse{
	// 		Student: []*pb.Student{},
	// 	}, nil
	// }

	studList, err := s.Q.SearchStudent(query.SearchData{
		District:     req.DistrictId,
		Upazillas:    req.UpazillaId,
		Unions:       req.UnonId,
		VillageRoads: req.VillageOrRoadId,
		SessionId:    req.SessionId,
	})
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("get list an error in query SearchStudent %+v", err)
		return &pb.SearchStudentResponse{
			Student: []*pb.Student{},
		}, nil
	}

	stdb := make([]*pb.Student, 0, 1)

	for _, k := range studList {
		stdb = append(stdb, &pb.Student{
			Id:            k.ID,
			StdId:         k.StdId,
			Name:          k.Name,
			NameBn:        k.NameBn,
			FathersName:   k.FathersName,
			MothersName:   k.MothersName,
			Dob:           timestamppb.New(k.DOB),
			Gender:        k.Gender,
			BloodGroup:    k.BloodGroup,
			MobileNumber:  k.MobileNumber,
			Session:       k.Session,
			ClassName:     k.ClassName,
			ClassSection:  k.ClassSection,
			ClassGroup:    "",
			ClassRoll:     "",
			Address:       k.Address,
			Religion:      "",
			CreatedAt:     &timestamppb.Timestamp{},
			UpdatedAt:     &timestamppb.Timestamp{},
			CreatedBy:     "",
			UpdatedBy:     "",
			Email:         "",
			District:      k.District,
			Upzilla:       k.Upzilla,
			UnionName:     k.Union,
			VillageOrRoad: k.VillageOrRoad,
		})
	}

	logger.WithFields(log.Fields{"Name": "SearchStudent"}).Info("Successfully find SearchStudent")

	return &pb.SearchStudentResponse{
		Student: stdb,
	}, nil

}
