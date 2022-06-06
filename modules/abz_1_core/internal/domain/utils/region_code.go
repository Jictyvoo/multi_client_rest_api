package utils

import (
	"fmt"
	"regexp"
	"strings"
)

var RegionCodeRegex = regexp.MustCompile(" \\d{2} ")

func ReplaceRegionCode(s string) string {
	return fmt.Sprintf(" (%s) ", strings.TrimSpace(s))
}
