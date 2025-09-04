// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/wujchao/sys-base/internal/model"
)

type (
	IFile interface {
		// Upload 上传文件
		Upload(ctx context.Context, in model.FileUploadInput) (*model.FileUploadOutput, error)
		DownloadImage(ctx context.Context, url string) (out *model.FileUploadOutput, err error)
	}
)

var (
	localFile IFile
)

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
