// This file is auto-generated, don't edit it. Thanks.
package client

import (
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  darabonbabase  "github.com/toolsetlink/darabonba-base-go/client"
  "github.com/alibabacloud-go/tea/tea"
)

type Config struct {
  AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty" require:"true"`
  AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty" require:"true"`
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty" require:"true"`
}

func (s Config) String() string {
  return tea.Prettify(s)
}

func (s Config) GoString() string {
  return s.String()
}

func (s *Config) SetAccessKey(v string) *Config {
  s.AccessKey = &v
  return s
}

func (s *Config) SetAccessSecret(v string) *Config {
  s.AccessSecret = &v
  return s
}

func (s *Config) SetProtocol(v string) *Config {
  s.Protocol = &v
  return s
}

func (s *Config) SetEndpoint(v string) *Config {
  s.Endpoint = &v
  return s
}

type UrlUpgradeRequest struct {
  UrlKey *string `json:"urlKey,omitempty" xml:"urlKey,omitempty" require:"true"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  AppointVersionCode *int `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
  DevModelKey *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
  DevKey *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
}

func (s UrlUpgradeRequest) String() string {
  return tea.Prettify(s)
}

func (s UrlUpgradeRequest) GoString() string {
  return s.String()
}

func (s *UrlUpgradeRequest) SetUrlKey(v string) *UrlUpgradeRequest {
  s.UrlKey = &v
  return s
}

func (s *UrlUpgradeRequest) SetVersionCode(v int) *UrlUpgradeRequest {
  s.VersionCode = &v
  return s
}

func (s *UrlUpgradeRequest) SetAppointVersionCode(v int) *UrlUpgradeRequest {
  s.AppointVersionCode = &v
  return s
}

func (s *UrlUpgradeRequest) SetDevModelKey(v string) *UrlUpgradeRequest {
  s.DevModelKey = &v
  return s
}

func (s *UrlUpgradeRequest) SetDevKey(v string) *UrlUpgradeRequest {
  s.DevKey = &v
  return s
}

type UrlUpgradeDataResponse struct {
  UrlKey *string `json:"urlKey,omitempty" xml:"urlKey,omitempty" require:"true"`
  VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  UrlPath *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
  UpgradeType *int `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
  PromptUpgradeContent *string `json:"promptUpgradeContent,omitempty" xml:"promptUpgradeContent,omitempty" require:"true"`
}

func (s UrlUpgradeDataResponse) String() string {
  return tea.Prettify(s)
}

func (s UrlUpgradeDataResponse) GoString() string {
  return s.String()
}

func (s *UrlUpgradeDataResponse) SetUrlKey(v string) *UrlUpgradeDataResponse {
  s.UrlKey = &v
  return s
}

func (s *UrlUpgradeDataResponse) SetVersionName(v string) *UrlUpgradeDataResponse {
  s.VersionName = &v
  return s
}

func (s *UrlUpgradeDataResponse) SetVersionCode(v int) *UrlUpgradeDataResponse {
  s.VersionCode = &v
  return s
}

func (s *UrlUpgradeDataResponse) SetUrlPath(v string) *UrlUpgradeDataResponse {
  s.UrlPath = &v
  return s
}

func (s *UrlUpgradeDataResponse) SetUpgradeType(v int) *UrlUpgradeDataResponse {
  s.UpgradeType = &v
  return s
}

func (s *UrlUpgradeDataResponse) SetPromptUpgradeContent(v string) *UrlUpgradeDataResponse {
  s.PromptUpgradeContent = &v
  return s
}

type UrlUpgradeResponse struct {
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
  Data *UrlUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UrlUpgradeResponse) String() string {
  return tea.Prettify(s)
}

func (s UrlUpgradeResponse) GoString() string {
  return s.String()
}

func (s *UrlUpgradeResponse) SetCode(v int) *UrlUpgradeResponse {
  s.Code = &v
  return s
}

func (s *UrlUpgradeResponse) SetMsg(v string) *UrlUpgradeResponse {
  s.Msg = &v
  return s
}

func (s *UrlUpgradeResponse) SetTraceId(v string) *UrlUpgradeResponse {
  s.TraceId = &v
  return s
}

func (s *UrlUpgradeResponse) SetData(v *UrlUpgradeDataResponse) *UrlUpgradeResponse {
  s.Data = v
  return s
}

type FileUpgradeRequest struct {
  FileKey *string `json:"fileKey,omitempty" xml:"fileKey,omitempty" require:"true"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  AppointVersionCode *int `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
  DevModelKey *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
  DevKey *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
}

func (s FileUpgradeRequest) String() string {
  return tea.Prettify(s)
}

func (s FileUpgradeRequest) GoString() string {
  return s.String()
}

func (s *FileUpgradeRequest) SetFileKey(v string) *FileUpgradeRequest {
  s.FileKey = &v
  return s
}

func (s *FileUpgradeRequest) SetVersionCode(v int) *FileUpgradeRequest {
  s.VersionCode = &v
  return s
}

func (s *FileUpgradeRequest) SetAppointVersionCode(v int) *FileUpgradeRequest {
  s.AppointVersionCode = &v
  return s
}

func (s *FileUpgradeRequest) SetDevModelKey(v string) *FileUpgradeRequest {
  s.DevModelKey = &v
  return s
}

func (s *FileUpgradeRequest) SetDevKey(v string) *FileUpgradeRequest {
  s.DevKey = &v
  return s
}

type FileUpgradeDataResponse struct {
  FileKey *string `json:"fileKey,omitempty" xml:"fileKey,omitempty" require:"true"`
  VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  UrlPath *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
  UpgradeType *int `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
  PromptUpgradeContent *string `json:"promptUpgradeContent,omitempty" xml:"promptUpgradeContent,omitempty" require:"true"`
}

