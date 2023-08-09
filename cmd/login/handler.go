package main

import (
	"context"
	login "github.com/Tiktok-Lite/kotkit/kitex_gen/login"
	"github.com/Tiktok-Lite/kotkit/internal/repository"
	"github.com/Tiktok-Lite/kotkit/pkg/zap"
	"github.com/Tiktok-Lite/kotkit/internal/tools"
	"github.com/Tiktok-Lite/kotkit/internal/model"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *login.UserRegisterRequest) (resp *login.UserRegisterResponse, err error) {
	// TODO: Your code here...
	logger := zap.InitLogger()

	// 检查用户名是否冲突
	usr, err := repository.QueryUserByName(req.Username)
	if err != nil {
		logger.Errorln(err.Error())
		res := &login.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, nil
	} else if usr != nil {
		logger.Errorf("该用户名已存在：%s", usr.UserName)
		res := &login.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "该用户名已存在，请更换",
		}
		return res, nil
	}

	// 创建login,user
	loginData := &model.Login{
		UserName: req.Username,
		Password: tools.Md5Encrypt(req.Password),
	}
	if err := repository.CreateLogin(loginData); err != nil {
		logger.Errorln(err.Error())
		res := &login.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, nil
	}

	newUser := &model.User{
		Name: req.Username,
		UserLogin: *loginData
	}
	if err := repository.Create(newUser); err != nil {
		logger.Errorln(err.Error())
		res := &login.UserRegisterResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败：服务器内部错误",
		}
		return res, nil
	}

	//生成token 
	//TODO
	

	res := &login.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(newUser.ID),
		Token:      "todo",
	}
	return res, nil

}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *login.UserLoginRequest) (resp *login.UserLoginResponse, err error) {
	logger := zap.InitLogger()

	// 根据用户名获取密码
	usr, err := repository.QueryUserByName(req.UserName)
	if err != nil {
		logger.Errorln(err.Error())
		res := &login.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "登录失败：服务器内部错误",
		}
		return res, nil
	} else if usr == nil {
		res := &login.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名不存在",
		}
		return res, nil
	}

	// 比较数据库中的密码和请求的密码
	if tool.Md5Encrypt(req.Password) != usr.Password {
		logger.Errorln("用户名或密码错误")
		res := &login.UserLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误",
		}
		return res, nil
	}

	// 密码认证通过,获取用户id并生成token 
	// TODO
	

	// 返回结果
	res := &login.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(usr.ID),
		Token:      "todo",
	}
	return res, nil
}
