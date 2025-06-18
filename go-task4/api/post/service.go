package post

import (
	"errors"
	"go-task4/models"
)

type PostService interface {
	CreatePost(userID uint, req *CreatePostRequest) (*models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	UpdatePost(userID, postID uint, req *UpdatePostRequest) (*models.Post, error)
	DeletePost(userID, postID uint) error
}

type postService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(userID uint, req *CreatePostRequest) (*models.Post, error) {
	post := &models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := s.repo.Create(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	return s.repo.GetByID(id)
}

func (s *postService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAll()
}

func (s *postService) UpdatePost(userID, postID uint, req *UpdatePostRequest) (*models.Post, error) {
	post, err := s.repo.GetByID(postID)
	if err != nil {
		return nil, errors.New("post not found")
	}

	if post.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if req.Title != "" {
		post.Title = req.Title
	}

	if req.Content != "" {
		post.Content = req.Content
	}

	if err := s.repo.Update(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) DeletePost(userID, postID uint) error {
	post, err := s.repo.GetByID(postID)
	if err != nil {
		return errors.New("post not found")
	}

	if post.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(postID)
}
