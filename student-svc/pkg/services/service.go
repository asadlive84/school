package services

import (
	q "github.com/asadlive84/school/student-svc/pkg/query"
	pb "github.com/asadlive84/school/pb/student"
)

type Server struct {
	// H db.DbHandler
	Q q.Query
	pb.UnimplementedStudentServiceServer 
}
