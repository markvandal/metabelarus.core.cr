package mbutils

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"os"
	"os/exec"

	ecies "github.com/ecies/go"

	"github.com/spf13/cobra"
)

const (
	MBFlagCrypt = "node-crypt"
)

func AddMbCryptFlags(cmd *cobra.Command) {
	cmd.Flags().String(MBFlagCrypt, "", "Specify absolute path to node crypt script")
}

// EncryptPayload - Asymetric event/message payload encryption
func EncryptPayload(script string, pubKeyBytes []byte, payload []byte) ([]byte, error) {
	if script == "" {
		pubKey, err := ecies.NewPublicKeyFromBytes(pubKeyBytes)
		if err != nil {
			return nil, err
		}

		val, err := ecies.Encrypt(pubKey, payload)
		if err != nil {
			return nil, err
		}

		return []byte(base64.StdEncoding.EncodeToString(val)), nil
	}

	return crypt(script, "encrypt", payload, pubKeyBytes)
}

// DecryptPayload - Asymetric event/message payload deciption
func DecryptPayload(script string, pkBytes []byte, ciphertext []byte) ([]byte, error) {
	if script == "" {
		return ecies.Decrypt(ecies.NewPrivateKeyFromBytes(pkBytes), ciphertext)
	}

	return crypt(script, "decrypt", ciphertext, pkBytes)
}

func crypt(script string, action string, payload []byte, key []byte) ([]byte, error) {
	if script == "" {
		return nil, errors.New("Crypt script location should be specified, probably with --node-crypt flag")
	}

	args := []string{
		script,
		action,
		base64.StdEncoding.EncodeToString(payload),
		hex.EncodeToString(key),
	}
	process := exec.Command("node", args...)
	stdin, err := process.StdinPipe()
	if err != nil {
		return nil, err
	}
	defer stdin.Close()
	buf := new(bytes.Buffer) // THIS STORES THE NODEJS OUTPUT
	process.Stdout = buf
	process.Stderr = os.Stderr

	if err = process.Start(); err != nil {
		return nil, err
	}

	process.Wait()

	return buf.Bytes(), nil
}
