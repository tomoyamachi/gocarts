package fetcher

import (
	"github.com/tomoyamachi/gocarts/models"
	"reflect"
	"testing"
)

func TestFindCveIDs(t *testing.T) {
	sample := "CVE-111-1111 hoge CVE-111-1112 CVE-hoge- CVE-111-1113"
	expected := []models.CveID{
		"CVE-111-1111",
		"CVE-111-1112",
		"CVE-111-1113",
	}
	returned := findCveIDs(sample)
	if !reflect.DeepEqual(expected, returned) {
		t.Fatalf("Should return %v, but got %v", expected, returned)
	}
}
