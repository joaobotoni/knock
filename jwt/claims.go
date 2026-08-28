package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const Issuer = "knock"
const DefaultTTL = 15 * time.Minute

type registered = jwt.RegisteredClaims

type Customer struct {
	ID    string   `json:"-"`
	Name  string   `json:"name,omitempty"`
	Email string   `json:"email,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type claims struct {
	registered
	Customer
}

func newClaims(customer Customer, now time.Time, ttl time.Duration) claims {
	return claims{
		registered: registered{
			Subject:   customer.ID,
			Issuer:    Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
		Customer: customer,
	}
}
