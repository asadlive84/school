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
	})

	if err != nil {
		logger.WithFields(log.Fields{"Name": req.Student.Name}).Errorf("insert an error in query Student %+v", err)
		return &pb.InsertResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	logger.WithFields(log.Fields{"Name": req.Student.Name}).Info("Successfully created Student")

	// req.GetDistrict()

	district, err := s.Q.GetDistrictByName(req.Student.GetDistrict())

	if err != q.NotFound && err != nil {
		logger.WithFields(log.Fields{"GetDistrict": req.Student.GetDistrict()}).Errorf("insert an error in query GetDistrict %+v", err)
		return &pb.InsertResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	if err == q.NotFound {
		districtId, err := s.Q.InsertDistrict(&q.District{
			Name: req.Student.GetDistrict(),
		})
		if err != nil {
			logger.WithFields(log.Fields{"District": req.Student.GetDistrict()}).Errorf("insert an error in query InsertDistrict %+v", err)
			return &pb.InsertResponse{
				Status: http.StatusBadRequest,
			}, nil
		}
		district.Id = districtId
	}

	upzillaList, err := s.Q.GetUpazillasByDistrictID(district.Id)
	if err != q.NotFound && err != nil {
		logger.WithFields(log.Fields{"Upzailla": req.Student.Upzilla}).Errorf("insert an error in query GetUpazillasByDistrictID %+v", err)
		return &pb.InsertResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	var upZillaId int32

	for _, k := range upzillaList {
		if k.Name == req.Student.Upzilla {
			upZillaId = k.Id
		}
	}

	if upzillaList == nil || upZillaId == 0 {
		upZillaId, err = s.Q.InsertUpazilla(&q.Upazilla{
			Name:       req.Student.Upzilla,
			DistrictId: district.Id,
		})
		if err != nil {
			logger.WithFields(log.Fields{"Name": req.Student.Name}).Errorf("insert an error in query InsertUpazilla %+v", err)
			return &pb.InsertResponse{
				Status: http.StatusBadRequest,
			}, nil
		}
	}

	unionList, err := s.Q.GetUnionsByUpazillaID(upZillaId)

	if err != q.NotFound && err == q.NotFound {
		logger.WithFields(log.Fields{"Name": req.Student.Name}).Errorf("insert an error in query GetUnionsByUpazillaID %+v", err)
		return &pb.InsertResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	var unionId int32

	for _, k := range unionList {
		if k.Name == req.Student.UnionName {
			unionId = k.Id
		}
	}

	if unionList == nil || unionId == 0 {
		unionId, err = s.Q.InsertUnion(&q.UnionInfo{
			Name:       req.Student.UnionName,
			UpazillaId: upZillaId,
		})
		if err != nil {
			logger.WithFields(log.Fields{"Name": req.Student.Name}).Errorf("insert an error in query InsertUnion %+v", err)
			return &pb.InsertResponse{
				Status: http.StatusBadRequest,
			}, nil
		}
	}

	villageList, err := s.Q.GetVillagesOrRoadsByUnionID(unionId)
	if err != q.NotFound && err != nil {
		logger.WithFields(log.Fields{"Name": req.Student.Name}).Errorf("insert an error in query GetUnionsByUpazillaID %+v", err)
		return &pb.InsertResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	var villageId int32

	for _, k := range villageList {
		if k.Name == req.Student.VillageOrRoad {
			villageId = k.Id
		}
	}

	if villageList == nil || villageId == 0 {
		villageId, err = s.Q.InsertVillageOrRoad(&q.VillageOrRoad{
			Name:    req.Student.GetVillageOrRoad(),
			UnionId: unionId,
		})
		if err != nil {
			logger.WithFields(log.Fields{"Name": req.Student.Name}).Errorf("insert an error in query InsertVillageOrRoad %+v", err)
			return &pb.InsertResponse{
				Status: http.StatusBadRequest,
			}, nil
		}
	}

	_, err = s.Q.InsertAddress(q.Address{
		StudentId:       stdId,
		VillageOrRoadId: villageId,
		UnionID:         unionId,
		UpazillaId:      upZillaId,
		DistrictId:      district.Id,
		IsPresent:       false,
		IsPermanent:     false,
	})
	if err != nil {
		logger.WithFields(log.Fields{"InsertAddress": req.Student.Name}).Errorf("insert an error in query InsertAddress %+v", err)
		return &pb.InsertResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	return &pb.InsertResponse{
		Status: http.StatusCreated,
		StdId:  stdId,
	}, nil
}
