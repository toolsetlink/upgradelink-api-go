package test

import (
	"fmt"
	"testing"

	"github.com/toolsetlink/upgradelink-api-go/client"
)

// 获取 url 应用升级内容
func TestGetUrlUpgrade(t *testing.T) {

	var config = client.Config{}
	config.SetAccessKey("mui2W50H1j-OC4xD6PgQag")
	config.SetAccessSecret("PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc")

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	urlKey := "uJ47NPeT7qjLa1gL3sVHqw"
	versionCode := 1
	appointVersionCode := 0
	devModelKey := ""
	devKey := ""

	// 接口调用
	request := &client.UrlUpgradeRequest{
		UrlKey:             &urlKey,
		VersionCode:        &versionCode,
		AppointVersionCode: &appointVersionCode,
		DevModelKey:        &devModelKey,
		DevKey:             &devKey,
	}
	Info, err := Client.UrlUpgrade(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 获取 file 应用升级内容
func TestGetFileUpgrade(t *testing.T) {

	accessKey := "mui2W50H1j-OC4xD6PgQag"
	accessSecret := "PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc"

	var config = client.Config{
		AccessKey:    &accessKey,
		AccessSecret: &accessSecret,
	}

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	fileKey := "LOYlLXNy7wV3ySuh0XgtSg"
	versionCode := 1
	appointVersionCode := 0
	devModelKey := ""
	devKey := ""

	// 接口调用
	request := &client.FileUpgradeRequest{
		FileKey:            &fileKey,
		VersionCode:        &versionCode,
		AppointVersionCode: &appointVersionCode,
		DevModelKey:        &devModelKey,
		DevKey:             &devKey,
	}
	Info, err := Client.FileUpgrade(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 上报事件
func TestPostAppReport(t *testing.T) {

	accessKey := "mui2W50H1j-OC4xD6PgQag"
	accessSecret := "PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc"
	var config = client.Config{
		AccessKey:    &accessKey,
		AccessSecret: &accessSecret,
	}

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	/* app_start 应用-启动事件 */
	//eventType := client.EventTypeAppStart
	//appKey := "LOYlLXNy7wV3ySuh0XgtSg"
	//devModelKey := ""
	//devKey := ""
	//versionCode := 1
	//timestamp := client.TimeRFC3339()
	//launchTime := client.TimeRFC3339()
	//eventData := &client.AppReportRequestEventData{
	//	LaunchTime: launchTime,
	//}

	/* app_upgrade_download 应用升级-下载事件 */
	//eventType := client.EventTypeAppUpgradeDownload
	//appKey := "LOYlLXNy7wV3ySuh0XgtSg"
	//devModelKey := ""
	//devKey := ""
	//versionCode := 1
	//timestamp := client.TimeRFC3339()
	//downloadVersionCode := 10
	//code := client.EventTypeCodeError
	//eventData := &client.AppReportRequestEventData{
	//	Code:                &code,
	//	DownloadVersionCode: &downloadVersionCode,
	//}

	/* app_upgrade_upgrade 应用升级-升级事件 */
	eventType := client.EventTypeAppUpgradeUpgrade
	appKey := "LOYlLXNy7wV3ySuh0XgtSg"
	devModelKey := ""
	devKey := ""
	versionCode := 1
	timestamp := client.TimeRFC3339()
	upgradeVersionCode := 10
	code := client.EventTypeCodeSuccess
	eventData := &client.AppReportRequestEventData{
		Code:               &code,
		UpgradeVersionCode: &upgradeVersionCode,
	}

	// 接口调用
	request := &client.AppReportRequest{
		EventType:   &eventType,
		AppKey:      &appKey,
		DevModelKey: &devModelKey,
		DevKey:      &devKey,
		VersionCode: &versionCode,
		Timestamp:   timestamp,
		EventData:   eventData,
	}

	Info, err := Client.AppReport(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}
