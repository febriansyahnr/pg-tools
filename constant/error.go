package constant

import "errors"

var (
	ErrTrxNotFound                      = errors.New("virtual account not found")
	ErrTrxExpired                       = errors.New("virtual account has expired")
	ErrUpdateRefID                      = errors.New("failed to update ref id")
	ErrPaidAmountGreaterThanTotalAmount = errors.New("paid amount greater than total amount")
	ErrAlreadyPaid                      = errors.New("already paid")
	ErrVANumberInUse                    = errors.New("va number still in use")
	ErrInvalidAmount                    = errors.New("invalid amount")
	ErrInvalidVAAndRefnumber            = errors.New("invalid va and ref number")
)
