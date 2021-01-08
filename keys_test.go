package opaque

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestToOpaque(t *testing.T) {
	legacyToOpaque := func(in string) string {
		out, err := ToOpaque(in)
		require.NoError(t, err, "impossible, legacy was never returning an error anyway")

		return out
	}

	tests := []struct {
		key    string
		opaque string
		encode func(in string) string
		decode func(in string) (string, error)
	}{
		{"super:secret:internalkey", "kEBU5KAkz32my1wU1QT_WrX0ZMN6DBk6BkmwbhsY28-jsiyUgcbzSQ==", EncodeString, DecodeToString},
		{"blah:super:secret", "xE77fJ_z6Z7UXwyx-e0kWqTtdc4yRR8vAEnvaUQSx96y", EncodeString, DecodeToString},
		{"", "9hd1ai3EX1dUhNPvo0dj0Q==", EncodeString, DecodeToString},

		// Legacy interface
		{"super:secret:internalkey", "kEBU5KAkz32my1wU1QT_WrX0ZMN6DBk6BkmwbhsY28-jsiyUgcbzSQ==", legacyToOpaque, FromOpaque},
		{"blah:super:secret", "xE77fJ_z6Z7UXwyx-e0kWqTtdc4yRR8vAEnvaUQSx96y", legacyToOpaque, FromOpaque},
		{"", "9hd1ai3EX1dUhNPvo0dj0Q==", legacyToOpaque, FromOpaque},
	}

	for _, test := range tests {
		outres := test.encode(test.key)
		assert.Equal(t, test.opaque, outres)

		inres, err := test.decode(test.opaque)
		assert.NoError(t, err)
		assert.Equal(t, test.key, inres)
	}
}
