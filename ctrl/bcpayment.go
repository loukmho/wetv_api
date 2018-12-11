package ctrl


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/loukmho/wetv_api/model"
)

func SearchPaymentByDocNo(c *gin.Context) {
	c.Keys = HeaderKeys

	docno := c.Request.URL.Query().Get("docno")
	pmt := new(model.Payment)
	err := pmt.SearchPaymentByDocNo(Dbc, docno)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = pmt
		c.JSON(http.StatusOK, rs)
	}
}

func SearchPaymentByKeyword(c *gin.Context) {
	c.Keys = HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	atd := new(model.Payment)
	atds, err := atd.SearchPaymentByKeyword(Dbc, keyword)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = atds
		c.JSON(http.StatusOK, rs)
	}

}

func InsertAndEditPayment(c *gin.Context) {
	c.Keys = HeaderKeys

	atd := &model.Payment{}
	err := c.BindJSON(atd)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = atd.InsertAndEditPayment(Dbc)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	} else {
		rs.Status = "success"
		rs.Data = atd
		c.JSON(http.StatusOK, rs)
	}

}