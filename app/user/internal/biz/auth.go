package biz

import (
	"context"
	v1 "kx-boutique/api/user/v1"
	"kx-boutique/app/user/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAuthUnspecified   = errors.BadRequest(v1.ErrorReason_AUTHENTICATION_UNSPECIFIED.String(), "auth unspecified")
	ErrAuthFailed        = errors.Unauthorized(v1.ErrorReason_AUTHENTICATION_FAILED.String(), "auth failed")
	ErrUserNotFound      = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
	ErrUserAlreadyExists = errors.Conflict(v1.ErrorReason_USER_ALREADY_EXISTS.String(), "user already exists")
)

type SignInRequest struct {
	UserName string
	Password string
}

type SignInReply struct {
	Token string
}

type SignUpRequest struct {
	UserName string
	Password string
	Email    string
}

type SignUpReply struct {
	Token string
}

type AuthUsecase struct {
	repo       UserRepo
	secret     string
	expiration time.Duration
	log        *log.Helper
}

func NewAuthUsecase(repo UserRepo, conf *conf.Auth, logger log.Logger) *AuthUsecase {
	return &AuthUsecase{repo: repo, secret: conf.Jwt.Secret, expiration: conf.Jwt.Expire.AsDuration(), log: log.NewHelper(logger)}
}

func (uc *AuthUsecase) SignIn(ctx context.Context, req *SignInRequest) (string, error) {
	uc.log.WithContext(ctx).Infof("Sign in user[%v]", req.UserName)

	user, err := uc.repo.GetUserByName(ctx, req.UserName)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Get user by name error: %v", err)
		return "", err
	}

	if user == nil {
		uc.log.WithContext(ctx).Errorf("User not found")
		return "", ErrUserNotFound
	}

	if !CheckPasswordHash(req.Password, user.PasswordHash) {
		return "", ErrAuthFailed
	}

	signed, err := uc.generateToken(ctx, user.Id)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Generate token error: %v", err)
		return "", ErrAuthUnspecified
	}

	return signed, nil
}

func (uc *AuthUsecase) SignUp(ctx context.Context, req *SignUpRequest) (string, error) {
	uc.log.WithContext(ctx).Infof("Sign up user[%v]", req.UserName)

	user, err := uc.repo.GetUserByName(ctx, req.UserName)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Get user by name error: %v", err)
		return "", ErrAuthUnspecified
	}

	if user != nil {
		uc.log.WithContext(ctx).Errorf("User already exists")
		return "", ErrUserAlreadyExists
	}

	hash, err := HashPassword(req.Password)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Hash password error: %v", err)
		return "", ErrAuthUnspecified
	}

	user = &User{
		UserName:     req.UserName,
		Email:        req.Email,
		PasswordHash: hash,
	}

	user, err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Create user error: %v", err)
		return "", ErrAuthUnspecified
	}

	uc.log.WithContext(ctx).Infof("Signed with secret: %v", uc.secret)
	signed, err := uc.generateToken(ctx, user.Id)
	if err != nil {
		uc.log.WithContext(ctx).Errorf("Generate token error: %v", err)
		return "", ErrAuthUnspecified
	}

	return signed, nil
}

func (uc *AuthUsecase) generateToken(ctx context.Context, userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(uc.expiration).Unix()
	claims["authorized"] = true
	claims["user"] = userId

	ss, err := token.SignedString([]byte(uc.secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
