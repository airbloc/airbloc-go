package types

import (
	"github.com/airbloc/airbloc-go/shared/types"
)

// ID is bridge type between airbloc-go and contract bind
type ID struct {
	types.ID
}

// RowId is bridge type between airbloc-go and contract bind
type RowId struct {
	types.RowId
}

// DataId is bridge type between airbloc-go and contract bind
type DataId struct {
	types.DataId
}
