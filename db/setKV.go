package db

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)


var _db *sql.DB

func InitDB()  {

	db, err := sql.Open("sqlite3", "setKey.db")

	_db = db
	checkErr(err)

}


func SetK(c *gin.Context, key string,value string)  {
	if _db == nil {
		return
	}
	//fmt.Sprintf("key : %s,value : %s", key, value,db)
	//INSERT INTO students (id, class_id, name, gender, score) VALUES (1, 1, '小明', 'F', 99) ON DUPLICATE KEY UPDATE name='小明', gender='F', score=99;

	//replace into setKey( k, v ,upTime ) VALUES ('key','{id = 60}',35)

	//// 插入数据
	stmt, err := _db.Prepare("replace into setKey(k, v, upTime) values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec(key, value,time.Now().Format("2006-01-02 15:04:05"))
	checkErr(err)
	_, err = res.LastInsertId() //返回新增的id号
	checkErr(err)
	//fmt.Println(id)
	//c.String(http.StatusOK, fmt.Sprintf("key : %s,value : %s", key, value))

	data := gin.H{"code": 0,"msg":"成功"}
	c.JSON(http.StatusOK, data)
}

func GetK(c *gin.Context, key string)  {
	if _db == nil {
		return
	}
	//fmt.Sprintf("key : %s,value : %s", key, value,db)
	//INSERT INTO students (id, class_id, name, gender, score) VALUES (1, 1, '小明', 'F', 99) ON DUPLICATE KEY UPDATE name='小明', gender='F', score=99;

	//replace into setKey( k, v ,upTime ) VALUES ('key','{id = 60}',35)

	//select * from setKey where k =  'passData'

	var query = "SELECT * FROM setKey where k ="+"'"+key+"'"

	var data  []map[string]interface{}

	rows, err := _db.Query(query)
	checkErr(err)
	for rows.Next() {
		var k string
		var v string
		var upTime string
		err = rows.Scan(&k, &v, &upTime)
		checkErr(err)
		log.Println(fmt.Sprintf("数据表中所有数据信息如下 \n %s %s %s",k,v,upTime))

		//fmt.Println("userinfo 数据表中所有数据信息如下：\n", uid, username, department, created)
		data = append(data, gin.H{"k": k, "v": v ,"upTime":upTime})

	}
	log.Println(fmt.Sprintf("返回的数据是 \n %s",data))

	var code = 0
	if len(data)>0 {
		code = 0
	}else{
		code = 1
	}
	returnData := gin.H{"code": code,"msg":"成功","data":data}
	c.JSON(http.StatusOK, returnData)
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}