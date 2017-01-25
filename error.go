package tgc

const (
	// Client Errors
	ResourceNotFound = 404

	// Parameter Errors
	ObjectNotFound           = 440
	MissingRequiredParameter = 441
	OutOfRange               = 442
	NotAvailable             = 443
	FileTypeNotSupported     = 444
	FileMismatch             = 445
	PaymentDeclined          = 446

	// Account Errors
	InsufficientPrivileges  = 450
	SessionExpired          = 451
	RPCRequestLimitExceeded = 452
	PrequisiteFailed        = 453
	PasswordIncorrect       = 454
	MustVerifyHumanity      = 455
	OfflineProcessing       = 499

	// Server Errors
	UndefinedError  = 500
	CoundNotConnect = 504
)

type Error struct {
	Error struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		Code    int         `json:"code"`
	} `json:"error"`
}
