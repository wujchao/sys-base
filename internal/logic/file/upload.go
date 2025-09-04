package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"os"
	"path"
	"strings"
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
	uploadPath := g.Cfg().MustGet(ctx, "server.upload.path", "./public/upload").String()
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

func (s *sFile) DownloadImage(ctx context.Context, url string) (out *model.FileUploadOutput, err error) {
	resp, err := g.Client().Get(ctx, url)
	if err != nil {
		return nil, err
	}
	defer resp.Close()

	if resp.StatusCode != 200 {
		return nil, gerror.NewCode(consts.SystemErrCode, "下载文件失败")
	}
	uploadPath := g.Cfg().MustGet(ctx, "server.upload.path").String()
	if uploadPath == "" {
		return nil, gerror.NewCode(consts.UnknownErrCode, g.I18n().T(ctx, "上传文件路径配置不存在"))
	}
	basePath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	uploadPath = path.Join(basePath, uploadPath)

	parts := strings.Split(url, ".")
	ext := ""
	if len(parts) > 1 {
		ext = "." + parts[len(parts)-1]
	}
	fileName := guid.S() + ext

	fullPath := gfile.Join(uploadPath, fileName)
	// 保存文件
	err = gfile.PutBytes(fullPath, resp.ReadAll())
	if err != nil {
		return nil, err
	}
	out = &model.FileUploadOutput{
		Name: fileName,
		Path: fullPath, // upload/20230302/cqfgfefdf.png
		Url:  "/upload/" + fileName,
	}
	return
}
