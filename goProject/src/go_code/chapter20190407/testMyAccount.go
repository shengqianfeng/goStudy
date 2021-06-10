package main
import (
	"fmt"
)

func main()  {
	//声明一个变量接收用户输入选项
	key := ""
	//声明一个变量是否退出for循环
	loop := true

	//定义账户余额
	 balance := 10000.0
	//每次收支的金额
	 money := 0.0
	//每次收支的说明
	 note := ""
	 //定义个变量记录是否有收支的行为
	 flag:=false
	//收支的详情，使用字符串来记录
	//当有收支时只需要对details进行拼接处理即可
	details:="收支\t账户金额\t收支金额\t说明"

	//显示主菜单
	for{
		fmt.Println("\n----------家庭收支记账软件----------------------")
		fmt.Println("----------1 收支明细----------------------")
		fmt.Println("----------2 登记收入----------------------")
		fmt.Println("----------3 登记支出----------------------")
		fmt.Println("----------4 退出----------------------")
		fmt.Println("请选择:[1-4]")
		
		fmt.Scanln(&key)

		switch key {
			case "1":
				fmt.Println("----------显示当前收支明细记录----------------------")
				if flag{
					fmt.Println(details)
				}else{
					fmt.Println("当前没有收支，请来一笔吧！")
				}
			case "2":
				fmt.Println("----------登记收入----------------------")
				fmt.Println("本次收入金额:")
				fmt.Scanln(&money)
				balance += money //修改账户余额
				fmt.Println("本次收入说明:")
				fmt.Scanln(&note)
				//拼接到details变量
				details+=fmt.Sprintf("\n收入\t%v\t%v\t%v",balance,money,note)
				flag=true
			case "3":
				fmt.Println("----------登记支出----------------------")
				fmt.Println("本次支出金额:")
				fmt.Scanln(&money)
				//这里需要做一个判断：支出必须小于余额
				if money > balance{
					fmt.Println("支出余额不足")
					break
				}
				balance -= money //修改账户余额
				fmt.Println("本次支出说明:")
				fmt.Scanln(&note)
				//拼接到details变量
				details+=fmt.Sprintf("\n支出\t%v\t%v\t%v",balance,money,note)
				flag=true
			case "4":
				fmt.Println("您确定要退出吗？y/n")
				choice:=""
				for{
					fmt.Scanln(&choice)
					if choice == "y" || choice=="n"{
						break//退出内层循环
					}
					fmt.Println("您的输入有误，请重新输入y/n")
				}
				if choice == "y"{
					loop = false
				}
			default:
				fmt.Println("请输入正确选项..")
		}
		if(!loop){//退出for循环
			break
		}
	}
	fmt.Println("你退出了家庭记账软件！")
}