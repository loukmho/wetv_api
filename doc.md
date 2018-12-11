#API Document wetv

##Table BCItem

###API Name

#####Table BCAr

######API Name

- ค้นหาลูกค้าโดยค้นหาจาก รหัสลูกค้า
- GET 
   - http://localhost:8001/customer?code=a00504
   - response
   
   ```json
   {
   "status": "success",
   "data": {
       "code": "A00504",
       "name_1": "คุณณรงค์ฤทธิ์ อินผูก",
       "name_2": "NARONGRIT INPOOK",
       "bill_address": "9/1 บ.หนองหล่ม ม.4 ถ.เชียงใหม่-ฝาง ต.สันมหาพน อ.แม่แตง จ.เชียงใหม่ 50150 ",
       "work_address": "70/9 บ.ชลประทาน ต.ฟ้าฮ่าม อ.เมือง จ.เชียงใหม่ 50000",
       "telephone": "053471210,019512657",
       "fax": "",
       "account_code": "",
       "id_card_no": "3520800348371",
       "sale_code": "",
       "credit_status": 0,
       "debt_limit_1": 1,
       "debt_limit_2": 0,
       "debt_limit_bal": 1,
       "tax_type": 1,
       "tax_no": "",
       "pic_file_name": "",
       "type_code": "11",
       "group_code": "060",
       "email_address": "",
       "group_of_debt": "10",
       "person_type": 0,
       "birth_day": "1967-03-08T00:00:00Z",
       "cust_age": 38,
       "cust_status": 0,
       "cust_credit_type": 0,
       "depart_code": "",
       "memberid": "2300600270123",
       "active_status": 1,
       "creator_code": "",
       "create_date_time": "2014-07-21T09:15:10.82Z",
       "last_editor_code": "55048",
       "last_edit_date_t": "2018-04-30T05:14:39.25Z",
       "user_code": ""
      }
   }
   
   ```
- ค้นหาลูกค้าโดยค้นหาจาก คำค้นหา
- GET 
   - http://localhost:8001/customers?keyword=คุณณรงค์ฤทธิ์ อินผูก
   - response
   ```json
   {
      "status": "success",
      "data": {
          "code": "A00504",
          "name_1": "คุณณรงค์ฤทธิ์ อินผูก",
          "name_2": "NARONGRIT INPOOK",
          "bill_address": "9/1 บ.หนองหล่ม ม.4 ถ.เชียงใหม่-ฝาง ต.สันมหาพน อ.แม่แตง จ.เชียงใหม่ 50150 ",
          "work_address": "70/9 บ.ชลประทาน ต.ฟ้าฮ่าม อ.เมือง จ.เชียงใหม่ 50000",
          "telephone": "053471210,019512657",
          "fax": "",
          "account_code": "",
          "id_card_no": "3520800348371",
          "sale_code": "",
          "credit_status": 0,
          "debt_limit_1": 1,
          "debt_limit_2": 0,
          "debt_limit_bal": 1,
          "tax_type": 1,
          "tax_no": "",
          "pic_file_name": "",
          "type_code": "11",
          "group_code": "060",
          "email_address": "",
          "group_of_debt": "10",
          "person_type": 0,
          "birth_day": "1967-03-08T00:00:00Z",
          "cust_age": 38,
          "cust_status": 0,
          "cust_credit_type": 0,
          "depart_code": "",
          "memberid": "2300600270123",
          "active_status": 1,
          "creator_code": "",
          "create_date_time": "2014-07-21T09:15:10.82Z",
          "last_editor_code": "55048",
          "last_edit_date_t": "2018-04-30T05:14:39.25Z",
          "user_code": ""
          }
      }
   ```
   
   - บันทึกและแก้ไขข้อมูลลูกค้า
   - POST
       - Request
```json
{
    "code": "A00504",
    "name_1": "คุณณรงค์ฤทธิ์ อินผูก",
    "name_2": "NARONGRIT INPOOK",
    "bill_address": "9/1 บ.หนองหล่ม ม.4 ถ.เชียงใหม่-ฝาง ต.สันมหาพน อ.แม่แตง จ.เชียงใหม่ 50150 ",
    "work_address": "70/9 บ.ชลประทาน ต.ฟ้าฮ่าม อ.เมือง จ.เชียงใหม่ 50000",
    "telephone": "053471210,019512657",
    "fax": "",
    "account_code": "",
    "id_card_no": "3520800348371",
    "credit_status": 0,
    "debt_limit_1": 1,
    "debt_limit_2": 0,
    "debt_limit_bal": 1,
    "tax_type": 1,
    "tax_no": "",
    "pic_file_name": "",
    "type_code": "11",
    "group_code": "060",
    "email_address": "",
    "group_of_debt": "10",
    "person_type": 0,
    "birth_day": "08/12/2530",
    "cust_age": 38,
    "cust_status": 0,
    "cust_credit_type": 0,
    "depart_code": "",
    "memberid": "2300600270123",
    "user_code": "admin"
}

```
#####Table BCArInvoice

######API Name      
- ค้นหาเอกสารบิลขายโดยค้นหาจาก เลขที่เอกสาร
- GET 
   - http://localhost:8001/arinvoice?docno=w01-icv6107-2616
   - response
