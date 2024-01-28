package handler

import (
	"fmt"
	config "github.com/DrNikita/People/internal/db"
	"github.com/DrNikita/People/internal/model"
	"github.com/DrNikita/People/internal/service"
	"github.com/DrNikita/People/internal/service/pagination"
	"github.com/DrNikita/People/internal/status"
	"github.com/DrNikita/People/internal/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get people by filters
// @Tags People
// @Description get people by filters
// @ID get-people
// @Accept  json
// @Produce  json
// @Success 200	{object} []model.Person
// @Param page query int false	"Page"
// @Param perPage query int false	"Size"
// @Param sortBy query string false	"Sort field"
// @Param sortDirection query string false	"Direction"
// @Param name query string false	"Name"
// @Param surname query string false	"Surname"
// @Param patronymic query string false	"Patronymic"
// @Param age query int false	"age"
// @Param gender query string false	"gender"
// @Param country query string false	"country"
// @Router /find-persons [get]
func FindPeople(c *gin.Context) {
	var pageableResponse model.PageableResponse
	var paginationInfo pagination.Pagination
	var filters model.Filters
	var emptyFilters model.Filters
	var people []model.PersonFullInfo
	dbConn := config.GetDBInstance()

	if err := paginationInfo.Bind(c); err != nil {
		c.JSON(http.StatusAccepted, pageableResponse.ErrorResponse(err))
		return
	}
	filters.Bind(c)

	if filters.Name != emptyFilters.Name {
		dbConn = dbConn.Where("name = ?", filters.Name)
	}
	if filters.Surname != emptyFilters.Surname {
		dbConn = dbConn.Where("surname = ?", filters.Surname)
	}
	if filters.Patronymic != emptyFilters.Surname {
		dbConn = dbConn.Where("patronymic = ?", filters.Patronymic)
	}
	if filters.Age != emptyFilters.Age {
		dbConn = dbConn.Where("age = ?", filters.Age)
	}
	if filters.Gender != emptyFilters.Gender {
		dbConn = dbConn.Where("gender = ?", filters.Gender)
	}
	if filters.Country != emptyFilters.Country {
		dbConn = dbConn.Where("country = ?", filters.Country)
	}

	paginateFunc := pagination.Paginate(model.Person{}, &paginationInfo, dbConn)
	result := dbConn.
		Debug().
		Table("persons").
		Scopes(paginateFunc).
		Find(&people)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, pageableResponse.ErrorResponse(status.PeopleNotFound()))
		return
	}

	c.JSON(http.StatusOK, pageableResponse.New(people, paginationInfo, nil))
}

// @Summary Create person
// @Tags People
// @Description create person
// @ID create-person
// @Accept  json
// @Produce  json
// @Success 202	{object} model.Response
// @Failure 404 {object} model.Response
// @Param person body model.Person true "Person"
// @Router /create-person [post]
func CreatePerson(c *gin.Context) {
	var response model.Response
	var person model.SupplementedPerson
	var personFullInfo model.PersonFullInfo
	dbConn := config.GetDBInstance()
	tx := dbConn.Begin()

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	validationErr, err := validation.ValidatePerson(person)
	if validationErr != "" {
		c.JSON(http.StatusBadRequest, response.MessageResponse(validationErr))
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	personFullInfo = model.PersonFullInfo{
		SupplementedPerson: person,
	}

	err = tx.Debug().Create(&personFullInfo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	go service.AddInfo(&personFullInfo, tx)

	c.JSON(http.StatusCreated, response.New(person, fmt.Sprintf("person (ID : %d) created", personFullInfo.ID)))
}

// @Summary Update person
// @Tags People
// @Description Update person
// @ID update-person
// @Accept  json
// @Produce  json
// @Success 202	{object} model.Response
// @Failure 404 {object} model.Response
// @Param id path int true "ID"
// @Param person body model.SupplementedPerson true "Person"
// @Router /update-person/{id} [patch]
func UpdatePerson(c *gin.Context) {
	var response model.Response
	var person model.SupplementedPerson
	var personFullInfo model.PersonFullInfo
	dbConn := config.GetDBInstance()

	pathId := c.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		return
	}

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	personFullInfo.SupplementedPerson = person

	result := dbConn.Debug().Where("id = ?", id).Updates(&personFullInfo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(result.Error))
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(status.NonExistentId(id)))
		return
	}

	c.JSON(http.StatusCreated, response.New(person, fmt.Sprintf("person (ID : %d) created", id)))
}

// @Summary Delete person
// @Tags People
// @Description delete person by id
// @ID delete-person
// @Accept  json
// @Produce  json
// @Success 202	{object} model.Response
// @Failure 404 {object} model.Response
// @Param id path int false	"ID"
// @Router /delete-person/{id} [delete]
func DeletePerson(c *gin.Context) {
	var response model.Response
	dbConn := config.GetDBInstance()

	pathId := c.Param("id")
	id, err := strconv.Atoi(pathId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	result := dbConn.Debug().
		Where("id = ?", id).
		Delete(&model.Person{})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(result.Error))
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(status.NonExistentId(id)))
		return
	}

	c.JSON(http.StatusAccepted, response.MessageResponse(fmt.Sprintf("%s; id : %d", status.DELETED, id)))
}
