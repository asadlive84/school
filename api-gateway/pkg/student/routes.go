package student

import (
	"github.com/asadlive84/school/api-gateway/pkg/config"
	"github.com/asadlive84/school/api-gateway/pkg/student/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/student")
	routes.POST("/insert", svc.Insert)
	routes.POST("/check", svc.Login)
	routes.POST("/uploadfile", svc.UploadFile)
	routes.GET("/list", svc.StudentList)
	routes.GET("/clslist", svc.GetClassList)
	routes.GET("/profile/:id", svc.GetStudentProfileById)
	routes.POST("/stdlistbysessionid", svc.GetStudentListBySession)

	routes.GET("/get-district-list", svc.GetDistrict)
	routes.GET("/get-upazilla-list/:id", svc.GetUpazillaListByDistrictId)
	routes.GET("/get-union-list/:id", svc.GetUnionListByUpazillaId)
	routes.GET("/get-village-or-road-list/:id", svc.GetVillageOrRoadListByUnionId)

	routes.POST("/search", svc.SearchStudent)

	return svc

}

func (svc *ServiceClient) Insert(ctx *gin.Context) {
	routes.Insert(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}

func (svc *ServiceClient) UploadFile(ctx *gin.Context) {
	routes.UploadFile(ctx, svc.Client)
}
func (svc *ServiceClient) StudentList(ctx *gin.Context) {
	routes.StudentList(ctx, svc.Client)
}
func (svc *ServiceClient) GetClassList(ctx *gin.Context) {
	routes.GetClassList(ctx, svc.Client)
}
func (svc *ServiceClient) GetStudentListBySession(ctx *gin.Context) {
	routes.GetStudentListBySession(ctx, svc.Client)
}

func (svc *ServiceClient) GetStudentProfileById(ctx *gin.Context) {
	routes.GetStudentProfileById(ctx, svc.Client)
}
func (svc *ServiceClient) GetDistrict(ctx *gin.Context) {
	routes.GetDistrict(ctx, svc.Client)
}
func (svc *ServiceClient) GetUpazillaListByDistrictId(ctx *gin.Context) {
	routes.GetUpazillaListByDistrictId(ctx, svc.Client)
}
func (svc *ServiceClient) GetUnionListByUpazillaId(ctx *gin.Context) {
	routes.GetUnionListByUpazillaId(ctx, svc.Client)
}
func (svc *ServiceClient) GetVillageOrRoadListByUnionId(ctx *gin.Context) {
	routes.GetVillageOrRoadListByUnionId(ctx, svc.Client)
}
func (svc *ServiceClient) SearchStudent(ctx *gin.Context) {
	routes.SearchStudent(ctx, svc.Client)
}
