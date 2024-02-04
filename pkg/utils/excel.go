package utils

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
)

// ExportExcel 导出Excel表格
/**
 * @param  name     {string}    导出的表名
 * @param  header   {[]string}  表头key，导出后显示的顺序
 * @param  headerKV {map[string]string}  表头、数据kv对照
 * @param  data     {[]map[string]interface{}} 数据集合
 * @param  filePath {string} 文件保存路径，需要带"/"
 * @return err      {error}                    异常
 */
func ExportExcel(name string, header []string, headerKV map[string]string, data []map[string]interface{}, filePath string) (fileNamePath string, err error) {
	f := excelize.NewFile()
	// Create a new sheet
	index := f.NewSheet("Sheet1")
	headerValue := make([]string, 0)
	for _, v := range header {
		headerValue = append(headerValue, headerKV[v])
	}
	f.SetSheetRow("Sheet1", "A1", &headerValue)
	var rowValue []interface{} //行数据对象
	for i, v := range data {
		rowValue = make([]interface{}, 0)
		rowNum := strconv.Itoa(i + 2) //表中的行顺序是从1开始的;  A1是表头,A2才是数据的开始行
		for _, key := range header {
			rowValue = append(rowValue, v[key])
		}
		f.SetSheetRow("Sheet1", "A"+rowNum, &rowValue)
	}

	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	fileName := name + ".xlsx"         //文件名称
	fileNamePath = filePath + fileName //保存文件的位置
	err = f.SaveAs(fileNamePath)
	return
}

// ReadExcel 读取excel数据
func ReadExcel(filename, sheetName string) (rows [][]string, err error) {
	// 打开 Excel 文件
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		return rows, err
	}

	// 读取所有行
	rows = xlsx.GetRows(sheetName)
	// 去掉 表头
	rows = rows[1:]
	return rows, nil
}
