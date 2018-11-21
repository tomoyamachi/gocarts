package fetcher

import (
	"github.com/tomoyamachi/gocarts/models"
	"reflect"
	"testing"
)

func TestReturnOrCreateCveSlice(t *testing.T) {
	cve1Data := []models.ArticleID{"fuga", "hoge"}
	cves := map[models.CveID][]models.ArticleID{
		"1": cve1Data,
	}
	returned1Data := returnOrCreateCveSlice(cves, "1")
	if reflect.DeepEqual(returned1Data, cve1Data) == false {
		t.Fatalf("Should return %q, but got %q", cve1Data, returned1Data)
	}

	returned3Data := returnOrCreateCveSlice(cves, "3")
	if reflect.DeepEqual([]models.ArticleID{}, returned3Data) == false {
		t.Fatalf("Should return null slice, but got %q", returned3Data)
	}
}

func TestFindCveIDs(t *testing.T) {
	string := "CVE-111-1111 hoge CVE-111-1112 CVE-hoge- CVE-111-1113"
	expected := []models.CveID{
		"CVE-111-1111",
		"CVE-111-1112",
		"CVE-111-1113",
	}
	returned := findCveIDs(string)
	if !reflect.DeepEqual(expected, returned) {
		t.Fatalf("Should return %v, but got %v", expected, returned)
	}
}

func TestAddArticleIDtoCveMap(t *testing.T) {
	cves := models.JpcertCveKeyMap{
		"CVE-111-1111": []models.ArticleID{"1", "2"},
		"CVE-111-1112": []models.ArticleID{"1", "2"},
	}
	expected := models.JpcertCveKeyMap{
		"CVE-111-1111": []models.ArticleID{"1", "2", "3"},
		"CVE-111-1112": []models.ArticleID{"1", "2", "3"},
		"CVE-111-1113": []models.ArticleID{"3"},
	}
	returned := addArticleIDtoCveMap(cves, models.ArticleID("3"), findCveIDs("CVE-111-1111 hoge CVE-111-1112 CVE-hoge- CVE-111-1113"))

	if !reflect.DeepEqual(cves, returned) {
		t.Fatalf("Should return %v, but got %v", expected, returned)
	}
}

func TestAddArticleIDtoCveMap2(t *testing.T) {
	cves := models.JpcertCveKeyMap{} //map[models.CveID][]models.ArticleID{}
	expected := models.JpcertCveKeyMap{
		"CVE-111-1111": []models.ArticleID{"3"},
		"CVE-111-1112": []models.ArticleID{"3"},
	}
	returned := addArticleIDtoCveMap(cves, models.ArticleID("3"), findCveIDs("CVE-111-1111 hoge CVE-111-1112"))

	if !reflect.DeepEqual(expected, returned) {
		t.Fatalf("Should return %v, but got %v", expected, returned)
	}
}
