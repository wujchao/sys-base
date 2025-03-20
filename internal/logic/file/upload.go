package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"os"
	"path"
	"sys-base/consts"
	"sys-base/internal/model"
	"sys-base/service"
)

type sFile struct {
}

func init() {
	service.RegisterFile(&sFile{})
}

// Upload 上传文件
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (*model.FileUploadOutput, error) {
	uploadPath := g.Cfg().MustGet(ctx, "server.upload.path").String()
	if uploadPath == "" {
		return nil, gerror.NewCode(consts.UnknownErrCode, g.I18n().T(ctx, "上传文件路径配置不存在"))
	}
	basePath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	uploadPath = path.Join(basePath, uploadPath)
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	dateDirName := gtime.Now().Format("Ymd") //年月日
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}
	return &model.FileUploadOutput{
		Name: fileName,
		Path: gfile.Join(uploadPath, dateDirName, fileName), // upload/20230302/cqfgfefdf.png
		Url:  "/upload/" + dateDirName + "/" + fileName,
	}, nil
}