```json
    {
        "status": "success",
        "data": {
            "save_from": 0,
            "source": 0,
            "doc_no": "W01-ICV6107-2616",
            "doc_date": "2018-07-20T00:00:00Z",
            "tax_no": "W01-ICV6107-2616",
            "tax_date": "",
            "book_code": "",
            "ar_code": "374216",
            "ar_name": "วัดผาลาด",
            "sale_code": "46095",
            "sale_name": "ทิพวรรณ อ้ายเสาร์",
            "creator_code": "tunwaratsp",
            "create_date_time": "2018-07-20T09:55:44Z",
            "last_editor_code": "",
            "last_edit_date_t": "",
            "tax_type": 1,
            "depart_code": "S03-00-00",
            "credit_day": 30,
            "delivery_day": 30,
            "delivery_date": "2018-08-19T00:00:00Z",
            "due_date": "2018-08-19T00:00:00Z",
            "pay_bill_date": "2018-08-19T00:00:00Z",
            "tax_rate": 7,
            "is_confirm": 1,
            "my_description": "1. สินค้าสั่งตัดพิเศษไม่สามารถคืนหรือเปลี่ยนได้ 2. มัดจำค่าสินค้า 50% ก่อนสั่งผลิต",
            "bill_type": 1,
            "bill_group": "",
            "ref_doc_no": "",
            "delivery_addr": "001",
            "contact_code": "",
            "sum_of_item_amount": 106000,
            "discount_word": "",
            "discount_amount": 0,
            "after_discount": 106000,
            "before_tax_amount": 99065.42,
            "tax_amount": 6934.58,
            "total_amount": 106000,
            "zero_tax_amount": 0,
            "except_tax_amount": 0,
            "sum_cash_amount": 0,
            "sum_chq_amount": 0,
            "sum_credit_amount": 0,
            "sum_bank_amount": 0,
            "deposit_inc_tax": 1,
            "sum_of_deposit_1": 0,
            "sum_of_deposit_2": 0,
            "sum_of_w_tax": 0,
            "net_debt_amount": 106000,
            "home_amount": 106000,
            "other_income": 0,
            "other_expense": 0,
            "excess_amount_1": 0,
            "excess_amount_2": 0,
            "bill_balance": 106000,
            "currency_code": "",
            "exchange_rate": 1,
            "gl_format": "B02",
            "is_cancel": 0,
            "is_complete_save": 1,
            "allocate_code": "",
            "project_code": "",
            "recur_name": "",
            "confirm_code": "",
            "confirm_date_time": "",
            "cancel_code": "",
            "cancel_date_time": "",
            "is_condition_send": 1,
            "pay_bill_amount": 106000,
            "so_ref_no": "W03-SCV6107-0060",
            "holding_status": 0,
            "pos_status": 0,
            "credit_base_amount": 0,
            "user_code": "",
            "shiftcode": "",
            "cashier_code": "",
            "shift_no": "",
            "machine_no": "",
            "machine_code": "",
            "bill_time": "",
            "coupong_amount": 0,
            "change_amount": 0,
            "charge_amount": 0,
            "grand_total": 0,
            "credit_type": "",
            "cofirm_no": "",
            "credit_no": "",
            "credit_due_date": "1900-01-01T00:00:00Z",
            "bank_code": "",
            "bank_branch_code": "",
            "bank_ref_no": "",
            "trans_bank_date": "",
            "ref_date": "",
            "subs": [
              {
                    "my_type": 4,
                    "item_code": "3050091",
                    "my_description": "รุ่น 18MFF(S) 181 ชนิดกันน้ำ สีพิเศษลายไม้แนวนอน WP-45",
                    "item_name": "ผนังห้องน้ำสำเร็จรูป Willy รุ่น Willy 18 MFF",
                    "wh_code": "S1-SPO",
                    "shelf_code": "-",
                    "cn_qty": 10,
                    "qty": 10,
                    "price": 10600,
                    "discount_word": "",
                    "discount_amount": 0,
                    "amount": 106000,
                    "net_amount": 99065.42,
                    "home_amount": 99065.42,
                    "sum_of_cost": 92000,
                    "balance_amount": 106000,
                    "unit_code": "ห้อง",
                    "so_ref_no": "W03-SCV6107-0060",
                    "po_ref_no": "",
                    "line_number": 0,
                    "ref_line_number": 0,
                    "is_cancel": 0,
                    "bar_code": "",
                    "posstatus": 0,
                    "is_condition_send": 0,
                    "averagecost": 9200,
                    "lot_number": "B02",
                    "stock_type": 0,
                    "packing_rate_1": 1,
                    "packing_rate_2": 1
            }
            ],
            "deps": null,
            "cdcs": null,
            "chqs": null
        }
    }
```   
- ค้นหาเอกสารบิลขายโดยค้นหาจาก คำค้นหา
- GET 
   - http://localhost:8001/arinvoices?keyword=วัดผาลาด
   - response
```json
    {
        "status": "success",
        "data": {
            "save_from": 0,
            "source": 0,
            "doc_no": "W01-ICV6107-2616",
            "doc_date": "2018-07-20T00:00:00Z",
            "tax_no": "W01-ICV6107-2616",
            "tax_date": "",
            "book_code": "",
            "ar_code": "374216",
            "ar_name": "วัดผาลาด",
            "sale_code": "46095",
            "sale_name": "ทิพวรรณ อ้ายเสาร์",
            "creator_code": "tunwaratsp",
            "create_date_time": "2018-07-20T09:55:44Z",
            "last_editor_code": "",
            "last_edit_date_t": "",
            "tax_type": 1,
            "depart_code": "S03-00-00",
            "credit_day": 30,
            "delivery_day": 30,
            "delivery_date": "2018-08-19T00:00:00Z",
            "due_date": "2018-08-19T00:00:00Z",
            "pay_bill_date": "2018-08-19T00:00:00Z",
            "tax_rate": 7,
            "is_confirm": 1,
            "my_description": "1. สินค้าสั่งตัดพิเศษไม่สามารถคืนหรือเปลี่ยนได้ 2. มัดจำค่าสินค้า 50% ก่อนสั่งผลิต",
            "bill_type": 1,
            "bill_group": "",
            "ref_doc_no": "",
            "delivery_addr": "001",
            "contact_code": "",
            "sum_of_item_amount": 106000,
            "discount_word": "",
            "discount_amount": 0,
            "after_discount": 106000,
            "before_tax_amount": 99065.42,
            "tax_amount": 6934.58,
            "total_amount": 106000,
            "zero_tax_amount": 0,
            "except_tax_amount": 0,
            "sum_cash_amount": 0,
            "sum_chq_amount": 0,
            "sum_credit_amount": 0,
            "sum_bank_amount": 0,
            "deposit_inc_tax": 1,
            "sum_of_deposit_1": 0,
            "sum_of_deposit_2": 0,
            "sum_of_w_tax": 0,
            "net_debt_amount": 106000,
            "home_amount": 106000,
            "other_income": 0,
            "other_expense": 0,
            "excess_amount_1": 0,
            "excess_amount_2": 0,
            "bill_balance": 106000,
            "currency_code": "",
            "exchange_rate": 1,
            "gl_format": "B02",
            "is_cancel": 0,
            "is_complete_save": 1,
            "allocate_code": "",
            "project_code": "",
            "recur_name": "",
            "confirm_code": "",
            "confirm_date_time": "",
            "cancel_code": "",
            "cancel_date_time": "",
            "is_condition_send": 1,
            "pay_bill_amount": 106000,
            "so_ref_no": "W03-SCV6107-0060",
            "holding_status": 0,
            "pos_status": 0,
            "credit_base_amount": 0,
            "user_code": "",
            "shiftcode": "",
            "cashier_code": "",
            "shift_no": "",
            "machine_no": "",
            "machine_code": "",
            "bill_time": "",
            "coupong_amount": 0,
            "change_amount": 0,
            "charge_amount": 0,
            "grand_total": 0,
            "credit_type": "",
            "cofirm_no": "",
            "credit_no": "",
            "credit_due_date": "1900-01-01T00:00:00Z",
            "bank_code": "",
            "bank_branch_code": "",
            "bank_ref_no": "",
            "trans_bank_date": "",
            "ref_date": "",
            "subs": [
              {
                    "my_type": 4,
                    "item_code": "3050091",
                    "my_description": "รุ่น 18MFF(S) 181 ชนิดกันน้ำ สีพิเศษลายไม้แนวนอน WP-45",
                    "item_name": "ผนังห้องน้ำสำเร็จรูป Willy รุ่น Willy 18 MFF",
                    "wh_code": "S1-SPO",
                    "shelf_code": "-",
                    "cn_qty": 10,
                    "qty": 10,
                    "price": 10600,
                    "discount_word": "",
                    "discount_amount": 0,
                    "amount": 106000,
                    "net_amount": 99065.42,
                    "home_amount": 99065.42,
                    "sum_of_cost": 92000,
                    "balance_amount": 106000,
                    "unit_code": "ห้อง",
                    "so_ref_no": "W03-SCV6107-0060",
                    "po_ref_no": "",
                    "line_number": 0,
                    "ref_line_number": 0,
                    "is_cancel": 0,
                    "bar_code": "",
                    "posstatus": 0,
                    "is_condition_send": 0,
                    "averagecost": 9200,
                    "lot_number": "B02",
                    "stock_type": 0,
                    "packing_rate_1": 1,
                    "packing_rate_2": 1
            }
            ],
            "deps": null,
            "cdcs": null,
            "chqs": null
        }
    }
```   

