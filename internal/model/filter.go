package model

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Filters struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Country    string `json:"country"`
}

func (f *Filters) Bind(c *gin.Context) {
	*f = Filters{
		Name:       c.Query("name"),
		Surname:    c.Query("surname"),
		Patronymic: c.Query("patronymic"),
		Age: func(strAge string) int {
			age, err := strconv.Atoi(strAge)
			if err != nil {
				return 0
			}
			return age
		}(c.Query("age")),
		Gender:  c.Query("gender"),
		Country: c.Query("country"),
	}
}
