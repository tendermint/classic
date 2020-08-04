package errors

type EncodingError struct{}
type BadNonce struct{}
type Unauthorized struct{}
type UnknownError struct{}

// Implements abci.Error
func (_ EncodingError) Error() string { return "EncodingError" }
func (_ BadNonce) Error() string      { return "BadNonce" }
func (_ Unauthorized) Error() string  { return "Unauthorized" }
func (_ UnknownError) Error() string  { return "UnknownError" }