- บันทึกและแก้ไขข้อมูลเอกสารบิลขาย 
- POST
    - Request
```json
{
        "doc_no": "CDO-610401-0001",
        "doc_date": "02/08/2018",
        "ar_code": "45040",
        "sale_code": "01",
        "user_code": "admin",
        "tax_type": 1,
        "depart_code": "SALE",
        "my_description": "Test",
        "bill_type": 0,
        "ref_doc_no": "S01-SHV6104-0001",
        "sum_of_item_amount": 2500,
        "discount_word": "",
        "discount_amount": 0,
        "after_discount": 2500,
        "sum_cash_amount": 300,
        "sum_chq_amount": 800,
        "sum_credit_amount": 500,
        "sum_bank_amount": 700,
        "sum_of_deposit_1": 200,
        "so_ref_no": "600107-0008",
        "credit_type": "MASTERCARD",
		"confirm_no":"1133",
		"credit_ref_no":"6721",
		"bank_code":"BBL",
		"bank_branch_code":"01",
		"credit_ref_no":"5678",
		"bank_ref_no":"533-0-03046-3",
        "subs": [
            {
                "my_type": 4,
                "item_code": "S20027",
                "my_description": "",
                "item_name": "พลาสติกโคน 9x10 มิล",
                "wh_code": "01",
                "shelf_code": "01",
                "qty": 500,
                "price": 7.75,
                "discount_word": "",
                "discount_amount": 0,
                "amount": 3875,
                "unit_code": "ตัว",
                "so_ref_no": "SO600107-0008",
                "ref_line_number": 0,
                "bar_code": "",
                "is_condition_send": 0,
                "packing_rate_1": 3
            },
            {
                "my_type": 4,
                "item_code": "S20028",
                "my_description": "",
                "item_name": "พลาสติกโคน 9x10 มิล",
                "wh_code": "01",
                "shelf_code": "01",
                "qty": 5,
                "price": 7.00,
                "discount_word": "",
                "discount_amount": 0,
                "amount": 35,
                "unit_code": "ตัว",
                "so_ref_no": "SO600107-0008",
                "ref_line_number": 0,
                "bar_code": "8852437100080"
            }
        ],
        "deps":[
		        {
		        	"deposit_no":"DEP-Test-0001",
					"balance":100,
					"amount":100,
					"net_amount":100
		        },
		        		        {
		        	"deposit_no":"DEP-Test-0002",
					"balance":100,
					"amount":100,
					"net_amount":100
		        }
		],
	     "chqs":[{
			"chq_number":"55457",
			"bank_code":"BBL",      
			"bank_branch_code":"01",   
			"amount":1000       
		  }]
    }
```
#####Table BCReceipt1

######API Name  

- ค้นหาเอกสารรับชำระโดยค้นหาจาก เลขที่เอกสาร
- GET 
   - localhost:8001/receipt1?docno=W01-RE6109-0021
   - response
```json
  
  {
      "status": "success",
      "data": {
          "source": 0,
          "save_from": 0,
          "doc_no": "W01-RE6109-0021",
          "tax_no": "",
          "tax_date": "",
          "book_code": "",
          "tax_rate": 0,
          "doc_date": "2018-09-01T00:00:00Z",
          "creator_code": "58112",
          "create_date_time": "2018-09-01T10:47:22Z",
          "last_editor_code": "",
          "last_edit_date_t": "1900-01-01T00:00:00Z",
          "ar_code": "027403604",
          "sale_code": "35026",
          "depart_code": "HO-FN-CD",
          "sum_of_invoice": 43770,
          "sum_of_debit_note": 0,
          "sum_of_credit_note": 0,
          "before_tax_amount": 0,
          "tax_amount": 0,
          "total_amount": 43770,
          "discount_amount": 0,
          "net_amount": 43770,
          "sum_cash_amount": 0,
          "sum_chq_amount": 0,
          "sum_credit_amount": 0,
          "charge_amount": 0,
          "change_amount": 0,
          "sum_bank_amount": 43707,
          "sum_of_w_tax": 0,
          "confirm_no": "",
          "credit_no": "",
          "gl_format": "R01",
          "is_cancel": 0,
          "my_description": "",
          "allocate_code": "",
          "project_code": "",
          "bill_group": "",
          "is_complete_save": 1,
          "sum_home_amount_1": 43770,
          "sum_home_amount_2": 43770,
          "debt_limit_return": 43770,
          "currency_code": "",
          "exchange_rate": 1,
          "other_income": 0,
          "other_expense": 63,
          "excess_amount_1": 0,
          "excess_amount_2": 0,
          "is_confirm": 1,
          "recur_name": "",
          "confirm_code": "",
          "confirm_date_time": "1900-01-01T00:00:00Z",
          "cancel_code": "",
          "cancel_date_time": "1900-01-01T00:00:00Z",
          "user_code": "",
          "credit_type": "",
          "cofirm_no": "",
          "credit_due_date": "",
          "bank_code": "",
          "bank_branch_code": "",
          "bank_ref_no": "",
          "trans_bank_date": "",
          "ref_date": "",
          "subs": [
            {
          "invoice_date": "2018-06-15T00:00:00Z",
          "invoice_no": "W01-ICV6106-1761",
          "inv_balance": 35190,
          "pay_amount": 35190,
          "debt_limit_return": 35190,
          "my_description": "ใบวางบิล",
          "is_cancel": 0,
          "line_number": 0,
          "bill_type": 0,
          "ref_no": "W01-BI6106-0744",
          "ref_type": 0,
          "home_amount_1": 35190,
          "home_amount_2": 35190
          },
            {
          "invoice_date": "2018-06-26T00:00:00Z",
          "invoice_no": "W01-ICV6106-3046",
          "inv_balance": 4290,
          "pay_amount": 4290,
          "debt_limit_return": 4290,
          "my_description": "ใบวางบิล",
          "is_cancel": 0,
          "line_number": 1,
          "bill_type": 0,
          "ref_no": "W01-BI6106-0744",
          "ref_type": 0,
          "home_amount_1": 4290,
          "home_amount_2": 4290
          },
            {
          "invoice_date": "2018-06-27T00:00:00Z",
          "invoice_no": "W01-ICV6106-3158",
          "inv_balance": 4290,
          "pay_amount": 4290,
          "debt_limit_return": 4290,
          "my_description": "ใบวางบิล",
          "is_cancel": 0,
          "line_number": 2,
          "bill_type": 0,
          "ref_no": "W01-BI6106-0744",
          "ref_type": 0,
          "home_amount_1": 4290,
          "home_amount_2": 4290
          }
          ],
          "cdcs": null,
          "chqs": null
          }
      }

  ```
   
