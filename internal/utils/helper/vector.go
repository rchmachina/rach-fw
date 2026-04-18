package helper

import (
	"fmt"
	"strings"
)

func VectorToString(vec []float32) string {
	strs := make([]string, len(vec))
	for i, v := range vec {
		strs[i] = fmt.Sprintf("%f", v)
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ","))
}
