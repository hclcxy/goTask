package comment

import (
	"go-task4/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByPostID(postID uint) ([]models.Comment, error)
	GetByID(id uint) (*models.Comment, error)
	Delete(id uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.Preload("User").Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}

func (r *commentRepository) GetByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.Preload("User").First(&comment, id).Error
	return &comment, err
}

func (r *commentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Comment{}, id).Error
}
