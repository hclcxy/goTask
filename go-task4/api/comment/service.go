package comment

import (
	"errors"
	"go-task4/api/post"
	"go-task4/models"
)

type CommentService interface {
	CreateComment(userID uint, req *CreateCommentRequest) (*models.Comment, error)
	GetCommentsByPostID(postID uint) ([]models.Comment, error)
	DeleteComment(userID, commentID uint) error
}

type commentService struct {
	commentRepo CommentRepository
	postRepo    post.PostRepository
}

func NewCommentService(commentRepo CommentRepository, postRepo post.PostRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		postRepo:    postRepo,
	}
}

func (s *commentService) CreateComment(userID uint, req *CreateCommentRequest) (*models.Comment, error) {
	// 检查文章是否存在
	if _, err := s.postRepo.GetByID(req.PostID); err != nil {
		return nil, errors.New("post not found")
	}

	comment := &models.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  req.PostID,
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	return s.commentRepo.GetByPostID(postID)
}

func (s *commentService) DeleteComment(userID, commentID uint) error {
	comment, err := s.commentRepo.GetByID(commentID)
	if err != nil {
		return errors.New("comment not found")
	}

	if comment.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.commentRepo.Delete(commentID)
}
