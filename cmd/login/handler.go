package main

import (
	"context"
	login "github.com/Tiktok-Lite/kotkit/kitex_gen/login"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/internal/model"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/tools"
	"github.com/Tiktok-Lite/kotkit/pkg/helper/constant"
)


type LoginServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *LoginServiceImpl) Register(ctx context.Context, req *login.UserRegisterRequest) (resp *login.UserRegisterResponse, err error) {
	repo := repository.NewRepository(repository.DB)
	loginRepo := repository.NewLoginRepository(repo)
	userRepo := repository.NewUserRepository(repo)
	
	// 检查用户名是否冲突
	usr, err := userRepo.QueryUserByName(req.Username)
	if err != nil {
		// TODO: 添加日志
		res := &login.UserRegisterResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "注册失败：服务器查询错误",
		}
		return res, nil
	} else if usr != nil {
		// TODO: 添加日志
		res := &login.UserRegisterResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "该用户名已存在，请更换",
		}
		return res, nil
	}

	// 创建login,user
	newUser := &model.User{
		Name:            req.Username,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}
	
	if err := userRepo.Create(newUser); err != nil {
		// TODO: 添加日志
		// 日志：TODO
		res := &login.UserRegisterResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "注册失败: 服务器新建user错误",
		}
		return res, nil
	}
	
	loginData := &model.Login{
		Username: req.Username,
		Password: tools.Md5Encrypt(req.Password),
		UserID:   newUser.ID,
	}
	
	if err := loginRepo.CreateLogin(loginData); err != nil {
		// TODO: 添加日志
		// 日志：TODO
		res := &login.UserRegisterResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "注册失败: 服务器新建login错误",
		}
		return res, nil
	}

	newUser = &model.User{
		Name:            req.Username,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
		UserLogin:       *loginData,
	}

	if err := userRepo.UpdateByUsername(req.Username,newUser); err != nil {
		// TODO: 添加日志
		// 日志：TODO
		res := &login.UserRegisterResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "注册失败: 服务器更新user错误",
		}
		return res, nil
	}


	//生成token TODO
	
	userLogin,err := loginRepo.QueryLoginByName(req.Username)
	if err != nil {
		// 日志：TODO
		res := &login.UserRegisterResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "注册失败：服务器错误",
		}
		return res, nil
	}
	res := &login.UserRegisterResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "success",
		UserId:     int64(userLogin.UserID),
		Token:      "todo",
	}
	return res, nil
}

// Login implements the UserServiceImpl interface.
func (s *LoginServiceImpl) Login(ctx context.Context, req *login.UserLoginRequest) (resp *login.UserLoginResponse, err error) {
	repo := repository.NewRepository(repository.DB)
	loginRepo := repository.NewLoginRepository(repo)
	// 根据用户名获取密码
	userLogin,err := loginRepo.QueryLoginByName(req.Username)
	if err != nil {
		// 日志：TODO
		res := &login.UserLoginResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "登录失败：服务器内部错误",
		}
		return res, nil
	} else if userLogin == nil {
		res := &login.UserLoginResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "用户名不存在",
		}
		return res, nil
	}

	
	// 比较数据库中的密码和请求的密码
	if tools.Md5Encrypt(req.Password) != userLogin.Password {
		// TODO: 添加日志
		res := &login.UserLoginResponse{
			StatusCode: constant.StatusErrorCode,
			StatusMsg:  "用户名或密码错误",
		}
		return res, nil
	}

	// 密码认证通过,获取用户id并生成token 
	// TODO
	

	// 返回结果
	res := &login.UserLoginResponse{
		StatusCode: constant.StatusOKCode,
		StatusMsg:  "success",
		UserId:     int64(userLogin.UserID),
		Token:      "todo",
	}
	return res, nil
}