func (s FileUpgradeDataResponse) String() string {
  return tea.Prettify(s)
}

func (s FileUpgradeDataResponse) GoString() string {
  return s.String()
}

func (s *FileUpgradeDataResponse) SetFileKey(v string) *FileUpgradeDataResponse {
  s.FileKey = &v
  return s
}

func (s *FileUpgradeDataResponse) SetVersionName(v string) *FileUpgradeDataResponse {
  s.VersionName = &v
  return s
}

func (s *FileUpgradeDataResponse) SetVersionCode(v int) *FileUpgradeDataResponse {
  s.VersionCode = &v
  return s
}

func (s *FileUpgradeDataResponse) SetUrlPath(v string) *FileUpgradeDataResponse {
  s.UrlPath = &v
  return s
}

func (s *FileUpgradeDataResponse) SetUpgradeType(v int) *FileUpgradeDataResponse {
  s.UpgradeType = &v
  return s
}

func (s *FileUpgradeDataResponse) SetPromptUpgradeContent(v string) *FileUpgradeDataResponse {
  s.PromptUpgradeContent = &v
  return s
}

type FileUpgradeResponse struct {
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
  Data *FileUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s FileUpgradeResponse) String() string {
  return tea.Prettify(s)
}

func (s FileUpgradeResponse) GoString() string {
  return s.String()
}

func (s *FileUpgradeResponse) SetCode(v int) *FileUpgradeResponse {
  s.Code = &v
  return s
}

func (s *FileUpgradeResponse) SetMsg(v string) *FileUpgradeResponse {
  s.Msg = &v
  return s
}

func (s *FileUpgradeResponse) SetTraceId(v string) *FileUpgradeResponse {
  s.TraceId = &v
  return s
}

func (s *FileUpgradeResponse) SetData(v *FileUpgradeDataResponse) *FileUpgradeResponse {
  s.Data = v
  return s
}

type ApkUpgradeRequest struct {
  ApkKey *string `json:"apkKey,omitempty" xml:"apkKey,omitempty" require:"true"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  AppointVersionCode *int `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
  DevModelKey *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
  DevKey *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
}

