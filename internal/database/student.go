package database

import (
	"fmt"
	"strings"
)

type Student struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	StudentID   string `gorm:"uniqueIndex;not null"`
	Name        string `gorm:"type:varchar(255);not null"`
	DateOfBirth string `gorm:"type:date;not null"`
	Gender      string `gorm:"type:varchar(10);not null"`
	Class       string `gorm:"type:varchar(50);not null"`

	Grades []Grade
}

func (s *Student) CreateStudent() (*Student, error) {
	result := db.Create(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

func GetAllStudents() ([]Student, error) {
	var students []Student
	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func GetStudentByID(id uint) (*Student, error) {
	var student Student
	if err := db.First(&student, id).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *Student) UpdateStudent(id uint) (*Student, error) {
	var student Student
	if err := db.First(&student, id).Error; err != nil {
		return nil, err
	}

	student.StudentID = s.StudentID
	student.Name = s.Name
	student.DateOfBirth = s.DateOfBirth
	student.Gender = s.Gender
	student.Class = s.Class
	result := db.Save(&student)
	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

func (s *Student) DeleteStudent(id uint) (bool, error) {
	result := db.Delete(&Student{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func SearchStudentByField(field string, value string) ([]Student, error) {
	var students []Student
	field = strings.ToLower(field)
	query := fmt.Sprintf("%s = ?", field)
	err := db.Where(query, value).Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
