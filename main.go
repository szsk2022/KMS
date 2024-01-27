package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	// 定义KMS服务器地址和端口
	kmsServer := "szsk.sunzishaokao.com"
	kmsPort := "1688"

	// 定义Office版本和对应的KMS客户端安装密钥
	officeVersions := map[string]string{
		"13": "YC7DK-G2NP3-2QQC3-J6H88-GVGXT",
		"16": "XQNVK-8JYDB-WJ9W3-YJ8YR-WFG99",
		"19": "NMMKJ-6RK4F-KMJVX-8D9MJ-6MWKP",
		"21": "FXYTK-NJJ8C-GB6DW-3DYQT-6F7TH",
		// 添加其他Office版本和对应的密钥
	}

	// 获取操作系统架构
	var arch string
	if runtime.GOARCH == "386" {
		arch = "x86"
	} else {
		arch = "x64"
	}

	// 遍历所有可能的Office版本
	for version, key := range officeVersions {
		// 构造可能的Office安装路径
		paths := []string{
			fmt.Sprintf(`C:\Program Files\Microsoft Office\Root\Office%s`, version),       // Click-to-Run安装路径
			fmt.Sprintf(`C:\Program Files\Microsoft Office\Office%s`, version),            // MSI安装路径（64位）
			fmt.Sprintf(`C:\Program Files (%s)\Microsoft Office\Office%s`, arch, version), // MSI安装路径（32位）
		}

		// 查找ospp.vbs文件
		var osppPath string
		for _, path := range paths {
			if _, err := exec.LookPath(fmt.Sprintf("%s\\ospp.vbs", path)); err == nil {
				osppPath = fmt.Sprintf("%s\\ospp.vbs", path)
				break
			}
		}

		// 如果找不到ospp.vbs，则跳过此版本
		if osppPath == "" {
			fmt.Printf("未找到Office %s 的 ospp.vbs 文件，跳过激活。\n", version)
			continue
		}

		// 设置KMS客户端安装密钥
		fmt.Printf("正在为Office %s 设置KMS客户端安装密钥...\n", version)
		cmd := exec.Command("cscript", osppPath, "/inpkey:"+key)
		cmd.Run()

		// 设置KMS服务器地址
		fmt.Printf("正在为Office %s 设置KMS服务器地址...\n", version)
		cmd = exec.Command("cscript", osppPath, "/sethst:"+kmsServer+":"+kmsPort)
		cmd.Run()

		// 尝试激活
		fmt.Printf("正在尝试激活Office %s ...\n", version)
		cmd = exec.Command("cscript", osppPath, "/act")
		cmd.Run()

		// 检查激活状态
		fmt.Printf("正在检查Office %s 的激活状态...\n", version)
		cmd = exec.Command("cscript", osppPath, "/dstatus")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("激活过程中发生错误:", err)
		} else {
			fmt.Println(string(output))
		}

		fmt.Println("-------------------------------------")
	}

	fmt.Println("激活过程完成。")
}
