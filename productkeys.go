package main

/*
	@Author:@szsk2022
	@Email：admin@sunzishaokao.com
	@URL:https://github.com/szsk2022/kms
*/

// ProductKey 结构体表示产品名称和密钥对
type ProductKey struct {
	Name string
	Key  string
}

// productKeys 是一个预定义的产品密钥列表
var productKeys = [...]ProductKey{
	{"Windows 11 Pro", "W269N-WFGWX-YVC9B-4J6C9-T83GX"},
	{"Windows 11 Enterprise", "NPPR9-FWDCX-D2C8J-H872K-2YT43"},
	{"Windows 10 Pro", "W269N-WFGWX-YVC9B-4J6C9-T83GX"},
	{"Windows 10 Enterprise", "NPPR9-FWDCX-D2C8J-H872K-2YT43"},
	{"Windows 8.1 Pro", "GCRJD-8NW9H-F2CDX-CCM8D-9D6T9"},
	{"Windows 8.1 Enterprise", "MHF9N-XY6XB-WVXMC-BTDCT-MKKG7"},
	{"Windows 7 Pro", "FJ82H-XT6CR-J8D7P-XQJJ2-GPDD4"},
	{"Windows 7 Enterprise", "33PXH-7Y6KF-2VJC9-XBBR8-HVTHH"},
	{"Windows Server 2022 Standard", "VDYBN-27WPP-V4HQT-9VMD4-VMK7H"},
	{"Windows Server 2022 Datacenter", "WX4NM-KYWYW-QJJR4-XV3QB-6VM33"},
	{"Windows Server 2019 Standard", "N69G4-B89J2-4G8F4-WWYCC-J464C"},
	{"Windows Server 2019 Datacenter", "WMDGN-G9PQG-XVVXX-R3X43-63DFG"},
	{"Windows Server 2016 Standard", "WC2BQ-8NRM3-FDDYY-2BFGV-KHKQY"},
	{"Windows Server 2016 Datacenter", "CB7KF-BWN84-R7R2Y-793K2-8XDDG"},
	{"Windows Server 2012 R2 Server Standard", "D2N9P-3P6X9-2R39C-7RTCD-MDVJX"},
	{"Windows Server 2012 R2 Datacenter", "W3GGN-FT8W3-Y4M27-J84CP-Q3VJ9"},
	{"Windows Server 2008 R2 Standard", "YC6KT-GKW9T-YTKYR-T4X34-R7VHC"},
	{"Windows Server 2008 R2 Datacenter", "74YFP-3QFB3-KQT8W-PMXWJ-7M648"},
	{"Windows Server 2008 R2 Enterprise", "489J6-VHDMP-X63PK-3K798-CPX3Y"},
}
