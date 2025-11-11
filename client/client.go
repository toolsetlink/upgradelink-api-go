// This file is auto-generated, don't edit it. Thanks.
package client

import (
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	darabonbabase "github.com/toolsetlink/darabonba-base-go/client"
)

type Config struct {
	AccessKey    *string `json:"accessKey,omitempty" xml:"accessKey,omitempty" require:"true"`
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty" require:"true"`
	Protocol     *string `json:"protocol,omitempty" xml:"protocol,omitempty" require:"true"`
	Endpoint     *string `json:"endpoint,omitempty" xml:"endpoint,omitempty" require:"true"`
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
	UrlKey             *string `json:"urlKey,omitempty" xml:"urlKey,omitempty" require:"true"`
	VersionCode        *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	AppointVersionCode *int    `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
	DevModelKey        *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
	DevKey             *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
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
	UrlKey               *string `json:"urlKey,omitempty" xml:"urlKey,omitempty" require:"true"`
	VersionName          *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode          *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	UrlPath              *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
	UpgradeType          *int    `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
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
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *UrlUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

type UrlVersionRequest struct {
	UrlKey      *string `json:"urlKey,omitempty" xml:"urlKey,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
}

func (s UrlVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UrlVersionRequest) GoString() string {
	return s.String()
}

func (s *UrlVersionRequest) SetUrlKey(v string) *UrlVersionRequest {
	s.UrlKey = &v
	return s
}

func (s *UrlVersionRequest) SetVersionCode(v int) *UrlVersionRequest {
	s.VersionCode = &v
	return s
}

type UrlVersionDataResponse struct {
	UrlKey      *string `json:"urlKey,omitempty" xml:"urlKey,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s UrlVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlVersionDataResponse) GoString() string {
	return s.String()
}

func (s *UrlVersionDataResponse) SetUrlKey(v string) *UrlVersionDataResponse {
	s.UrlKey = &v
	return s
}

func (s *UrlVersionDataResponse) SetVersionName(v string) *UrlVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *UrlVersionDataResponse) SetVersionCode(v int) *UrlVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *UrlVersionDataResponse) SetDescription(v string) *UrlVersionDataResponse {
	s.Description = &v
	return s
}

type UrlVersionResponse struct {
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *UrlVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s UrlVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UrlVersionResponse) GoString() string {
	return s.String()
}

func (s *UrlVersionResponse) SetCode(v int) *UrlVersionResponse {
	s.Code = &v
	return s
}

func (s *UrlVersionResponse) SetMsg(v string) *UrlVersionResponse {
	s.Msg = &v
	return s
}

func (s *UrlVersionResponse) SetTraceId(v string) *UrlVersionResponse {
	s.TraceId = &v
	return s
}

func (s *UrlVersionResponse) SetData(v *UrlVersionDataResponse) *UrlVersionResponse {
	s.Data = v
	return s
}

type FileUpgradeRequest struct {
	FileKey            *string `json:"fileKey,omitempty" xml:"fileKey,omitempty" require:"true"`
	VersionCode        *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	AppointVersionCode *int    `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
	DevModelKey        *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
	DevKey             *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
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
	FileKey              *string `json:"fileKey,omitempty" xml:"fileKey,omitempty" require:"true"`
	VersionName          *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode          *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	UrlPath              *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
	UpgradeType          *int    `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
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
	Code    *int                     `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                  `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                  `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *FileUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

type FileVersionRequest struct {
	FileKey     *string `json:"fileKey,omitempty" xml:"fileKey,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
}

func (s FileVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s FileVersionRequest) GoString() string {
	return s.String()
}

func (s *FileVersionRequest) SetFileKey(v string) *FileVersionRequest {
	s.FileKey = &v
	return s
}

func (s *FileVersionRequest) SetVersionCode(v int) *FileVersionRequest {
	s.VersionCode = &v
	return s
}

type FileVersionDataResponse struct {
	FileKey     *string `json:"fileKey,omitempty" xml:"fileKey,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s FileVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s FileVersionDataResponse) GoString() string {
	return s.String()
}

func (s *FileVersionDataResponse) SetFileKey(v string) *FileVersionDataResponse {
	s.FileKey = &v
	return s
}

func (s *FileVersionDataResponse) SetVersionName(v string) *FileVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *FileVersionDataResponse) SetVersionCode(v int) *FileVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *FileVersionDataResponse) SetDescription(v string) *FileVersionDataResponse {
	s.Description = &v
	return s
}

type FileVersionResponse struct {
	Code    *int                     `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                  `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                  `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *FileVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s FileVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s FileVersionResponse) GoString() string {
	return s.String()
}

func (s *FileVersionResponse) SetCode(v int) *FileVersionResponse {
	s.Code = &v
	return s
}

func (s *FileVersionResponse) SetMsg(v string) *FileVersionResponse {
	s.Msg = &v
	return s
}

func (s *FileVersionResponse) SetTraceId(v string) *FileVersionResponse {
	s.TraceId = &v
	return s
}

func (s *FileVersionResponse) SetData(v *FileVersionDataResponse) *FileVersionResponse {
	s.Data = v
	return s
}

