package post

import (
	"go-task4/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *models.Post) error
	GetByID(id uint) (*models.Post, error)
	GetAll() ([]models.Post, error)
	Update(post *models.Post) error
	Delete(id uint) error
	GetByUserID(userID uint) ([]models.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) GetByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.db.Preload("User").Preload("Comments.User").First(&post, id).Error
	return &post, err
}

func (r *postRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("User").Find(&posts).Error
	return posts, err
}

func (r *postRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(id uint) error {
	return r.db.Delete(&models.Post{}, id).Error
}

func (r *postRepository) GetByUserID(userID uint) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}
