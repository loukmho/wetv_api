package main

import (
	"github.com/gin-gonic/gin"
	c "github.com/loukmho/bcaccount_api/ctrl"
)

func main() {
	r := gin.New()
	r.Use(cors.Default())

	//Module Buy ///////////////////////////////////////////////////////////////////

	r.GET("/apinvoice", b.SearchApInvoiceByDocNo)
	r.GET("/apinvoices", b.SearchApInvoiceByKeyword)
	r.POST("/apinvoice", b.InsertAndEditApinvoice)

	//Module Customer ///////////////////////////////////////////////////////////////////

	r.GET("/customer", c.SearchCustomerByCode)
	r.GET("/customers", c.SearchCustomerByKeyword)
	r.POST("/customer", c.InsertAndEditCustomer)

	r.GET("/receipt1", c.SearchReceipt1ByDocNo)
	r.GET("/receipts1", c.SearchReceipt1ByKeyword)
	r.POST("/receipt1", c.InsertAndEditReceipt1)

	//Module Sale ///////////////////////////////////////////////////////////////////

	r.GET("/arinvoice", s.SearchArInvoiceByDocNo)
	r.GET("/arinvoices", s.SearchArInvoiceByKeyword)
	r.POST("/arinvoice", s.InsertAndEditArInvoice)

	//Module Inventory ///////////////////////////////////////////////////////////////////

	r.GET("/item", i.SearchItemByCode)
	r.GET("/items", i.SearchItemByKeyword)
	r.POST("/item", i.InsertAndEditItem)

	r.GET("/item", i.SearchItemByCode)
	r.GET("/items", i.SearchItemByKeyword)
	r.POST("/item", i.InsertAndEditItem)


	r.Run(":8001")
}