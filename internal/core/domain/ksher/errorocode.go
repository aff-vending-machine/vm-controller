package ksher

type KsherErrorCode = string

const (
	SUCCESS    KsherErrorCode = "SUCCESS"
	FAIL       KsherErrorCode = "FAIL"
	DUPLICATED KsherErrorCode = "DUPLICATED"
	PENDING    KsherErrorCode = "PENDING"
	REFUNDED   KsherErrorCode = "REFUNDED"
	SIGNERROR  KsherErrorCode = "SIGNERROR"
)
