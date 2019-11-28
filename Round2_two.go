package main

import (
	"Round2_two/yep"
	"log"
	"net/http"
)

func main() {

	//路由
	//访问前面的路径/   --》执行后面的方法
	http.HandleFunc("/", yep.Interesting)
	http.HandleFunc("/init", yep.Build_up)
	http.HandleFunc("/delete", yep.Delete_one)
	http.HandleFunc("/query", yep.Query_one)
	http.HandleFunc("/response", yep.Response)
	errr := http.ListenAndServe(":8080", nil)
	if errr != nil {
		log.Fatal("ListenAndserve:", errr)
	}
}

// dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
// db, err := sql.Open("mysql", dsn)
// if err != nil {
// 	fmt.Printf("Open mysql failed,err:%v\n", err)
// 	return
// }
// db.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
// db.SetMaxOpenConns(100)                  //设置最大连接数
// db.SetMaxIdleConns(16)                   //设置闲置连接数

//多行查询
// rows, err := db.Query("select id,name,story from withwater")
// if err != nil {
// 	fmt.Println(err)
// }
// for rows.Next() {
// 	rows.Scan(&id, &name, &story)
// 	fmt.Println(name, story)
// }
// defer rows.Close()

//单行查询
// var id int
// var name string
// var story string
// roww := db.QueryRow("SELECT * FROM withwater ORDER BY RAND() limit 1")
// roww.Scan(&id, &name, &story)
// fmt.Println(name, story)

//插入元素
// ret, _ := db.Exec("insert into withwater(id,name,story) values(3,'秋','我自关山点酒')")
// //LastInsertId返回一个数据库生成的回应命令的整数。
// //返回插入的ID
// insID, _ := ret.LastInsertId()
// fmt.Println(insID)

//删除元素
// ret3, _ := db.Exec("delete from withwater where id = ?", 3)
// delNums, _ := ret3.RowsAffected()
// fmt.Println(delNums)

// //关闭数据库
// defer db.Close()