type ApkUpgradeRequest struct {
	ApkKey             *string `json:"apkKey,omitempty" xml:"apkKey,omitempty" require:"true"`
	VersionCode        *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	AppointVersionCode *int    `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
	DevModelKey        *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
	DevKey             *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
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
	ApkKey               *string `json:"apkKey,omitempty" xml:"apkKey,omitempty" require:"true"`
	PackageName          *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
	VersionName          *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode          *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	UrlPath              *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
	UrlFileSize          *int    `json:"urlFileSize,omitempty" xml:"urlFileSize,omitempty" require:"true"`
	UrlFileMd5           *string `json:"urlFileMd5,omitempty" xml:"urlFileMd5,omitempty" require:"true"`
	UpgradeType          *int    `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
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
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *ApkUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

type ApkVersionRequest struct {
	ApkKey      *string `json:"apkKey,omitempty" xml:"apkKey,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
}

func (s ApkVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ApkVersionRequest) GoString() string {
	return s.String()
}

func (s *ApkVersionRequest) SetApkKey(v string) *ApkVersionRequest {
	s.ApkKey = &v
	return s
}

func (s *ApkVersionRequest) SetVersionCode(v int) *ApkVersionRequest {
	s.VersionCode = &v
	return s
}

type ApkVersionDataResponse struct {
	ApkKey      *string `json:"apkKey,omitempty" xml:"apkKey,omitempty" require:"true"`
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ApkVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ApkVersionDataResponse) GoString() string {
	return s.String()
}

func (s *ApkVersionDataResponse) SetApkKey(v string) *ApkVersionDataResponse {
	s.ApkKey = &v
	return s
}

func (s *ApkVersionDataResponse) SetPackageName(v string) *ApkVersionDataResponse {
	s.PackageName = &v
	return s
}

func (s *ApkVersionDataResponse) SetVersionName(v string) *ApkVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *ApkVersionDataResponse) SetVersionCode(v int) *ApkVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *ApkVersionDataResponse) SetDescription(v string) *ApkVersionDataResponse {
	s.Description = &v
	return s
}

type ApkVersionResponse struct {
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *ApkVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ApkVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ApkVersionResponse) GoString() string {
	return s.String()
}

func (s *ApkVersionResponse) SetCode(v int) *ApkVersionResponse {
	s.Code = &v
	return s
}

func (s *ApkVersionResponse) SetMsg(v string) *ApkVersionResponse {
	s.Msg = &v
	return s
}

func (s *ApkVersionResponse) SetTraceId(v string) *ApkVersionResponse {
	s.TraceId = &v
	return s
}

func (s *ApkVersionResponse) SetData(v *ApkVersionDataResponse) *ApkVersionResponse {
	s.Data = v
	return s
}

type ConfigurationUpgradeRequest struct {
	ConfigurationKey   *string `json:"configurationKey,omitempty" xml:"configurationKey,omitempty" require:"true"`
	VersionCode        *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	AppointVersionCode *int    `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
	DevModelKey        *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
	DevKey             *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
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
	ConfigurationKey     *string     `json:"configurationKey,omitempty" xml:"configurationKey,omitempty" require:"true"`
	VersionName          *string     `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode          *int        `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	UpgradeType          *int        `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
	PromptUpgradeContent *string     `json:"promptUpgradeContent,omitempty" xml:"promptUpgradeContent,omitempty" require:"true"`
	Content              interface{} `json:"content,omitempty" xml:"content,omitempty" require:"true"`
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
	Code    *int                              `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                           `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                           `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *ConfigurationUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
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

type ConfigurationVersionRequest struct {
	ConfigurationKey *string `json:"configurationKey,omitempty" xml:"configurationKey,omitempty" require:"true"`
	VersionCode      *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
}

func (s ConfigurationVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationVersionRequest) GoString() string {
	return s.String()
}

func (s *ConfigurationVersionRequest) SetConfigurationKey(v string) *ConfigurationVersionRequest {
	s.ConfigurationKey = &v
	return s
}

func (s *ConfigurationVersionRequest) SetVersionCode(v int) *ConfigurationVersionRequest {
	s.VersionCode = &v
	return s
}

type ConfigurationVersionDataResponse struct {
	ConfigurationKey *string `json:"configurationKey,omitempty" xml:"configurationKey,omitempty" require:"true"`
	VersionName      *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode      *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Description      *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ConfigurationVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationVersionDataResponse) GoString() string {
	return s.String()
}

func (s *ConfigurationVersionDataResponse) SetConfigurationKey(v string) *ConfigurationVersionDataResponse {
	s.ConfigurationKey = &v
	return s
}

func (s *ConfigurationVersionDataResponse) SetVersionName(v string) *ConfigurationVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *ConfigurationVersionDataResponse) SetVersionCode(v int) *ConfigurationVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *ConfigurationVersionDataResponse) SetDescription(v string) *ConfigurationVersionDataResponse {
	s.Description = &v
	return s
}

type ConfigurationVersionResponse struct {
	Code    *int                              `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                           `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                           `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *ConfigurationVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ConfigurationVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationVersionResponse) GoString() string {
	return s.String()
}

func (s *ConfigurationVersionResponse) SetCode(v int) *ConfigurationVersionResponse {
	s.Code = &v
	return s
}

func (s *ConfigurationVersionResponse) SetMsg(v string) *ConfigurationVersionResponse {
	s.Msg = &v
	return s
}

func (s *ConfigurationVersionResponse) SetTraceId(v string) *ConfigurationVersionResponse {
	s.TraceId = &v
	return s
}

func (s *ConfigurationVersionResponse) SetData(v *ConfigurationVersionDataResponse) *ConfigurationVersionResponse {
	s.Data = v
	return s
}

type TauriVersionRequest struct {
	TauriKey    *string `json:"tauriKey,omitempty" xml:"tauriKey,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	Target      *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
	Arch        *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
}

