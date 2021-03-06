package ctrl

import(
	"github.com/gin-gonic/gin"
	"github.com/loukmho/wetv_api/model"
	"net/http"
	"fmt"
)

func InsertAndEditArInvoice(c *gin.Context){
	c.Keys = HeaderKeys

	inv := &model.ArInvoice{}
	err := c.BindJSON(inv)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = inv.InsertAndEditArInvoice(Dbc)
	rs := Response{}
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
	c.Keys = HeaderKeys

	docno := c.Request.URL.Query().Get("docno")
	inv := new(model.ArInvoice)
	err := inv.SearchArInvoiceByDocNo(Dbc, docno)

	rs := Response{}
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
	c.Keys = HeaderKeys

	keyword := c.Request.URL.Query().Get("keyword")

	inv := new(model.ArInvoice)
	ins, err := inv.SearchArInvoiceByKeyword(Dbc, keyword)
	rs := Response{}
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
