package database

type Subject struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	SubjectID string `gorm:"uniqueIndex;not null"`
	Name      string `gorm:"type:varchar(255);not null"`

	Grades []Grade
}

func (s *Subject) CreateSubject() (*Subject, error) {
	result := db.Create(s)
	if result.Error != nil {
		return nil, result.Error
	}
	return s, nil
}

func GetAllSubjects() ([]Subject, error) {
	var subjects []Subject
	if err := db.Find(&subjects).Error; err != nil {
		return nil, err
	}
	return subjects, nil
}

func GetSubjectByID(id uint) (*Subject, error) {
	var subject Subject
	if err := db.First(&subject, id).Error; err != nil {
		return nil, err
	}
	return &subject, nil
}

func (s *Subject) UpdateSubject(id string) (*Subject, error) {
	var subject Subject
	if err := db.First(&subject, id).Error; err != nil {
		return nil, err
	}

	subject.SubjectID = s.SubjectID
	subject.Name = s.Name
	result := db.Save(&subject)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subject, nil
}
func (s *Subject) DeleteSubject(id string) (bool, error) {
	var subject Subject
	if err := db.First(&subject, id).Error; err != nil {
		return false, err
	}

	result := db.Delete(&subject)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