func (s ApkUpgradeRequest) String() string {
  return tea.Prettify(s)
}

func (s ApkUpgradeRequest) GoString() string {
  return s.String()
}

func (s *ApkUpgradeRequest) SetApkKey(v string) *ApkUpgradeRequest {
  s.ApkKey = &v
  return s
}

func (s *ApkUpgradeRequest) SetVersionCode(v int) *ApkUpgradeRequest {
  s.VersionCode = &v
  return s
}

func (s *ApkUpgradeRequest) SetAppointVersionCode(v int) *ApkUpgradeRequest {
  s.AppointVersionCode = &v
  return s
}

func (s *ApkUpgradeRequest) SetDevModelKey(v string) *ApkUpgradeRequest {
  s.DevModelKey = &v
  return s
}

func (s *ApkUpgradeRequest) SetDevKey(v string) *ApkUpgradeRequest {
  s.DevKey = &v
  return s
}

type ApkUpgradeDataResponse struct {
  ApkKey *string `json:"apkKey,omitempty" xml:"apkKey,omitempty" require:"true"`
  PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
  VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  UrlPath *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
  UrlFileSize *int `json:"urlFileSize,omitempty" xml:"urlFileSize,omitempty" require:"true"`
  UrlFileMd5 *string `json:"urlFileMd5,omitempty" xml:"urlFileMd5,omitempty" require:"true"`
  UpgradeType *int `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
  PromptUpgradeContent *string `json:"promptUpgradeContent,omitempty" xml:"promptUpgradeContent,omitempty" require:"true"`
}

func (s ApkUpgradeDataResponse) String() string {
  return tea.Prettify(s)
}

func (s ApkUpgradeDataResponse) GoString() string {
  return s.String()
}

func (s *ApkUpgradeDataResponse) SetApkKey(v string) *ApkUpgradeDataResponse {
  s.ApkKey = &v
  return s
}

func (s *ApkUpgradeDataResponse) SetPackageName(v string) *ApkUpgradeDataResponse {
  s.PackageName = &v
  return s
}

func (s *ApkUpgradeDataResponse) SetVersionName(v string) *ApkUpgradeDataResponse {
  s.VersionName = &v
  return s
}

func (s *ApkUpgradeDataResponse) SetVersionCode(v int) *ApkUpgradeDataResponse {
  s.VersionCode = &v
  return s
}

func (s *ApkUpgradeDataResponse) SetUrlPath(v string) *ApkUpgradeDataResponse {
  s.UrlPath = &v
  return s
}

func (s *ApkUpgradeDataResponse) SetUrlFileSize(v int) *ApkUpgradeDataResponse {
  s.UrlFileSize = &v
  return s
}

func (s *ApkUpgradeDataResponse) SetUrlFileMd5(v string) *ApkUpgradeDataResponse {
  s.UrlFileMd5 = &v
  return s
}

func (s *ApkUpgradeDataResponse) SetUpgradeType(v int) *ApkUpgradeDataResponse {
  s.UpgradeType = &v
  return s
}

func (s *ApkUpgradeDataResponse) SetPromptUpgradeContent(v string) *ApkUpgradeDataResponse {
  s.PromptUpgradeContent = &v
  return s
}

type ApkUpgradeResponse struct {
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
  Data *ApkUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ApkUpgradeResponse) String() string {
  return tea.Prettify(s)
}

func (s ApkUpgradeResponse) GoString() string {
  return s.String()
}

func (s *ApkUpgradeResponse) SetCode(v int) *ApkUpgradeResponse {
  s.Code = &v
  return s
}

func (s *ApkUpgradeResponse) SetMsg(v string) *ApkUpgradeResponse {
  s.Msg = &v
  return s
}

func (s *ApkUpgradeResponse) SetTraceId(v string) *ApkUpgradeResponse {
  s.TraceId = &v
  return s
}

func (s *ApkUpgradeResponse) SetData(v *ApkUpgradeDataResponse) *ApkUpgradeResponse {
  s.Data = v
  return s
}

type ConfigurationUpgradeRequest struct {
  ConfigurationKey *string `json:"configurationKey,omitempty" xml:"configurationKey,omitempty" require:"true"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  AppointVersionCode *int `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
  DevModelKey *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
  DevKey *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
}

func (s ConfigurationUpgradeRequest) String() string {
  return tea.Prettify(s)
}

func (s ConfigurationUpgradeRequest) GoString() string {
  return s.String()
}

func (s *ConfigurationUpgradeRequest) SetConfigurationKey(v string) *ConfigurationUpgradeRequest {
  s.ConfigurationKey = &v
  return s
}

func (s *ConfigurationUpgradeRequest) SetVersionCode(v int) *ConfigurationUpgradeRequest {
  s.VersionCode = &v
  return s
}

func (s *ConfigurationUpgradeRequest) SetAppointVersionCode(v int) *ConfigurationUpgradeRequest {
  s.AppointVersionCode = &v
  return s
}

func (s *ConfigurationUpgradeRequest) SetDevModelKey(v string) *ConfigurationUpgradeRequest {
  s.DevModelKey = &v
  return s
}

func (s *ConfigurationUpgradeRequest) SetDevKey(v string) *ConfigurationUpgradeRequest {
  s.DevKey = &v
  return s
}

type ConfigurationUpgradeDataResponse struct {
  ConfigurationKey *string `json:"configurationKey,omitempty" xml:"configurationKey,omitempty" require:"true"`
  VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  UpgradeType *int `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
  PromptUpgradeContent *string `json:"promptUpgradeContent,omitempty" xml:"promptUpgradeContent,omitempty" require:"true"`
  Content interface{} `json:"content,omitempty" xml:"content,omitempty" require:"true"`
}

