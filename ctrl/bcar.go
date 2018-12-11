package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/loukmho/wetv_api/model"
	"net/http"
	"fmt"
)

func SearchCustomerByCode(c *gin.Context){
	c.Keys = HeaderKeys

	code := c.Request.URL.Query().Get("code")

	ctm := new(model.Customer)
	err := ctm.SearchCustomerByCode(Dbc, code)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = ctm
		c.JSON(http.StatusOK ,rs)
	}
}

func SearchCustomerByKeyword(c *gin.Context){
	c.Keys = HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	ctm := new(model.Customer)
	ctms, err := ctm.SearchCustomerByKeyword(Dbc, keyword)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = ctms
		c.JSON(http.StatusOK ,rs)
	}
}

func InsertAndEditCustomer(c *gin.Context){
	c.Keys = HeaderKeys

	ctm := &model.Customer{}
	err := c.BindJSON(ctm)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = ctm.InsertAndEditCustomer(Dbc)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = ctm
		c.JSON(http.StatusOK ,rs)
	}
}
