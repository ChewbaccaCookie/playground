package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	createUser := User{Name: "jinzhu"}

	DB.Create(&createUser)

	var result User
	if err := DB.First(&result, createUser.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	// Preload nested
	var user User
	DB.Preload(clause.Associations).Preload("Languages." + clause.Associations).Find(&user)

}
