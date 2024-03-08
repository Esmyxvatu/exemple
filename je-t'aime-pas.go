package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func encryptFile(key []byte, inFile string, outFile string) error {
	plaintext, err := ioutil.ReadFile(inFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	out, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.Write(ciphertext)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	key := []byte("your-32-byte-key-for-AES-encryption")

	inputDir := "input_directory"
	outputDir := "output_directory"

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				return err
			}
			outFilePath := filepath.Join(outputDir, relPath)
			if err := os.MkdirAll(filepath.Dir(outFilePath), 0755); err != nil {
				return err
			}
			if err := encryptFile(key, path, outFilePath); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
