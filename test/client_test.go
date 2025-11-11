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

	urlKey := "OpggWISrLVRFa5y04LzkwA"
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

// 获取 url 应用版本信息
func TestGetUrlVersion(t *testing.T) {

	var config = client.Config{}
	config.SetAccessKey("mui2W50H1j-OC4xD6PgQag")
	config.SetAccessSecret("PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc")

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	urlKey := "OpggWISrLVRFa5y04LzkwA"
	versionCode := 110

	// 接口调用
	request := &client.UrlVersionRequest{
		UrlKey:      &urlKey,
		VersionCode: &versionCode,
	}
	Info, err := Client.UrlVersion(request)
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

// 获取 file 应用版本信息
func TestGetFileVersion(t *testing.T) {

	var config = client.Config{}
	config.SetAccessKey("mui2W50H1j-OC4xD6PgQag")
	config.SetAccessSecret("PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc")

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	fileKey := "LOYlLXNy7wV3ySuh0XgtSg"
	versionCode := 2

	// 接口调用
	request := &client.FileVersionRequest{
		FileKey:     &fileKey,
		VersionCode: &versionCode,
	}
	Info, err := Client.FileVersion(request)
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

// 获取 apk 应用版本信息
func TestGetApkVersion(t *testing.T) {

	var config = client.Config{}
	config.SetAccessKey("mui2W50H1j-OC4xD6PgQag")
	config.SetAccessSecret("PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc")

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	apkKey := "isVZBUvkFhv6oHxk_X-D0Q"
	versionCode := 2

	// 接口调用
	request := &client.ApkVersionRequest{
		ApkKey:      &apkKey,
		VersionCode: &versionCode,
	}
	Info, err := Client.ApkVersion(request)
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

// 获取 配置应用版本信息
func TestGetConfigurationVersion(t *testing.T) {

	var config = client.Config{}
	config.SetAccessKey("mui2W50H1j-OC4xD6PgQag")
	config.SetAccessSecret("PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc")

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	configurationKey := "q1hfB1VUQaK9VksTZGPU1Q"
	versionCode := 2

	// 接口调用
	request := &client.ConfigurationVersionRequest{
		ConfigurationKey: &configurationKey,
		VersionCode:      &versionCode,
	}
	Info, err := Client.ConfigurationVersion(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 获取Tauri应用版本信息
func TestGetTauriVersion(t *testing.T) {

	var config = client.Config{}
	config.SetAccessKey("mui2W50H1j-OC4xD6PgQag")
	config.SetAccessSecret("PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc")

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	tauriKey := "a0jtz0HUwL66r7gCGvbMKQ"
	VersionName := "0.1.30"
	target := "linux"
	arch := "x86_64"

	// 接口调用
	request := &client.TauriVersionRequest{
		TauriKey:    &tauriKey,
		VersionName: &VersionName,
		Target:      &target,
		Arch:        &arch,
	}

	Info, err := Client.TauriVersion(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 获取Electron应用版本信息
func TestGetElectronVersion(t *testing.T) {

	var config = client.Config{}
	config.SetAccessKey("mui2W50H1j-OC4xD6PgQag")
	config.SetAccessSecret("PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc")

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	electronKey := "kPUtUMDIjBhS48q5771pow"
	versionName := "1.5.2"
	platform := "linux"
	arch := "x64"

	// 接口调用
	request := &client.ElectronVersionRequest{
		ElectronKey: &electronKey,
		VersionName: &versionName,
		Platform:    &platform,
		Arch:        &arch,
	}

	Info, err := Client.ElectronVersion(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 获取 Windows应用升级内容
func TestGetWinUpgrade(t *testing.T) {

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

	winKey := "npJi367lttpwmD1goZ1yOQ"
	arch := "x64"
	versionCode := 1
	appointVersionCode := 0
	devModelKey := ""
	devKey := ""

	// 接口调用
	request := &client.WinUpgradeRequest{
		WinKey:             &winKey,
		Arch:               &arch,
		VersionCode:        &versionCode,
		AppointVersionCode: &appointVersionCode,
		DevModelKey:        &devModelKey,
		DevKey:             &devKey,
	}
	Info, err := Client.WinUpgrade(request)
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("info: ", Info)
	}

}

// 获取 Windows 应用版本信息
func TestGetWinVersion(t *testing.T) {

	var config = client.Config{}
	config.SetAccessKey("mui2W50H1j-OC4xD6PgQag")
	config.SetAccessSecret("PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc")

	Client, err := client.NewClient(&config)
	if err != nil {
		return
	}

	winKey := "npJi367lttpwmD1goZ1yOQ"
	arch := "x64"
	versionCode := 1

	// 接口调用
	request := &client.WinVersionRequest{
		WinKey:      &winKey,
		Arch:        &arch,
		VersionCode: &versionCode,
	}
	Info, err := Client.WinVersion(request)
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
	eventType := "app_start"
	appKey := "LOYlLXNy7wV3ySuh0XgtSg"
	target := "darwin"
	arch := "x86_64"
	devModelKey := ""
	devKey := ""
	versionCode := 0
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