func (s ConfigurationUpgradeDataResponse) String() string {
  return tea.Prettify(s)
}

func (s ConfigurationUpgradeDataResponse) GoString() string {
  return s.String()
}

func (s *ConfigurationUpgradeDataResponse) SetConfigurationKey(v string) *ConfigurationUpgradeDataResponse {
  s.ConfigurationKey = &v
  return s
}

func (s *ConfigurationUpgradeDataResponse) SetVersionName(v string) *ConfigurationUpgradeDataResponse {
  s.VersionName = &v
  return s
}

func (s *ConfigurationUpgradeDataResponse) SetVersionCode(v int) *ConfigurationUpgradeDataResponse {
  s.VersionCode = &v
  return s
}

func (s *ConfigurationUpgradeDataResponse) SetUpgradeType(v int) *ConfigurationUpgradeDataResponse {
  s.UpgradeType = &v
  return s
}

func (s *ConfigurationUpgradeDataResponse) SetPromptUpgradeContent(v string) *ConfigurationUpgradeDataResponse {
  s.PromptUpgradeContent = &v
  return s
}

func (s *ConfigurationUpgradeDataResponse) SetContent(v interface{}) *ConfigurationUpgradeDataResponse {
  s.Content = v
  return s
}

type ConfigurationUpgradeResponse struct {
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
  Data *ConfigurationUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ConfigurationUpgradeResponse) String() string {
  return tea.Prettify(s)
}

func (s ConfigurationUpgradeResponse) GoString() string {
  return s.String()
}

func (s *ConfigurationUpgradeResponse) SetCode(v int) *ConfigurationUpgradeResponse {
  s.Code = &v
  return s
}

func (s *ConfigurationUpgradeResponse) SetMsg(v string) *ConfigurationUpgradeResponse {
  s.Msg = &v
  return s
}

func (s *ConfigurationUpgradeResponse) SetTraceId(v string) *ConfigurationUpgradeResponse {
  s.TraceId = &v
  return s
}

func (s *ConfigurationUpgradeResponse) SetData(v *ConfigurationUpgradeDataResponse) *ConfigurationUpgradeResponse {
  s.Data = v
  return s
}