- ค้นหาเอกสารรับชำระโดยค้นหาจาก คำค้นหา
- GET 
   - localhost:8001/receipt1s?keycode=W01-RE6109-0021
   - response
```json
    {
    "status": "success",
    "data": {
        "source": 0,
        "save_from": 0,
        "doc_no": "W01-RE6109-0021",
        "tax_no": "",
        "tax_date": "",
        "book_code": "",
        "tax_rate": 0,
        "doc_date": "2018-09-01T00:00:00Z",
        "creator_code": "58112",
        "create_date_time": "2018-09-01T10:47:22Z",
        "last_editor_code": "",
        "last_edit_date_t": "1900-01-01T00:00:00Z",
        "ar_code": "027403604",
        "sale_code": "35026",
        "depart_code": "HO-FN-CD",
        "sum_of_invoice": 43770,
        "sum_of_debit_note": 0,
        "sum_of_credit_note": 0,
        "before_tax_amount": 0,
        "tax_amount": 0,
        "total_amount": 43770,
        "discount_amount": 0,
        "net_amount": 43770,
        "sum_cash_amount": 0,
        "sum_chq_amount": 0,
        "sum_credit_amount": 0,
        "charge_amount": 0,
        "change_amount": 0,
        "sum_bank_amount": 43707,
        "sum_of_w_tax": 0,
        "confirm_no": "",
        "credit_no": "",
        "gl_format": "R01",
        "is_cancel": 0,
        "my_description": "",
        "allocate_code": "",
        "project_code": "",
        "bill_group": "",
        "is_complete_save": 1,
        "sum_home_amount_1": 43770,
        "sum_home_amount_2": 43770,
        "debt_limit_return": 43770,
        "currency_code": "",
        "exchange_rate": 1,
        "other_income": 0,
        "other_expense": 63,
        "excess_amount_1": 0,
        "excess_amount_2": 0,
        "is_confirm": 1,
        "recur_name": "",
        "confirm_code": "",
        "confirm_date_time": "1900-01-01T00:00:00Z",
        "cancel_code": "",
        "cancel_date_time": "1900-01-01T00:00:00Z",
        "user_code": "",
        "credit_type": "",
        "cofirm_no": "",
        "credit_due_date": "",
        "bank_code": "",
        "bank_branch_code": "",
        "bank_ref_no": "",
        "trans_bank_date": "",
        "ref_date": "",
        "subs": [
          {
        "invoice_date": "2018-06-15T00:00:00Z",
        "invoice_no": "W01-ICV6106-1761",
        "inv_balance": 35190,
        "pay_amount": 35190,
        "debt_limit_return": 35190,
        "my_description": "ใบวางบิล",
        "is_cancel": 0,
        "line_number": 0,
        "bill_type": 0,
        "ref_no": "W01-BI6106-0744",
        "ref_type": 0,
        "home_amount_1": 35190,
        "home_amount_2": 35190
        },
          {
        "invoice_date": "2018-06-26T00:00:00Z",
        "invoice_no": "W01-ICV6106-3046",
        "inv_balance": 4290,
        "pay_amount": 4290,
        "debt_limit_return": 4290,
        "my_description": "ใบวางบิล",
        "is_cancel": 0,
        "line_number": 1,
        "bill_type": 0,
        "ref_no": "W01-BI6106-0744",
        "ref_type": 0,
        "home_amount_1": 4290,
        "home_amount_2": 4290
        },
          {
        "invoice_date": "2018-06-27T00:00:00Z",
        "invoice_no": "W01-ICV6106-3158",
        "inv_balance": 4290,
        "pay_amount": 4290,
        "debt_limit_return": 4290,
        "my_description": "ใบวางบิล",
        "is_cancel": 0,
        "line_number": 2,
        "bill_type": 0,
        "ref_no": "W01-BI6106-0744",
        "ref_type": 0,
        "home_amount_1": 4290,
        "home_amount_2": 4290
        }
        ],
        "cdcs": null,
        "chqs": null
        }
    }
```
#####Table BCItem

######API Name  
- ค้นหาสินค้าโดยค้นหาจาก รหัสสินค้า
- GET 
   - localhost:8001/item?code=2120250
   - response
