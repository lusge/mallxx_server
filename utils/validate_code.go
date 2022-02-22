package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func CreateValidateCode() string {
	mun := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	r := len(mun)
	rand.Seed(time.Now().UnixNano())

	var sbuild strings.Builder

	for i := 0; i < 6; i++ {
		fmt.Fprintf(&sbuild, "%d", mun[rand.Intn(r)])
	}

	return sbuild.String()
}
