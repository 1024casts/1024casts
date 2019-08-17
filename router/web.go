package router

import (
	"html/template"

	"github.com/1024casts/1024casts/handler/qiniu"

	"github.com/1024casts/1024casts/handler/web/github"

	_ "github.com/1024casts/1024casts/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/1024casts/1024casts/router/middleware"

	webComment "github.com/1024casts/1024casts/handler/web/comment"
	webCourse "github.com/1024casts/1024casts/handler/web/course"
	webPlan "github.com/1024casts/1024casts/handler/web/plan"
	webTopic "github.com/1024casts/1024casts/handler/web/topic"
	webUser "github.com/1024casts/1024casts/handler/web/user"
	webVideo "github.com/1024casts/1024casts/handler/web/video"
	"github.com/1024casts/1024casts/handler/web/wiki"

	"github.com/1024casts/1024casts/handler/web"

	"time"

	"github.com/1024casts/1024casts/handler/web/notification"
	"github.com/1024casts/1024casts/pkg/flash"
	"github.com/foolin/gin-template"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

// Load loads the middlewares, routes, handlers.
func LoadWebRouter(g *gin.Engine) *gin.Engine {
	router := g

	// Middlewares.

	// 404 Handler.
	router.NoRoute(func(c *gin.Context) {
		web.Error404(c)
	})

	router.Use(static.Serve("/static", static.LocalFile(viper.GetString("static"), false)))
	router.Use(static.Serve("/uploads/avatar", static.LocalFile(viper.GetString("avatar"), false)))
	router.Use(static.Serve("/uploads/images", static.LocalFile(viper.GetString("images"), false)))

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "templates",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{},
		Funcs: template.FuncMap{
			// 判断是否是当前链接
			"isActive": func(ctx *gin.Context, currentUri string) string {
				if ctx.Request.RequestURI == currentUri {
					return "is-active"
				} else {
					return ""
				}
			},
			// 全局消息
			"flashMessage": func(ctx *gin.Context) string {
				errorMessage, err := flash.GetMessage(ctx.Writer, ctx.Request)
				if err != nil {
					log.Warnf("[router] get flash message err: %v", err)
					return ""
				}
				return string(errorMessage)
			},
			"hasFlash": func(ctx *gin.Context) bool {
				return flash.HasFlash(ctx.Request)
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	})

	router.GET("/", web.Index)

	// auth route
	router.GET("/login", webUser.GetLogin)
	router.POST("/login", webUser.DoLogin)
	router.GET("/logout", webUser.Logout)

	// github login
	router.GET("/login/oauth/github", github.RedirectHandler)
	router.GET("/login/oauth/github/callback", github.CallbackHandler)

	// register route
	router.GET("/register", webUser.GetRegister)
	router.POST("/register", webUser.DoRegister)
	// reset route
	router.GET("/password/reset", webUser.ShowLinkRequestForm)  // 忘记密码	输入要重置密码的form
	router.POST("/password/email", webUser.SendResetLinkEmail)  // 发送重置密码邮件
	router.GET("/password/reset/:token", webUser.ShowResetForm) // 重置密码form
	router.POST("/password/reset", webUser.ResetPassword)       // 重置密码
	// user profile route
	router.GET("/users/:username", webUser.Index)                        // 个人首页
	router.GET("/users/:username/activation/:token", webUser.ActiveUser) // 通过发送到邮箱中的链接激活
	router.GET("/users/:username/topics", webUser.TopicList)             // 发表过的主题
	router.GET("/users/:username/replies", webUser.ReplyList)            // 回复过的
	router.GET("/users/:username/favorites", webUser.TopicList)          // 收藏过的
	router.GET("/users/:username/following", webUser.Following)          // 正在关注的人
	router.GET("/users/:username/followers", webUser.Follower)           // 关注者(粉丝)

	settings := router.Group("/settings")
	settings.Use(middleware.CookieMiddleware())
	{
		settings.GET("/basic", webUser.Basic)
		settings.POST("/basic", webUser.DoBasic)
		settings.GET("/profile", webUser.Profile)
		settings.POST("/profile", webUser.DoProfile)
		settings.GET("/password", webUser.Password)
		settings.POST("/password", webUser.DoPassword)
	}

	notice := router.Group("/notifications")
	notice.Use(middleware.CookieMiddleware())
	{
		notice.GET("", notification.List)
	}

	u := router.Group("/users")
	u.Use(middleware.CookieMiddleware())
	{
		u.GET("/:username/orders", webUser.OrderList)
	}

	// courses
	router.GET("/courses", webCourse.Index)
	router.GET("/courses/:slug", webCourse.Detail)
	router.GET("/courses/:slug/episodes/:episode_id", webVideo.Detail)

	// topics
	router.GET("/topics", webTopic.Index)
	router.GET("/topics/:id", webTopic.Detail)
	ts := router.Group("/topics")
	ts.Use(middleware.CookieMiddleware())
	{
		ts.POST("/reply", webTopic.Reply)
		ts.POST("/like/:reply_id", webTopic.Like)
	}
	t := router.Group("/topic")
	t.Use(middleware.CookieMiddleware())
	{
		t.GET("/new", webTopic.Create)
		t.POST("/new", webTopic.DoCreate)
		t.GET("/edit/:id", webTopic.Edit)
		t.POST("/edit/:id", webTopic.DoEdit)
	}

	// comment
	cmt := router.Group("/comments")
	cmt.Use(middleware.CookieMiddleware())
	{
		cmt.POST("", webComment.Create)
		cmt.POST("/:comment_id/like", webComment.Like)
	}

	router.GET("/vip", webPlan.Index)
	p := router.Group("/plans")
	p.Use(middleware.CookieMiddleware())
	{
		p.GET("/:alias/purchase", webPlan.Purchase)
		p.GET("/:alias/pay", webPlan.Pay)
		p.GET("/:alias/success", webPlan.Success)
	}

	pay := router.Group("/pay")
	pay.Use(middleware.CookieMiddleware())
	{
		pay.GET("/:order_id/check", webPlan.Check)
	}
	// wiki
	router.GET("/wiki", wiki.Index)
	router.GET("/wiki/:slug", wiki.Detail)
	//router.GET("/wiki/:slug/comments")

	// web upload
	image := router.Group("/image")
	image.Use(middleware.CookieMiddleware())
	{
		image.POST("/upload", qiniu.WebUpload)
	}

	//router.GET("/test", wiki.Test)

	return router
}
