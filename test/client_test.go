package test

import (
	"fmt"
	"testing"

	"github.com/toolsetlink/upgradelink-api-go/client"
)

// 获取 url 应用升级内容
func TestGetUrlUpgrade(t *testing.T) {

	accessKeyId := "accessKeyId"
	accessKeySecret := "accessKeySecret"
	Client, err := client.NewClient(&accessKeyId, &accessKeySecret)
	if err != nil {
		return
	}

	urlKey := "urlkey1"
	versionCode := 1
	appointVersionCode := 0
	devModelKey := "tv1"
	devKey := "macmacmacmac"

	// 接口调用
	request := &client.UrlUpgradeRequest{
		UrlKey:             &urlKey,
		VersionCode:        &versionCode,
		AppointVersionCode: &appointVersionCode,
		DevModelKey:        &devModelKey,
		DevKey:             &devKey,
	}
	Info, err := Client.GetUrlUpgrade(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 获取 file 应用升级内容
func TestGetFileUpgrade(t *testing.T) {

	accessKeyId := "accessKeyId"
	accessKeySecret := "accessKeySecret"
	Client, err := client.NewClient(&accessKeyId, &accessKeySecret)
	if err != nil {
		return
	}

	fileKey := "filekey1"
	versionCode := 1
	appointVersionCode := 0
	devModelKey := "tv1"
	devKey := "macmacmacmac"

	// 接口调用
	request := &client.FileUpgradeRequest{
		FileKey:            &fileKey,
		VersionCode:        &versionCode,
		AppointVersionCode: &appointVersionCode,
		DevModelKey:        &devModelKey,
		DevKey:             &devKey,
	}
	Info, err := Client.GetFileUpgrade(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}
