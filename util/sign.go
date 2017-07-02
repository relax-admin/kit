package util

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func GetMD5Hash(text string) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(text)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", hasher.Sum(nil)), nil //this is my choice
}

func MakeMd5Sign(url, key string) (sign string) {
	str := url + "&key=" + key
	sign, _ = GetMD5Hash(str)
	return
}

func GetSha1Hash(text string, priKey string) (string, error) {

	preKey := "-----BEGIN RSA PRIVATE KEY-----\n"
	sufKey := "\n-----END RSA PRIVATE KEY-----"

	signer, err := ParsePrivateKey([]byte(preKey + priKey + sufKey))
	if err != nil {
		return "", fmt.Errorf("signer is damaged: %v", err)
	}
	return Sign(signer, []byte(text))
}

func CheckPubKey(text string, signed string, pubKey string) bool {
	var err error
	prePubKey := "-----BEGIN PUBLIC KEY-----\n"
	sufPubKey := "\n-----END PUBLIC KEY-----"

	parser, err := ParsePublicKey([]byte(prePubKey + pubKey + sufPubKey))
	if err != nil {
		return false
		//fmt.Errorf("could not sign request: %v", err)
	}
	err = Verify(parser, []byte(text), signed)
	if err != nil {
		return false
	}
	return true
}

func ParsePrivateKey(data []byte) (*rsa.PrivateKey, error) {
	pemData, err := pemParse(data, "RSA PRIVATE KEY")
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(pemData)
}

func ParsePublicKey(data []byte) (*rsa.PublicKey, error) {
	pemData, err := pemParse(data, "PUBLIC KEY")
	if err != nil {
		return nil, err
	}

	keyInterface, err := x509.ParsePKIXPublicKey(pemData)
	if err != nil {
		return nil, err
	}

	pubKey, ok := keyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("Could not cast parsed key to *rsa.PublickKey")
	}

	return pubKey, nil
}

func pemParse(data []byte, pemType string) ([]byte, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("No PEM block found")
	}
	if pemType != "" && block.Type != pemType {
		return nil, fmt.Errorf("Key's type is '%s', expected '%s'", block.Type, pemType)
	}
	return block.Bytes, nil
}

func LoadPublicKey(publicKeyPath string) (*rsa.PublicKey, error) {
	certPEMBlock, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}

	return ParsePublicKey(certPEMBlock)
}

func LoadPrivateKey(prikeyPath string) (*rsa.PrivateKey, error) {
	certPEMBlock, err := ioutil.ReadFile(prikeyPath)
	if err != nil {
		return nil, err
	}

	return ParsePrivateKey(certPEMBlock)
}

func Sign(priKey *rsa.PrivateKey, data []byte) (string, error) {
	hash := crypto.SHA1
	h := hash.New()
	h.Write(data)
	hashed := h.Sum(nil)

	bs, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bs), nil
}

func Verify(pubKey *rsa.PublicKey, data []byte, sig string) error {
	bs, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		return err
	}

	hash := crypto.SHA1
	h := hash.New()
	h.Write(data)
	hashed := h.Sum(nil)

	return rsa.VerifyPKCS1v15(pubKey, hash, hashed, bs)
}
