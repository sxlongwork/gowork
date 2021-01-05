package account

import (
	"fmt"
)

type FamilyAccount struct{
	//定义变量记录用户的选择
	sel int
	//定义变量记录是否退出
	flag bool
	//当前余额
	bal float64
	//支出/收入金额
	money float64
	//收入/支出说明
	note string
	//定义变量记录收支明细
	details string
	//定义变量记录使用有收支记录
	record bool
	//密码
	pwd string
}

func (account *FamilyAccount) SetPwd(pwd string){
	account.pwd = pwd
}

func (account *FamilyAccount) GetPwd() string{
	 return account.pwd
}

func GetFamilyAccount() *FamilyAccount{
	return &FamilyAccount{
		sel: 0,
		flag: false,
		bal: 0.0,
		money: 0.0,
		note: "",
		details: "收支\t余额\t金额\t说明",
		record: false,
		pwd: "",
	}
}

func (account *FamilyAccount) detail(){
	account.flag = false
	if account.record{
		fmt.Println(account.details)
	} else {
		fmt.Println("您还没有收支记录，来记一笔吧!")
	}
	fmt.Println()
}

func (account *FamilyAccount) comeIn(){
	account.flag = false
	fmt.Printf("本次收入金额:")
	fmt.Scanln(&account.money)
	fmt.Printf("本次收入说明:")
	fmt.Scanln(&account.note)

	account.bal += account.money
	account.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", account.bal, account.money, account.note)
	account.record = true
}

func (account *FamilyAccount) pay(){
	account.flag = false
	fmt.Printf("本次支出金额:")
	fmt.Scanln(&account.money)
	if account.money > account.bal {
		fmt.Println("余额不足,无法支出!")
		return
	}
	fmt.Printf("本次支出说明:")
	fmt.Scanln(&account.note)

	account.bal -=account.money
	account.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", account.bal, account.money, account.note)
	account.record = true
}

func (account *FamilyAccount) exit(){
	var key string
	for {
		fmt.Printf("您确定退出吗? (y/n):")
		fmt.Scanln(&key)
		if key == "y"{
			account.flag = true
			break
		} else if key == "n" {
			break
		} else{
			fmt.Println("您的输入有误，请重新输入")
		}
	}
}


func (account *FamilyAccount) MainMenu(){
	for {
		fmt.Println("----------家庭收支记账软件----------")
		fmt.Println("          1 收支明细")
		fmt.Println("          2 登记收入")
		fmt.Println("          3 登记支出")
		fmt.Println("          4 退出")
		fmt.Println()
		fmt.Printf("          请选择(1-4):")
		fmt.Scanln(&account.sel)
		switch account.sel{
		case 1:
			account.detail()
		case 2:
			account.comeIn()
		case 3:
			account.pay()
		case 4:
			account.exit()
		default:
			fmt.Println("您的选择有误，请重新选择!")
		}
		if account.flag {
			break
		}

	}
}



