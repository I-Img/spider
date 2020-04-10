package common

import "time"

//Source 各个图片源的控制字段值
type Source struct {
	ID       int
	Name     string
	LastTime time.Time
	PID      int
}
