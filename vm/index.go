package vm

import "blog/model"

type IndexViewModel struct{
	BaseViewModel
	// model.User
	Posts []model.Post
    Flash string

    BasePageViewModel
}

type IndexViewModelOp struct{}

// GetVM func
func (IndexViewModelOp) GetVM(username string, flash string, page, limit int) IndexViewModel {
    // mock data by myself:
    // u1 := model.User{Username: "bonfy"}
    // u2 := model.User{Username: "rene"}

    // posts := []model.Post{
    //     model.Post{User: u1, Body: "Beautiful day in Portland!"},
    //     model.Post{User: u2, Body: "The Avengers movie was so cool!"},
    // }

    // v := IndexViewModel{BaseViewModel{Title: "Homepage"}, u1, posts}
    // return v

    // Get data from database:
    u, _ := model.GetUserByUsername(username)
    // posts, _ := model.GetPostsByUserID(u1.ID)
    // posts, _ := u.FollowingPosts()
    // v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts, flash}
    posts, total, _ := u.FollowingPostsByPageAndLimit(page, limit)
    v := IndexViewModel{}
    v.SetTitle("Homepage")
    v.SetCurrentUser(username)
    v.Posts = *posts
    v.Flash = flash
    v.SetBasePageViewModel(total, page, limit)
    return v
}

// CreatePost func
func CreatePost(username, post string) error {
    u, _ := model.GetUserByUsername(username)
    return u.CreatePost(post)
}