func (s TauriVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s TauriVersionRequest) GoString() string {
	return s.String()
}

func (s *TauriVersionRequest) SetTauriKey(v string) *TauriVersionRequest {
	s.TauriKey = &v
	return s
}

func (s *TauriVersionRequest) SetVersionName(v string) *TauriVersionRequest {
	s.VersionName = &v
	return s
}

func (s *TauriVersionRequest) SetTarget(v string) *TauriVersionRequest {
	s.Target = &v
	return s
}

func (s *TauriVersionRequest) SetArch(v string) *TauriVersionRequest {
	s.Arch = &v
	return s
}

type TauriVersionDataResponse struct {
	TauriKey    *string `json:"tauriKey,omitempty" xml:"tauriKey,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Target      *string `json:"target,omitempty" xml:"target,omitempty" require:"true"`
	Arch        *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s TauriVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s TauriVersionDataResponse) GoString() string {
	return s.String()
}

func (s *TauriVersionDataResponse) SetTauriKey(v string) *TauriVersionDataResponse {
	s.TauriKey = &v
	return s
}

func (s *TauriVersionDataResponse) SetVersionName(v string) *TauriVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *TauriVersionDataResponse) SetVersionCode(v int) *TauriVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *TauriVersionDataResponse) SetTarget(v string) *TauriVersionDataResponse {
	s.Target = &v
	return s
}

func (s *TauriVersionDataResponse) SetArch(v string) *TauriVersionDataResponse {
	s.Arch = &v
	return s
}

func (s *TauriVersionDataResponse) SetDescription(v string) *TauriVersionDataResponse {
	s.Description = &v
	return s
}

type TauriVersionResponse struct {
	Code    *int                      `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                   `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                   `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *TauriVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s TauriVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s TauriVersionResponse) GoString() string {
	return s.String()
}

func (s *TauriVersionResponse) SetCode(v int) *TauriVersionResponse {
	s.Code = &v
	return s
}

func (s *TauriVersionResponse) SetMsg(v string) *TauriVersionResponse {
	s.Msg = &v
	return s
}

func (s *TauriVersionResponse) SetTraceId(v string) *TauriVersionResponse {
	s.TraceId = &v
	return s
}

func (s *TauriVersionResponse) SetData(v *TauriVersionDataResponse) *TauriVersionResponse {
	s.Data = v
	return s
}

type ElectronVersionRequest struct {
	ElectronKey *string `json:"electronKey,omitempty" xml:"electronKey,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	Platform    *string `json:"platform,omitempty" xml:"platform,omitempty" require:"true"`
	Arch        *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
}

func (s ElectronVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ElectronVersionRequest) GoString() string {
	return s.String()
}

func (s *ElectronVersionRequest) SetElectronKey(v string) *ElectronVersionRequest {
	s.ElectronKey = &v
	return s
}

func (s *ElectronVersionRequest) SetVersionName(v string) *ElectronVersionRequest {
	s.VersionName = &v
	return s
}

func (s *ElectronVersionRequest) SetPlatform(v string) *ElectronVersionRequest {
	s.Platform = &v
	return s
}

func (s *ElectronVersionRequest) SetArch(v string) *ElectronVersionRequest {
	s.Arch = &v
	return s
}

type ElectronVersionDataResponse struct {
	ElectronKey *string `json:"electronKey,omitempty" xml:"electronKey,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Platform    *string `json:"platform,omitempty" xml:"platform,omitempty" require:"true"`
	Arch        *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s ElectronVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s ElectronVersionDataResponse) GoString() string {
	return s.String()
}

func (s *ElectronVersionDataResponse) SetElectronKey(v string) *ElectronVersionDataResponse {
	s.ElectronKey = &v
	return s
}

func (s *ElectronVersionDataResponse) SetVersionName(v string) *ElectronVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *ElectronVersionDataResponse) SetVersionCode(v int) *ElectronVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *ElectronVersionDataResponse) SetPlatform(v string) *ElectronVersionDataResponse {
	s.Platform = &v
	return s
}

func (s *ElectronVersionDataResponse) SetArch(v string) *ElectronVersionDataResponse {
	s.Arch = &v
	return s
}

func (s *ElectronVersionDataResponse) SetDescription(v string) *ElectronVersionDataResponse {
	s.Description = &v
	return s
}

type ElectronVersionResponse struct {
	Code    *int                         `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                      `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                      `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *ElectronVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s ElectronVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ElectronVersionResponse) GoString() string {
	return s.String()
}

func (s *ElectronVersionResponse) SetCode(v int) *ElectronVersionResponse {
	s.Code = &v
	return s
}

func (s *ElectronVersionResponse) SetMsg(v string) *ElectronVersionResponse {
	s.Msg = &v
	return s
}

func (s *ElectronVersionResponse) SetTraceId(v string) *ElectronVersionResponse {
	s.TraceId = &v
	return s
}

func (s *ElectronVersionResponse) SetData(v *ElectronVersionDataResponse) *ElectronVersionResponse {
	s.Data = v
	return s
}

type LnxUpgradeRequest struct {
	LnxKey             *string `json:"lnxKey,omitempty" xml:"lnxKey,omitempty" require:"true"`
	Arch               *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
	VersionCode        *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	AppointVersionCode *int    `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
	DevModelKey        *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
	DevKey             *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
}