type AppReportRequest struct {
  EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty" require:"true"`
  AppKey *string `json:"appKey,omitempty" xml:"appKey,omitempty" require:"true"`
  Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
  EventData *AppReportRequestEventData `json:"eventData,omitempty" xml:"eventData,omitempty" require:"true" type:"Struct"`
}

func (s AppReportRequest) String() string {
  return tea.Prettify(s)
}

func (s AppReportRequest) GoString() string {
  return s.String()
}

func (s *AppReportRequest) SetEventType(v string) *AppReportRequest {
  s.EventType = &v
  return s
}

func (s *AppReportRequest) SetAppKey(v string) *AppReportRequest {
  s.AppKey = &v
  return s
}

func (s *AppReportRequest) SetTimestamp(v string) *AppReportRequest {
  s.Timestamp = &v
  return s
}

func (s *AppReportRequest) SetEventData(v *AppReportRequestEventData) *AppReportRequest {
  s.EventData = v
  return s
}

type AppReportRequestEventData struct {
  LaunchTime *string `json:"launchTime,omitempty" xml:"launchTime,omitempty"`
  VersionCode *int `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
  DevModelKey *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty"`
  DevKey *string `json:"devKey,omitempty" xml:"devKey,omitempty"`
  Target *string `json:"target,omitempty" xml:"target,omitempty"`
  Arch *string `json:"arch,omitempty" xml:"arch,omitempty"`
  DownloadVersionCode *int `json:"downloadVersionCode,omitempty" xml:"downloadVersionCode,omitempty"`
  UpgradeVersionCode *int `json:"upgradeVersionCode,omitempty" xml:"upgradeVersionCode,omitempty"`
  Code *int `json:"code,omitempty" xml:"code,omitempty"`
}

func (s AppReportRequestEventData) String() string {
  return tea.Prettify(s)
}

func (s AppReportRequestEventData) GoString() string {
  return s.String()
}

func (s *AppReportRequestEventData) SetLaunchTime(v string) *AppReportRequestEventData {
  s.LaunchTime = &v
  return s
}

func (s *AppReportRequestEventData) SetVersionCode(v int) *AppReportRequestEventData {
  s.VersionCode = &v
  return s
}

func (s *AppReportRequestEventData) SetDevModelKey(v string) *AppReportRequestEventData {
  s.DevModelKey = &v
  return s
}

func (s *AppReportRequestEventData) SetDevKey(v string) *AppReportRequestEventData {
  s.DevKey = &v
  return s
}

func (s *AppReportRequestEventData) SetTarget(v string) *AppReportRequestEventData {
  s.Target = &v
  return s
}

func (s *AppReportRequestEventData) SetArch(v string) *AppReportRequestEventData {
  s.Arch = &v
  return s
}

func (s *AppReportRequestEventData) SetDownloadVersionCode(v int) *AppReportRequestEventData {
  s.DownloadVersionCode = &v
  return s
}

func (s *AppReportRequestEventData) SetUpgradeVersionCode(v int) *AppReportRequestEventData {
  s.UpgradeVersionCode = &v
  return s
}

func (s *AppReportRequestEventData) SetCode(v int) *AppReportRequestEventData {
  s.Code = &v
  return s
}

type AppReportResponse struct {
  Code *int `json:"code,omitempty" xml:"code,omitempty" require:"true"`
  Msg *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
  Docs *string `json:"docs,omitempty" xml:"docs,omitempty" require:"true"`
  TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
}

func (s AppReportResponse) String() string {
  return tea.Prettify(s)
}

func (s AppReportResponse) GoString() string {
  return s.String()
}

func (s *AppReportResponse) SetCode(v int) *AppReportResponse {
  s.Code = &v
  return s
}

func (s *AppReportResponse) SetMsg(v string) *AppReportResponse {
  s.Msg = &v
  return s
}

func (s *AppReportResponse) SetDocs(v string) *AppReportResponse {
  s.Docs = &v
  return s
}

