package code

const (
	OK uint32 = iota
	ERROR
	UNKNOWN_TRANSACTION
	INVALID_TRANSACTION
	INVALID_SIGNATURE
	UNKNOWN_QUERY
	NOT_FOUND
)
