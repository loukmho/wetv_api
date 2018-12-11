package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/loukmho/wetv_api/model"
	"net/http"
	"fmt"
)

func SearchItemByCode(c *gin.Context){
	c.Keys = HeaderKeys

	code := c.Request.URL.Query().Get("code")

	fmt.Println("ctrl code = ",code)
	itm := new(model.Item)
	err := itm.SearchItemByCode(Dbc, code)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = itm
		c.JSON(http.StatusOK, rs)
	}
}

func SearchItemByKeyword(c *gin.Context){
	c.Keys = HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	itm := new(model.Item)
	itms, err := itm.SearchItemByKeycode(Dbc, keyword)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = itms
		c.JSON(http.StatusOK, rs)
	}
}

func InsertAndEditItem(c *gin.Context){
	c.Keys = HeaderKeys

	itm := &model.Item{}
	err := c.BindJSON(itm)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = itm.InsertAndEditItem(Dbc)

	rs := Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = itm
		c.JSON(http.StatusOK ,rs)
	}
}