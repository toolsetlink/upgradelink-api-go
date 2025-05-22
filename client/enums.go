package client

const (
	EventTypeAppStart           string = "app_start"
	EventTypeAppUpgradeDownload string = "app_upgrade_download"
	EventTypeAppUpgradeUpgrade  string = "app_upgrade_upgrade"
)

const (
	EventTypeCodeSuccess                    = 0
	EventTypeCodeError                      = 1
	EventTypeCodeDownloadHttpError          = 1001
	EventTypeCodeDownloadNoEnoughSpaceError = 1002
	EventTypeCodeDownloadFileOperateError   = 1003
	EventTypeCodeDownloadMd5Error           = 1004
)
