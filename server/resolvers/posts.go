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

// Handle post creation
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

// Handle post delete request
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

// handle post update by id
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
		return nil, fmt.Errorf("post not found")
	}

	return post, nil
}

// get all posts
func (q *queryResolver) Posts(ctx context.Context) ([]models.Post, error) {
	posts := postTable.GetAllPost()
	return posts, nil
}