func (s LnxUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s LnxUpgradeRequest) GoString() string {
	return s.String()
}

func (s *LnxUpgradeRequest) SetLnxKey(v string) *LnxUpgradeRequest {
	s.LnxKey = &v
	return s
}

func (s *LnxUpgradeRequest) SetArch(v string) *LnxUpgradeRequest {
	s.Arch = &v
	return s
}

func (s *LnxUpgradeRequest) SetVersionCode(v int) *LnxUpgradeRequest {
	s.VersionCode = &v
	return s
}

func (s *LnxUpgradeRequest) SetAppointVersionCode(v int) *LnxUpgradeRequest {
	s.AppointVersionCode = &v
	return s
}

func (s *LnxUpgradeRequest) SetDevModelKey(v string) *LnxUpgradeRequest {
	s.DevModelKey = &v
	return s
}

func (s *LnxUpgradeRequest) SetDevKey(v string) *LnxUpgradeRequest {
	s.DevKey = &v
	return s
}

type LnxUpgradeDataResponse struct {
	LnxKey               *string `json:"lnxKey,omitempty" xml:"lnxKey,omitempty" require:"true"`
	PackageName          *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
	VersionName          *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode          *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	UrlPath              *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
	UrlFileSize          *int    `json:"urlFileSize,omitempty" xml:"urlFileSize,omitempty" require:"true"`
	UrlFileMd5           *string `json:"urlFileMd5,omitempty" xml:"urlFileMd5,omitempty" require:"true"`
	UpgradeType          *int    `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
	PromptUpgradeContent *string `json:"promptUpgradeContent,omitempty" xml:"promptUpgradeContent,omitempty" require:"true"`
}

func (s LnxUpgradeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s LnxUpgradeDataResponse) GoString() string {
	return s.String()
}

func (s *LnxUpgradeDataResponse) SetLnxKey(v string) *LnxUpgradeDataResponse {
	s.LnxKey = &v
	return s
}

func (s *LnxUpgradeDataResponse) SetPackageName(v string) *LnxUpgradeDataResponse {
	s.PackageName = &v
	return s
}

func (s *LnxUpgradeDataResponse) SetVersionName(v string) *LnxUpgradeDataResponse {
	s.VersionName = &v
	return s
}

func (s *LnxUpgradeDataResponse) SetVersionCode(v int) *LnxUpgradeDataResponse {
	s.VersionCode = &v
	return s
}

func (s *LnxUpgradeDataResponse) SetUrlPath(v string) *LnxUpgradeDataResponse {
	s.UrlPath = &v
	return s
}

func (s *LnxUpgradeDataResponse) SetUrlFileSize(v int) *LnxUpgradeDataResponse {
	s.UrlFileSize = &v
	return s
}

func (s *LnxUpgradeDataResponse) SetUrlFileMd5(v string) *LnxUpgradeDataResponse {
	s.UrlFileMd5 = &v
	return s
}

func (s *LnxUpgradeDataResponse) SetUpgradeType(v int) *LnxUpgradeDataResponse {
	s.UpgradeType = &v
	return s
}

func (s *LnxUpgradeDataResponse) SetPromptUpgradeContent(v string) *LnxUpgradeDataResponse {
	s.PromptUpgradeContent = &v
	return s
}

type LnxUpgradeResponse struct {
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *LnxUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s LnxUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s LnxUpgradeResponse) GoString() string {
	return s.String()
}

func (s *LnxUpgradeResponse) SetCode(v int) *LnxUpgradeResponse {
	s.Code = &v
	return s
}

func (s *LnxUpgradeResponse) SetMsg(v string) *LnxUpgradeResponse {
	s.Msg = &v
	return s
}

func (s *LnxUpgradeResponse) SetTraceId(v string) *LnxUpgradeResponse {
	s.TraceId = &v
	return s
}

func (s *LnxUpgradeResponse) SetData(v *LnxUpgradeDataResponse) *LnxUpgradeResponse {
	s.Data = v
	return s
}

type LnxVersionRequest struct {
	LnxKey      *string `json:"lnxKey,omitempty" xml:"lnxKey,omitempty" require:"true"`
	Arch        *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
}

func (s LnxVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s LnxVersionRequest) GoString() string {
	return s.String()
}

func (s *LnxVersionRequest) SetLnxKey(v string) *LnxVersionRequest {
	s.LnxKey = &v
	return s
}

func (s *LnxVersionRequest) SetArch(v string) *LnxVersionRequest {
	s.Arch = &v
	return s
}

func (s *LnxVersionRequest) SetVersionCode(v int) *LnxVersionRequest {
	s.VersionCode = &v
	return s
}

type LnxVersionDataResponse struct {
	LnxKey      *string `json:"lnxKey,omitempty" xml:"lnxKey,omitempty" require:"true"`
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s LnxVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s LnxVersionDataResponse) GoString() string {
	return s.String()
}

func (s *LnxVersionDataResponse) SetLnxKey(v string) *LnxVersionDataResponse {
	s.LnxKey = &v
	return s
}

func (s *LnxVersionDataResponse) SetPackageName(v string) *LnxVersionDataResponse {
	s.PackageName = &v
	return s
}

func (s *LnxVersionDataResponse) SetVersionName(v string) *LnxVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *LnxVersionDataResponse) SetVersionCode(v int) *LnxVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *LnxVersionDataResponse) SetDescription(v string) *LnxVersionDataResponse {
	s.Description = &v
	return s
}

type LnxVersionResponse struct {
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *LnxVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s LnxVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s LnxVersionResponse) GoString() string {
	return s.String()
}

func (s *LnxVersionResponse) SetCode(v int) *LnxVersionResponse {
	s.Code = &v
	return s
}

func (s *LnxVersionResponse) SetMsg(v string) *LnxVersionResponse {
	s.Msg = &v
	return s
}

func (s *LnxVersionResponse) SetTraceId(v string) *LnxVersionResponse {
	s.TraceId = &v
	return s
}

func (s *LnxVersionResponse) SetData(v *LnxVersionDataResponse) *LnxVersionResponse {
	s.Data = v
	return s
}

type WinUpgradeRequest struct {
	WinKey             *string `json:"winKey,omitempty" xml:"winKey,omitempty" require:"true"`
	Arch               *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
	VersionCode        *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	AppointVersionCode *int    `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
	DevModelKey        *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
	DevKey             *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
}

