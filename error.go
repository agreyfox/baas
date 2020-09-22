package baas

// General errors.
const (
	ErrUnauthorized = Error("unauthorized")
	ErrInternal     = Error("internal error")
	ErrInvalidDate  = Error("invalid date")
	ErrNotFound     = Error("resource not found")
	ErrBadRequest   = Error("bad request")
	ErrInvalidJSON  = Error("invalid json")
)

const (
	ErrAuthorizationHeaderMissing = Error("authorization header is missing")
)

// Usage errors
const (
	ErrUsageNotFound = Error("usage not found")
)

// Application errors
const (
	ErrApplicationNotFound  = Error("application_id not found")
	ErrApplicationNameTaken = Error("application name taken")
)

// Storage errors
const (
	ErrBucketNotFound    = Error("bucket does not exist")
	ErrInvalidBucketName = Error("invalid bucket name")
	ErrInvalidStorageKey = Error("invalid api keys or api keys do not have correct permissions")
)

// File errors
const (
	ErrFileNotFound            = Error("file not found")
	ErrFileToLargeToTransform  = Error("file is to large to transform")
	ErrTransformationNotUnique = Error("transformation is not unique")
)

// Blockchain baas error
const (
	ErrBaasNoSuchUser            = Error("userId does not exist")
	ErrBaasUserNameTaken         = Error("userId is already exist")
	ErrBaasInvalidPassword       = Error("Invalid password")
	ErrBaasParameterNotFound     = Error("Parameter is no valid")
	ErrBaasApplicationIDRequired = Error("Application ID is required")
	ErrBaasCipherTextRequired    = Error("invalid cipherText")
	ErrBaasNotEnoughMoney        = Error("Account doesn't have enough gas")
)

// Error represents a Mahi error.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}
