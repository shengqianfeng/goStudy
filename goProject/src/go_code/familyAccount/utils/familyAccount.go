package utils
import (
	"fmt"
)

type FamilyAccount struct{
	//声明必要的字段

	//声明一个变量接收用户输入选项
	key string
	//声明一个变量是否退出for循环
	loop bool

	//定义账户余额
	 balance float64
	//每次收支的金额
	 money float64
	//每次收支的说明
	 note string
	 //定义个变量记录是否有收支的行为
	 flag bool
	//收支的详情，使用字符串来记录
	//当有收支时只需要对details进行拼接处理即可 "收支\t账户金额\t收支金额\t说明"
	details string

}


func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key : "",
		//声明一个变量是否退出for循环
		loop : true,

		//定义账户余额
		balance : 10000.0,
		//每次收支的金额
		money : 0.0,
		//每次收支的说明
		note : "",
		//定义个变量记录是否有收支的行为
		flag:false,
		//收支的详情，使用字符串来记录
		//当有收支时只需要对details进行拼接处理即可
		details:"收支\t账户金额\t收支金额\t说明",
	}
}
//将显示明细写成一个方法
func (this *FamilyAccount) showDeatils()  {
	fmt.Println("----------显示当前收支明细记录----------------------")
	if this.flag{
		fmt.Println(this.details)
	}else{
		fmt.Println("当前没有收支，请来一笔吧！")
	}
}

//将登记收入写成一个方法和*FamilyAccount绑定
func (this *FamilyAccount) income(){
	fmt.Println("----------登记收入----------------------")
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	//拼接到details变量
	this.details+=fmt.Sprintf("\n收入\t%v\t%v\t%v",this.balance,this.money,this.note)
	this.flag=true
}

//将登记支出也成一个防范和*FamilyAccount绑定
func (this *FamilyAccount) pay(){
	fmt.Println("----------登记支出----------------------")
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	//这里需要做一个判断：支出必须小于余额
	if this.money > this.balance{
		fmt.Println("支出余额不足")
		// break
	}else{
		this.balance -= this.money //修改账户余额
		fmt.Println("本次支出说明:")
		fmt.Scanln(&this.note)
		//拼接到details变量
		this.details+=fmt.Sprintf("\n支出\t%v\t%v\t%v",this.balance,this.money,this.note)
		this.flag=true
	}
}

//将推出系统也成一个防范和*FamilyAccount绑定
func (this *FamilyAccount) exit(){
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
					this.loop = false
				}
}
//给该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu()  {
	//显示主菜单
	for{
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDeatils()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("请输入正确选项..")
		}
		if(!this.loop){//退出for循环
			break
		}

	}
}
