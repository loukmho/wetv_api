package ctrl

import(
	"github.com/gin-gonic/gin"
	"github.com/loukmho/wetv_api/model"
	"net/http"
	"fmt"
	ct "github.com/loukmho/wetv_api/ctrl"
)

func InsertAndEditArInvoice(c *gin.Context){
	c.Keys = ct.HeaderKeys

	inv := &model.ArInvoice{}
	err := c.BindJSON(inv)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = inv.InsertAndEditArInvoice(ct.Dbc)
	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound, rs)
	}else{
		rs.Status = "success"
		rs.Data = inv
		c.JSON(http.StatusNotFound, rs)
	}
}

func SearchArInvoiceByDocNo(c *gin.Context) {
	c.Keys = ct.HeaderKeys

	docno := c.Request.URL.Query().Get("docno")
	inv := new(model.ArInvoice)
	err := inv.SearchArInvoiceByDocNo(ct.Dbc, docno)

	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = inv
		c.JSON(http.StatusOK, rs)
	}
}

func SearchArInvoiceByKeyword(c *gin.Context){
	c.Keys = ct.HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	inv := new(model.ArInvoice)
	ins, err := inv.SearchArInvoiceByKeyword(ct.Dbc, keyword)
	rs := ct.Response{}
	if err != nil {
		rs.Status = "error"
		rs.Message = "No Content and Error : " + err.Error()
		c.JSON(http.StatusNotFound,rs)
	}else{
		rs.Status = "success"
		rs.Data = ins
		c.JSON(http.StatusOK, rs)
	}
}
