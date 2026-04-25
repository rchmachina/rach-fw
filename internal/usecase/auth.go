package usecase

import (
	"context"
	"fmt"

	"github.com/rchmachina/rach-fw/configs"
	modelAuth "github.com/rchmachina/rach-fw/internal/dto/model/auth"
	modelJwt "github.com/rchmachina/rach-fw/internal/dto/model/jwt"
	request "github.com/rchmachina/rach-fw/internal/dto/request/auth"
	"github.com/rchmachina/rach-fw/internal/utils/helper"
	"github.com/rchmachina/rach-fw/internal/utils/jwt"

	logger "github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	repo "github.com/rchmachina/rach-fw/internal/repository"
)

type AuthUsecase interface {
	LoginAuth(ctx context.Context, req *request.LoginUser) (*modelAuth.GetToken, error)
	CreateUser(ctx context.Context, req request.CreateAuth) (*modelAuth.GetToken, error)
	RevokeAuth(ctx context.Context, req request.Token) error
	CreateNewAccessToken(ctx context.Context, req request.Token) (*modelAuth.GetToken, error)
}

type authUsecase struct {
	repo   repo.Repository
	logger logger.Logger
	cfg    configs.Configs
}

func NewAuthUserCase(r *repo.Repository, logger logger.Logger, cfg configs.Configs) AuthUsecase {
	return &authUsecase{
		repo:   *r,
		logger: logger,
		cfg:    cfg,
	}
}

func (u *authUsecase) LoginAuth(ctx context.Context, req *request.LoginUser) (*modelAuth.GetToken, error) {
	log := u.logger.WithCtx(ctx)
	userModel, err := u.repo.Sql.UserSchema.AuthRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Info("LoginAuth", logger.ToField("error_LoginAuth_get_password", err))
		return nil, err
	}

	err = helper.CheckPassword(req.Password, userModel.Password)

	if err != nil {
		count, err := u.repo.NoSql.Auth.GetLoginAttempt(ctx, userModel.Email)
		if err != nil {
			log.Error("LoginAuth", logger.ToField("error_LoginAuth_GetLoginAttemp", err))
			return nil, fmt.Errorf("uppsss there is something wrong")
		}
		if count > 3 {
			log.Error("LoginAuth", logger.ToField("error_LoginAuth_GetLoginAttemp", fmt.Errorf("already max attemp, please try again letter")))
			return nil, fmt.Errorf("already max attemp, please try again letter")
		}
		_, err = u.repo.NoSql.Auth.IncrementLoginAttempt(ctx, userModel.Email, u.cfg.TTlRedis.TTLAttemp)
		if err != nil {
			log.Error("LoginAuth", logger.ToField("error_LoginAuth_IncrementLoginAttempt", err))
			return nil, fmt.Errorf("uppsss there is something wrong")
		}
		log.Error("LoginAuth", logger.ToField("error_LoginAuth_check_password", fmt.Errorf("pass and hashed pass not same")))
		return nil, fmt.Errorf("uppsss there is something wrong")
	}

	err = u.repo.NoSql.Auth.RemoveLoginAttempt(ctx, userModel.Email, u.cfg.TTlRedis.TTLAttemp)
	if err != nil {
		log.Error("LoginAuth", logger.ToField("error_LoginAuth_RemoveLoginAttempt", err))
		return nil, fmt.Errorf("uppsss there is something wrong")
	}

	modelToken := modelJwt.TokenValue{
		Name:  userModel.Name,
		Id:    userModel.Id,
		Email: userModel.Email,
	}

	return u.createLoginToken(ctx, modelToken)
}

func (u *authUsecase) CreateUser(ctx context.Context, req request.CreateAuth) (*modelAuth.GetToken, error) {

	log := u.logger.WithCtx(ctx)

	hashedPassword, err := helper.HashPassword(req.PasswordHash)
	if err != nil {
		log.Error("error_CreateUser", logger.ToField("error_CreateUser_HashPassword", err))
		return nil, fmt.Errorf("Upss somethingwrong")
	}

	req.PasswordHash = hashedPassword
	id, err := u.repo.Sql.UserSchema.AuthRepo.CreateUser(ctx, req)

	if err != nil {
		return nil, err
	}

	modelToken := modelJwt.TokenValue{
		Name:  req.Name,
		Id:    id,
		Email: req.Email,
	}

	return u.createLoginToken(ctx, modelToken)
}

func (u *authUsecase) RevokeAuth(ctx context.Context, req request.Token) error {
	getInfoUser, err := jwt.GetUserInfo(ctx)
	if err != nil {
		return fmt.Errorf("upss something wrong")
	}

	err = u.repo.NoSql.Auth.RevokeRefreshToken(ctx, getInfoUser.Email, req.RefreshToken)
	if err != nil {
		return fmt.Errorf("upss something wrong")
	}

	return nil
}

func (u *authUsecase) CreateNewAccessToken(ctx context.Context, req request.Token) (*modelAuth.GetToken, error) {
	log := u.logger.WithCtx(ctx)
	modelToken, err := u.repo.NoSql.Auth.IsTokenExists(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("upss something wrong")
	}
	fmt.Println("isii",modelToken)
	acessToken, err := jwt.GenerateAccessToken(*modelToken, u.cfg.AccesTokenKey, u.cfg.TTlRedis.TTLAccessToken)
	if err != nil {
		log.Error("CreateNewAccessToken", logger.ToField("error_createToken_GenerateAccessToken", err))
		return nil, fmt.Errorf("error create token")
	}
	token := modelAuth.GetToken{
		AccessToken: acessToken,
	}

	return &token, nil
}

func (u *authUsecase) createLoginToken(ctx context.Context, modelToken modelJwt.TokenValue) (*modelAuth.GetToken, error) {

	log := u.logger.WithCtx(ctx)
	acessToken, err := jwt.GenerateAccessToken(modelToken, u.cfg.AccesTokenKey, u.cfg.TTlRedis.TTLAccessToken)
	if err != nil {
		log.Error("LoginAuth", logger.ToField("error_createToken_GenerateAccessToken", err))
		return nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken()
	if err != nil {
		log.Error("error_CreateUser", logger.ToField("error_createToken_GenearaterefreshToken", err))
		return nil, fmt.Errorf("Upss somethingwrong")
	}

	err = u.repo.NoSql.Auth.UpdateLogin(ctx, modelToken, refreshToken, configs.LoadConfig().TTlRedis.TTLRefreshToken)
	if err != nil {
		log.Error("error_CreateUser", logger.ToField("error_createToken_UpdateLogin", err))
		return nil, err
	}
	return &modelAuth.GetToken{
		AccessToken:  acessToken,
		RefreshToken: refreshToken,
	}, nil
}
