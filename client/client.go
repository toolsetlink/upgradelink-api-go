// This file is auto-generated, don't edit it. Thanks.
package client

import (
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  darabonbabase  "github.com/toolsetlink/darabonba-base-go/client"
  "github.com/alibabacloud-go/tea/tea"
)

type Config struct {
  AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty" require:"true"`
  AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty" require:"true"`
  Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
  Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty" require:"true"`
}

func (s Config) String() string {
  return tea.Prettify(s)
}

func (s Config) GoString() string {
  return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
  s.AccessKeyId = &v
  return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
  s.AccessKeySecret = &v
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

type ClientInterface interface {
}

type Client struct {
  AccessKeyId  *string
  AccessKeySecret  *string
  Protocol  *string
  Endpoint  *string
}

func NewClient(config *Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *Config)(_err error) {
  client.AccessKeyId = config.AccessKeyId
  client.AccessKeySecret = config.AccessKeySecret
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


func (client *Client) GetUrlUpgrade(request *UrlUpgradeRequest) (_result *UrlUpgradeResponse, _err error) {
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
      accessKeySecret := client.AccessKeySecret
      accessKeyId := client.AccessKeyId
      // 生成签名
      signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessKeySecret, timestamp, uri)
      request_.Protocol = tea.String("HTTP")
      request_.Method = tea.String("POST")
      request_.Pathname = tea.String("/v1/url/upgrade")
      request_.Headers = map[string]*string{
        "host": client.Endpoint,
        "content-type": tea.String("application/json"),
        "x-timestamp": timestamp,
        "x-nonce": nonce,
        "x-accesskey": accessKeyId,
        "x-signature": signature,
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
          "trace_id": tea.ToString(result["trace_id"]),
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

func (client *Client) GetFileUpgrade(request *FileUpgradeRequest) (_result *FileUpgradeResponse, _err error) {
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
      accessKeySecret := client.AccessKeySecret
      accessKeyId := client.AccessKeyId
      // 生成签名
      signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessKeySecret, timestamp, uri)
      request_.Protocol = tea.String("HTTP")
      request_.Method = tea.String("POST")
      request_.Pathname = tea.String("/v1/file/upgrade")
      request_.Headers = map[string]*string{
        "host": client.Endpoint,
        "content-type": tea.String("application/json"),
        "x-timestamp": timestamp,
        "x-nonce": nonce,
        "x-accesskey": accessKeyId,
        "x-signature": signature,
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
          "trace_id": tea.ToString(result["trace_id"]),
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


