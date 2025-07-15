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

// 获取 apk 应用升级内容
func TestGetApkUpgrade(t *testing.T) {

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

	apkKey := "isVZBUvkFhv6oHxk_X-D0Q"
	versionCode := 1
	appointVersionCode := 0
	devModelKey := ""
	devKey := ""

	// 接口调用
	request := &client.ApkUpgradeRequest{
		ApkKey:             &apkKey,
		VersionCode:        &versionCode,
		AppointVersionCode: &appointVersionCode,
		DevModelKey:        &devModelKey,
		DevKey:             &devKey,
	}
	Info, err := Client.ApkUpgrade(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 获取 配置 升级内容
func TestGetConfigurationUpgrade(t *testing.T) {

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

	configurationKey := "q1hfB1VUQaK9VksTZGPU1Q"
	versionCode := 1
	appointVersionCode := 0
	devModelKey := ""
	devKey := ""

	// 接口调用
	request := &client.ConfigurationUpgradeRequest{
		ConfigurationKey:   &configurationKey,
		VersionCode:        &versionCode,
		AppointVersionCode: &appointVersionCode,
		DevModelKey:        &devModelKey,
		DevKey:             &devKey,
	}
	Info, err := Client.ConfigurationUpgrade(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 上报事件
// /* app_start 应用-启动事件 */
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
	eventType := client.EventTypeAppStart
	appKey := "LOYlLXNy7wV3ySuh0XgtSg"
	target := "darwin"
	arch := "x86_64"
	devModelKey := ""
	devKey := ""
	versionCode := 1
	timestamp := client.TimeRFC3339()

	eventData := &client.AppReportRequestEventData{
		LaunchTime:  timestamp,
		VersionCode: &versionCode,
		Target:      &target,
		Arch:        &arch,
		DevModelKey: &devModelKey,
		DevKey:      &devKey,
	}

	// 接口调用
	request := &client.AppReportRequest{
		EventType: &eventType,
		AppKey:    &appKey,
		Timestamp: timestamp,
		EventData: eventData,
	}

	Info, err := Client.AppReport(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}
}
