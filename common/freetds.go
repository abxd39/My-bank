/**
 * @Author: wangyingwen
 * @Description:
 * @File:  freetds
 * @Version: 1.0.0
 * @Date: 2019-12-23 15:37
 */

package common

import (
	"log"

	freetds "github.com/minus5/gofreetds"
)

func InitFreeTds() error {
	connStr := "sybase://sa:123456@10.100.100.63:5000/benefit?charset=utf8"
	connStr = "host=egServer50;database=benefit;user=sa;pwd=123456;compatibility_mode=Sybase"
	pool, err := freetds.NewConnPool(connStr)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	//get connection
	conn, err := pool.Get()
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	//rst, err := conn.ExecSp("sp_help", "authors")
	//returnValue := rst.Status()
	//log.Println(returnValue)
	//var param1, param2 int
	//err = rst.ParamScan(&param1, &param2)
	//if err != nil {
	//	log.Println(err)
	//}
	//if rst.NextResult() {
	//	for rst.Next() {
	//		var v1, v2 string
	//		err = rst.Scan(&v1, &v2)
	//		if err != nil {
	//			log.Println(err)
	//		}
	//		log.Println(v1, v2)
	//	}
	//}
	//

	result, err := conn.Exec(`select * from LC_BATCHTRANS`)
	if err != nil {
		log.Println(err)
		return err
	}

	//batch := &LC_BATCHTRANS{}
	batch := &LC_BATCHTRANS{}
	//rst := result[0]
	//err = rst.Scan(&batch.BATCHID, &batch.AMOUNT, &batch.STATUS, &batch.FAMEBCUSNO, &batch.SUBMITTIME, &batch.SUBMITTLR, &batch.TRANSTYPE, &batch.ABS)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//log.Println(batch)
	for i := 0; i < len(result); i++ {
		rst := result[i]
		for rst.Next() {
			err = rst.Scan(&batch.BATCHID, &batch.AMOUNT, &batch.STATUS, &batch.FAMEBCUSNO, &batch.SUBMITTIME, &batch.SUBMITTLR, &batch.TRANSTYPE, &batch.ABS)
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println("2", batch)
		}
		for _, v := range rst.Rows {
			log.Println("2.2", v)
		}
		log.Println("1")
	}
	//log.Printf("%+v", result[0].Rows)
	col := result[0].Columns
	for _, v := range col {
		log.Printf("%+v", v.Name)
	}

	//SybaseDB, err := sql.Open("Sybase", connStr)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//if err := SybaseDB.Ping(); err != nil {
	//	log.Fatal(err)
	//}
	////execute query
	//batch := &LC_BATCHTRANS{}
	//result, err := SybaseDB.Query(`select * from LC_BATCHTRANS where  BATCHID='test001'`)
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//
	//defer result.Close()
	//value, err := result.Columns()
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("www", value)
	//for result.Next() {
	//	if err := result.Scan(&batch.BATCHID, &batch.AMOUNT, &batch.STATUS, &batch.FAMEBCUSNO, &batch.SUBMITTIME, &batch.SUBMITTLR, &batch.TRANSTYPE, &batch.ABS); err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//	log.Println("next")
	//}
	return nil
}
