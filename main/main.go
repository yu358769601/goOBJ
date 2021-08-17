package main

import (
	"log"
	"passManger/db"
	"strings"
)
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

)

func main() {
	db.InitDB()
	r := gin.Default()
	r.Use(Cors())
	r.GET("/", hello )
	r.Static("/h5", "./assets")
	r.POST("/postTest", postTest )
	r.POST("/setKey", setKey )
	r.POST("/getKey", getKey )
	//监听端口默认为8080
	r.Run(":10086")
}
////// 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method      //请求方法
		origin := c.Request.Header.Get("Origin")        //请求头部
		var headerKeys []string                             // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")       // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next()        //  处理请求
	}
}

func hello(c *gin.Context)  {



	var m = make(map[string]interface{}, 1)
	m["txt"] = "hello word"
	data := gin.H{"name": "hello word"}
	c.JSON(http.StatusOK, data)
}

func postTest(c *gin.Context)  {
	types := c.DefaultPostForm("type", "post")

	log.Println(fmt.Sprintf("Params : %s", c.Params))
	username := c.PostForm("username")
	password := c.PostForm("userpassword")
	// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))

	log.Println(fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	data := gin.H{"code": 0,"msg":"成功"}
	c.JSON(http.StatusOK, data)

}

func setKey(c *gin.Context)  {
	_ = c.DefaultPostForm("type", "post")
	key := c.PostForm("key")
	value := c.PostForm("value")
	//c.String(http.StatusOK, fmt.Sprintf("key:%s,type:%s", key, types))
	db.SetK(c,key,value)
}
func getKey(c *gin.Context)  {
	_ = c.DefaultPostForm("type", "post")
	key := c.PostForm("key")
	//c.String(http.StatusOK, fmt.Sprintf("key:%s,type:%s", key, types))
	db.GetK(c,key)
}