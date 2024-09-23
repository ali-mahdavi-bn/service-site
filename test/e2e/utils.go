package e2e

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func RandomSuffix() string {
	return uuid.NewString()[:6]
}

func RandomSku(name ...string) string {
	return fmt.Sprintf("sku-%s-%s", strings.Join(name, "-"), RandomSuffix())
}

func RandomBatchref(name ...string) string {
	return fmt.Sprintf("batchref-%s-%s", strings.Join(name, "-"), RandomSuffix())
}

func RandomOrderid(name ...string) string {
	return fmt.Sprintf("orderid-%s-%s", strings.Join(name, "-"), RandomSuffix())
}
