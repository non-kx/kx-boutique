package data

import (
	"context"

	"kx-boutique/app/user/internal/biz"
	"kx-boutique/app/user/internal/data/ent"
	"kx-boutique/app/user/internal/data/ent/user"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Implementation of biz.UserRepo interface
func (r *userRepo) GetUserByID(ctx context.Context, id string) (*biz.User, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	user, err := r.data.db.User.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:           user.ID.String(),
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil
}

func (r *userRepo) GetUserByName(ctx context.Context, name string) (*biz.User, error) {
	user, err := r.data.db.User.Query().Where(user.UserName(name)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &biz.User{
		Id:           user.ID.String(),
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	user, err := r.data.db.User.Query().Where(user.Email(email)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &biz.User{
		Id:           user.ID.String(),
		UserName:     user.UserName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil
}

func (r *userRepo) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	saved, err := r.data.db.User.Create().SetUserName(user.UserName).SetEmail(user.Email).SetPasswordHash(user.PasswordHash).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:           saved.ID.String(),
		UserName:     saved.UserName,
		Email:        saved.Email,
		PasswordHash: saved.PasswordHash,
		CreatedAt:    saved.CreatedAt,
		UpdatedAt:    saved.UpdatedAt,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	uuid, err := uuid.Parse(user.Id)
	if err != nil {
		return nil, err
	}

	updated, err := r.data.db.User.UpdateOneID(uuid).SetUserName(user.UserName).SetEmail(user.Email).SetPasswordHash(user.PasswordHash).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		Id:           updated.ID.String(),
		UserName:     updated.UserName,
		Email:        updated.Email,
		PasswordHash: updated.PasswordHash,
		CreatedAt:    updated.CreatedAt,
		UpdatedAt:    updated.UpdatedAt,
	}, nil
}

func (r *userRepo) DeleteUser(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}