func (s WinUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s WinUpgradeRequest) GoString() string {
	return s.String()
}

func (s *WinUpgradeRequest) SetWinKey(v string) *WinUpgradeRequest {
	s.WinKey = &v
	return s
}

func (s *WinUpgradeRequest) SetArch(v string) *WinUpgradeRequest {
	s.Arch = &v
	return s
}

func (s *WinUpgradeRequest) SetVersionCode(v int) *WinUpgradeRequest {
	s.VersionCode = &v
	return s
}

func (s *WinUpgradeRequest) SetAppointVersionCode(v int) *WinUpgradeRequest {
	s.AppointVersionCode = &v
	return s
}

func (s *WinUpgradeRequest) SetDevModelKey(v string) *WinUpgradeRequest {
	s.DevModelKey = &v
	return s
}

func (s *WinUpgradeRequest) SetDevKey(v string) *WinUpgradeRequest {
	s.DevKey = &v
	return s
}

type WinUpgradeDataResponse struct {
	WinKey               *string `json:"winKey,omitempty" xml:"winKey,omitempty" require:"true"`
	PackageName          *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
	VersionName          *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode          *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	UrlPath              *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
	UrlFileSize          *int    `json:"urlFileSize,omitempty" xml:"urlFileSize,omitempty" require:"true"`
	UrlFileMd5           *string `json:"urlFileMd5,omitempty" xml:"urlFileMd5,omitempty" require:"true"`
	UpgradeType          *int    `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
	PromptUpgradeContent *string `json:"promptUpgradeContent,omitempty" xml:"promptUpgradeContent,omitempty" require:"true"`
}

func (s WinUpgradeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s WinUpgradeDataResponse) GoString() string {
	return s.String()
}

func (s *WinUpgradeDataResponse) SetWinKey(v string) *WinUpgradeDataResponse {
	s.WinKey = &v
	return s
}

func (s *WinUpgradeDataResponse) SetPackageName(v string) *WinUpgradeDataResponse {
	s.PackageName = &v
	return s
}

func (s *WinUpgradeDataResponse) SetVersionName(v string) *WinUpgradeDataResponse {
	s.VersionName = &v
	return s
}

func (s *WinUpgradeDataResponse) SetVersionCode(v int) *WinUpgradeDataResponse {
	s.VersionCode = &v
	return s
}

func (s *WinUpgradeDataResponse) SetUrlPath(v string) *WinUpgradeDataResponse {
	s.UrlPath = &v
	return s
}

func (s *WinUpgradeDataResponse) SetUrlFileSize(v int) *WinUpgradeDataResponse {
	s.UrlFileSize = &v
	return s
}

func (s *WinUpgradeDataResponse) SetUrlFileMd5(v string) *WinUpgradeDataResponse {
	s.UrlFileMd5 = &v
	return s
}

func (s *WinUpgradeDataResponse) SetUpgradeType(v int) *WinUpgradeDataResponse {
	s.UpgradeType = &v
	return s
}

func (s *WinUpgradeDataResponse) SetPromptUpgradeContent(v string) *WinUpgradeDataResponse {
	s.PromptUpgradeContent = &v
	return s
}

type WinUpgradeResponse struct {
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *WinUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s WinUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s WinUpgradeResponse) GoString() string {
	return s.String()
}

func (s *WinUpgradeResponse) SetCode(v int) *WinUpgradeResponse {
	s.Code = &v
	return s
}

func (s *WinUpgradeResponse) SetMsg(v string) *WinUpgradeResponse {
	s.Msg = &v
	return s
}

func (s *WinUpgradeResponse) SetTraceId(v string) *WinUpgradeResponse {
	s.TraceId = &v
	return s
}

func (s *WinUpgradeResponse) SetData(v *WinUpgradeDataResponse) *WinUpgradeResponse {
	s.Data = v
	return s
}

type WinVersionRequest struct {
	WinKey      *string `json:"winKey,omitempty" xml:"winKey,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Arch        *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
}

