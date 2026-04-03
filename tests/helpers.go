package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"
)

// GenerateUUID generates a random UUID.
func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	if _, err := rand.Read(uuid); err != nil {
		return "", err
	}
	uuid[6] = (uuid[6] & 0x0F) | 0x40
	uuid[8] = (uuid[8] & 0x3F) | 0x80
	return hex.EncodeToString(uuid), nil
}

// ValidateEmail checks if the email is valid.
func ValidateEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// FormatCurrency formats the amount as a currency string.
func FormatCurrency(amount float64) string {
	return "$" + strings.TrimRight(strings.TrimRight(formatFloat(amount), "0"), ".")
}

// formatFloat formats a float64 to two decimal places.
func formatFloat(value float64) string {
	return strings.TrimRight(strings.TrimRight(formatFloat(value), "0"), ".")
}

// ParseDate parses a date string into a time.Time object.
func ParseDate(dateString string) (time.Time, error) {
	return time.Parse("2006-01-02", dateString)
}

// IsExpired checks if a given date is in the past.
func IsExpired(date time.Time) bool {
	return date.Before(time.Now())
}

// MaskCardNumber masks the middle digits of a card number.
func MaskCardNumber(cardNumber string) (string, error) {
	if len(cardNumber) < 12 {
		return "", errors.New("invalid card number length")
	}
	return cardNumber[:4] + "****" + cardNumber[len(cardNumber)-4:], nil
}

// GenerateRandomString generates a random string of a given length.
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}