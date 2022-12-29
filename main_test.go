package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type resultDTO struct {
	Rows []*BusinessMetric
}

type genericDTO struct {
	Rows interface{}
}

func TestGORM(t *testing.T) {

	event1 := Event{Name: "event1", Description: "event 1"}
	DB.Create(&event1)
	event2 := Event{Name: "event2", Description: "event 2"}
	DB.Create(&event2)
	event3 := Event{Name: "event3", Description: "event 3"}
	DB.Create(&event3)

	bm1 := BusinessMetric{Name: "bm_1", Events: []*Event{&event1, &event2}}
	DB.Create(&bm1)
	bm2 := BusinessMetric{Name: "bm_2", Events: []*Event{&event2, &event3}}
	DB.Create(&bm2)

	var result BusinessMetric
	if err := DB.First(&result, bm1.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result2 []BusinessMetric
	if err := DB.Preload("Events").Find(&result2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result2) <= 0 || result2[0].Events == nil {
		t.Errorf("Failed to retrieve user languages")
	}

	printResult(result2)

	result3 := &resultDTO{}
	if err := DB.Preload("Events").Find(&result3.Rows).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result2) != len(result3.Rows[0].Events) {
		t.Error("Expected results to be the same")
	}

	printResult(result3)

	result4 := &genericDTO{
		Rows: []*BusinessMetric{},
	}
	if err := DB.Preload("Events").Find(&result4.Rows).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	printResult(result4.Rows)

}

func printResult(data interface{}) {
	str, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(str))
}
