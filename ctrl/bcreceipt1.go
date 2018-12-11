package ctrl

import (
	"github.com/loukmho/wetv_api/model"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SearchReceipt1ByDocNo(c *gin.Context) {
	c.Keys = HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	rcp := new(model.Receipt1)
	err := rcp.SearchReceipt1ByDocNo(Dbc, docno)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rcp
		c.JSON(http.StatusOK, rs)
	}
}

func SearchReceipt1ByKeyword(c *gin.Context) {
	c.Keys = HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	rcp := new(model.Receipt1)
	rcps, err := rcp.SearchReceipt1ByKeyword(Dbc, keyword)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rcps
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditReceipt1(c *gin.Context) {
	c.Keys = HeaderKeys

	rcp := &model.Receipt1{}
	err := c.BindJSON(rcp)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = rcp.InsertAndEditReceipt1(Dbc)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = rcp
		c.JSON(http.StatusOK, rs)
	}

}
