package repository

import (
	"depixen/backend/database" 
	"time"
	"gorm.io/gorm"
)

type CaseStudyRepository interface {
	CreateCard(card *database.CaseStudy) error
	GetAllCards() ([]database.CaseStudy, error)
}

type GormCaseStudyRepository struct {
	db *gorm.DB
}

func NewGormCaseStudyRepository(db *gorm.DB) CaseStudyRepository {
	return &GormCaseStudyRepository{db: db}
}

func (r *GormCaseStudyRepository) CreateCard(card *database.CaseStudy) error {
	card.CreatedDate = time.Now()
	result := r.db.Table("tb_casestudy").Create(card)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormCaseStudyRepository) GetAllCards() ([]database.CaseStudy, error) {
	var cards []database.CaseStudy
	result := r.db.Table("tb_casestudy").Find(&cards)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return cards, nil
}