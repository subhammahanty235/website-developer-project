package utils

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var privateKey *ecdsa.PrivateKey
var publicKey *ecdsa.PublicKey

type PortfolioMakerAppClaims struct {
	jwt.RegisteredClaims
	Email string
	ID    primitive.ObjectID
}

func loadPrivateKeyFromPEM() *ecdsa.PrivateKey {
	privKeyPEM, err := os.ReadFile("private.key")
	if err != nil {
		panic(fmt.Errorf("error reading private key file: %v", err))
	}

	block, _ := pem.Decode(privKeyPEM)
	if block == nil {
		panic("failed to decode PEM block containing private key")
	}

	privKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Errorf("error parsing private key: %v", err))
	}

	return privKey
}
func loadPublicKeyFromPEM() *ecdsa.PublicKey {
	pubKeyPEM, err := os.ReadFile("public-key.pem")
	if err != nil {
		panic(fmt.Errorf("error reading public key file: %v", err))
	}

	block, _ := pem.Decode(pubKeyPEM)
	if block == nil {
		panic("failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(fmt.Errorf("error parsing public key: %v", err))
	}

	ecdsaPubKey, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid public key type")
	}

	return ecdsaPubKey
}

func Generate(email string, id primitive.ObjectID) (string, error) {

	// privateKey, generr := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// if generr != nil {
	// 	println("error generating the key, ", generr.Error())
	// 	return "", nil
	// }
	// loadPrivateKeyFromPEM()
	privateKey = loadPrivateKeyFromPEM()

	portfolioMakerAppClaims := PortfolioMakerAppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "PortfolioMakerAppUser",
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(24 * time.Hour),
			},
		},
		Email: email,
		ID:    id,
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, portfolioMakerAppClaims).SignedString(privateKey)
	if err != nil {
		println(err.Error())
		return "", err
	}

	return token, nil

}
