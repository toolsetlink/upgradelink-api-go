package client

import (
	"time"
)

// TimeRFC3339 生成 RFC3339 格式的时间戳
func TimeRFC3339() *string {
	timestamp := time.Now().Format(time.RFC3339)
	return &timestamp
}
