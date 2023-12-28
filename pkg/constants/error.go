package constants

import "errors"

var (
	BearerTokenHasError = errors.New("Bearer token catch error")
	BearerTokenInvalid  = errors.New("Invalid token")

	DuplicateStoreServer = errors.New("Duplicate data server")

	ErrorNoFoundDataApkSetting = errors.New("No data found for APK setting")
)
