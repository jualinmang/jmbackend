package auth

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"
)

func VerifyIDToken(idToken string, audience string) (*idtoken.Payload, error) {
	payload, err := idtoken.Validate(context.Background(), idToken, audience)
	if err != nil {
		return nil, fmt.Errorf("id token validation failed: %v", err)
	}
	return payload, nil
}