func (s *AppReportResponse) SetTraceId(v string) *AppReportResponse {
  s.TraceId = &v
  return s
}

type ClientInterface interface {
}

type Client struct {
  AccessKey  *string
  AccessSecret  *string
  Protocol  *string
  Endpoint  *string
}

func NewClient(config *Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *Config)(_err error) {
  client.AccessKey = config.AccessKey
  client.AccessSecret = config.AccessSecret
  if tea.BoolValue(util.EqualString(config.Protocol, tea.String("HTTPS"))) {
    client.Protocol = tea.String("HTTPS")
  } else {
    client.Protocol = tea.String("HTTP")
  }

  if tea.BoolValue(util.Empty(config.Endpoint)) {
    client.Endpoint = tea.String("api.upgrade.toolsetlink.com")
  } else {
    client.Endpoint = config.Endpoint
  }

  return nil
}


func (client *Client) UrlUpgrade(request *UrlUpgradeRequest) (_result *UrlUpgradeResponse, _err error) {
  _err = tea.Validate(request)
  if _err != nil {
    return _result, _err
  }
  _runtime := map[string]interface{}{
    "timeout": 10000,
    // 10s 的过期时间
  }

  _resp := &UrlUpgradeResponse{}
  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _resp, _err = func()(*UrlUpgradeResponse, error){
      request_ := tea.NewRequest()
      // 序列化请求体
      bodyStr := util.ToJSONString(request)
      // 生成请求参数
      timestamp := darabonbabase.TimeRFC3339()
      nonce := darabonbabase.GenerateNonce()
      uri := tea.String("/v1/url/upgrade")
      accessKey := client.AccessKey
      accessSecret := client.AccessSecret
      // 生成签名
      signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
      request_.Protocol = client.Protocol
      request_.Method = tea.String("POST")
      request_.Pathname = tea.String("/v1/url/upgrade")
      request_.Headers = map[string]*string{
        "host": client.Endpoint,
        "content-type": tea.String("application/json"),
        "x-Timestamp": timestamp,
        "x-Nonce": nonce,
        "x-AccessKey": accessKey,
        "x-Signature": signature,
      }
      request_.Body = tea.ToReader(bodyStr)
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }
      assertAsMapTmp, err := util.ReadAsJSON(response_.Body)
      if err != nil {
        _err = err
        return _result, _err
      }
      result, _err := util.AssertAsMap(assertAsMapTmp)
      if _err != nil {
        return _result, _err
      }

      if !tea.BoolValue(util.EqualNumber(response_.StatusCode, tea.Int(200))) {
        _err = tea.NewSDKError(map[string]interface{}{
          "statusCode": tea.ToString(tea.IntValue(response_.StatusCode)),
          "code": tea.ToString(result["code"]),
          "message": tea.ToString(result["msg"]),
          "docs": tea.ToString(result["docs"]),
          "traceId": tea.ToString(result["traceId"]),
        })
        return _result, _err
      }

      _result = &UrlUpgradeResponse{}
      _err = tea.Convert(tea.ToMap(result), &_result)
      return _result, _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _resp, _err
}

