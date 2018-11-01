package blockchain

import "github.com/pkg/errors"

var (
	ErrTxFailed     = errors.New("tx failed")
	ErrTxTimeout    = errors.New("tx timeout")
	ErrTxNoContract = errors.New("tx is not contract creation")

	ErrZeroAddress = errors.New("zero address")
)
