package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/loukmho/wetv_api/model"
	"fmt"
)


func InsertAndEditApinvoice(c *gin.Context){
	c.Keys = HeaderKeys

	apv := &model.ApInvoice{}
	err := c.BindJSON(apv)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = apv.InsertAndEditApInvoice(Dbc)
	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = apv
		c.JSON(http.StatusNotFound, rs)
	}
}

func CancelApinvoice(c *gin.Context){
	c.Keys = HeaderKeys

	apv := &model.ApInvoice{}
	err := c.BindJSON(apv)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = apv.InsertAndEditApInvoice(Dbc)
	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = apv
		c.JSON(http.StatusNotFound, rs)
	}
}

func SearchApInvoiceByDocNo(c *gin.Context){
	c.Keys = HeaderKeys

	docno := c.Request.URL.Query().Get("docno")
	apv := new(model.ApInvoice)
	err := apv.SearchApInvoiceByDocNo(Dbc, docno)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = apv
		c.JSON(http.StatusOK, rs)
	}
}

func SearchApInvoiceByKeyword(c *gin.Context){
	c.Keys = HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")
	apv := new(model.ApInvoice)
	apvs, err := apv.SearchApInvoiceByKeyword(Dbc, keyword)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = apvs
		c.JSON(http.StatusOK, rs)
	}
}
