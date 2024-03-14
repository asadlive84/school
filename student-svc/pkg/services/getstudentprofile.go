package services

import (
	"context"
	"os"
	"time"

	pb "github.com/asadlive84/school/pb/student"
	"google.golang.org/protobuf/types/known/timestamppb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) GetStudentProfileById(ctx context.Context, req *pb.GetStudentProfileByIdRequest) (*pb.GetStudentProfileByIdResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"Name": req}).Info("Service: GetStudentProfileById method")

	student, err := s.Q.GetStudentProfileById(req.StudentId)
	if err != nil {
		logger.WithFields(log.Fields{"Name": err}).Errorf("get list an error in query Student %+v", err)
		return &pb.GetStudentProfileByIdResponse{
			Student: nil,
		}, nil
	}

	logger.WithFields(log.Fields{"Name": "req.Student.Name"}).Info("Successfully find Student")

	return &pb.GetStudentProfileByIdResponse{
		Student: &pb.Student{
			Id:            student.ID,
			StdId:         student.StdId,
			Name:          student.Name,
			NameBn:        student.NameBn,
			FathersName:   student.FathersName,
			MothersName:   student.MothersName,
			Dob:           timestamppb.New(student.DOB),
			Gender:        student.Gender,
			BloodGroup:    student.BloodGroup,
			MobileNumber:  student.MobileNumber,
			Session:       student.Session,
			ClassName:     student.ClassName,
			ClassSection:  student.ClassSection,
			ClassGroup:    "",
			ClassRoll:     "",
			Address:       "",
			Religion:      student.Religion,
			CreatedAt:     &timestamppb.Timestamp{},
			UpdatedAt:     &timestamppb.Timestamp{},
			VillageOrRoad: student.VillageOrRoad,
			UnionName:     student.Union,
			Upzilla:       student.Upzilla,
			District:      student.District,
		},
	}, nil

}