func (client *Client) FileUpgrade(request *FileUpgradeRequest) (_result *FileUpgradeResponse, _err error) {
  _err = tea.Validate(request)
  if _err != nil {
    return _result, _err
  }
  _runtime := map[string]interface{}{
    "timeout": 10000,
    // 10s 的过期时间
  }

  _resp := &FileUpgradeResponse{}
  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _resp, _err = func()(*FileUpgradeResponse, error){
      request_ := tea.NewRequest()
      // 序列化请求体
      bodyStr := util.ToJSONString(request)
      // 生成请求参数
      timestamp := darabonbabase.TimeRFC3339()
      nonce := darabonbabase.GenerateNonce()
      uri := tea.String("/v1/file/upgrade")
      accessKey := client.AccessKey
      accessSecret := client.AccessSecret
      // 生成签名
      signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
      request_.Protocol = client.Protocol
      request_.Method = tea.String("POST")
      request_.Pathname = tea.String("/v1/file/upgrade")
      request_.Headers = map[string]*string{
        "host": client.Endpoint,
        "content-type": tea.String("application/json"),
        "x-Timestamp": timestamp,
        "x-Nonce": nonce,
        "x-AccessKey": accessKey,
        "x-Signature": signature,
      }
      request_.Body = tea.ToReader(bodyStr)
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }
      assertAsMapTmp, err := util.ReadAsJSON(response_.Body)
      if err != nil {
        _err = err
        return _result, _err
      }
      result, _err := util.AssertAsMap(assertAsMapTmp)
      if _err != nil {
        return _result, _err
      }

      if !tea.BoolValue(util.EqualNumber(response_.StatusCode, tea.Int(200))) {
        _err = tea.NewSDKError(map[string]interface{}{
          "statusCode": tea.ToString(tea.IntValue(response_.StatusCode)),
          "code": tea.ToString(result["code"]),
          "message": tea.ToString(result["msg"]),
          "docs": tea.ToString(result["docs"]),
          "traceId": tea.ToString(result["traceId"]),
        })
        return _result, _err
      }

      _result = &FileUpgradeResponse{}
      _err = tea.Convert(tea.ToMap(result), &_result)
      return _result, _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _resp, _err
}

func (client *Client) ApkUpgrade(request *ApkUpgradeRequest) (_result *ApkUpgradeResponse, _err error) {
  _err = tea.Validate(request)
  if _err != nil {
    return _result, _err
  }
  _runtime := map[string]interface{}{
    "timeout": 10000,
    // 10s 的过期时间
  }

  _resp := &ApkUpgradeResponse{}
  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _resp, _err = func()(*ApkUpgradeResponse, error){
      request_ := tea.NewRequest()
      // 序列化请求体
      bodyStr := util.ToJSONString(request)
      // 生成请求参数
      timestamp := darabonbabase.TimeRFC3339()
      nonce := darabonbabase.GenerateNonce()
      uri := tea.String("/v1/apk/upgrade")
      accessKey := client.AccessKey
      accessSecret := client.AccessSecret
      // 生成签名
      signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
      request_.Protocol = client.Protocol
      request_.Method = tea.String("POST")
      request_.Pathname = tea.String("/v1/apk/upgrade")
      request_.Headers = map[string]*string{
        "host": client.Endpoint,
        "content-type": tea.String("application/json"),
        "x-Timestamp": timestamp,
        "x-Nonce": nonce,
        "x-AccessKey": accessKey,
        "x-Signature": signature,
      }
      request_.Body = tea.ToReader(bodyStr)
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }
      assertAsMapTmp, err := util.ReadAsJSON(response_.Body)
      if err != nil {
        _err = err
        return _result, _err
      }
      result, _err := util.AssertAsMap(assertAsMapTmp)
      if _err != nil {
        return _result, _err
      }

      if !tea.BoolValue(util.EqualNumber(response_.StatusCode, tea.Int(200))) {
        _err = tea.NewSDKError(map[string]interface{}{
          "statusCode": tea.ToString(tea.IntValue(response_.StatusCode)),
          "code": tea.ToString(result["code"]),
          "message": tea.ToString(result["msg"]),
          "docs": tea.ToString(result["docs"]),
          "traceId": tea.ToString(result["traceId"]),
        })
        return _result, _err
      }

      _result = &ApkUpgradeResponse{}
      _err = tea.Convert(tea.ToMap(result), &_result)
      return _result, _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _resp, _err
}