```json
   {
   "status": "success",
   "data": {
       "code": "2120250",
       "name_1": "น้ำยาเชื่อมท่อ PVC 250 กรัม",
       "name_2": "PVC 250 g",
       "creator_code": "Satit",
       "create_date_time": "1900-01-01T00:00:00Z",
       "last_editor_code": "59218",
       "last_edit_date_t": "2018-05-02T12:31:54Z",
       "short_name": "",
       "category_code": "7050800",
       "group_code": "70",
       "brand_code": "062",
       "type_code": "7050",
       "format_code": "705080010",
       "color_code": "",
       "unit_type": 1,
       "def_stk_unit_code": "กระป๋อง",
       "def_sale_unit_code": "กระป๋อง",
       "def_buy_unit_code": "กระป๋อง",
       "stock_type": 0,
       "item_status": 1,
       "order_point": 500,
       "stock_min": 200,
       "stock_max": 800,
       "stock_qty": 536,
       "average_cost": 92.114,
       "pic_file_name_1": "http://qserver.nopadol.com/pictures/item/2120250.jpg",
       "pic_file_name_2": "",
       "my_description": "กล่องบรรจุ 20 กระป๋อง",
       "def_buy_wh_code": "S1-B",
       "def_buy_shelf": "-",
       "def_sale_wh_code": "S1-B",
       "def_sale_shelf": "-",
       "active_status": 1,
       "sale_price_1": 116,
       "sale_price_2": 0,
       "sale_price_3": 0,
       "sale_price_4": 0,
       "def_fix_unit_code": "กระป๋อง",
       "user_code": ""
     }
   }
```

- ค้นหาสินค้าค้า โดยค้นหาจากคำค้น
- GET
   - localhost:8001/items?keyword=น้ำยาเชื่อมท่อ 
   - reponse 
```json
   {
      "status": "success",
      "data": {
          "code": "2120250",
          "name_1": "น้ำยาเชื่อมท่อ PVC 250 กรัม",
          "name_2": "PVC 250 g",
          "creator_code": "Satit",
          "create_date_time": "1900-01-01T00:00:00Z",
          "last_editor_code": "59218",
          "last_edit_date_t": "2018-05-02T12:31:54Z",
          "short_name": "",
          "category_code": "7050800",
          "group_code": "70",
          "brand_code": "062",
          "type_code": "7050",
          "format_code": "705080010",
          "color_code": "",
          "unit_type": 1,
          "def_stk_unit_code": "กระป๋อง",
          "def_sale_unit_code": "กระป๋อง",
          "def_buy_unit_code": "กระป๋อง",
          "stock_type": 0,
          "item_status": 1,
          "order_point": 500,
          "stock_min": 200,
          "stock_max": 800,
          "stock_qty": 536,
          "average_cost": 92.114,
          "pic_file_name_1": "http://qserver.nopadol.com/pictures/item/2120250.jpg",
          "pic_file_name_2": "",
          "my_description": "กล่องบรรจุ 20 กระป๋อง",
          "def_buy_wh_code": "S1-B",
          "def_buy_shelf": "-",
          "def_sale_wh_code": "S1-B",
          "def_sale_shelf": "-",
          "active_status": 1,
          "sale_price_1": 116,
          "sale_price_2": 0,
          "sale_price_3": 0,
          "sale_price_4": 0,
          "def_fix_unit_code": "กระป๋อง",
          "user_code": ""
        }
      }
```      
- บันทึกและแก้ไขข้อมูลสินค้า
- POST
    - Request
```json
{
   	"Code":"99999",
   	"Name1":"xxxx",
   	"Name2":"xxxx",
   	"ShortName":"xxx",
   	"CategoryCode":"",
   	"GroupCode":"",
   	"BrandCode":"",
   	"TypeCode":"",
   	"FormatCode":"",
   	"ColorCode":"",
   	"UnitType":0,
   	"DefStkUnitCode":"xxx",
   	"DefSaleUnitCode":"xx",
   	"DefBuyUnitCode":"xx",
   	"StockType":0,
   	"ItemStatus":1,
   	"OrderPoint":0,
   	"StockMin":0,
   	"StockMax":0,
   	"StockQty":0,
   	"PicFileName1":"xxx",
   	"PicFileName2":"",
    "MyDescription":"",
   	"DefBuyWHCode":"xxx",
   	"DefBuyShelf":"xxx",
   	"DefSaleWHCode":"xxx",
   	"DefSaleShelf":"xxx",
   	"SalePrice1":100,
   	"SalePrice2":120,
    "SalePrice3":0,
   	"SalePrice4":0,
   	"UserCode":"admin"
   	}
```
#####Table BCAp

######API Name

- ค้นหาเจ้าหนี้โดยค้นหาจาก รหัสเจ้าหนี้
- GET 
   - http://localhost:8001/vendor?code=a00504
   - response
```json
    {
     "status": "success",
    "data": {
        "code": "036240000",
        "name_1": "บริษัท สยามมอร์ตาร์ จำกัด",
        "name_2": "บริษัท สยามมอร์ตาร์ จำกัด",
        "def_delivery_addr": "",
        "def_contact_code": "",
        "address": "110 หมู่ 1 ต.บ้านป่า อ.แก่งคอย จ.สระบุรี 18110",
        "telephone": "036-240000",
        "fax": "036-240083",
        "account_code": "",
        "id_card_no": "",
        "bank_acc_no": "",
        "credit_day": 7,
        "default_tax_rate": 7,
        "pic_file_name": "",
        "type_code": "92",
        "email_address": "",
        "group_code": "06",
        "group_of_debt": "",
        "person_type": 1,
        "active_status": 1,
        "creator_code": "kanjirapt",
        "create_date_time": "2009-05-25T00:00:00Z",
        "last_editor_code": "",
        "last_edit_date_t": "1900-01-01T00:00:00Z",
        "user_code": ""
        }
    }
```
- ค้นหาเจ้าหนี้โดยค้นหาจาก คำค้นหา
- GET 
   - http://localhost:8001/vendors?keyword=สยามมอร์ตาร์ 
   - response
```json
    {
     "status": "success",
    "data": {
        "code": "036240000",
        "name_1": "บริษัท สยามมอร์ตาร์ จำกัด",
        "name_2": "บริษัท สยามมอร์ตาร์ จำกัด",
        "def_delivery_addr": "",
        "def_contact_code": "",
        "address": "110 หมู่ 1 ต.บ้านป่า อ.แก่งคอย จ.สระบุรี 18110",
        "telephone": "036-240000",
        "fax": "036-240083",
        "account_code": "",
        "id_card_no": "",
        "bank_acc_no": "",
        "credit_day": 7,
        "default_tax_rate": 7,
        "pic_file_name": "",
        "type_code": "92",
        "email_address": "",
        "group_code": "06",
        "group_of_debt": "",
        "person_type": 1,
        "active_status": 1,
        "creator_code": "kanjirapt",
        "create_date_time": "2009-05-25T00:00:00Z",
        "last_editor_code": "",
        "last_edit_date_t": "1900-01-01T00:00:00Z",
        "user_code": ""
        }
    }
```