func (s WinVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s WinVersionRequest) GoString() string {
	return s.String()
}

func (s *WinVersionRequest) SetWinKey(v string) *WinVersionRequest {
	s.WinKey = &v
	return s
}

func (s *WinVersionRequest) SetVersionCode(v int) *WinVersionRequest {
	s.VersionCode = &v
	return s
}

func (s *WinVersionRequest) SetArch(v string) *WinVersionRequest {
	s.Arch = &v
	return s
}

type WinVersionDataResponse struct {
	WinKey      *string `json:"winKey,omitempty" xml:"winKey,omitempty" require:"true"`
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s WinVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s WinVersionDataResponse) GoString() string {
	return s.String()
}

func (s *WinVersionDataResponse) SetWinKey(v string) *WinVersionDataResponse {
	s.WinKey = &v
	return s
}

func (s *WinVersionDataResponse) SetPackageName(v string) *WinVersionDataResponse {
	s.PackageName = &v
	return s
}

func (s *WinVersionDataResponse) SetVersionName(v string) *WinVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *WinVersionDataResponse) SetVersionCode(v int) *WinVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *WinVersionDataResponse) SetDescription(v string) *WinVersionDataResponse {
	s.Description = &v
	return s
}

type WinVersionResponse struct {
	Code    *int                     `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                  `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                  `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *FileVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s WinVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s WinVersionResponse) GoString() string {
	return s.String()
}

func (s *WinVersionResponse) SetCode(v int) *WinVersionResponse {
	s.Code = &v
	return s
}

func (s *WinVersionResponse) SetMsg(v string) *WinVersionResponse {
	s.Msg = &v
	return s
}

func (s *WinVersionResponse) SetTraceId(v string) *WinVersionResponse {
	s.TraceId = &v
	return s
}

func (s *WinVersionResponse) SetData(v *FileVersionDataResponse) *WinVersionResponse {
	s.Data = v
	return s
}

type MacUpgradeRequest struct {
	MacKey             *string `json:"macKey,omitempty" xml:"macKey,omitempty" require:"true"`
	Arch               *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
	VersionCode        *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	AppointVersionCode *int    `json:"appointVersionCode,omitempty" xml:"appointVersionCode,omitempty" require:"true"`
	DevModelKey        *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty" require:"true"`
	DevKey             *string `json:"devKey,omitempty" xml:"devKey,omitempty" require:"true"`
}

func (s MacUpgradeRequest) String() string {
	return tea.Prettify(s)
}

func (s MacUpgradeRequest) GoString() string {
	return s.String()
}

func (s *MacUpgradeRequest) SetMacKey(v string) *MacUpgradeRequest {
	s.MacKey = &v
	return s
}

func (s *MacUpgradeRequest) SetArch(v string) *MacUpgradeRequest {
	s.Arch = &v
	return s
}

func (s *MacUpgradeRequest) SetVersionCode(v int) *MacUpgradeRequest {
	s.VersionCode = &v
	return s
}

func (s *MacUpgradeRequest) SetAppointVersionCode(v int) *MacUpgradeRequest {
	s.AppointVersionCode = &v
	return s
}

func (s *MacUpgradeRequest) SetDevModelKey(v string) *MacUpgradeRequest {
	s.DevModelKey = &v
	return s
}

func (s *MacUpgradeRequest) SetDevKey(v string) *MacUpgradeRequest {
	s.DevKey = &v
	return s
}

type MacUpgradeDataResponse struct {
	MacKey               *string `json:"macKey,omitempty" xml:"macKey,omitempty" require:"true"`
	PackageName          *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
	VersionName          *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode          *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	UrlPath              *string `json:"urlPath,omitempty" xml:"urlPath,omitempty" require:"true"`
	UrlFileSize          *int    `json:"urlFileSize,omitempty" xml:"urlFileSize,omitempty" require:"true"`
	UrlFileMd5           *string `json:"urlFileMd5,omitempty" xml:"urlFileMd5,omitempty" require:"true"`
	UpgradeType          *int    `json:"upgradeType,omitempty" xml:"upgradeType,omitempty" require:"true"`
	PromptUpgradeContent *string `json:"promptUpgradeContent,omitempty" xml:"promptUpgradeContent,omitempty" require:"true"`
}

func (s MacUpgradeDataResponse) String() string {
	return tea.Prettify(s)
}

func (s MacUpgradeDataResponse) GoString() string {
	return s.String()
}

func (s *MacUpgradeDataResponse) SetMacKey(v string) *MacUpgradeDataResponse {
	s.MacKey = &v
	return s
}

func (s *MacUpgradeDataResponse) SetPackageName(v string) *MacUpgradeDataResponse {
	s.PackageName = &v
	return s
}

func (s *MacUpgradeDataResponse) SetVersionName(v string) *MacUpgradeDataResponse {
	s.VersionName = &v
	return s
}

func (s *MacUpgradeDataResponse) SetVersionCode(v int) *MacUpgradeDataResponse {
	s.VersionCode = &v
	return s
}

