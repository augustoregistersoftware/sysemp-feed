package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Service struct {
	secret   []byte
	tokenTTL time.Duration
}

type Claims struct {
	Subject   string `json:"sub"`
	IssuedAt  int64  `json:"iat"`
	ExpiresAt int64  `json:"exp"`
}

type tokenHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

func NewService(secret string, tokenTTL time.Duration) *Service {
	return &Service{
		secret:   []byte(secret),
		tokenTTL: tokenTTL,
	}
}

func (s *Service) GenerateToken(subject string) (string, error) {
	now := time.Now().UTC()

	header := tokenHeader{
		Algorithm: "HS256",
		Type:      "JWT",
	}

	claims := Claims{
		Subject:   subject,
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(s.tokenTTL).Unix(),
	}

	encodedHeader, err := encodeJSON(header)
	if err != nil {
		return "", err
	}

	encodedClaims, err := encodeJSON(claims)
	if err != nil {
		return "", err
	}

	signingInput := encodedHeader + "." + encodedClaims
	signature := s.sign(signingInput)

	return signingInput + "." + signature, nil
}

func (s *Service) ValidateToken(token string) (*Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("token invalido")
	}

	signingInput := parts[0] + "." + parts[1]
	expectedSignature := s.sign(signingInput)

	if !hmac.Equal([]byte(expectedSignature), []byte(parts[2])) {
		return nil, errors.New("assinatura invalida")
	}

	var header tokenHeader
	if err := decodeJSON(parts[0], &header); err != nil {
		return nil, err
	}

	if header.Algorithm != "HS256" || header.Type != "JWT" {
		return nil, errors.New("algoritmo de token invalido")
	}

	var claims Claims
	if err := decodeJSON(parts[1], &claims); err != nil {
		return nil, err
	}

	if claims.Subject == "" {
		return nil, errors.New("token sem usuario")
	}

	if time.Now().UTC().Unix() > claims.ExpiresAt {
		return nil, errors.New("token expirado")
	}

	return &claims, nil
}

func (s *Service) ExpiresInSeconds() int64 {
	return int64(s.tokenTTL.Seconds())
}

func (s *Service) sign(value string) string {
	mac := hmac.New(sha256.New, s.secret)
	mac.Write([]byte(value))

	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func encodeJSON(value any) (string, error) {
	body, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(body), nil
}

func decodeJSON(value string, destination any) error {
	body, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return fmt.Errorf("base64 invalido: %w", err)
	}

	if err := json.Unmarshal(body, destination); err != nil {
		return fmt.Errorf("json invalido: %w", err)
	}

	return nil
}
