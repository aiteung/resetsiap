package reset

import (
	"fmt"
	"time"
)

func GetCurrentAcademicYear() string {
	now := time.Now()
	year := now.Year()
	nextYear := year + 1
	tahunAkademik := fmt.Sprintf("%d-%d", year, nextYear)
	return tahunAkademik
}
