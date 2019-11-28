package yep

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Query_one(w http.ResponseWriter, r *http.Request) {
	//第一个是往客户端回数据，第二个是读取客户端的数据
	// 解析url传递的参数
	if r.Method == "POST" {
		//parseform 通过header 的格式解析r,放到form,是map类型
		// 输出到客户端
		db, _ := sql.Open("mysql", "root:66ccff@tcp(127.0.0.1:3306)/waterloo")
		var id int
		var name string
		var story string
		var response string
		roww := db.QueryRow("SELECT * FROM withwater ORDER BY RAND() limit 1")
		roww.Scan(&id, &name, &story, &response)
		//fmt.Println("ddd")
		fmt.Println(id, name, story, response)
		//LastInsertId返回一个数据库生成的回应命令的整数。
		//返回插入的ID

		defer db.Close()
	} else {
		return
	}
}
