package database

import (
	"fmt"
	"graphql-hasura-demo/graph/model"
	"graphql-hasura-demo/internal/utils"
	"strconv"
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

// func ExportStudentByClass(class string) (*model.BaseResponseView, error) {
// 	var students []Student
// 	if err := db.Where("class = ?", class).Find(&students); err != nil {
// 		return &model.BaseResponseView{
// 			Success: false,
// 			Message: fmt.Sprintf("fetch list failed: %v", err),
// 			Data:    nil,
// 		}, nil
// 	}

// 	f := excelize.NewFile()
// 	sheet := "Sheet1"
// 	f.SetCellValue(sheet, "A1", "Mã học sinh")
// 	f.SetCellValue(sheet, "B1", "Tên")
// 	f.SetCellValue(sheet, "C1", "Ngày sinh")
// 	f.SetCellValue(sheet, "D1", "Giới tính")
// 	f.SetCellValue(sheet, "E1", "Lớp")
// 	f.SetCellValue(sheet, "F1", "Điểm trung bình")
// 	f.SetCellValue(sheet, "G1", "Học lực")

// 	for i, student := range students {
// 		row := i + 2
// 		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), student.StudentID)
// 		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), student.Name)
// 		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), student.DateOfBirth)
// 		f.SetCellValue(sheet, fmt.Sprintf("D%d", row), student.Gender)
// 		f.SetCellValue(sheet, fmt.Sprintf("E%d", row), student.Class)
// 		f.SetCellValue(sheet, fmt.Sprintf("F%d", row), student.OverallAverage)
// 		f.SetCellValue(sheet, fmt.Sprintf("G%d", row), student.AcademicPerformance)

// 		// Thêm cột cho điểm trung bình từng môn
// 		for j, avg := range student.SubjectAverages {
// 			col := string(rune('H') + rune(j))
// 			if i == 0 {
// 				f.SetCellValue(sheet, fmt.Sprintf("%s1", col), avg.Subject.Name)
// 			}
// 			f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), avg.AverageScore)
// 		}
// 	}

// 	// Lưu file Excel
// 	filePath := fmt.Sprintf("students_class_%s.xlsx", class)
// 	if err := f.SaveAs(filePath); err != nil {
// 		return &model.BaseResponseView{
// 			Success: false,
// 			Message: fmt.Sprintf("lưu file Excel thất bại: %v", err),
// 		}, nil
// 	}

// 	return &model.BaseResponseView{
// 		Success: true,
// 		Message: fmt.Sprintf("File Excel được lưu tại: %s", filePath),
// 		Data:    nil,
// 	}, nil
// }

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

func BuildStudentModel(student Student) (*model.Student, error) {
	var grade Grade
	grades, err := grade.GetGradesByStudent(student.ID)
	if err != nil {
		return nil, fmt.Errorf("lấy điểm thất bại: %w", err)
	}

	subjectAverages := make([]*model.SubjectAverage, 0)
	subjectScores := make(map[uint][]float64)
	for _, grade := range grades {
		subjectScores[grade.SubjectID] = append(subjectScores[grade.SubjectID], grade.Score)
	}
	totalScore := 0.0
	totalGrades := 0
	for subjectID, scores := range subjectScores {
		subject, err := GetSubjectByID(subjectID)
		if err != nil {
			return nil, fmt.Errorf("lấy môn học thất bại: %w", err)
		}

		avgScore := 0.0
		if len(scores) > 0 {
			sum := 0.0
			for _, score := range scores {
				sum += score
			}
			avgScore = sum / float64(len(scores))
			totalScore += sum
			totalGrades += len(scores)
		}

		subjectAverages = append(subjectAverages, &model.SubjectAverage{
			Subject: &model.Subject{
				ID:        strconv.FormatUint(uint64(subject.ID), 10),
				SubjectID: subject.SubjectID,
				Name:      subject.Name,
			},
			AverageScore: avgScore,
		})
	}

	overallAverage := 0.0
	if totalGrades > 0 {
		overallAverage = totalScore / float64(totalGrades)
	}

	academicPerformance := utils.CalculateAcademicPerformance(overallAverage)
	return &model.Student{
		ID:                  strconv.FormatUint(uint64(student.ID), 10),
		StudentID:           student.StudentID,
		Name:                student.Name,
		DateOfBirth:         student.DateOfBirth,
		Gender:              model.Gender(student.Gender),
		Class:               student.Class,
		SubjectAverages:     subjectAverages,
		OverallAverage:      overallAverage,
		AcademicPerformance: academicPerformance,
	}, nil
}
