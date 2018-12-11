package ctrl

import (
	"github.com/loukmho/wetv_api/model"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func SearchStkIssuetByDocNo(c *gin.Context) {
	c.Keys = HeaderKeys


	docno := c.Request.URL.Query().Get("docno")
	iss := new(model.StkIssue)
	err := iss.SearchStkIssueByDocNo(Dbc, docno)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = iss
		c.JSON(http.StatusOK, rs)
	}
}

func SearchStkIssueByKeyword(c *gin.Context) {
	c.Keys = HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	iss := new(model.StkIssue)
	isss, err := iss.SearchStkIssueByKeyword(Dbc, keyword)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = isss
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditStkIssue(c *gin.Context) {
	c.Keys = HeaderKeys

	iss := &model.StkIssue{}
	err := c.BindJSON(iss)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = iss.InsertAndEditStkIssue(Dbc)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = iss
		c.JSON(http.StatusOK, rs)
	}

}