- บันทึกและแก้ไขข้อมูลเจ้าหนี้
- POST
    - Request
```json
{
   	"code": "036240000",
    "name_1": "บริษัท สยามมอร์ตาร์ จำกัด",
    "name_2": "บริษัท สยามมอร์ตาร์ จำกัด",
    "def_delivery_addr": "",
    "def_contact_code": "",
    "address": "110 หมู่ 1 ต.บ้านป่า อ.แก่งคอย จ.สระบุรี 18110",
    "telephone": "036-240000",
    "fax": "036-240083",
    "account_code": "xxx",
    "id_card_no": "xxx",
    "bank_acc_no": "",
    "credit_day": 7,
    "default_tax_rate": 7,
    "pic_file_name": "",
    "type_code": "92",
    "email_address": "",
    "group_code": "06",
    "group_of_debt": "",
    "person_type": 1,
    "active_status": 1,
    "user_code": "admin"
   	}
```
#####Table BCApInvoice

######API Name

- ค้นหาเอกสารซื้อสินค้าโดยค้นหาจาก เลขที่เอกสาร
- GET 
   - http://localhost:8001/apinvoice?docno=S01-RV6105-0796
   - response
 ```json
        {
        "status": "success",
        "data": {
        "save_from": 0,
        "source": 0,
        "doc_no": "S01-RV6105-0796",
        "tax_no": "",
        "book_code": "",
        "doc_date": "2018-05-19T00:00:00Z",
        "tax_type": 0,
        "ap_code": "270063",
        "ap_name": "",
        "depart_code": "HO-MC-01",
        "credit_day": 60,
        "due_date": "2018-07-18T00:00:00Z",
        "statement_date": "1900-01-01T00:00:00Z",
        "tax_rate": 7,
        "is_confirm": 1,
        "my_description": "",
        "bill_type": 1,
        "bill_group": "",
        "contact_code": "217",
        "sum_of_item_amount": 20800,
        "discount_word": "",
        "discount_amount": 0,
        "after_discount": 20800,
        "before_tax_amount": 20800,
        "tax_amount": 1456,
        "total_amount": 22256,
        "except_tax_amount": 0,
        "zero_tax_amount": 0,
        "petty_cash_amount": 0,
        "sum_cash_amount": 0,
        "sum_chq_amount": 0,
        "sum_bank_amount": 0,
        "deposit_inc_tax": 1,
        "sum_of_deposit_1": 0,
        "sum_of_deposit_2": 0,
        "sum_of_w_tax": 0,
        "net_debt_amount": 22256,
        "home_amount": 22256,
        "other_income": 0,
        "other_expense": 0,
        "excess_amount_1": 0,
        "excess_amount_2": 0,
        "bill_balance": 22256,
        "currency_code": "",
        "exchange_rate": 1,
        "gl_format": "P02-3",
        "gl_start_posting": 1,
        "gr_bill_status": 1,
        "grir_bill_status": 1,
        "is_cancel": 0,
        "is_credit_note": 0,
        "is_debit_note": 0,
        "is_complete_save": 1,
        "allocate_code": "",
        "project_code": "",
        "recur_name": "",
        "exchange_profit": 0,
        "confirm_code": "",
        "confirm_date_time": "1900-01-01T00:00:00Z",
        "cancel_code": "",
        "cancel_date_time": "1900-01-01T00:00:00Z",
        "pay_bill_amount": 0,
        "ref_doc_no": "",
        "creator_code": "55044",
        "create_date_time": "2018-05-19T14:30:56Z",
        "last_editor_code": "",
        "last_edit_date_t": "1900-01-01T00:00:00Z",
        "user_code": "",
        "credit_type": "",
        "confirm_no": "",
        "credit_ref_no": "",
        "credit_due_date": "",
        "bank_code": "",
        "bank_branch_code": "",
        "bank_ref_no": "",
        "trans_bank_date": "",
        "ref_date": "",
        "subs": [
          {
        "my_type": 2,
        "stock_type": 0,
        "item_code": "8852437300060",
        "my_description": "",
        "item_name": "ปูนฉาบทั่วไปสำเร็จรูปถุง 50 กก.(เสือมอร์ตาร์)ถุงเขียว",
        "wh_code": "S0-PASS",
        "shelf_code": "-",
        "cn_qty": 16,
        "gr_remain_qty": 0,
        "qty": 16,
        "price": 1300,
        "discount_word": "",
        "discount_amount": 0,
        "amount": 20800,
        "net_amount": 20800,
        "home_amount": 20800,
        "balance_amount": 20800,
        "sum_of_exp_cost": 0,
        "unit_code": "ตัน-20",
        "po_ref_no": "3381321361",
        "ir_ref_no": "",
        "is_cancel": 0,
        "line_number": 0,
        "bar_code": "8852437300060",
        "po_ref_linenum": 0,
        "averagecost": 65,
        "lot_number": "A02",
        "sum_of_cost": 20800,
        "packing_rate_1": 20,
        "packing_rate_2": 1
        }
        ],
        "deps": null,
        "cdcs": null,
        "chqs": null
        }
       }
```
- ค้นหาเอกสารซื้อสินค้าโดยค้นหาจาก คำค้นหา
- GET 
   - http://localhost:8001/apinvoices?keyword=S01-RV6105-0796
   - response
 ```json
        {
        "status": "success",
        "data": {
        "save_from": 0,
        "source": 0,
        "doc_no": "S01-RV6105-0796",
        "tax_no": "",
        "book_code": "",
        "doc_date": "2018-05-19T00:00:00Z",
        "tax_type": 0,
        "ap_code": "270063",
        "ap_name": "",
        "depart_code": "HO-MC-01",
        "credit_day": 60,
        "due_date": "2018-07-18T00:00:00Z",
        "statement_date": "1900-01-01T00:00:00Z",
        "tax_rate": 7,
        "is_confirm": 1,
        "my_description": "",
        "bill_type": 1,
        "bill_group": "",
        "contact_code": "217",
        "sum_of_item_amount": 20800,
        "discount_word": "",
        "discount_amount": 0,
        "after_discount": 20800,
        "before_tax_amount": 20800,
        "tax_amount": 1456,
        "total_amount": 22256,
        "except_tax_amount": 0,
        "zero_tax_amount": 0,
        "petty_cash_amount": 0,
        "sum_cash_amount": 0,
        "sum_chq_amount": 0,
        "sum_bank_amount": 0,
        "deposit_inc_tax": 1,
        "sum_of_deposit_1": 0,
        "sum_of_deposit_2": 0,
        "sum_of_w_tax": 0,
        "net_debt_amount": 22256,
        "home_amount": 22256,
        "other_income": 0,
        "other_expense": 0,
        "excess_amount_1": 0,
        "excess_amount_2": 0,
        "bill_balance": 22256,
        "currency_code": "",
        "exchange_rate": 1,
        "gl_format": "P02-3",
        "gl_start_posting": 1,
        "gr_bill_status": 1,
        "grir_bill_status": 1,
        "is_cancel": 0,
        "is_credit_note": 0,
        "is_debit_note": 0,
        "is_complete_save": 1,
        "allocate_code": "",
        "project_code": "",
        "recur_name": "",
        "exchange_profit": 0,
        "confirm_code": "",
        "confirm_date_time": "1900-01-01T00:00:00Z",
        "cancel_code": "",
        "cancel_date_time": "1900-01-01T00:00:00Z",
        "pay_bill_amount": 0,
        "ref_doc_no": "",
        "creator_code": "55044",
        "create_date_time": "2018-05-19T14:30:56Z",
        "last_editor_code": "",
        "last_edit_date_t": "1900-01-01T00:00:00Z",
        "user_code": "",
        "credit_type": "",
        "confirm_no": "",
        "credit_ref_no": "",
        "credit_due_date": "",
        "bank_code": "",
        "bank_branch_code": "",
        "bank_ref_no": "",
        "trans_bank_date": "",
        "ref_date": "",
        "subs": [
          {
        "my_type": 2,
        "stock_type": 0,
        "item_code": "8852437300060",
        "my_description": "",
        "item_name": "ปูนฉาบทั่วไปสำเร็จรูปถุง 50 กก.(เสือมอร์ตาร์)ถุงเขียว",
        "wh_code": "S0-PASS",
        "shelf_code": "-",
        "cn_qty": 16,
        "gr_remain_qty": 0,
        "qty": 16,
        "price": 1300,
        "discount_word": "",
        "discount_amount": 0,
        "amount": 20800,
        "net_amount": 20800,
        "home_amount": 20800,
        "balance_amount": 20800,
        "sum_of_exp_cost": 0,
        "unit_code": "ตัน-20",
        "po_ref_no": "3381321361",
        "ir_ref_no": "",
        "is_cancel": 0,
        "line_number": 0,
        "bar_code": "8852437300060",
        "po_ref_linenum": 0,
        "averagecost": 65,
        "lot_number": "A02",
        "sum_of_cost": 20800,
        "packing_rate_1": 20,
        "packing_rate_2": 1
        }
        ],
        "deps": null,
        "cdcs": null,
        "chqs": null
        }
       }
```
- บันทึกและแก้ไขข้อมูลซื้อสินค้า
- POST
    - Request
 ```json
{
       "doc_no": "S01-RV6105-0796",
       "doc_date": "18/05/2018",
       "tax_type": 0,
       "ap_code": "270063",
       "ap_name": "",
       "depart_code": "HO-MC-01",
       "credit_day": 60,
       "due_date": "2018-07-18",
       "statement_date": "",
       "tax_rate": 7,
       "my_description": "",
       "bill_type": 1,
       "contact_code": "217",
       "sum_of_item_amount": 20800,
       "discount_word": "",
       "discount_amount": 0,
       "after_discount": 20800,
       "before_tax_amount": 20800,
       "tax_amount": 1456,
       "total_amount": 22256,
       "except_tax_amount": 0,
       "zero_tax_amount": 0,
       "petty_cash_amount": 0,
       "sum_cash_amount": 0,
       "sum_chq_amount": 0,
       "sum_bank_amount": 0,
       "deposit_inc_tax": 1,
       "sum_of_deposit_1": 0,
       "sum_of_deposit_2": 0,
       "sum_of_w_tax": 0,
       "net_debt_amount": 22256,
       "home_amount": 22256,
       "other_income": 0,
       "other_expense": 0,
       "excess_amount_1": 0,
       "excess_amount_2": 0,
       "bill_balance": 22256,
       "allocate_code": "",
       "project_code": "",
       "recur_name": "",
       "ref_doc_no": "",
       "user_code": "admin",
       "subs": [
         {
       "item_code": "8852437300060",
       "my_description": "",
       "item_name": "ปูนฉาบทั่วไปสำเร็จรูปถุง 50 กก.(เสือมอร์ตาร์)ถุงเขียว",
       "wh_code": "S0-PASS",
       "shelf_code": "-",
       "qty": 16,
       "price": 1300,
       "discount_word": "",
       "discount_amount": 0,
       "amount": 20800,
       "unit_code": "ตัน-20",
       "po_ref_no": "3381321361",
       "is_cancel": 0,
       "line_number": 0,
       "po_ref_linenum": 0,
       "packing_rate_1": 20
       }
       ]
    }
```
#####Table BCPayment

