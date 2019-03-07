package types

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	s := time.Now()
	d := Time{s}

	require.Equal(t, d.Timestamp(), ParseTimestamp(d.Timestamp()).Timestamp())
}
