package gdlib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"io"
)

var (
	key = []byte("123344555555123344555555")
)

func serialize(data any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)

	if err != nil {
		return nil, fmt.Errorf("error encoding data: %v", err)
	}

	return buf.Bytes(), nil
}

func deserialize(data []byte, dest any) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	err := dec.Decode(dest)
	if err != nil {
		return fmt.Errorf("error decoding data: %v", err)
	}

	return nil
}

func Encrypt(data any) ([]byte, error) {
	serialized, libErr := serialize(data)
	if libErr != nil {
		return nil, libErr
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(serialized))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("error reading iv: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], serialized)

	return cipherText, nil
}

func Decrypt(cipherText []byte, dest any) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error creating new decrypt cipher: %v", err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	libErr := deserialize(cipherText, dest)
	if libErr != nil {
		return libErr
	}

	return nil
}
