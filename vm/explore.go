package vm

import (
	"blog/model"
)

// ExploreViewModel struct
type ExploreViewModel struct {
	BaseViewModel
	Posts []model.Post
	BasePageViewModel
}

// ExploreViewModelOp struct
type ExploreViewModelOP struct {}

// GetVM
func (ExploreViewModelOP) GetVM(username string, page, limit int) ExploreViewModel {
	posts, total, _ := model.GetPostsByPageAndLimit(page, limit)
	v := ExploreViewModel{}
	v.SetTitle("搜索")
	v.Posts = *posts
	v.SetBasePageViewModel(total, page, limit)
	v.SetCurrentUser(username)
	return v
}