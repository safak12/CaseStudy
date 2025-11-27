package service

import (
	"context"
	"fmt"
	
	
	"depixen/backend/database" 
	"depixen/backend/repository" 
)

type CaseStudyService interface {
	CreateCard(ctx context.Context, title, description, imageBase64 string) (database.CaseStudy, error)
	ListCards(ctx context.Context) ([]database.CaseStudy, error)
}


type caseStudyService struct {
	repo repository.CaseStudyRepository
	uploader *Uploader 
}

func NewCaseStudyService(repo repository.CaseStudyRepository, uploader *Uploader) CaseStudyService { 
	return &caseStudyService{repo: repo, uploader: uploader}
}


func (s *caseStudyService) CreateCard(ctx context.Context, title, description, imageBase64 string) (database.CaseStudy, error) {
    
    imageURL := ""
    if imageBase64 != "" {
        var err error
        imageURL, err = s.uploader.UploadBase64Image(ctx, imageBase64)
        if err != nil {
            return database.CaseStudy{}, fmt.Errorf("görsel yüklenemedi: %w", err)
        }
    }
    
    
	card := database.CaseStudy{
		Title: title,
		Description: description,
        ImageURI: imageURL, 
	}
	
    
	err := s.repo.CreateCard(&card)
	if err != nil {
		return database.CaseStudy{}, err
	}
	return card, nil
}


func (s *caseStudyService) ListCards(ctx context.Context) ([]database.CaseStudy, error) {
	return s.repo.GetAllCards()
}