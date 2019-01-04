package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"errors"
	"time"
)

type SaleOrder struct {
	DocNo           string         `json:"doc_no" db:"DocNo"`
	DocDate         string         `json:"doc_date" db:"DocDate"`
	ArCode          string         `json:"ar_code" db:"ArCode"`
	ArName          string         `json:"ar_name" db:"ArName"`
	SaleCode        string         `json:"sale_code" db:"SaleCode"`
	SaleName        string         `json:"sale_name" db:"SaleName"`
	BillType        int64          `json:"bill_type" db:"BillType"`
	TaxType         int            `json:"tax_type" db:"TaxType"`
	TaxRate         float64        `json:"tax_rate" db:"TaxRate"`
	DepartCode      string         `json:"depart_code" db:"DepartId"`
	RefDocNo        string         `json:"ref_doc_no" db:"RefDocNo"`
	IsConfirm       int64          `json:"is_confirm" db:"IsConfirm"`
	BillStatus      int64          `json:"bill_status" db:"BillStatus"`
	SOStatus        int64          `json:"so_status" db:"SOStatus"`
	CreditDay       int            `json:"credit_day" db:"CreditDay"`
	DueDate         string         `json:"due_date" db:"DueDate"`
	DeliveryDay     int            `json:"delivery_day" db:"DeliveryDay"`
	DeliveryDate    string         `json:"delivery_date" db:"DeliveryDate"`
	IsConditionSend int64          `json:"is_condition_send" db:"IsConditionSend"`
	MyDescription   string         `json:"my_description" db:"MyDescription"`
	SumOfItemAmount float64        `json:"sum_of_item_amount" db:"SumOfItemAmount"`
	DiscountWord    string         `json:"discount_word" db:"DiscountWord"`
	DiscountAmount  float64        `json:"discount_amount" db:"DiscountAmount"`
	AfterDiscount   float64        `json:"after_discount" db:"AfterDiscount"`
	BeforeTaxAmount float64        `json:"before_tax_amount" db:"BeforeTaxAmount"`
	TaxAmount       float64        `json:"tax_amount" db:"TaxAmount"`
	TotalAmount     float64        `json:"total_amount" db:"TotalAmount"`
	ExchangeRate    float64        `json:"exchange_rate" db:"exchange_rate"`
	NetAmount       float64        `json:"net_amount" db:"NetDebtAmount"`
	ProjectCode     int64          `json:"project_code" db:"ProjectId"`
	AllocateCode    int64          `json:"allocate_code" db:"AllocateId"`
	IsCancel        int64          `json:"allocate_code" db:"IsCancel"`
	IsCompleteSave  int64          `json:"is_complete_save" db:"IsCompleteSave"`
	CreatorCode     string         `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string         `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string         `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string         `json:"last_edit_date_t" db:"LastEditDateT"`
	ConfirmCode     string         `json:"confirm_code" db:"confirm_code"`
	ConfirmDateTime string         `json:"confirm_date_time" db:"ConfirmDateTime"`
	CancelCode      string         `json:"cancel_code" db:"CancelCode"`
	CancelDateTime  string         `json:"cancel_date_time" db:"CancelDateTime"`
	UserCode        string         `json:"user_code" db:"UserCode"`
	Subs            []SaleOrderSub `json:"subs"`
}

type SaleOrderSub struct {
	ItemCode       string  `json:"item_code" db:"ItemCode"`
	WHCode         string  `json:"wh_code" db:"WHCode"`
	ShelfCode      string  `json:"shelf_code" db:"ShelfCode"`
	ItemName       string  `json:"item_name" db:"ItemName"`
	Qty            float64 `json:"qty" db:"Qty"`
	RemainQty      float64 `json:"remain_qty" db:"RemainQty"`
	Price          float64 `json:"price" db:"Price"`
	DiscountWord   string  `json:"discount_word" db:"DiscountWord"`
	DiscountAmount float64 `json:"discount_amount" db:"DiscountAmount"`
	UnitCode       string  `json:"unit_code" db:"UnitCode"`
	Amount         float64 `json:"amount" db:"Amount"`
	NetAmount      float64 `json:"net_amount" db:"NetAmount"`
	HomeAmount     float64 `json:"home_amount" db:"HomeAmount"`
	MyDescription  string  `json:"my_description" db:"MyDescription"`
	StockType      int64   `json:"stock_type" db:"StockType"`
	AverageCost    float64 `json:"average_cost" db:"AverageCost"`
	SumOfCost      float64 `json:"sum_of_cost" db:"SumOfCost"`
	PackingRate1   float64 `json:"packing_rate_1" db:"PackingRate1"`
	RefNo          string  `json:"ref_no" db:"RefNo"`
	LineNumber     int     `json:"line_number" db:"LineNumber"`
	RefLineNUmber  int     `json:"ref_line_n_umber" db:"RefLineNUmber"`
	IsCancel       int64   `json:"is_cancel" db:"IsCancel"`
}

func (sod *SaleOrder) InsertAndEditSaleOrder(db *sqlx.DB) error {
	var check_exist int
	var count_item int

	now := time.Now()

	for _, sub_item := range sod.Subs {
		if (sub_item.Qty != 0) {
			count_item = count_item + 1
		}
	}

	sqlexist := `select count(docno) as check_exist from dbo.bcsaleorder where docno = ?`
	err := db.Get(&check_exist, sqlexist, sod.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil
	}
	fmt.Println("Docno = ", sod.DocNo)
	switch {
	case sod.DocNo == "":
		return errors.New("Docno is null")
	case sod.ArCode == "":
		return errors.New("Arcode is null")
	case sod.DocDate == "":
		return errors.New("Docdate is null")
	case sod.IsCancel == 1:
		return errors.New("Docno is cancel")
	case sod.IsConfirm == 1:
		return errors.New("Docno is confirm")
	case sod.SumOfItemAmount == 0 && count_item == 0:
		return errors.New("Docno not have SumOfItemAmount")
	}

	def := Default{}
	def = LoadDefaultData("bcdata.json")

	fmt.Println("TaxRate = ", def.TaxRateDefault)

	if sod.TaxRate == 0 {
		sod.TaxRate = def.TaxRateDefault
	}
	if sod.ExchangeRate == 0 {
		sod.ExchangeRate = def.ExchangeRateDefault
	}

	sod.BeforeTaxAmount, sod.TaxAmount, sod.TotalAmount = CalcTaxItem(sod.TaxType, sod.TaxRate, sod.AfterDiscount)
	fmt.Println("Befor,Tax", sod.BeforeTaxAmount, sod.TaxAmount)

	fmt.Println("UserCode = ", sod.UserCode)

	if (sod.IsCompleteSave == 0) {
		sod.IsCompleteSave = 1
	}

	if (sod.CreditDay == 0 || sod.DueDate == "") {
		sod.DueDate = sod.DocDate
	} else {
		sod.DueDate = now.AddDate(0, 0, sod.CreditDay).Format("2006-01-02")
	}
	if (sod.DeliveryDay == 0 && sod.DeliveryDate == "") {
		sod.DeliveryDate = sod.DocDate
	} else {
		sod.DeliveryDate = now.AddDate(0, 0, sod.DeliveryDay).Format("2006-01-02")
	}

	if check_exist == 0 {
		sql := `set dateformat dmy      insert into dbo.bcsaleorder(DocNo, DocDate, ArCode, SaleCode, BillType, TaxType, TaxRate, DepartCode, RefDocNo, IsConfirm, BillStatus, SOStatus, CreditDay, DueDate, DeliveryDay, DeliveryDate, IsConditionSend, MyDescription, SumOfItemAmount, DiscountWord, DiscountAmount, AfterDiscount, BeforeTaxAmount, TaxAmount, TotalAmount, ExchangeRate, NetAmount, ProjectCode, AllocateCode, IsCancel, IsCompleteSave, CreatorCode, CreateDateTime) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, getdate())`
		_, err = db.Exec(sql, sod.DocNo, sod.DocDate, sod.ArCode, sod.SaleCode, sod.BillType, sod.TaxType, sod.TaxRate, sod.DepartCode, sod.RefDocNo, sod.IsConfirm, sod.BillStatus, sod.SOStatus, sod.CreditDay, sod.DueDate, sod.DeliveryDay, sod.DeliveryDate, sod.IsConditionSend, sod.MyDescription, sod.SumOfItemAmount, sod.DiscountWord, sod.DiscountAmount, sod.AfterDiscount, sod.BeforeTaxAmount, sod.TaxAmount, sod.TotalAmount, sod.ExchangeRate, sod.NetAmount, sod.ProjectCode, sod.AllocateCode, sod.IsCancel, sod.IsCompleteSave, sod.UserCode)
		if err != nil {
			fmt.Println("Insert Credit =", err.Error())
			return err
		}
	} else {
		sql := `set dateformat dmy      update dbo.bcsaleorder set DocDate=?, ArCode=?, SaleCode=?, BillType=?, TaxType=?, TaxRate=?, DepartCode=?, RefDocNo=?, IsConfirm=?, BillStatus=?, SOStatus=?, CreditDay=?, DueDate=?, DeliveryDay=?, DeliveryDate=?, IsConditionSend=?, MyDescription=?, SumOfItemAmount=?, DiscountWord=?, DiscountAmount=?, AfterDiscount=?, BeforeTaxAmount=?, TaxAmount=?, TotalAmount=?, ExchangeRate=?, NetAmount=?, ProjectCode=?, AllocateCode=?, IsCancel=?, LastEditorCode=?, LastEditDateT=getdate() where docno = ? `
		_, err = db.Exec(sql, sod.DocDate, sod.ArCode, sod.SaleCode, sod.BillType, sod.TaxType, sod.TaxRate, sod.DepartCode, sod.RefDocNo, sod.IsConfirm, sod.BillStatus, sod.SOStatus, sod.CreditDay, sod.DueDate, sod.DeliveryDay, sod.DeliveryDate, sod.IsConditionSend, sod.MyDescription, sod.SumOfItemAmount, sod.DiscountWord, sod.DiscountAmount, sod.AfterDiscount, sod.BeforeTaxAmount, sod.TaxAmount, sod.TotalAmount, sod.ExchangeRate, sod.NetAmount, sod.ProjectCode, sod.AllocateCode, sod.IsCancel, sod.UserCode, sod.DocNo)
		if err != nil {
			fmt.Println("Update Credit =", err.Error())
			return err
		}
	}

	sql_del_sub := `delete dbo.bcsaleordersub where docno = ?`
	_, err = db.Exec(sql_del_sub, sod.DocNo)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}

	var vLineNumber int

	for _, item := range sod.Subs {

		item.LineNumber = vLineNumber

		item.IsCancel = 0
		if (item.PackingRate1 == 0) {
			item.PackingRate1 = 1
		}

		if item.Amount == 0 {
			item.Amount = (item.Qty * (item.Price - (item.DiscountAmount / item.Qty)))
		}

		switch {
		case sod.TaxType == 0:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		case sod.TaxType == 1:
			taxamount := ToFixed(item.Amount-((item.Amount*100)/(100+float64(sod.TaxRate))), 2)
			beforetaxamount := ToFixed(item.Amount-taxamount, 2)
			item.HomeAmount = beforetaxamount
			item.NetAmount = beforetaxamount
		case sod.TaxType == 2:
			item.HomeAmount = item.Amount
			item.NetAmount = item.Amount
		}

		item.RemainQty = item.Qty

		sqlsub := `set dateformat dmy       insert into dbo.BCSaleOrderSub(DocNo,TaxType,ItemCode,DocDate,ArCode,DepartCode,SaleCode,MyDescription,ItemName,WHCode,ShelfCode,Qty,RemainQty, Price, DiscountWord, DiscountAmount, Amount,NetAmount,HomeAmount,UnitCode,SOStatus,AllocateCode,ProjectCode,ExchangeRate,IsCancel,LineNumber,IsConditionSend,PACKINGRATE1,PACKINGRATE2) values(DocNo,TaxType,ItemCode,DocDate,ArCode,DepartCode,SaleCode,MyDescription,ItemName,WHCode,ShelfCode,Qty,RemainQty, Price, DiscountWord, DiscountAmount, Amount,NetAmount,HomeAmount,UnitCode,SOStatus,HoldingStatus,RefType,AllocateCode,ProjectCode,ExchangeRate,IsCancel,LineNumber,IsConditionSend,PACKINGRATE1,PACKINGRATE2)`
		_, err = db.Exec(sqlsub, sod.DocNo, sod.TaxType, item.ItemCode, sod.DocDate, sod.ArCode, sod.DepartCode, sod.SaleCode, item.MyDescription, item.ItemName, item.WHCode, item.ShelfCode, item.Qty, item.RemainQty, item.Price, item.DiscountWord, item.DiscountAmount, item.Amount, item.NetAmount, item.HomeAmount, item.UnitCode, sod.SOStatus, sod.AllocateCode, sod.ProjectCode, sod.ExchangeRate, item.IsCancel, item.LineNumber, sod.IsConditionSend, item.PackingRate1, 1)
		fmt.Println("sqltax = ", sqlsub)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return err
		}

		vLineNumber = vLineNumber + 1
	}

	return nil
}

func (inv *ArInvoice) SearchSaleOrderByDocNo(db *sqlx.DB, docno string) error {
	sql := `set dateformat dmy     select a.DocNo, a.DocDate, a.ArCode, isnull(b.name1,'') as ArName, isnull(a.SaleCode,'')  as SaleCode, isnull(c.name,'') as SaleName, BillType, TaxType, TaxRate, DepartCode, RefDocNo, IsConfirm, BillStatus, SOStatus, CreditDay, DueDate, DeliveryDay, DeliveryDate, IsConditionSend, MyDescription, SumOfItemAmount, DiscountWord, DiscountAmount, AfterDiscount, BeforeTaxAmount, TaxAmount, TotalAmount, ExchangeRate, NetAmount, ProjectCode, AllocateCode, IsCancel, IsCompleteSave, CreatorCode, CreateDateTime from dbo.bcsaleorder a left join dbo.bcar b with (nolock) on a.arcode = b.code left join dbo.bcsale c with (nolock) on a.salecode = c.code where docno = ? `
	fmt.Println("sql=", sql)
	err := db.Get(inv, sql, docno)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sqlsub := `set dateformat dmy    select MyType, ItemCode, StockType,isnull(MyDescription,'') as MyDescription, ItemName, WHCode, ShelfCode, CNQty, Qty, Price, isnull(DiscountWord,'') as DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, BalanceAmount, UnitCode , isnull(SORefNo,'') as SORefNo, isnull(PORefNo,'') as PORefNo, LineNumber, RefLineNumber, IsCancel, isnull(BarCode,'') as BarCode, IsConditionSend, AverageCost , isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.bcarinvoicesub  where docno = ? order by LineNumber`
	fmt.Println("sqlsub=", sqlsub)
	err = db.Select(&inv.Subs, sqlsub, docno)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return err
	}
	return nil
}

func (inv *ArInvoice) SearchSaleOrderByKeyword(db *sqlx.DB, keyword string) (invs []*ArInvoice, err error) {
	sql := `set dateformat dmy     select DocNo,DocDate,isnull(a.TaxNo,'') as TaxNo,ArCode,isnull(b.name1,'') as ArName,isnull(a.SaleCode,'') as SaleCode,isnull(c.name,'') as SaleName,a.TaxType,isnull(a.DepartCode,'') as DepartCode,CreditDay,DeliveryDay,isnull(DeliveryDate,'') as DeliveryDate,isnull(DueDate,'') as DueDate,isnull(PayBillDate,'') as PayBillDate,TaxRate,IsConfirm,isnull(MyDescription,'') as MyDescription,BillType,isnull(BillGroup,'') as BillGroup,isnull(RefDocNo,'') as RefDocNo,isnull(DeliveryAddr,'') as DeliveryAddr,isnull(ContactCode,'') as ContactCode,SumOfItemAmount,isnull(a.DiscountWord,'') as DiscountWord,isnull(DiscountAmount,'') as DiscountAmount,AfterDiscount,BeforeTaxAmount,TaxAmount,TotalAmount,ZeroTaxAmount,ExceptTaxAmount,SumCashAmount,SumChqAmount,SumCreditAmount,SumBankAmount,DepositIncTax,SumOfDeposit1,SumOfDeposit2,SumOfWTax,NetDebtAmount,HomeAmount,OtherIncome,OtherExpense,ExcessAmount1,ExcessAmount2,BillBalance,isnull(CurrencyCode,'') as CurrencyCode,ExchangeRate,isnull(GLFormat,'') as GLFormat,IsCancel,IsCompleteSave,isnull(AllocateCode,'') as AllocateCode,isnull(ProjectCode,'') as ProjectCode,isnull(RecurName,'') as RecurName,IsConditionSend,PayBillAmount,isnull(SORefNo,'') as SORefNo,HoldingStatus,PosStatus,a.CreatorCode,a.CreateDateTime,isnull(ShiftCode,'') as ShiftCode,isnull(CashierCode,'') as CashierCode,'' as ShiftNo,isnull(MachineNo,'') as MachineNo,isnull(MachineCode,'') as MachineCode,isnull(BillTime,'') as BillTime,isnull(CreditType,'') as CreditType,isnull(CreditDueDate,'') as CreditDueDate,isnull(CreditNo,'') as CreditNo,isnull(CofirmNo,'') as CofirmNo,CreditBaseAmount,CoupongAmount,ChangeAmount,ChargeAmount,GrandTotal from dbo.bcarinvoice a with (nolock) left join dbo.bcar b with (nolock) on a.arcode = b.code left join dbo.bcsale c with (nolock) on a.salecode = c.code where (docno  like '%'+?+'%' or arcode like '%'+?+'%' or a.salecode like '%'+?+'%' or b.name1 like '%'+?+'%' or c.name like '%'+?+'%' ) order by docno`
	fmt.Println("sql=", sql)
	err = db.Select(&invs, sql, keyword, keyword, keyword, keyword, keyword)
	if err != nil {
		fmt.Println("Error = ", err.Error())
		return nil, err
	}
	for _, sub := range invs {
		sqlsub := `set dateformat dmy    select MyType, ItemCode, StockType, isnull(MyDescription,'') as MyDescription, isnull(ItemName,'') as ItemName, isnull(WHCode,'') as WHCode, isnull(ShelfCode,'') as ShelfCode, CNQty, Qty, Price, isnull(DiscountWord,'') as DiscountWord, DiscountAmount, Amount, NetAmount, HomeAmount, SumOfCost, BalanceAmount, isnull(UnitCode,'') as UnitCode, isnull(SORefNo,'') as SORefNo, isnull(PORefNo,'') as PORefNo, LineNumber, RefLineNumber, IsCancel, isnull(BarCode,'') as BarCode, IsConditionSend, AverageCost , isnull(LotNumber,'') as LotNumber, isnull(PackingRate1,1) as PackingRate1, isnull(PackingRate2,1) as PackingRate2 from dbo.bcarinvoicesub  where docno = ? order by LineNumber`
		fmt.Println("sqlsub=", sqlsub)
		err = db.Select(&sub.Subs, sqlsub, sub.DocNo)
		if err != nil {
			fmt.Println("Error = ", err.Error())
			return nil, err
		}
	}
	return invs, nil
}
