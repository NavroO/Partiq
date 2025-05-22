package shared

import (
	"fmt"
	"strconv"
	"strings"
)

func GetIDFromURL(urlPath string) (int64, error) {
	segments := strings.Split(urlPath, "/")
	if len(segments) == 0 {
		return 0, fmt.Errorf("invalid URL path")
	}

	idStr := segments[len(segments)-1]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid ID: %v", err)
	}

	return id, nil
}