func (s *MacUpgradeDataResponse) SetUrlPath(v string) *MacUpgradeDataResponse {
	s.UrlPath = &v
	return s
}

func (s *MacUpgradeDataResponse) SetUrlFileSize(v int) *MacUpgradeDataResponse {
	s.UrlFileSize = &v
	return s
}

func (s *MacUpgradeDataResponse) SetUrlFileMd5(v string) *MacUpgradeDataResponse {
	s.UrlFileMd5 = &v
	return s
}

func (s *MacUpgradeDataResponse) SetUpgradeType(v int) *MacUpgradeDataResponse {
	s.UpgradeType = &v
	return s
}

func (s *MacUpgradeDataResponse) SetPromptUpgradeContent(v string) *MacUpgradeDataResponse {
	s.PromptUpgradeContent = &v
	return s
}

type MacUpgradeResponse struct {
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *MacUpgradeDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s MacUpgradeResponse) String() string {
	return tea.Prettify(s)
}

func (s MacUpgradeResponse) GoString() string {
	return s.String()
}

func (s *MacUpgradeResponse) SetCode(v int) *MacUpgradeResponse {
	s.Code = &v
	return s
}

func (s *MacUpgradeResponse) SetMsg(v string) *MacUpgradeResponse {
	s.Msg = &v
	return s
}

func (s *MacUpgradeResponse) SetTraceId(v string) *MacUpgradeResponse {
	s.TraceId = &v
	return s
}

func (s *MacUpgradeResponse) SetData(v *MacUpgradeDataResponse) *MacUpgradeResponse {
	s.Data = v
	return s
}

type MacVersionRequest struct {
	MacKey      *string `json:"macKey,omitempty" xml:"macKey,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Arch        *string `json:"arch,omitempty" xml:"arch,omitempty" require:"true"`
}

func (s MacVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s MacVersionRequest) GoString() string {
	return s.String()
}

func (s *MacVersionRequest) SetMacKey(v string) *MacVersionRequest {
	s.MacKey = &v
	return s
}

func (s *MacVersionRequest) SetVersionCode(v int) *MacVersionRequest {
	s.VersionCode = &v
	return s
}

func (s *MacVersionRequest) SetArch(v string) *MacVersionRequest {
	s.Arch = &v
	return s
}

type MacVersionDataResponse struct {
	MacKey      *string `json:"macKey,omitempty" xml:"macKey,omitempty" require:"true"`
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty" require:"true"`
	VersionName *string `json:"versionName,omitempty" xml:"versionName,omitempty" require:"true"`
	VersionCode *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	Description *string `json:"description,omitempty" xml:"description,omitempty" require:"true"`
}

func (s MacVersionDataResponse) String() string {
	return tea.Prettify(s)
}

func (s MacVersionDataResponse) GoString() string {
	return s.String()
}

func (s *MacVersionDataResponse) SetMacKey(v string) *MacVersionDataResponse {
	s.MacKey = &v
	return s
}

func (s *MacVersionDataResponse) SetPackageName(v string) *MacVersionDataResponse {
	s.PackageName = &v
	return s
}

func (s *MacVersionDataResponse) SetVersionName(v string) *MacVersionDataResponse {
	s.VersionName = &v
	return s
}

func (s *MacVersionDataResponse) SetVersionCode(v int) *MacVersionDataResponse {
	s.VersionCode = &v
	return s
}

func (s *MacVersionDataResponse) SetDescription(v string) *MacVersionDataResponse {
	s.Description = &v
	return s
}

