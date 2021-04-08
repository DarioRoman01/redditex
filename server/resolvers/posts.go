package resolvers

import (
	"context"
	"fmt"
	"lireddit/db"
	"lireddit/models"
	"lireddit/utils"
)

var postTable db.PostTable

func init() {
	psql, _ := db.Connect()
	postTable = db.PostTable{
		Table: *psql,
	}
}

// resolve textSnippet field in post resolver return only 50 characters
func (p *PostResolver) TextSnippet(ctx context.Context, obj *models.Post) (string, error) {
	if len(obj.Text) > 50 {
		strings := utils.SplitSubN(obj.Text, 50)
		return strings[0], nil
	}

	return obj.Text, nil
}

// Handle post creation and validate user is logged in
func (m *mutationResolver) CreatePost(ctx context.Context, options models.PostInput) (*models.Post, error) {
	userId, err := utils.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}

	post := postTable.Postcreation(models.Post{
		Title:     options.Title,
		Text:      options.Text,
		CreatorId: userId,
	})

	if post == nil {
		return nil, fmt.Errorf("unable to create post")
	}

	return post, nil
}

// Handle post delete request and validate user is logged in
func (m *mutationResolver) DeletePost(ctx context.Context, id int) (bool, error) {
	userId, err := utils.GetUserSession(ctx)
	if err != nil {
		return false, err
	}

	if err := postTable.PostDelete(id, userId); err != nil {
		return false, err
	}

	return true, nil
}

// handle post update by id and validate user is logged in
func (m *mutationResolver) UpdatePost(ctx context.Context, id int, options models.PostInput) (*models.Post, error) {
	userId, err := utils.GetUserSession(ctx)
	if err != nil {
		return nil, err
	}

	post, err := postTable.PostUpdate(id, userId, options)
	if err != nil {
		return nil, err
	}

	return post, nil

}

// return post by id
func (q *queryResolver) Post(ctx context.Context, id int) (*models.Post, error) {
	post := postTable.GetPostById(id)

	if post == nil {
		return nil, nil
	}

	return post, nil
}

// get all posts
func (q *queryResolver) Posts(ctx context.Context, limit int, cursor *string) (*models.PaginatedPosts, error) {
	userId, _ := utils.GetUserSession(ctx)
	posts, isMore := postTable.GetAllPost(limit, userId, cursor)
	return &models.PaginatedPosts{Posts: posts, HasMore: isMore}, nil
}

func (p *mutationResolver) Vote(ctx context.Context, postID int, value int) (bool, error) {
	userId, err := utils.GetUserSession(ctx)
	if err != nil {
		return false, err
	}

	voted := postTable.SetVote(postID, userId, value)
	if !voted {
		return false, nil
	}

	return true, nil
}
