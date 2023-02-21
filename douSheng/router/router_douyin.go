package router

import (
	"douSheng/controller/user_info"
	"douSheng/controller/user_login"
	"douSheng/controller/video"
	"douSheng/models"
	"github.com/gin-gonic/gin"
)

func InitDouyinRouter() *gin.Engine {
	//初始化数据库表
	models.InitDB()

	// 获取Engine
	r := gin.Default()

	//设置静态文件夹 存储视频图片
	r.Static("static", "./static")

	//设置统一入口
	baseGroup := r.Group("/douyin")
	//根据灵活性考虑是否加入JWT中间件来进行鉴权，还是在之后再做鉴权
	// basic apis 基础api
	//视频推荐
	//baseGroup.GET("/feed/", video.FeedVideoListHandler)
	//用户信息
	baseGroup.GET("/user/", user_info.UserInfoHandler)
	//登录
	baseGroup.POST("/user/login/", user_login.UserLoginHandler)
	//注册
	baseGroup.POST("/user/register/", user_login.UserRegisterHandler)
	//发布视频
	baseGroup.POST("/publish/action/", video.PublishVideoHandler)
	////视频列表
	baseGroup.GET("/publish/list/", video.QueryVideoListHandler)

	return r
}
