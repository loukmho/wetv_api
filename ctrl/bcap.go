package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/loukmho/wetv_api/model"
)

func SearchVendorByCode(c *gin.Context){
	c.Keys = HeaderKeys

	code := c.Request.URL.Query().Get("code")

	vdm := new(model.Vendor)
	err := vdm.SearchVendorByCode(Dbc, code)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = vdm
		c.JSON(http.StatusOK ,rs)
	}
}

func SearchVendorByKeyword(c *gin.Context){
	c.Keys = HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	vdm := new(model.Vendor)
	vdms, err := vdm.SearchVendorByKeyword(Dbc, keyword)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = vdms
		c.JSON(http.StatusOK ,rs)
	}
}

func InsertAndEditVendor(c *gin.Context){
	c.Keys = HeaderKeys

	vdm := &model.Vendor{}
	err := c.BindJSON(vdm)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = vdm.InsertAndEditVendor(Dbc)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = vdm
		c.JSON(http.StatusOK ,rs)
	}
}
