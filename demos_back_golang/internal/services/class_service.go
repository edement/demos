package services

import (
	"context"
	"demos_back_golang/internal/models"
	"demos_back_golang/internal/storage"
)

type ClassService interface {
	CreateClass(ctx context.Context, req models.CreateClassRequest, trainerID int64) (models.ClassResponse, error)
	GetClassById(ctx context.Context, classID int64) (models.ClassResponse, error)
	GetClasses(ctx context.Context) ([]models.ClassResponse, error)
	DeleteClass(ctx context.Context, classID int64) error
	UpdateClass(ctx context.Context, classID int64, req models.UpdateClassRequest) (models.ClassResponse, error)
}

type classService struct {
	classRepo storage.ClassRepository
}

func NewClassService(classRepo storage.ClassRepository) ClassService {
	return &classService{classRepo: classRepo}
}

func (s *classService) CreateClass(ctx context.Context, req models.CreateClassRequest, trainerID int64) (models.ClassResponse, error) {
	return s.classRepo.CreateClass(ctx, req, trainerID)
}

func (s *classService) GetClassById(ctx context.Context, classID int64) (models.ClassResponse, error) {
	return s.classRepo.GetClassById(ctx, classID)
}

func (s *classService) GetClasses(ctx context.Context) ([]models.ClassResponse, error) {
	return s.classRepo.GetClasses(ctx)
}

func (s *classService) DeleteClass(ctx context.Context, classID int64) error {
	return s.classRepo.DeleteClass(ctx, classID)
}

func (s *classService) UpdateClass(ctx context.Context, classID int64, req models.UpdateClassRequest) (models.ClassResponse, error) {
	return s.classRepo.UpdateClass(ctx, classID, req)
}
