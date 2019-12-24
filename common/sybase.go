/**
 * @Author: wangyingwen
 * @Description:
 * @File:  sybase
 * @Version: 1.0.0
 * @Date: 2019-12-22 10:54
 */

package common

import (
	"database/sql"
	"log"
	"time"
)

var SybaseDB *sql.DB

type LC_BATCHTRANS struct {
	BATCHID    string
	AMOUNT     float64
	STATUS     string
	FAMEBCUSNO string
	SUBMITTIME time.Time
	SUBMITTLR  string
	TRANSTYPE  string
	ABS        string
}

func OpenSybase() error {
	var err error
	//cnxStr := "tds://my_user:my_password@dbhost.com:5000/pubs?charset=utf8"
	cnxStr := "tds://sa:123456@10.100.100.63:5000/benefit?charset=utf8"

	SybaseDB, err = sql.Open("tds", cnxStr)
	if err != nil {
		log.Println(err)
		return err
	}
	if err := SybaseDB.Ping(); err != nil {
		log.Fatal(err)
	}

	result, err := SybaseDB.Query(`select * from LC_BATCHTRANS where  BATCHID='100001'`)
	if err != nil {
		log.Println(err)
		return err
	}
	defer result.Close()
	value, err := result.Columns()
	if err != nil {
		log.Println(err)
	}
	log.Println("www", value)
	batch := &LC_BATCHTRANS{}
	for result.Next() {
		if err := result.Scan(&batch.BATCHID, &batch.AMOUNT, &batch.STATUS, &batch.FAMEBCUSNO, &batch.SUBMITTIME, &batch.SUBMITTLR, &batch.TRANSTYPE, &batch.ABS); err != nil {
			log.Println(err)
			return err
		}
		log.Println("next")
	}
	param := &LC_BATCHTRANS{}
	param.BATCHID = "H10000003"
	param.AMOUNT = 0.01
	//rat := big.NewRat(1, 2)
	//var num tds.Num
	//num.Scan(10.235555)
	//log.Println(num.Rat())

	res, err := SybaseDB.Exec("insert into LC_BATCHTRANS(BATCHID,AMOUNT,STATUS,FAMEBCUSNO,SUBMITTIME,SUBMITTLR,TRANSTYPE,ABS) values(?,?,?,?,?,?,?,?)", param.BATCHID, param.AMOUNT, "N", "100003", time.Now(), "1", "0", "222")
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res.LastInsertId())
	log.Printf("%+v", batch)
	log.Printf("%+v", batch)

	return nil
}
