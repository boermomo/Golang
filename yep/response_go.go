package yep

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Response(w http.ResponseWriter, r *http.Request) {
	//第一个是往客户端回数据，第二个是读取客户端的数据
	// 解析url传递的参数
	if r.Method == "POST" {
		//parseform 通过header 的格式解析r,放到form,是map类型
		// 输出到客户端
		r.ParseForm()
		na := r.FormValue("name") //解析
		re := r.FormValue("response")
		db, _ := sql.Open("mysql", "root:66ccff@tcp(127.0.0.1:3306)/waterloo")
		ret, err := db.Exec("update withwater set response=? where name=?", re, na)

		//LastInsertId返回一个数据库生成的回应命令的整数。
		//返回插入的ID
		if err != nil {
			fmt.Println("失败:", err)
			return
		} else {
			rows, _ := ret.RowsAffected()
			fmt.Println("执行成功,影响行数", rows, "行")
		}
		defer db.Close()
	} else {
		return
	}
}
