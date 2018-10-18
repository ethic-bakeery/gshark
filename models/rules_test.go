package models_test

import (
	"x-patrol/models"

	"testing"
)

func TestLoadRuleFromFile(t *testing.T) {
	filename := "/data/code/golang/src/x-patrol/conf/gitrob.json"
	t.Log(models.LoadRuleFromFile(filename))
}

func TestInsertRules(t *testing.T) {
	filename := "/data/code/golang/src/x-patrol/conf/gitrob.json"
	rules, err := models.GetRules()
	t.Log(rules, err)
	if err == nil && len(rules) == 0 {
		t.Logf("Init rules, err: %v", models.InsertRules(filename))
	}

}