type MacVersionResponse struct {
	Code    *int                    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string                 `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	TraceId *string                 `json:"traceId,omitempty" xml:"traceId,omitempty" require:"true"`
	Data    *MacVersionDataResponse `json:"data,omitempty" xml:"data,omitempty" require:"true"`
}

func (s MacVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s MacVersionResponse) GoString() string {
	return s.String()
}

func (s *MacVersionResponse) SetCode(v int) *MacVersionResponse {
	s.Code = &v
	return s
}

func (s *MacVersionResponse) SetMsg(v string) *MacVersionResponse {
	s.Msg = &v
	return s
}

func (s *MacVersionResponse) SetTraceId(v string) *MacVersionResponse {
	s.TraceId = &v
	return s
}

func (s *MacVersionResponse) SetData(v *MacVersionDataResponse) *MacVersionResponse {
	s.Data = v
	return s
}

type AppReportRequest struct {
	EventType *string                    `json:"eventType,omitempty" xml:"eventType,omitempty" require:"true"`
	AppKey    *string                    `json:"appKey,omitempty" xml:"appKey,omitempty" require:"true"`
	Timestamp *string                    `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
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
	LaunchTime          *string `json:"launchTime,omitempty" xml:"launchTime,omitempty"`
	VersionCode         *int    `json:"versionCode,omitempty" xml:"versionCode,omitempty" require:"true"`
	DevModelKey         *string `json:"devModelKey,omitempty" xml:"devModelKey,omitempty"`
	DevKey              *string `json:"devKey,omitempty" xml:"devKey,omitempty"`
	Target              *string `json:"target,omitempty" xml:"target,omitempty"`
	Arch                *string `json:"arch,omitempty" xml:"arch,omitempty"`
	DownloadVersionCode *int    `json:"downloadVersionCode,omitempty" xml:"downloadVersionCode,omitempty"`
	UpgradeVersionCode  *int    `json:"upgradeVersionCode,omitempty" xml:"upgradeVersionCode,omitempty"`
	Code                *int    `json:"code,omitempty" xml:"code,omitempty"`
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
	Code    *int    `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	Msg     *string `json:"msg,omitempty" xml:"msg,omitempty" require:"true"`
	Docs    *string `json:"docs,omitempty" xml:"docs,omitempty" require:"true"`
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
	AccessKey    *string
	AccessSecret *string
	Protocol     *string
	Endpoint     *string
}

func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
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

		_resp, _err = func() (*UrlUpgradeResponse, error) {
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
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
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

func (client *Client) UrlVersion(request *UrlVersionRequest) (_result *UrlVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &UrlVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*UrlVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/url/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/url/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &UrlVersionResponse{}
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

		_resp, _err = func() (*FileUpgradeResponse, error) {
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
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
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

func (client *Client) FileVersion(request *FileVersionRequest) (_result *FileVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &FileVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*FileVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/file/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/file/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &FileVersionResponse{}
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

		_resp, _err = func() (*ApkUpgradeResponse, error) {
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
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
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

func (client *Client) ApkVersion(request *ApkVersionRequest) (_result *ApkVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &ApkVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*ApkVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/apk/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/apk/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &ApkVersionResponse{}
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

		_resp, _err = func() (*ConfigurationUpgradeResponse, error) {
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
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
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

func (client *Client) ConfigurationVersion(request *ConfigurationVersionRequest) (_result *ConfigurationVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &ConfigurationVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*ConfigurationVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/configuration/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/configuration/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &ConfigurationVersionResponse{}
			_err = tea.Convert(tea.ToMap(result), &_result)
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) TauriVersion(request *TauriVersionRequest) (_result *TauriVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &TauriVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*TauriVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/tauri/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/tauri/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &TauriVersionResponse{}
			_err = tea.Convert(tea.ToMap(result), &_result)
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) ElectronVersion(request *ElectronVersionRequest) (_result *ElectronVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &ElectronVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*ElectronVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/electron/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/electron/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &ElectronVersionResponse{}
			_err = tea.Convert(tea.ToMap(result), &_result)
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) LnxUpgrade(request *LnxUpgradeRequest) (_result *LnxUpgradeResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &LnxUpgradeResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*LnxUpgradeResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/lnx/upgrade")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/lnx/upgrade")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &LnxUpgradeResponse{}
			_err = tea.Convert(tea.ToMap(result), &_result)
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) LnxVersion(request *LnxVersionRequest) (_result *LnxVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &LnxVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*LnxVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/lnx/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/lnx/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &LnxVersionResponse{}
			_err = tea.Convert(tea.ToMap(result), &_result)
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) WinUpgrade(request *WinUpgradeRequest) (_result *WinUpgradeResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &WinUpgradeResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*WinUpgradeResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/win/upgrade")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/win/upgrade")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &WinUpgradeResponse{}
			_err = tea.Convert(tea.ToMap(result), &_result)
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) WinVersion(request *WinVersionRequest) (_result *WinVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &WinVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*WinVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/win/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/win/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &WinVersionResponse{}
			_err = tea.Convert(tea.ToMap(result), &_result)
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) MacUpgrade(request *MacUpgradeRequest) (_result *MacUpgradeResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &MacUpgradeResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*MacUpgradeResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/mac/upgrade")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/mac/upgrade")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &MacUpgradeResponse{}
			_err = tea.Convert(tea.ToMap(result), &_result)
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

func (client *Client) MacVersion(request *MacVersionRequest) (_result *MacVersionResponse, _err error) {
	_err = tea.Validate(request)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeout": 10000,
		// 10s 的过期时间
	}

	_resp := &MacVersionResponse{}
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (*MacVersionResponse, error) {
			request_ := tea.NewRequest()
			// 序列化请求体
			bodyStr := util.ToJSONString(request)
			// 生成请求参数
			timestamp := darabonbabase.TimeRFC3339()
			nonce := darabonbabase.GenerateNonce()
			uri := tea.String("/v1/mac/version")
			accessKey := client.AccessKey
			accessSecret := client.AccessSecret
			// 生成签名
			signature := darabonbabase.GenerateSignature(bodyStr, nonce, accessSecret, timestamp, uri)
			request_.Protocol = client.Protocol
			request_.Method = tea.String("POST")
			request_.Pathname = tea.String("/v1/mac/version")
			request_.Headers = map[string]*string{
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
				})
				return _result, _err
			}

			_result = &MacVersionResponse{}
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

		_resp, _err = func() (*AppReportResponse, error) {
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
				"host":         client.Endpoint,
				"content-type": tea.String("application/json"),
				"x-Timestamp":  timestamp,
				"x-Nonce":      nonce,
				"x-AccessKey":  accessKey,
				"x-Signature":  signature,
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
					"code":       tea.ToString(result["code"]),
					"message":    tea.ToString(result["msg"]),
					"docs":       tea.ToString(result["docs"]),
					"traceId":    tea.ToString(result["traceId"]),
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

func TimeRFC3339() (_result *string) {
	_body := darabonbabase.TimeRFC3339()
	_result = _body
	return _result
}
