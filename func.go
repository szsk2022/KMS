package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"os/exec"
)

/*
	@Author:@szsk2022
	@Email：admin@sunzishaokao.com
	@URL:https://github.com/szsk2022/kms
*/

// getProductName 从注册表获取当前Windows版本名称
func getProductName() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	productName, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return "", err
	}

	return productName, nil
}

// findProductKey 根据产品名称查找匹配的密钥
func findProductKey(name string) (string, bool) {
	for _, pk := range productKeys {
		if pk.Name == name {
			return pk.Key, true
		}
	}
	return "", false
}

// activateProduct 模拟执行激活命令
func activateProduct(key string) {
	cmd := exec.Command("cscript.exe",
		"//Nologo", "C:\\Windows\\System32\\slmgr.vbs", "/ipk", key,
		"/skms", "szsk.sunzishaokao.com", "/ato")
	// 创建或打开日志文件
	logFile, err := os.OpenFile("activation.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("无法创建或打开日志文件: %v\n", err)
		pause()
		return
	}
	defer logFile.Close()

	// 创建一个基于文件的日志对象
	logger := log.New(logFile, "", log.LstdFlags)

	// 执行命令，并捕获输出和错误
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Printf("激活失败: %v\n", err)
		fmt.Printf("无法执行激活命令，请发Issues求助： %v\n", err)
		pause()
	} else {
		logger.Printf("激活过程输出: %s\n", output)
		fmt.Printf("激活过程输出: %s", output)
		pause()
	}
}

// 初始化termbox
func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
}

// pause 按任意键继续(termbox)
func pause() {
	fmt.Print("请按任意键继续...")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break Loop
		}
	}
}

// GetWindowsVersion 获取系统版本
func GetWindowsVersion() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	productName, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return "", err
	}

	return productName, nil
}
