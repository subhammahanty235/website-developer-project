package middlewares

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PortfolioMakerAppClaims struct {
	jwt.RegisteredClaims
	Email string             `json:"email"`
	ID    primitive.ObjectID `json:"id"`
}

// Example public key (replace with your actual public key)
var publicKey *ecdsa.PublicKey

func init() {
	println("Inside the init function ------------------------------------------------>")
	// Load your ECDSA public key from PEM file
	publicKeyFile := "public.pem" // Replace with your PEM file path
	publicKeyBytes, err := os.ReadFile(publicKeyFile)
	if err != nil {
		panic(fmt.Sprintf("failed to read public key file: %v", err))
	}

	println(publicKeyBytes)

	block, _ := pem.Decode(publicKeyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		panic("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(fmt.Sprintf("failed to parse DER encoded public key: %v", err))
	}

	var ok bool
	publicKey, ok = pub.(*ecdsa.PublicKey)

	if !ok {
		panic("not a valid ECDSA public key")
	}
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		print(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		publicKeyFile := "public.pem" // Replace with your PEM file path
		publicKeyBytes, err := os.ReadFile(publicKeyFile)
		if err != nil {
			panic(fmt.Sprintf("failed to read public key file: %v", err))
		}

		println(publicKeyBytes)

		block, _ := pem.Decode(publicKeyBytes)
		if block == nil || block.Type != "PUBLIC KEY" {
			panic("failed to decode PEM block containing public key")
		}

		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			panic(fmt.Sprintf("failed to parse DER encoded public key: %v", err))
		}

		var ok bool
		publicKey, ok = pub.(*ecdsa.PublicKey)

		if !ok {
			panic("not a valid ECDSA public key")
		}

		// The token normally comes as "Bearer <token>", so split it
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &PortfolioMakerAppClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil {
			println()
			println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*PortfolioMakerAppClaims); ok && token.Valid {
			println("Token is correct ------------------>")
			println(claims.ID.Hex())
			// Token is valid, set the claims in the context
			c.Set("userEmail", claims.Email)
			c.Set("userID", claims.ID.Hex())

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
