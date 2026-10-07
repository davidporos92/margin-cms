package server

import (
	"context"

	"github.com/davidporos92/margin-cms/api/internal/api"
)

type server struct {
}

func (s server) ListActivity(ctx context.Context, request api.ListActivityRequestObject) (api.ListActivityResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) Login(ctx context.Context, request api.LoginRequestObject) (api.LoginResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) Logout(ctx context.Context, request api.LogoutRequestObject) (api.LogoutResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) GetMe(ctx context.Context, request api.GetMeRequestObject) (api.GetMeResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) RefreshTokens(ctx context.Context, request api.RefreshTokensRequestObject) (api.RefreshTokensResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) ListPosts(ctx context.Context, request api.ListPostsRequestObject) (api.ListPostsResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) CreatePost(ctx context.Context, request api.CreatePostRequestObject) (api.CreatePostResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) BulkPosts(ctx context.Context, request api.BulkPostsRequestObject) (api.BulkPostsResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) DeletePost(ctx context.Context, request api.DeletePostRequestObject) (api.DeletePostResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) GetPost(ctx context.Context, request api.GetPostRequestObject) (api.GetPostResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) UpdatePost(ctx context.Context, request api.UpdatePostRequestObject) (api.UpdatePostResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) ExportPost(ctx context.Context, request api.ExportPostRequestObject) (api.ExportPostResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) ListRevisions(ctx context.Context, request api.ListRevisionsRequestObject) (api.ListRevisionsResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) RestoreRevision(ctx context.Context, request api.RestoreRevisionRequestObject) (api.RestoreRevisionResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) GetStats(ctx context.Context, request api.GetStatsRequestObject) (api.GetStatsResponseObject, error) {
	return nil, ErrNotImplemented
}

func (s server) ListTags(ctx context.Context, request api.ListTagsRequestObject) (api.ListTagsResponseObject, error) {
	return nil, ErrNotImplemented
}

func New() api.StrictServerInterface {
	return server{}
}
