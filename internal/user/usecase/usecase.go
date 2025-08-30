package usecase

import (
	"os"
	"time"

	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/user/repository"
	"github.com/MingPV/PostService/pkg/apperror"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// PostService struct
type PostService struct {
	repo repository.UserRepository
}

// Init PostService
func NewPostService(repo repository.UserRepository) UserUseCase {
	return &PostService{repo: repo}
}

// PostService Methods - 1 Register user (hash password)
func (s *PostService) Register(user *entities.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return apperror.ErrAlreadyExists
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPwd)

	return s.repo.Save(user)
}

// PostService Methods - 2 Login user (check email + password)
func (s *PostService) Login(email string, password string) (string, *entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, err
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

// PostService Methods - 3 Get user by id
func (s *PostService) FindUserByID(id string) (*entities.User, error) {
	return s.repo.FindByID(id)
}

// PostService Methods - 4 Get all users
func (s *PostService) FindAllUsers() ([]*entities.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// PostService Methods - 5 Get user by email
func (s *PostService) GetUserByEmail(email string) (*entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// PostService Methods - 6 Patch
func (s *PostService) PatchUser(id string, user *entities.User) (*entities.User, error) {
	if err := s.repo.Patch(id, user); err != nil {
		return nil, err
	}
	updatedUser, _ := s.repo.FindByID(id)

	return updatedUser, nil
}

// PostService Methods - 7 Delete
func (s *PostService) DeleteUser(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
