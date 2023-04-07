package crypto

import "github.com/google/uuid"

func MakeUUID() string {
	u4 := uuid.New()
	return u4.String()
}
