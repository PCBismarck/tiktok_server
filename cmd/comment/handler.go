package main

import (
	"context"
	"github.com/PCBismarck/tiktok_server/cmd/comment/consts"
	dbconfig "github.com/PCBismarck/tiktok_server/cmd/comment/db_config"
	comment "github.com/PCBismarck/tiktok_server/cmd/comment/kitex_gen/comment"
	"time"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

var success string
var failed string
var commentInfo *comment.CommentInfo

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	// TODO: Your code here...
	userId := req.UserId
	videoId := req.VideoId
	actionType := req.ActionType
	commentText := req.CommentText
	commentId := req.CommentId

	resp = new(comment.CommentActionResponse)

	switch actionType {
	case 1:
		success = "Create comment success"
		failed = "Create comment failed"
		commentid, err := dbconfig.CreateComment(videoId, userId, *commentText)
		if err != nil {
			resp.Base = &comment.BaseResp{StatusCode: consts.FailureCode, StatusMsg: &failed}
			resp.Comment = commentInfo
			return resp, nil
		}
		usr := dbconfig.GetUserInfoByUid(userId)
		resp.Base = &comment.BaseResp{StatusCode: consts.SuccessCode, StatusMsg: &success}
		resp.Comment = &comment.CommentInfo{
			CommentId:  commentid,
			User:       usr,
			Content:    *commentText,
			CreateDate: time.Now().Format("01-02"),
		}
		return resp, nil
	case 2:
		success = "Delete comment success"
		failed = "Delete comment failed"
		ok := dbconfig.DeleteComment(videoId, *commentId)
		if ok {
			resp.Base = &comment.BaseResp{StatusCode: consts.SuccessCode, StatusMsg: &success}
			resp.Comment = commentInfo
		} else {
			resp.Base = &comment.BaseResp{StatusCode: consts.FailureCode, StatusMsg: &failed}
			resp.Comment = commentInfo
		}
		return resp, nil
	}
	failed = "ActionType is illegal"
	resp.Base = &comment.BaseResp{StatusCode: consts.FailureCode, StatusMsg: &failed}
	resp.Comment = commentInfo
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	videoId := req.VideoId

	resp = new(comment.CommentListResponse)

	success = "Get comment list success"
	failed = "Get comment list failed"

	comments, err := dbconfig.GetCommentsByVid(videoId)

	var commentList []*comment.CommentInfo
	if err != nil {
		resp.Base = &comment.BaseResp{StatusCode: consts.FailureCode, StatusMsg: &failed}
		resp.CommentList = commentList
		return resp, nil
	}

	for _, v := range comments {
		newComment := &comment.CommentInfo{
			CommentId:  int64(v.ID),
			User:       dbconfig.GetUserInfoByUid(v.UserId),
			Content:    v.CommentText,
			CreateDate: v.CreatedAt.Format("01-02"),
		}
		commentList = append(commentList, newComment)
	}
	resp.Base = &comment.BaseResp{StatusCode: consts.SuccessCode, StatusMsg: &success}
	resp.CommentList = commentList
	return resp, nil
}
