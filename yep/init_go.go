package yep

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Build_up(w http.ResponseWriter, r *http.Request) {
	//第一个是往客户端回数据，第二个是读取客户端的数据
	// 解析url传递的参数
	if r.Method == "POST" {
		//parseform 通过header 的格式解析r,放到form,是map类型
		// 输出到客户端
		r.ParseForm()
		na := r.FormValue("你的名字") //解析
		pa := r.FormValue("你的故事")

		db, _ := sql.Open("mysql", "root:66ccff@tcp(127.0.0.1:3306)/waterloo")
		ret, _ := db.Exec("insert into withwater(name,story) values(?,?)", na, pa)

		//LastInsertId返回一个数据库生成的回应命令的整数。
		//返回插入的ID
		insID, _ := ret.LastInsertId()
		fmt.Println(insID)

		jsonbytes1, _ := json.Marshal(na)
		w.Write(jsonbytes1)
		jsonbytes, _ := json.Marshal(pa)
		w.Write(jsonbytes)

		defer db.Close()
	} else {
		return
	}
}
