package constant

const (
	DefaultDBConfigName       = "db"
	DefaultLogConfigName      = "log"
	DefaultAPIConfigName      = "api"
	DefaultUserConfigName     = "user"
	DefaultVideoConfigName    = "video"
	DefaultLoginConfigName    = "login"
	DefaultMinioConfigName    = "minio"
	DefaultRelationConfigName = "relation"
	DefaultFavoriteConfigName = "favorite"
)

const (
	StatusOKCode    = 0  // （这里默认）响应返回0即正确
	StatusErrorCode = -1 // 响应返回-1即错误
)

const (
	FavoriteCode   = 1
	UnFavoriteCode = 2
)
