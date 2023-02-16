package conf

import (
	"os"

	routine "github.com/0xunion/exercise_back/src/routine"
	beego "github.com/beego/beego/v2/server/web"
)

func getConf(name string) string {
	v, err := beego.AppConfig.String(name)
	if err != nil {
		// panic will cause the program exit
		routine.Panic("get config failed, name: %s, err: %v", name, err)
	}

	return v
}

func getWorkingPath() string {
	path, err := os.Getwd()
	if err != nil {
		routine.Panic("get working path failed, err: %v", err)
	}

	return path
}

var (
	// AppName is the name of the application.
	appName = getConf("appname")
	// HTTPPort is the port of the HTTP server.
	httpPort = getConf("httpport")
	// RunMode is the running mode of the application.
	runMode = getConf("runmode")

	// mysql port
	mysqlPort = getConf("mysqlport")
	// mysql user
	mysqlUser = getConf("mysqluser")
	// mysql password
	mysqlPass = getConf("mysqlpass")
	// mysql database
	mysqlName = getConf("mysqlname")
	// mysql host
	mysqlHost = getConf("mysqlhost")
	// mysql charset
	mysqlCharset = getConf("mysqlcharset")
	// redis port
	redisPort = getConf("redisport")
	// redis host
	redisHost = getConf("redishost")
	// redis password
	redisPass = getConf("redispass")
	// redis expire time
	redisExpire = getConf("redisexpire")
	// mongodb host
	mongodbHost = getConf("mongodbhost")
	// mongodb port
	mongodbPort = getConf("mongodbport")
	// mongodb user
	mongodbUser = getConf("mongodbuser")
	// mongodb password
	mongodbPass = getConf("mongodbpass")
	// mongodb database
	mongodbName = getConf("mongodbname")
	// working directory
	workDir = getWorkingPath()
	// temprary directory
	temppath = getConf("temppath")

	// oss address
	ossAddr = getConf("unionossaddress")
	// oss access key
	ossAccessKey = getConf("unionossaccesskey")
	// oss secret key
	ossSecretKey = getConf("unionosssecretkey")
	// oss bucket
	ossBucket = getConf("unionossbucket")
	// seucre_token_padding
	secureTokenPadding = getConf("securetokenpadding")
	// secure_aes_key
	secureAesKey = getConf("secureaeskey")
	// secure_aes_iv
	secureAesIv = getConf("secureaesiv")
	// cors
	corsOrigin      = getConf("corsorigin")
	corsHeaders     = getConf("corsheaders")
	corsMethods     = getConf("corsmethods")
	corsCredentials = getConf("corscredentials")
	corsAge         = getConf("corsage")
)

func IsDev() bool {
	return runMode == "dev"
}

func MySQLHost() string {
	return mysqlHost
}

func MySQLPort() string {
	return mysqlPort
}

func MySQLUser() string {
	return mysqlUser
}

func MySQLPass() string {
	return mysqlPass
}

func MySQLName() string {
	return mysqlName
}

func MySQLCharset() string {
	return mysqlCharset
}

func RedisHost() string {
	return redisHost
}

func RedisPort() string {
	return redisPort
}

func RedisPass() string {
	return redisPass
}

func RedisExpire() string {
	return redisExpire
}

func MongoDBHost() string {
	return mongodbHost
}

func MongoDBPort() string {
	return mongodbPort
}

func MongoDBUser() string {
	return mongodbUser
}

func MongoDBPass() string {
	return mongodbPass
}

func MongoDBName() string {
	return mongodbName
}

func WorkDir() string {
	return workDir
}

func OSSAddr() string {
	return ossAddr
}

func OSSAccessKey() string {
	return ossAccessKey
}

func OSSSecretKey() string {
	return ossSecretKey
}

func OSSBucket() string {
	return ossBucket
}

func TempPath() string {
	return temppath
}

func SecureTokenPadding() string {
	return secureTokenPadding
}

func SecureAesKey() string {
	return secureAesKey
}

func SecureAesIv() string {
	return secureAesIv
}

func CorsOrigin() string {
	return corsOrigin
}

func CorsHeaders() string {
	return corsHeaders
}

func CorsMethods() string {
	return corsMethods
}

func CorsCredentials() string {
	return corsCredentials
}

func CorsAge() string {
	return corsAge
}
