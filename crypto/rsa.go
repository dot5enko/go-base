package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func GetPrivateKey(path string) (*rsa.PrivateKey, error) {
	key, _ := ioutil.ReadFile(path)

	keyObj, err := ssh.ParseRawPrivateKey(key)
	if err != nil {
		return nil, err
	}

	result, ok := keyObj.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("got unknown type of key %#", keyObj)
	} else {
		return result, nil
	}
}

func GetPublicKey(path string) (*rsa.PublicKey, error) {
	pubBytes, errRead := ioutil.ReadFile(path)

	if errRead != nil {
		return nil, errRead
	}

	block, _ := pem.Decode([]byte(pubBytes))
	if block == nil {
		return nil, UnableToDecodePublicKey
	}

	pubKey, pubKeyError := x509.ParsePKIXPublicKey(block.Bytes)
	if pubKeyError != nil {
		return nil, pubKeyError
	} else {
		keyFormat, ok := pubKey.(*rsa.PublicKey)
		if !ok {
			return nil, DecodedKeyForamIssue
		} else {
			return keyFormat, nil
		}
	}
}

var (
	UnableToDecodePublicKey error = errors.New("failed to parse PEM block containing the public key")
	DecodedKeyForamIssue    error = errors.New("decoded key is not *rsa.PublicKey")
)

var (
	PrivateKeyPath string = ""
	PubKeyPath     string = ""
)

func SignString(algo crypto.Hash, input string) (result, hash string, err error) {

	key, err0 := GetPrivateKey(PrivateKeyPath)

	if err0 != nil {
		err = err0
		return
	} else {
		message := []byte(input)

		hf := algo.New()
		hf.Write(message)

		hashed := hf.Sum(nil)

		hashToReturn := hashed

		signature, err1 := rsa.SignPKCS1v15(rand.Reader, key, algo, hashed[:])
		if err != nil {
			err = err1
			return
		}

		result = base64.StdEncoding.EncodeToString(signature)
		hash = hex.EncodeToString(hashToReturn[:])
		return
	}
}

func VerifySignature(algo crypto.Hash, hash, signb64 string) error {
	pkey, _ := GetPublicKey(PubKeyPath)

	signBytes, err := base64.StdEncoding.DecodeString(signb64)
	if err != nil {
		return err
	}

	decodedHashBytes, err := hex.DecodeString(hash)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(pkey, algo, decodedHashBytes, signBytes)
}
