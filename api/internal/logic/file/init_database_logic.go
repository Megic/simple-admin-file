package file

import (
	"context"
	"net/http"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/suyuan32/simple-admin-core/pkg/enum"
	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/zeromicro/go-zero/core/errorx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-file/api/internal/svc"
	"github.com/suyuan32/simple-admin-file/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewInitDatabaseLogic(r *http.Request, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	err = l.initApi()
	if err != nil {
		if status.Code(err) == codes.InvalidArgument {
			return nil, errorx.NewCodeError(enum.InvalidArgument,
				l.svcCtx.Trans.Trans(l.lang, "init.alreadyInit"))
		}
		return nil, err
	}

	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewCodeError(enum.Internal, err.Error())
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.lang, i18n.Success),
	}, nil
}

func (l *InitDatabaseLogic) initApi() error {
	// create API in core service
	_, err := l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/upload",
		Description: "apiDesc.uploadFile",
		ApiGroup:    "file",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/file/list",
		Description: "apiDesc.fileList",
		ApiGroup:    "file",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/file",
		Description: "apiDesc.updateFileInfo",
		ApiGroup:    "file",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/file/status",
		Description: "apiDesc.setPublicStatus",
		ApiGroup:    "file",
		Method:      "POST",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/file",
		Description: "apiDesc.deleteFile",
		ApiGroup:    "file",
		Method:      "DELETE",
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		Path:        "/file/download/:id",
		Description: "apiDesc.downloadFile",
		ApiGroup:    "file",
		Method:      "GET",
	})

	if err != nil {
		return err
	}

	return nil
}
