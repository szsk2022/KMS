package main

import (
	"fmt"
)

/*
	@Author:@szsk2022
	@Email：admin@sunzishaokao.com
	@URL:https://github.com/szsk2022/kms
*/

func main() {
	fmt.Println("Author:szsk2022")
	fmt.Println("Ver:1.0.1")
	fmt.Println("Email:admin@sunzishaokao.com")
	fmt.Println("URL:https://www.github.com/szsk2022/kms")
	fmt.Printf("注意：请以管理员身份运行此程序，否则可能会激活失败！\n")
	productName, err := getProductName()
	if err != nil {
		fmt.Println("getProductName函数获取Windows系统版本失败:", err)
		return
	}
	version, err := GetWindowsVersion()
	if err != nil {
		fmt.Println("GetWindowsVersion函数获取Windows版本时失败:", err)
		return
	}
	key, found := findProductKey(productName)
	if found {
		fmt.Printf("检测到是%s，开始激活......\n", version)
		fmt.Println("已执行激活命令，请耐心等待！")
		activateProduct(key)
	} else {
		fmt.Println("Error：找不到或不支持的Windows版本或产品密钥")
		pause()
	}
}
