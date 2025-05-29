package database

import "fmt"

type Grade struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	StudentID uint    `gorm:"index;not null"`
	SubjectID uint    `gorm:"index;not null"`
	GradeType string  `gorm:"type:varchar(20);not null"`
	Score     float64 `gorm:"not null"`

	Student Student //`gorm:"foreignKey:StudentID"`
	Subject Subject // `gorm:"foreignKey:SubjectID"`
}

func (g *Grade) CreateGrade() (*Grade, error) {
	result := db.Create(g)
	if result.Error != nil {
		return nil, fmt.Errorf("create grade failed: %v", result.Error)
	}
	// if err := db.Preload("Student").Preload("Subject").First(g, g.ID).Error; err != nil {
	// 	return nil, fmt.Errorf("load failed: %v", err)
	// }
	return g, result.Error
}

func (g *Grade) UpdateGrade(id uint) (*Grade, error) {
	var updateGrade Grade
	if err := db.First(&updateGrade, id).Error; err != nil {
		return nil, err
	}

	updateGrade.StudentID = g.StudentID
	updateGrade.SubjectID = g.SubjectID
	updateGrade.GradeType = g.GradeType
	updateGrade.Score = g.Score

	result := db.Save(&updateGrade)
	if result != nil {
		return nil, result.Error
	}

	return &updateGrade, nil
}

func (g *Grade) DeleteGrade(id uint) (bool, error) {
	var deleteGrade Grade
	if err := db.First(&deleteGrade, id).Error; err != nil {
		return false, err
	}

	result := db.Delete(&deleteGrade)

	if result.Error != nil {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}

func (g *Grade) GetGradesByStudent(studentID uint) ([]Grade, error) {
	var grades []Grade
	result := db.Preload("Student").Preload("Subject").Where("student_id = ?", studentID).Find(&grades)
	if result.Error != nil {
		return nil, fmt.Errorf("can get list grade: %v", result.Error)
	}

	return grades, nil
}
