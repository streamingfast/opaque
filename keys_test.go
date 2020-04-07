package opaque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToOpaque(t *testing.T) {
	tests := []struct {
		key    string
		opaque string
	}{
		{"super:secret:internalkey", "kEBU5KAkz32my1wU1QT_WrX0ZMN6DBk6BkmwbhsY28-jsiyUgcbzSQ=="},
		{"blah:super:secret", "xE77fJ_z6Z7UXwyx-e0kWqTtdc4yRR8vAEnvaUQSx96y"},
		{"", "9hd1ai3EX1dUhNPvo0dj0Q=="},
	}

	for _, test := range tests {
		outres, err := ToOpaque(test.key)
		assert.NoError(t, err)
		assert.Equal(t, test.opaque, outres)

		inres, err := FromOpaque(test.opaque)
		assert.NoError(t, err)
		assert.Equal(t, test.key, inres)
	}
}
