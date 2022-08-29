package file

import (
	"net/http"

	"github.com/suyuan32/simple-admin-file/api/internal/logic/file"
	"github.com/suyuan32/simple-admin-file/api/internal/svc"
	"github.com/suyuan32/simple-admin-file/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route POST /file/list file fileList
// Get file list | 获取文件列表
// Responses:
//   200: FileListResp
//   401: SimpleMsg
//   500: SimpleMsg

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
