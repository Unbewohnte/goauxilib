/*
The MIT License (MIT)

Copyright © 2022 Kasyanov Nikolay Alexeyevich (Unbewohnte)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package goauxilib

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Encodes given bytes via Base64 encoding (with padding)
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Decodes given bytes via Base64 encoding (with padding)
func Base64Decode(encodedBytes []byte) string {
	decodedBytes := make([]byte, len(encodedBytes))
	base64.StdEncoding.Decode(decodedBytes, encodedBytes)
	return string(decodedBytes)
}

// Encodes given string via Base64 encoding (with padding)
func Base64EncodeStr(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Decodes given string via Base64 encoding (with padding)
func Base64DecodeStr(encodedStr string) string {
	decodedBytes, _ := base64.StdEncoding.DecodeString(encodedStr)
	return string(decodedBytes)
}

// Encrypts given data using AES encryption with key of 16, 24 or 32 bytes
func Encrypt(key, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	encryptedData := aesGCM.Seal(nonce, nonce, data, nil)

	return encryptedData, nil
}

// Decrypts data encrypted with AES encryption with given key of 16, 24 or 32 bytes
func Decrypt(key, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce, encryptedBytes := data[:aesGCM.NonceSize()], data[aesGCM.NonceSize():]
	decryptedData, err := aesGCM.Open(nil, nonce, encryptedBytes, nil)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}
