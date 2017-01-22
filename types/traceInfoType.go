package types

import "time"

type TraceInfo struct {
	TraceMsg  string    `json:"traceMsg"`
	TimeStamp time.Time `json:"traceTimeStamp"`
}

type TraceInfos []TraceInfo
