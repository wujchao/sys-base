// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sys_upload

import (
	"context"

	"github.com/wujchao/sys-base/api/sys_upload/v1"
)

type ISysUploadV1 interface {
	FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error)
	FileUploadForEditor(ctx context.Context, req *v1.FileUploadForEditorReq) (res *v1.FileUploadForEditorRes, err error)
}