######API Name

- ค้นหาเอกสารจ่ายชำระโดยค้นหาจาก เลขที่เอกสาร
- GET 
   - http://localhost:8001/payment?docno=S01-PX6112-0030
   - response
```json
{
    "status": "success",
    "data": {
        "save_from": 0,
        "source": 0,
        "doc_no": "S01-PX6112-0030",
        "tax_no": "",
        "tax_date": "",
        "book_code": "",
        "doc_date": "2018-12-06T00:00:00Z",
        "ap_code": "053904044",
        "creator_code": "57147",
        "create_date_time": "2018-12-06T12:07:39Z",
        "last_editor_code": "",
        "last_edit_date_t": "1900-01-01T00:00:00Z",
        "depart_code": "HO-FN-00",
        "sum_of_invoice": 56175,
        "sum_of_debit_note": 0,
        "sum_of_credit_note": 0,
        "before_tax_amount": 0,
        "tax_amount": 0,
        "total_amount": 56175,
        "discount_amount": 0,
        "net_amount": 54600,
        "petty_cash_amount": 0,
        "sum_cash_amount": 0,
        "sum_chq_amount": 54600,
        "sum_bank_amount": 0,
        "sum_of_w_tax": 1575,
        "gl_format": "P07",
        "is_cancel": 0,
        "my_description": "ค่าบริการรักษาความปลอดภัย รปภ. 3 นาย สาขาสันกำแพง (อัตรา 17,500 บาท/คน) ประจำเดือน พ.ย.61 งวดที่ 7/12 (สัญญาเลขที่ CGG6104-006 ครบวันที่ 30/4/2562) ตามใบวางบิลเลขที่ GCMTBC6111-0008 ลว.30/10/61",
        "allocate_code": "",
        "project_code": "",
        "bill_group": "",
        "sum_home_amount_1": 56175,
        "sum_home_amount_2": 56175,
        "other_income": 0,
        "other_expense": 0,
        "excess_amount_1": 0,
        "excess_amount_2": 0,
        "currency_code": "",
        "exchange_rate": 1,
        "is_complete_save": 1,
        "is_confirm": 1,
        "recur_name": "",
        "confirm_code": "",
        "confirm_date_time": "1900-01-01T00:00:00Z",
        "cancel_code": "",
        "cancel_date_time": "1900-01-01T00:00:00Z",
        "sum_of_deposit_1": 0,
        "sum_of_deposit_2": 0,
        "home_amount_inv": 56175,
        "home_amount_debt": 0,
        "home_amount_crd": 0,
        "user_code": "",
        "credit_type": "",
        "confirm_no": "",
        "credit_ref_no": "",
        "credit_due_date": "",
        "bank_code": "",
        "bank_branch_code": "",
        "bank_ref_no": "",
        "trans_bank_date": "",
        "ref_date": "",
        "deps": null,
        "chqs": null,
        "subs": [
          {
        "invoice_date": "2018-11-30T00:00:00Z",
        "invoice_no": "AT6111-0034",
        "inv_balance": 56175,
        "pay_amount": 56175,
        "my_description": "ใบรับวางบิล",
        "bill_type": 4,
        "paybill_no": "S01-AIX6112-0013",
        "ref_type": 0,
        "is_cancel": 0,
        "line_number": 0,
        "home_amount_1": 56175,
        "home_amount_2": 56175
        }
        ]
    }
}
```