func (client *Client) ConfigurationUpgrade(request *ConfigurationUpgradeRequest) (_result *ConfigurationUpgradeResponse, _err error) {
  _err = tea.Validate(request)
  if _err != nil {
    return _result, _err
  }
  _runtime := map[string]interface{}{
    "timeout": 10000,
    // 10s 的过期时间
  }

  _resp := &ConfigurationUpgradeResponse{}
  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _resp, _err = func()(*ConfigurationUpgradeResponse, error){
      request_ := tea.NewRequest()
      // 序列化请求体
      bodyStr := util.ToJSONString(request)
      // 生成请求参数
      timestamp := darabonbabase.TimeRFC3339()
      nonce := darabonbabase.GenerateNonce()
      uri := tea.String("/v1/configuration/upgrade")
      accessKey := client.AccessKey
      accessSecret := client.AccessSecret
      // 生成签名
      signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
      request_.Protocol = client.Protocol
      request_.Method = tea.String("POST")
      request_.Pathname = tea.String("/v1/configuration/upgrade")
      request_.Headers = map[string]*string{
        "host": client.Endpoint,
        "content-type": tea.String("application/json"),
        "x-Timestamp": timestamp,
        "x-Nonce": nonce,
        "x-AccessKey": accessKey,
        "x-Signature": signature,
      }
      request_.Body = tea.ToReader(bodyStr)
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }
      assertAsMapTmp, err := util.ReadAsJSON(response_.Body)
      if err != nil {
        _err = err
        return _result, _err
      }
      result, _err := util.AssertAsMap(assertAsMapTmp)
      if _err != nil {
        return _result, _err
      }

      if !tea.BoolValue(util.EqualNumber(response_.StatusCode, tea.Int(200))) {
        _err = tea.NewSDKError(map[string]interface{}{
          "statusCode": tea.ToString(tea.IntValue(response_.StatusCode)),
          "code": tea.ToString(result["code"]),
          "message": tea.ToString(result["msg"]),
          "docs": tea.ToString(result["docs"]),
          "traceId": tea.ToString(result["traceId"]),
        })
        return _result, _err
      }

      _result = &ConfigurationUpgradeResponse{}
      _err = tea.Convert(tea.ToMap(result), &_result)
      return _result, _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _resp, _err
}

func (client *Client) AppReport(request *AppReportRequest) (_result *AppReportResponse, _err error) {
  _err = tea.Validate(request)
  if _err != nil {
    return _result, _err
  }
  _runtime := map[string]interface{}{
    "timeout": 10000,
    // 10s 的过期时间
  }

  _resp := &AppReportResponse{}
  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _resp, _err = func()(*AppReportResponse, error){
      request_ := tea.NewRequest()
      // 序列化请求体
      bodyStr := util.ToJSONString(request)
      // 生成请求参数
      timestamp := darabonbabase.TimeRFC3339()
      nonce := darabonbabase.GenerateNonce()
      uri := tea.String("/v1/app/report")
      accessKey := client.AccessKey
      accessSecret := client.AccessSecret
      // 生成签名
      signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
      request_.Protocol = client.Protocol
      request_.Method = tea.String("POST")
      request_.Pathname = tea.String("/v1/app/report")
      request_.Headers = map[string]*string{
        "host": client.Endpoint,
        "content-type": tea.String("application/json"),
        "x-Timestamp": timestamp,
        "x-Nonce": nonce,
        "x-AccessKey": accessKey,
        "x-Signature": signature,
      }
      request_.Body = tea.ToReader(bodyStr)
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }
      assertAsMapTmp, err := util.ReadAsJSON(response_.Body)
      if err != nil {
        _err = err
        return _result, _err
      }
      result, _err := util.AssertAsMap(assertAsMapTmp)
      if _err != nil {
        return _result, _err
      }

      if !tea.BoolValue(util.EqualNumber(response_.StatusCode, tea.Int(200))) {
        _err = tea.NewSDKError(map[string]interface{}{
          "statusCode": tea.ToString(tea.IntValue(response_.StatusCode)),
          "code": tea.ToString(result["code"]),
          "message": tea.ToString(result["msg"]),
          "docs": tea.ToString(result["docs"]),
          "traceId": tea.ToString(result["traceId"]),
        })
        return _result, _err
      }

      _result = &AppReportResponse{}
      _err = tea.Convert(tea.ToMap(result), &_result)
      return _result, _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _resp, _err
}


