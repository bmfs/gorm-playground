package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type resultDTO struct {
	Rows []*User
}

func TestGORM(t *testing.T) {

	ptLang := Language{Code: "pt", Name: "Portuguese"}
	DB.Create(&ptLang)
	enLang := Language{Code: "en", Name: "English"}
	DB.Create(&enLang)
	itLang := Language{Code: "it", Name: "Italian"}
	DB.Create(&itLang)

	user1 := User{Name: "jinzhu", Languages: []Language{ptLang, enLang}}
	DB.Create(&user1)

	user2 := User{Name: "john", Languages: []Language{enLang, itLang}}
	DB.Create(&user2)

	var result User
	if err := DB.First(&result, user1.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var result2 []User
	if err := DB.Preload("Languages").Find(&result2).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result2) <= 0 || result2[0].Languages == nil {
		t.Errorf("Failed to retrieve user languages")
	}

	fmt.Println(result2[0].Languages)

	var result3 resultDTO
	if err := DB.Preload("Languages").Find(&result3.Rows).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if len(result2[0].Languages) != len(result3.Rows[0].Languages) {
		t.Error("Expected results to be the same")
	}
	fmt.Println(result3.Rows[0].Languages)

}
