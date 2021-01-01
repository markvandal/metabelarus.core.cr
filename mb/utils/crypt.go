package mbutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"

	"github.com/decred/dcrd/dcrec/secp256k1"
)

// EncryptPayload - Asymetric event/message payload encryption
func EncryptPayload(pubKeyBytes []byte, payload []byte) ([]byte, error) {
	cihperPubKey, err := secp256k1.ParsePubKey(pubKeyBytes)
	ephemeralPrivKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return nil, err
	}
	ephemeralPubKey := ephemeralPrivKey.PubKey().SerializeCompressed()
	cipherKey := sha256.Sum256(
		secp256k1.GenerateSharedSecret(ephemeralPrivKey, cihperPubKey),
	)

	aead, err := newAEAD(cipherKey[:])
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	cipherPayload := make([]byte, 4+len(ephemeralPubKey))
	binary.LittleEndian.PutUint32(cipherPayload, uint32(len(ephemeralPubKey)))
	copy(cipherPayload[4:], ephemeralPubKey)

	return aead.Seal(cipherPayload, nonce, payload, ephemeralPubKey), nil
}

// DecryptPayload - Asymetric event/message payload deciption
func DecryptPayload(pkBytes []byte, ciphertext []byte) ([]byte, error) {
	privKey, _ := secp256k1.PrivKeyFromBytes(pkBytes)

	// Read the sender's ephemeral public key from the start of the message.
	// Error handling for inappropriate pubkey lengths is elided here for
	// brevity.
	pubKeyLen := binary.LittleEndian.Uint32(ciphertext[:4])
	senderPubKeyBytes := ciphertext[4 : 4+pubKeyLen]
	senderPubKey, err := secp256k1.ParsePubKey(senderPubKeyBytes)
	if err != nil {
		return nil, err
	}

	// Derive the key used to seal the message, this time from the
	// recipient's private key and the sender's public key.
	recoveredCipherKey := sha256.Sum256(secp256k1.GenerateSharedSecret(privKey, senderPubKey))

	// Open the sealed message.
	aead, err := newAEAD(recoveredCipherKey[:])
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aead.NonceSize())
	return aead.Open(nil, nonce, ciphertext[4+pubKeyLen:], senderPubKeyBytes)
}

func newAEAD(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}
