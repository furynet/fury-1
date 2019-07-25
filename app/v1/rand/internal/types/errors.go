//nolint
package types

import (
	sdk "github.com/irisnet/irishub/types"
)

// Rand errors reserve 100 ~ 199.
const (
	DefaultCodespace sdk.CodespaceType = "rand"

	CodeInvalidConsumer sdk.CodeType = 100
	CodeInvalidReqID    sdk.CodeType = 101
)

//----------------------------------------
// Rand error constructors

func ErrInvalidConsumer(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidConsumer, msg)
}

func ErrInvalidReqID(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidReqID, msg)
}
