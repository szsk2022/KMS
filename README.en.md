# SZSK-KMS Activation Tool
![](https://www.sunzishaokao.com/wp-content/uploads/2022/06/73c6e81c1c33.jpg)

[![](https://img.shields.io/badge/Author-孙子烧烤-orange.svg)]()
[![](https://img.shields.io/badge/version-v1.0-brightgreen.svg)](https://gitee.com/szsk/kms)
[![star](https://gitee.com/szsk/kms/badge/star.svg?theme=dark)](https://gitee.com/szsk/kms/stargazers)
[![fork](https://gitee.com/szsk/kms/badge/fork.svg?theme=dark)](https://gitee.com/szsk/kms/members)

#### Introduction
The KMS activation tool is written in Go language, which is not yet fully developed. It is expected to be completed by March 2024!

#### Usage instructions
1. Download the latest compiled program from [Gitee-Releases](https://gitee.com/szsk/kms/releases "Releases")
2.Simply run the program, and it will automatically detect the system version and activate it

#### Compilation instructions
1.Clone the source code to your local machine
	`git clone https://gitee.com/szsk/kms.git`  
2.Open `main.go` and search for `kms.sunzishaokao.com` within the file, replace it with your own KMS server address
3.Compile the source code  
	`go build`

#### Contribution

1. Fork this repository
2. Create a new branch named Feat_xxx
3. Commit your code
4. Create a Pull Request