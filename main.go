package main

import (
	"github.com/gin-gonic/gin"
	we "github.com/loukmho/wetv_api/ctrl"
)

func main() {
	r := gin.New()

	//Module Buy ///////////////////////////////////////////////////////////////////

	r.GET("/vendor", we.SearchVendorByCode)
	r.GET("/vendors", we.SearchVendorByKeyword)
	r.POST("/vendor", we.InsertAndEditVendor)

	r.GET("/apinvoice", we.SearchApInvoiceByDocNo)
	r.GET("/apinvoices", we.SearchApInvoiceByKeyword)
	r.POST("/apinvoice", we.InsertAndEditApinvoice)

	r.GET("/payment", we.SearchPaymentByDocNo)
	r.GET("/payments", we.SearchPaymentByKeyword)
	r.POST("/payment", we.InsertAndEditPayment)

	//Module Customer ///////////////////////////////////////////////////////////////////

	r.GET("/customer", we.SearchCustomerByCode)
	r.GET("/customers", we.SearchCustomerByKeyword)
	r.POST("/customer", we.InsertAndEditCustomer)

	r.GET("/arinvoice", we.SearchArInvoiceByDocNo)
	r.GET("/arinvoices", we.SearchArInvoiceByKeyword)
	r.POST("/arinvoice", we.InsertAndEditArInvoice)

	r.GET("/receipt1", we.SearchReceipt1ByDocNo)
	r.GET("/receipt1s", we.SearchReceipt1ByKeyword)
	r.POST("/receipt1", we.InsertAndEditReceipt1)

	//Module Inventory ///////////////////////////////////////////////////////////////////

	r.GET("/item", we.SearchItemByCode)
	r.GET("/items", we.SearchItemByKeyword)
	r.POST("/item", we.InsertAndEditItem)

	r.GET("/issue", we.SearchStkIssuetByDocNo)
	r.GET("/issues", we.SearchStkIssueByKeyword)
	r.POST("/issue", we.InsertAndEditStkIssue)


	r.Run(":8003")
}