package consts

const (
	AccountCateAdmin = 1    //管理员
	AccountCateDirector = 2 //园长
)


const (
	JWTSecretKey = "boqujiaoyu2024!!!"
)
var CodeMsg = map[int]string {
	0 : "成功",
	-1 : "账号不存在，请找园长注册账号",
	-2 : "密码错误",
    -3 : "数据更新异常",
    -4 : "Authorization needed",
    -6 : "参数错误",
	-20: "服务器内部错误",
}

var TeacherRole = map[uint]string {
	1 : "园长",
	2 : "教师",
}

const (
	TimeFormatLayout = "2006-01-02 15:04:05"
)

const (
	EnvPrd = "prd"
	EnvTest = "test"
)