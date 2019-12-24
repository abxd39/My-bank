package main

import (
	"log"
	"regexp"

	"github.com/arthurkiller/rollingwriter"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-oci8"
	"github.com/tealeg/xlsx"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {
	var conn string = "C##subsidies/123456@10.100.100.42:1521/orcl?charset=utf-8"
	Engine, err := xorm.NewEngine("oci8", conn)
	if err != nil {
		panic("mysql connect fail")
	}
	config := rollingwriter.Config{
		LogPath:                "./logs",                    //日志路径
		TimeTagFormat:          "060102150405",              //时间格式串
		FileName:               "oracle_sql",                //日志文件名
		MaxRemain:              3,                           //配置日志最大存留数
		RollingPolicy:          rollingwriter.VolumeRolling, //配置滚动策略 norolling t imerolling volumerolling
		RollingTimePattern:     "* * * * * *",               //配置时间滚动策略
		RollingVolumeSize:      "1M",                        //配置截断文件下限大小
		WriterMode:             "none",
		BufferWriterThershould: 256,
		// Compress will compress log file with gzip
		Compress: true,
	}

	writer, err := rollingwriter.NewWriterFromConfig(&config)
	if err != nil {
		panic(err)
	}

	var logger *xorm.SimpleLogger = xorm.NewSimpleLogger(writer)

	Engine.SetLogger(logger)
	Engine.ShowSQL(true)

	logger.Info("ts info code")

	Engine.Exec("select * from cp_order where id = ? ", 3)
	//time.Sleep(1 * time.Second)
	//Myxlsx()
	Strregexp()
}

func Myxlsx() {
	var faileFile *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	faileFile = xlsx.NewFile()
	sheet, err = faileFile.AddSheet("Sheet1")
	if err != nil {
		log.Println(err)
		return
	}
	row = sheet.AddRow()
	row.SetHeight(77.25)
	cell = row.AddCell()
	cell.Value = `表格填写说明：
1、带*的信息请按照示例学生格式正确填写
2、“证件号码”默认作为学生信息唯一标识
注：“李雷”为示例，上传不会新增该条学生信息`

	sty := xlsx.NewStyle()
	al := xlsx.DefaultAlignment()
	//al.Horizontal = "left"
	//al.Vertical = "botten"
	fill := xlsx.DefaultFill()
	fill.BgColor = "8DB4E4"
	sty.Alignment = *al
	sty.Fill = *fill
	cell.SetStyle(sty)
	cell.Merge(14, 0)
	headList := []string{"*学校名称", "*入学年份", "*年级", "*班级", "*姓名", "*证件类型", "*证件号码", "*补助金额", "*监护人1姓名", "*监护人1证件类型", "*监护人1证件号码", "监护人2姓名", "监护人2证件类型", "监护人2证件号码", "导入失败原因"}
	row = sheet.AddRow()
	for _, v := range headList {
		cell = row.AddCell()
		//sty := xlsx.NewStyle()
		fill.BgColor = "E26B00"
		cell.Value = v
		cell.SetStyle(sty)
		//f := xlsx.DefaultFont()
		//f.Color = "E26B00"
		//sty.Font = *f
		//cell.SetStyle(sty)
	}
	err = faileFile.Save(`/Users/wangyingwen/job/myBank/fail.xlsx`)
	if err != nil {
		log.Println(err)
	}

}

func Strregexp() {

	pat := "[0-9]+"
	rex := regexp.MustCompile(pat)
	if rex != nil {
		s := rex.FindAllString("（12）", -1)
		log.Println(s[0])
	}
}
