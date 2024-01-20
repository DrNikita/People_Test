package handler

import (
	config "github.com/DrNikita/People/internal/db"
	"github.com/DrNikita/People/internal/model"
	"github.com/DrNikita/People/internal/service/pagination"
	"github.com/DrNikita/People/internal/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get users by filters
// @Tags Get users
// @Description get users by filters
// @ID get-users
// @Accept  json
// @Produce  json
// @Success 200	{object} []model.Persons
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
// @Router /find-users [get]
func FindPeople(c *gin.Context) {
	var pageableResponse model.PageableResponse
	var paginationInfo pagination.Pagination
	var filters model.Filters
	var emptyFilters model.Filters
	var people []model.Persons
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

	paginateFunc := pagination.Paginate(model.Persons{}, &paginationInfo, dbConn)
	result := dbConn.
		Debug().
		Model(model.Persons{}).
		Scopes(paginateFunc).
		Find(&people)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, pageableResponse.ErrorResponse(status.PeopleNotFound()))
		return
	}

	c.JSON(http.StatusOK, pageableResponse.New(people, paginationInfo, nil))
}
