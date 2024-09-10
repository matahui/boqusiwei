package routers

import (
	"github.com/gin-gonic/gin"
	"homeschooledu/controllers"
	"homeschooledu/middlewares"
)

func SetupRouters(r *gin.Engine) {
	// 路由组：公共路由，无需认证
	public := r.Group("/api")
	{
		public.POST("/login", controllers.Login)
		public.POST("/microLogin", controllers.StudentLogin)
	}

	// 路由组: 按需要添加......
	school := r.Group("/api/school")
	school.Use(middlewares.JWTAuthMiddleware())
	{
		school.GET("/list", controllers.SchoolList)
		school.POST("/update", controllers.SchoolUpdate)
		school.POST("/delete", controllers.SchoolDelete)
		school.POST("/add", controllers.SchoolAdd)
		school.GET("/regionList", controllers.RegionList)
		school.POST("/regionAdd", controllers.RegionAdd)
		school.GET("/class", controllers.SchoolClass)
	}

	student := r.Group("/api/student")
	student.Use(middlewares.JWTAuthMiddleware())
	{
		student.GET("/list", controllers.StudentList)
		student.POST("/update", controllers.StudentUpdate)
		student.POST("/delete", controllers.StudentDelete)
		student.POST("/add", controllers.StudentAdd)
		student.POST("/batchAdd", controllers.StudentBatchAdd)
	}

	teacher := r.Group("/api/teacher")
	teacher.Use(middlewares.JWTAuthMiddleware())
	{
		teacher.GET("/list", controllers.TeacherList)
		teacher.POST("/update", controllers.TeacherUpdate)
		teacher.POST("/delete", controllers.TeacherDelete)
		teacher.POST("/add", controllers.TeacherAdd)
		teacher.POST("/batchAdd", controllers.TeacherBatchAdd)
	}

	class := r.Group("/api/class")
	class.Use(middlewares.JWTAuthMiddleware())
	{
		class.GET("/list", controllers.ClassList)
		class.POST("/update", controllers.ClassUpdate)
		class.POST("/delete", controllers.ClassDelete)
		class.POST("/add", controllers.ClassAdd)
		class.POST("/importStudent", controllers.StudentBatchAdd)
		class.POST("/bindTeacher", controllers.ClassBindTeacher)
		class.GET("/detail", controllers.ClassDetail)
	}


	resource := r.Group("/api/resource")
	resource.Use(middlewares.JWTAuthMiddleware())
	{
		resource.GET("/list", controllers.ResourceList)
		resource.POST("/delete", controllers.ResourceDelete)
		resource.POST("/batchAdd", controllers.ResourceBatchAdd)
		resource.GET("/cate", controllers.ResourceCate)
	}

	schedule := r.Group("/api/schedule")
	schedule.Use(middlewares.JWTAuthMiddleware())
	{
		schedule.GET("/list", controllers.ScheduleList)
		schedule.POST("/add", controllers.ScheduleAdd)
		schedule.GET("/detail", controllers.ScheduleDetail)
		schedule.POST("/update", controllers.ScheduleUpdate)
	}


	micro := r.Group("/api/micro")
	micro.Use(middlewares.JWTAuthMiddleware())
	{
		micro.GET("/home", controllers.MicroHome)
		micro.GET("/self", controllers.MicroSelf)
		micro.GET("/rank", controllers.MicroRank)
		micro.POST("/task", controllers.MicroTask)
	}
}
