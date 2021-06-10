package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisController struct {
	beego.Controller
}


func (this *RedisController) Get(){
	//连接redis      //使用redis封装的Dial进行tcp连接.设置长连接时间为1s,连接超时时间为5s,读写超时均为1s，并设置密码进行连接
	conn,_:=redis.Dial("tcp",":6379",redis.DialKeepAlive(1*time.Second),redis.DialPassword("123456"),redis.DialConnectTimeout(5*time.Second),redis.DialReadTimeout(1*time.Second),redis.DialWriteTimeout(1*time.Second))

	//操作数据库
	//send和flush结合一起使用
	//conn.Send("set","name","zhangsan")
	//conn.Send("mset","age",11,"score","90")
	//conn.Flush()
	//reply,err:=conn.Receive();//接收执行结果
	//if err!=nil{
	//	this.Ctx.WriteString("设置内容错误")
	//	return
	//}
	//logs.Info(reply)//输出 OK
	//
	////conn.do（）可以更简单的使用这一条语句执行并返回结果，无需flush
	//reply,err=conn.Do("set","sex","women")
	//logs.Info(reply)//OK


	/*conn.Send("MULTI")
	conn.Send("get","name")
	conn.Send("set","class","oenClass")
	reply,err:=conn.Do("EXEC")
	if err!=nil{
		this.Ctx.WriteString("事务执行错误")
		return
	}
	logs.Info(reply)//[[122 104 97 110 103 115 97 110] OK]
	*/
	//回复助手函数的使用 Values结合Scan的使用
	reply,err:=redis.Values(conn.Do("mget","name","age"))
	if err!=nil{
		this.Ctx.WriteString("事务执行错误")
		return
	}
	var s string
	var i int
	redis.Scan(reply, &s, &i)
	logs.Info(s,i)//zhangsan 11

	this.Ctx.WriteString("执行成功")

}