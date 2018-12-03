package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
)

type Vendor struct {
	Code            string  `json:"code" db:"Code"`
	Name1           string  `json:"name_1" db:"Name1"`
	Name2           string  `json:"name_2" db:"Name2"`
	DefDeliveryAddr string  `json:"def_delivery_addr" db:"DefDeliveryAddr"`
	DefContactCode  string  `json:"def_contact_code" db:"DefContactCode"`
	Address         string  `json:"address" db:"Address"`
	Telephone       string  `json:"telephone" db:"Telephone"`
	Fax             string  `json:"fax" db:"Fax"`
	AccountCode     string  `json:"account_code" db:"AccountCode"`
	IDCardNo        string  `json:"id_card_no" db:"IDCardNo"`
	BankAccNo       string  `json:"bank_acc_no" db:"BankAccNo"`
	CreditDay       int64   `json:"credit_day" db:"CreditDay"`
	LeadTime        int64   `json:"lead_time" db:"LeadTime"`
	DefaultTaxRate  float64 `json:"default_tax_rate" db:"DefaultTaxRate"`
	TaxNo           string  `json:"lead_time" db:"TaxNo"`
	PicFileName     string  `json:"pic_file_name" db:"PicFileName"`
	TypeCode        string  `json:"type_code" db:"TypeCode"`
	EmailAddress    string  `json:"email_address" db:"EmailAddress"`
	GroupCode       string  `json:"group_code" db:"GroupCode"`
	GroupOfDebt     string  `json:"group_of_debt" db:"GroupOfDebt"`
	PersonType      int64   `json:"person_type" db:"PersonType"`
	ActiveStatus    int64   `json:"active_status" db:"ActiveStatus"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	UserCode        string  `json:"user_code" db:"UserCode"`
}


func (vnd *Vendor) SearchVendorByCode(db *sqlx.DB, code string) error {

	sql := `set dateformat dmy     select Code, Name1, isnull(Name2,'') as Name2, isnull(DefDeliveryAddr,'') as DefDeliveryAddr, isnull(DefContactCode,'') as DefContactCode,isnull(Address,'') as Address, isnull(Telephone,'') as Telephone, isnull(Fax,'') as Fax, isnull(AccountCode,'') as AccountCode, isnull(IDCardNo,'') as IDCardNo, isnull(BankAccNo,'') as BankAccNo, CreditDay, LeadTime, DefaultTaxRate,isnull(TaxNo,'') as TaxNo, isnull(PicFileName,'') as PicFileName, isnull(TypeCode,'') as TypeCode, isnull(EmailAddress,'') as EmailAddress, isnull(GroupCode,'') as GroupCode, isnull(GroupOfDebt,,'') as GroupOfDebt,PersonType, ActiveStatus, CreatorCode, CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT from dbo.bcap where code = ?`
	err := db.Get(vnd, sql, code)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (ctm *Vendor) SearchVendorByKeyword(db *sqlx.DB, keyword string) (vndList []*Vendor, err error) {
	sql := `set dateformat dmy     select Code, Name1, isnull(Name2,'') as Name2, isnull(DefDeliveryAddr,'') as DefDeliveryAddr, isnull(DefContactCode,'') as DefContactCode,isnull(Address,'') as Address, isnull(Telephone,'') as Telephone, isnull(Fax,'') as Fax, isnull(AccountCode,'') as AccountCode, isnull(IDCardNo,'') as IDCardNo, isnull(BankAccNo,'') as BankAccNo, CreditDay, LeadTime, DefaultTaxRate,isnull(TaxNo,'') as TaxNo, isnull(PicFileName,'') as PicFileName, isnull(TypeCode,'') as TypeCode, isnull(EmailAddress,'') as EmailAddress, isnull(GroupCode,'') as GroupCode, isnull(GroupOfDebt,,'') as GroupOfDebt,PersonType, ActiveStatus, CreatorCode, CreateDateTime, isnull(LastEditorCode,'') as LastEditorCode, isnull(LastEditDateT,'') as LastEditDateT from dbo.bcap where (code  like '%'+?+'%' or name1 like '%'+?+'%' or Address like '%'+?+'%')`
	err = db.Select(&vndList, sql, keyword, keyword, keyword, keyword)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return vndList, nil
}


func (vnd *Vendor) InsertAndEditVendor(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(code) as check_exist from dbo.bcap where code = ?`
	err := db.Get(&check_exist, sqlexist, vnd.Code)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case vnd.Code == "":
		return errors.New("code is null")
	case vnd.Name1 == "":
		return errors.New("Name1 is null")
	case vnd.Address == "":
		return errors.New("address is null")
	}

	if (vnd.DefDeliveryAddr == "") {
		vnd.DefDeliveryAddr = vnd.Address
	}

	if (check_exist == 0) {
		vnd.ActiveStatus = 1
		vnd.CreatorCode = vnd.UserCode

		sql := `Insert into dbo.bcap(Code, Name1, Name2, DefDeliveryAddr, DefContactCode, Address, Telephone, Fax, AccountCode, IDCardNo, BankAccNo, CreditDay, LeadTime, DefaultTaxRate, TaxNo, PicFileName, TypeCode, EmailAddress, GroupCode, GroupOfDebt, PersonType, ActiveStatus, CreatorCode, CreateDateTime) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, getdate())`
		_, err = db.Exec(sql, vnd.Code, vnd.Name1, vnd.Name2, vnd.DefDeliveryAddr, vnd.DefContactCode, vnd.Address, vnd.Telephone, vnd.Fax, vnd.AccountCode, vnd.IDCardNo, vnd.BankAccNo, vnd.CreditDay, vnd.LeadTime, vnd.DefaultTaxRate, vnd.TaxNo, vnd.PicFileName, vnd.TypeCode, vnd.EmailAddress, vnd.GroupCode, vnd.GroupOfDebt, vnd.PersonType, vnd.ActiveStatus, vnd.CreatorCode)
	} else {
		vnd.LastEditorCode = vnd.UserCode

		sql := `Update dbo.bcap set Name1=?, Name2=?, DefDeliveryAddr=?, DefContactCode=?, Address=?, Telephone=?, Fax=?, AccountCode=?, IDCardNo=?, BankAccNo=?, CreditDay=?, LeadTime=?, DefaultTaxRate=?, TaxNo=?, PicFileName=?, TypeCode=?, EmailAddress=?, GroupCode=?, GroupOfDebt=?, PersonType=?, ActiveStatus=?, LastEditorCode=?, LastEditDateT = getdate() where code = ?`
		_, err = db.Exec(sql, vnd.Name1, vnd.Name2, vnd.DefDeliveryAddr, vnd.DefContactCode, vnd.Address, vnd.Telephone, vnd.Fax, vnd.AccountCode, vnd.IDCardNo, vnd.BankAccNo, vnd.CreditDay, vnd.LeadTime, vnd.DefaultTaxRate, vnd.TaxNo, vnd.PicFileName, vnd.TypeCode, vnd.EmailAddress, vnd.GroupCode, vnd.GroupOfDebt, vnd.PersonType, vnd.ActiveStatus, vnd.LastEditorCode, vnd.Code)
	}
	if err != nil {
		fmt.Println("error = ", err.Error())
		return err
	}

	return nil
}




