package cryptoUtils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"github.com/samber/mo"
	"golang.org/x/crypto/pbkdf2"
	"os"
)

// Functional API --------

type FpCyUtil struct{}
type CyFn func([]byte, []byte) ([]byte, error)

func eitherCyOp(key, data []byte, fn CyFn) mo.Either[error, []byte] {
	_fn := func() ([]byte, error) {
		return fn(key, data)
	}
	return mo.Try(_fn).ToEither()
}

// todo: replace with sha256
func (r FpCyUtil) GetMD5Hash(filepath string) mo.Result[[16]byte] {
	data, _ := os.ReadFile(filepath)
	hash := md5.Sum(data)
	return mo.Ok(hash)
}

func (r FpCyUtil) EncryptAES(key, data []byte) mo.Either[error, []byte] {
	return eitherCyOp(key, data, encryptAES)
}

func (r FpCyUtil) DecryptAES(key, data []byte) mo.Either[error, []byte] {
	return eitherCyOp(key, data, decryptAES)
}

// Private ---------

func encryptAES(key, data []byte) ([]byte, error) {
	key, salt := DeriveKey(key, nil)

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func decryptAES(key, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]

	key, salt = DeriveKey(key, salt)

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func DeriveKey(passphrase []byte, salt []byte) ([]byte, []byte) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil
		}
	}
	return pbkdf2.Key(passphrase, salt, 1000, 32, sha256.New), salt
}
