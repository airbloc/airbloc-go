package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestParseTimestamp(t *testing.T) {
	s := time.Now()
	d := Time{s}

	require.Equal(t, d.Timestamp(), ParseTimestamp(d.Timestamp()).Timestamp())
}
