package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func GenerateOrderNumber() string {

	now := time.Now()

	return fmt.Sprintf(
		"ORD-%s-%s",
		now.Format("20060102"),
		uuid.New().String()[:8],
	)

}
