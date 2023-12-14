package repository

import (
	"context"

	"github.com/lyonnee/go-gin-template/ent"
	"github.com/lyonnee/go-gin-template/ent/user"
	"github.com/lyonnee/go-gin-template/infra/database/entorm"
)

type UserRepo struct {
	entCli *ent.Client
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		entCli: entorm.GetClient(),
	}
}

func (repo *UserRepo) GetById(ctx context.Context, id uint64) (*ent.User, error) {
	return repo.entCli.User.Query().Where(user.ID(id)).Only(ctx)
}

func (repo *UserRepo) GetByPhoneNumber(ctx context.Context, phoneNum string) (*ent.User, error) {
	return repo.entCli.User.Query().Where(user.PhoneNumber(phoneNum)).Only(ctx)
}

func (repo *UserRepo) InsertOne(
	ctx context.Context,
	name string,
	age uint8,
	phoneNum string,
	authHash string,
) (*ent.User, error) {
	return repo.entCli.User.Create().SetName(name).SetAge(age).SetPhoneNumber(phoneNum).SetAuthHash(authHash).Save(ctx)
}
