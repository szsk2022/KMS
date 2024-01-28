package main

import (
	"bytes"
	"fmt"
	"github.com/nsf/termbox-go"
	"golang.org/x/sys/windows/registry"
	"os"
	"os/exec"
)

/*
	@Author:@szsk2022
	@Email：admin@sunzishaokao.com
	@URL:https://github.com/szsk2022/kms
*/

// 初始化termbox
func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
}

// pause 按任意键继续
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

// 定义一个函数runCommand，接收命令名和参数作为输入，执行该命令并处理其输出和错误。
func runCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		// 打印错误信息
		fmt.Printf("执行命令失败: %v\n", err)
		fmt.Println("详细错误信息:", stderr.String())
		// 暂停程序以便查看错误信息
		pause()
		// 退出程序
		os.Exit(1)
	}
}
