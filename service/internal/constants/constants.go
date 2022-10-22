package constants

const RESPONSE_LIMIT_DB = 3

const (
	STATUS_RESERVED = iota
	STATUS_CONFIRM
	STATUS_CANCEL
)

type key string

const (
	BASE_PATH key = "base_path"
)
