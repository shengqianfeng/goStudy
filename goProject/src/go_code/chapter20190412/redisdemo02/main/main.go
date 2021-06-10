package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis" //引入redis包
)

//go操作redis的hash
func main() {
	//通过go 向redis 写入数据和读取数据
	//1. 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return 
	}

	//如果不设置密码则set  err= NOAUTH Authentication required.
	if _, err := conn.Do("AUTH", "123456"); err != nil {
	 	fmt.Println("redis.passwd err=", err)
	 	conn.Close()
	 	return
	}

	defer conn.Close() //关闭..

	//2. 通过go 向redis写入数据 string [key-val]
	_, err = conn.Do("HSet","user01", "name", "john")
	if err != nil {
		fmt.Println("hset  err=", err)
		return 
	}
	_, err = conn.Do("HSet","user01", "age", "20")
	if err != nil {
		fmt.Println("hset  err=", err)
		return 
	}
	//3. 通过go 向redis读取数据 string [key-val]
	//因为返回 r是 interface{}
	//因为 name 对应的值是string ,因此我们需要转换
	//nameString := r.(string)
	name, err := redis.String(conn.Do("HGet","user01", "name"))
	if err != nil {
		fmt.Println("hget  err=", name)
		return 
	}
	age, err := redis.Int(conn.Do("HGet", "user01","age"))
	if err != nil {
		fmt.Println("hget  err=", age)
		return 
	}
	
	
	fmt.Println("操作ok ", name,age)//操作ok  john 20
}