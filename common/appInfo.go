package common

type AppInfo struct {
	AppId   string `json:"appId"`   // 程序唯一ID
	AppName string `json:"appName"` // 程序名称
	Author  string `json:"author"`  // 作者
	Version string `json:"version"` // 版本
	WebSite string `json:"webSite"` // 网站
}

// 基础应用信息

var App = AppInfo{
	AppId:   "com.xingcxb.next",
	AppName: "next",
	Author:  "xingcxb",
	Version: "0.0.1",
	WebSite: "https://xingcxb.com",
}
