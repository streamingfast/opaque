package opaque

import (
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/nacl/secretbox"
)

// Yes, this is hardcoded, no big deal.
// We do not really care about Confidentiality - this is merely for obfuscation
// Was generated with `openssl rand -hex 32`
var opaqueCursorEncryptionKey = [32]byte{0x7b, 0xfc, 0xac, 0xee, 0x25, 0x74, 0x09, 0x09, 0x9e, 0xdd, 0x2b, 0xb6, 0xa4, 0x42, 0x63, 0x8b, 0x55, 0x9a, 0x80, 0xbf, 0xbf, 0xc0, 0xb9, 0xac, 0xde, 0xa0, 0xd8, 0x34, 0x4b, 0x10, 0xeb, 0x00}

// **WARNING** - DO NOT copy this unless you know what you are doing
// We are using a fixed nonce by design here.
// Using a fixed nonce is FATAL cryptography security flaw in normal cases
// But in this case we mostly care of obscuring / making opaque the key
var fixedNonce = [24]byte{0x26, 0x15, 0x54, 0xc4, 0x5a, 0xb9, 0xb7, 0x52, 0xab, 0xad, 0x4f, 0x19, 0xc2, 0x42, 0x60, 0x57, 0x2, 0xd5, 0x5a, 0x0d, 0x91, 0x61, 0x6a, 0x1b}

// ToOpaque obfuscates (encrypts) internal keys to be used as
// pagination cursors sent to frontend
func ToOpaque(internalKey string) (string, error) {
	var nonce [24]byte
	copy(nonce[:], fixedNonce[:])
	ciphertext := secretbox.Seal(nonce[:], []byte(internalKey), &nonce, &opaqueCursorEncryptionKey)
	return base64.URLEncoding.EncodeToString(ciphertext[24:]), nil
}

// FromOpaque de-obfuscates (decrypts) internal keys to be used as
// pagination cursors sent to frontend
func FromOpaque(opaqueKey string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(opaqueKey)
	if err != nil {
		return "", err
	}

	plaintext, ok := secretbox.Open(nil, ciphertext, &fixedNonce, &opaqueCursorEncryptionKey)
	if !ok {
		return "", errors.New("decryption failed")
	}
	return string(plaintext), nil
}
