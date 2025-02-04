package constants

type KeyType string

const (
	CORRELATION_ID_KEY       KeyType = "correlation-id"
	REQUEST_ID_KEY           KeyType = "request-id"
	SERVICE_KEY              KeyType = "service"
	REQUEST_CONTEXT_KEY      KeyType = "x-request-info"
	AUTHORIZATION_KEY        KeyType = "Authorization"
	CONTENT_TYPE_HEADER_KEY  KeyType = "Content-Type"
	LANGUAGE_CODE_HEADER_KEY KeyType = "x-language-code"
	REFRESH_TOKEN_HEADER_KEY KeyType = "x-refresh-token"
)
