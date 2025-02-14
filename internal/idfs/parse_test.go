package idfs_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/maniartech/gotime/internal/idfs"
)

func TestParse(t *testing.T) {
	// Test case for parsing the date 24-01-1984
	format, _ := idfs.Parse("dd-mm-yyyy", "24-01-1984")
	correctTime := time.Date(1984, 1, 24, 0, 0, 0, 0, time.UTC)
	if !format.Equal(correctTime) {
		t.Errorf("Expected %v, got, %v", correctTime, format)
	}
}

func TestParseWithLocation(t *testing.T) {
	// Test case for parsing the date 24-01-1984
	format, _ := idfs.ParseInLocation("dd-mm-yyyy", "24-01-1984", time.FixedZone("IST", 5.5*60*60))
	correctTime := time.Date(1984, 1, 24, 0, 0, 0, 0, time.FixedZone("IST", 5.5*60*60))
	if !format.Equal(correctTime) {
		t.Errorf("Expected %v, got, %v", correctTime, format)
	}
}

func BenchmarkParse(b *testing.B) {

	// Benchmarking for Parse function
	for i := 0; i < b.N; i++ {
		idfs.Parse("24-01-1984", "dd-mm-yyyy")
	}
}

func TestTrial(t *testing.T) {
	tm := time.Now().Add(8 * time.Hour)
	tmu := tm.UTC()

	// truncate time from tm and tmu and store them in tru and truu
	// tru := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, tm.Location())
	truu := time.Date(tmu.Year(), tmu.Month(), tmu.Day(), 0, 0, 0, 0, tmu.Location())

	fmt.Println(truu, truu.Local())
}
