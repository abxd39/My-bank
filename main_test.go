/**
 * @Author: wangyingwen
 * @Description:
 * @File:  main_test
 * @Version: 1.0.0
 * @Date: 2019-12-22 10:51
 */

package main

import (
	"myBank/common"
	"testing"
)

func TestSybaseDB(t *testing.T) {
	//SybaseDB.Exec("")
	//common.OpenSybase()
	common.InitFreeTds()
	//id := 0
	//rows, err := SybaseDB.Query("select * from authors where id = ?", id)
}
