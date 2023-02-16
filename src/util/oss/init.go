package oss

import (
	"math/rand"
	"time"

	"github.com/0xunion/exercise_back/src/const/conf"
	"github.com/0xunion/exercise_back/src/routine"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	endpoint  = conf.OSSAddr()
	accessKey = conf.OSSAccessKey()
	secretKey = conf.OSSSecretKey()
	bucket    = conf.OSSBucket()
	useSSL    = false

	client *minio.Client
)

func init() {
	rand.Seed(time.Now().Unix())

	routine.Info("[OSS] Connecting to OSS service...")
	var err error
	client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		routine.Panic("[OSS] Failed to connect to OSS service" + err.Error())
	}
	routine.Info("[OSS] Connected to OSS service")
}
