package oss

import (
	"bytes"
	"context"
	"errors"
	"io"
	"math/rand"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
)

type UploadOption struct {
	Type  string
	Value interface{}
}

func UploadWithPrivate() UploadOption {
	return UploadOption{
		Type:  "policy",
		Value: "private",
	}
}

func UploadWithContentType(contentType string) UploadOption {
	return UploadOption{
		Type:  "content-type",
		Value: contentType,
	}
}

func UploadWithFileName(fileName string) UploadOption {
	return UploadOption{
		Type: "file-name",
		Value: "upload/" +
			time.Now().Format("2006/01/02") +
			"/" + strconv.Itoa(int(time.Now().UnixMilli())) +
			"/" + strconv.Itoa(rand.Intn(65536)) +
			"/" + fileName,
	}
}

func UploadGeneric(file io.Reader, fileSize int64, options ...UploadOption) (
	string, error,
) {
	fileName := ""
	_options := minio.PutObjectOptions{}
	for _, option := range options {
		switch option.Type {
		case "policy":
			continue
		case "content-type":
			_options.ContentType = option.Value.(string)
		case "file-name":
			fileName = option.Value.(string)
		}
	}

	if fileName == "" {
		return "", errors.New("file name not specified")
	}

	_, err := client.PutObject(context.Background(), bucket, fileName, file, fileSize, _options)
	if err != nil {
		return "", err
	}

	return bucket + "/" + fileName, nil
}

func UploadGenericFromPath(filePath string, options ...UploadOption) (
	string, error,
) {
	fileName := ""
	_options := minio.PutObjectOptions{}
	for _, option := range options {
		switch option.Type {
		case "policy":
			continue
		case "content-type":
			_options.ContentType = option.Value.(string)
		case "file-name":
			fileName = option.Value.(string)
		}
	}

	if fileName == "" {
		return "", errors.New("file name not specified")
	}

	_, err := client.FPutObject(context.Background(), bucket, fileName, filePath, _options)
	if err != nil {
		return "", err
	}

	return bucket + "/" + fileName, nil
}

func UploadGenericFromContent(content []byte, options ...UploadOption) (
	string, error,
) {
	fileName := ""
	_options := minio.PutObjectOptions{}
	for _, option := range options {
		switch option.Type {
		case "policy":
			continue
		case "content-type":
			_options.ContentType = option.Value.(string)
		case "file-name":
			fileName = option.Value.(string)
		}
	}

	if fileName == "" {
		return "", errors.New("file name not specified")
	}

	_, err := client.PutObject(context.Background(), bucket, fileName,
		bytes.NewReader(content), int64(len(content)), _options)
	if err != nil {
		return "", err
	}

	return bucket + "/" + fileName, nil
}
