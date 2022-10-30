package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-file/api/internal/logic/file"
	"github.com/suyuan32/simple-admin-file/api/internal/svc"
	"github.com/suyuan32/simple-admin-file/api/internal/types"
)

// swagger:route post /file/list file FileList
//
// Get file list | 获取文件列表
//
// Get file list | 获取文件列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: FileListReq
//
// Responses:
//  200: FileListResp
//  401: SimpleMsg
//  500: SimpleMsg

func FileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := file.NewFileListLogic(r.Context(), svcCtx)
		resp, err := l.FileList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
