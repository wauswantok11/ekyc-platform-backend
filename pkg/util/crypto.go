package util

import (
	"crypto/sha256"
	"fmt"
)

// func PHPAes256CbcDecrypt(cipherText string, key string) (*string, error) {
// 	// Decode key
// 	b64Key, err := base64.StdEncoding.DecodeString(key)
// 	if err != nil {
// 		return nil, errors.New("decode key error")
// 	}

// 	// Decode crypt block
// 	cryptBlock := struct {
// 		Iv    string `json:"iv"`
// 		Value string `json:"value"`
// 		Mac   string `json:"mac"`
// 	}{}
// 	b64CryptBlock, err := base64.StdEncoding.DecodeString(cipherText)
// 	if err != nil {
// 		return nil, errors.New("decode crypt block error")
// 	}
// 	if err := sonic.Unmarshal(b64CryptBlock, &cryptBlock); err != nil {
// 		return nil, errors.New("invalid crypt block")
// 	}

// 	// Decode IV
// 	iv, err := base64.StdEncoding.DecodeString(cryptBlock.Iv)
// 	if err != nil {
// 		return nil, errors.New("decode iv error")
// 	}

// 	// Decode Cipher
// 	cipherByte, err := base64.StdEncoding.DecodeString(cryptBlock.Value)
// 	if err != nil {
// 		return nil, errors.New("decode cipher text error")
// 	}
// 	if len(cipherByte) < aes.BlockSize {
// 		return nil, errors.New("cipher text too short")
// 	}

// 	// Start decrypt
// 	block, err := aes.NewCipher(b64Key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	mode := cipher.NewCBCDecrypter(block, iv)
// 	mode.CryptBlocks(cipherByte, cipherByte)

// 	cipherByte, _ = pkcs7.Unpad(cipherByte, aes.BlockSize)
// 	plainText := fmt.Sprintf("%s", cipherByte)
// 	return &plainText, nil
// }

// func PHPAes256CbcEncrypt(text string, key string) (string, error) {

// 	b64Key, err := base64.StdEncoding.DecodeString(key)
// 	if err != nil {
// 		return "", errors.New("decode key error")
// 	}
// 	block, err := aes.NewCipher(b64Key)
// 	if err != nil {
// 		return "", err
// 	}

// 	textBytes := []byte(text)
// 	textBytes = padPKCS7(textBytes, aes.BlockSize)

// 	iv := make([]byte, aes.BlockSize)
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		return "", err
// 	}

// 	ivBase64 := base64.StdEncoding.EncodeToString(iv)

// 	mode := cipher.NewCBCEncrypter(block, iv)
// 	ciphertext := make([]byte, len(textBytes))
// 	mode.CryptBlocks(ciphertext, textBytes)

// 	encodedData := base64.StdEncoding.EncodeToString(ciphertext)
// 	ivCipher := ivBase64 + encodedData

// 	hmac := generateHMAC(b64Key, []byte(ivCipher))
// 	hmacHex := hex.EncodeToString(hmac)
// 	jsonData, err := json.Marshal(map[string]string{
// 		"iv":    ivBase64,
// 		"value": encodedData,
// 		"mac":   hmacHex,
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	return base64.StdEncoding.EncodeToString(jsonData), nil
// }

// func padPKCS7(data []byte, blockSize int) []byte {
// 	padding := blockSize - len(data)%blockSize
// 	padText := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(data, padText...)
// }

// func generateHMAC(key, data []byte) []byte {
// 	mac := hmac.New(sha256.New, key)
// 	mac.Write(data)
// 	return mac.Sum(nil)
// }

func HashSHA256(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}

// func PHPHashBcrypt(data string) string {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
// 	if err != nil {
// 		return data
// 	}
// 	hashedData := string(hash)
// 	// Change prefix from $2a$ to $2y$ if required
// 	if len(hashedData) > 4 && hashedData[:4] == "$2a$" {
// 		hashedData = "$2y$" + hashedData[4:]
// 	}
// 	return hashedData

// }
