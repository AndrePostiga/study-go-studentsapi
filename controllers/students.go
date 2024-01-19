package controllers

import (
	"net/http"
	"strconv"
	"strings"

	domainservices "github.com/andrepostiga/api-go-gin/domain/domainServices"
	"github.com/andrepostiga/api-go-gin/domain/repository"
	"github.com/andrepostiga/api-go-gin/models"
	"github.com/andrepostiga/api-go-gin/models/request"
	"github.com/andrepostiga/api-go-gin/models/response"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	r   repository.IStudentRepository
	svc domainservices.ICreateStudentService
}

func NewStudentController(
	studentRepository repository.IStudentRepository,
	createStudentDomainService domainservices.ICreateStudentService) *StudentController {

	return &StudentController{
		r:   studentRepository,
		svc: createStudentDomainService,
	}
}

func (ctrl *StudentController) SetupStudentController(rg *gin.RouterGroup) {
	rg = rg.Group("/students")
	{
		rg.GET("/:id", ctrl.GetStudentById)
		rg.GET("/", ctrl.GetStudents)
		rg.POST("/", ctrl.CreateStudent)
	}
}

func (ctrl *StudentController) GetStudents(c *gin.Context) {
	students, err := ctrl.r.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var studentsResponse []response.GetStudentResponse
	for _, s := range students {
		studentsResponse = append(studentsResponse, response.GetStudentResponse{
			Id:      strconv.Itoa(int(s.ID)),
			Name:    s.GetName(),
			CPF:     s.GetCPF(),
			Address: response.NewAddressResponse(s.Address),
		})
	}

	c.JSON(http.StatusOK, studentsResponse)
}

func (ctrl *StudentController) GetStudentById(c *gin.Context) {
	studentId, found := c.Params.Get("id")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":      "The request is invalid",
			"requested_id": studentId,
		})
	}

	student, err := ctrl.r.Get(studentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.GetStudentResponse{
		Id:      strconv.Itoa(int(student.ID)),
		Name:    student.GetName(),
		CPF:     student.GetCPF(),
		Address: response.NewAddressResponse(student.Address),
	})
}

func (ctrl *StudentController) CreateStudent(c *gin.Context) {
	var request *request.CreateStudentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ProblemDetails{
			Type:   "validation_error",
			Title:  "Invalid Request",
			Status: http.StatusBadRequest,
			Detail: strings.Split(err.Error(), "\n"),
		})
		return
	}

	// TODO tentar montar essa validação em um middleware futuramente
	if _, err := govalidator.ValidateStruct(request); err != nil {
		var errorDetails []string

		if validationErrors, ok := err.(govalidator.Errors); ok {
			for _, validationErr := range validationErrors.Errors() {
				errorDetails = append(errorDetails, validationErr.Error())
			}
		}

		c.JSON(http.StatusBadRequest, models.ProblemDetails{
			Type:   "validation_error",
			Title:  "Invalid Request",
			Status: http.StatusBadRequest,
			Detail: errorDetails,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
	studentDomain, err := ctrl.svc.Create(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ProblemDetails{
			Type:   "validation_error",
			Title:  "Invalid Request",
			Status: http.StatusBadRequest,
			Detail: strings.Split(err.Error(), "\n"),
		})
		return
	}

	c.JSON(http.StatusCreated, response.GetStudentResponse{
		Id:      strconv.Itoa(int(studentDomain.ID)),
		Name:    studentDomain.GetName(),
		CPF:     studentDomain.GetCPF(),
		Address: response.NewAddressResponse(studentDomain.Address),
	})
}