- ค้นหาเอกสารจ่ายชำระโดยค้นหาจาก คำค้นหา
- GET 
   - http://localhost:8001/payments?keyword=S01-PX6112-0030
   - response
```json
{
    "status": "success",
    "data": {
        "save_from": 0,
        "source": 0,
        "doc_no": "S01-PX6112-0030",
        "tax_no": "",
        "tax_date": "",
        "book_code": "",
        "doc_date": "2018-12-06T00:00:00Z",
        "ap_code": "053904044",
        "creator_code": "57147",
        "create_date_time": "2018-12-06T12:07:39Z",
        "last_editor_code": "",
        "last_edit_date_t": "1900-01-01T00:00:00Z",
        "depart_code": "HO-FN-00",
        "sum_of_invoice": 56175,
        "sum_of_debit_note": 0,
        "sum_of_credit_note": 0,
        "before_tax_amount": 0,
        "tax_amount": 0,
        "total_amount": 56175,
        "discount_amount": 0,
        "net_amount": 54600,
        "petty_cash_amount": 0,
        "sum_cash_amount": 0,
        "sum_chq_amount": 54600,
        "sum_bank_amount": 0,
        "sum_of_w_tax": 1575,
        "gl_format": "P07",
        "is_cancel": 0,
        "my_description": "ค่าบริการรักษาความปลอดภัย รปภ. 3 นาย สาขาสันกำแพง (อัตรา 17,500 บาท/คน) ประจำเดือน พ.ย.61 งวดที่ 7/12 (สัญญาเลขที่ CGG6104-006 ครบวันที่ 30/4/2562) ตามใบวางบิลเลขที่ GCMTBC6111-0008 ลว.30/10/61",
        "allocate_code": "",
        "project_code": "",
        "bill_group": "",
        "sum_home_amount_1": 56175,
        "sum_home_amount_2": 56175,
        "other_income": 0,
        "other_expense": 0,
        "excess_amount_1": 0,
        "excess_amount_2": 0,
        "currency_code": "",
        "exchange_rate": 1,
        "is_complete_save": 1,
        "is_confirm": 1,
        "recur_name": "",
        "confirm_code": "",
        "confirm_date_time": "1900-01-01T00:00:00Z",
        "cancel_code": "",
        "cancel_date_time": "1900-01-01T00:00:00Z",
        "sum_of_deposit_1": 0,
        "sum_of_deposit_2": 0,
        "home_amount_inv": 56175,
        "home_amount_debt": 0,
        "home_amount_crd": 0,
        "user_code": "",
        "credit_type": "",
        "confirm_no": "",
        "credit_ref_no": "",
        "credit_due_date": "",
        "bank_code": "",
        "bank_branch_code": "",
        "bank_ref_no": "",
        "trans_bank_date": "",
        "ref_date": "",
        "deps": null,
        "chqs": null,
        "subs": [
          {
        "invoice_date": "2018-11-30T00:00:00Z",
        "invoice_no": "AT6111-0034",
        "inv_balance": 56175,
        "pay_amount": 56175,
        "my_description": "ใบรับวางบิล",
        "bill_type": 4,
        "paybill_no": "S01-AIX6112-0013",
        "ref_type": 0,
        "is_cancel": 0,
        "line_number": 0,
        "home_amount_1": 56175,
        "home_amount_2": 56175
        }
        ]
    }
}
```
- บันทึกและแก้ไขข้อมูลจ่ายชำระ
- POST
    - Request
 ```json
   {
   "doc_no": "S01-PX6112-0030",
   "doc_date": "06/12/2018",
   "ap_code": "053904044",
   "depart_code": "HO-FN-00",
   "sum_of_invoice": 56175,
   "sum_of_debit_note": 0,
   "sum_of_credit_note": 0,
   "total_amount": 56175,
   "discount_amount": 0,
   "petty_cash_amount": 0,
   "sum_cash_amount": 0,
   "sum_chq_amount": 54600,
   "sum_bank_amount": 0,
   "sum_of_w_tax": 1575,
   "my_description": "",
   "allocate_code": "",
   "project_code": "",
   "other_income": 0,
   "other_expense": 0,
   "excess_amount_1": 0,
   "excess_amount_2": 0,
   "recur_name": "",
   "sum_of_deposit_1": 0,
   "sum_of_deposit_2": 0,
   "user_code": "admin",
   "deps": null,
   "chqs": null,
   "subs": [
     {
        "invoice_date": "30/11/2018",
        "invoice_no": "AT6111-0034",
        "inv_balance": 56175,
        "pay_amount": 56175,
        "my_description": "ใบรับวางบิล",
        "bill_type": 4,
        "paybill_no": "S01-AIX6112-0013",
        "ref_type": 0,
        "line_number": 0
   }
   ]
   }
       
```
#####Table BCStkIssue

######API Name

- ค้นหาเอกสารจ่ายชำระโดยค้นหาจาก เลขที่เอกสาร
- GET 
   - http://localhost:8001/stkissue?docno=S01-PX6112-0030
   - response