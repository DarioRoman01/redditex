package resolvers

import (
	"context"
	"fmt"
	"lireddit/db"
	"lireddit/models"
)

var postTable db.PostTable

func init() {
	psql, _ := db.Connect()
	postTable = db.PostTable{
		Table: *psql,
	}
}

func (m *mutationResolver) CreatePost(ctx context.Context, title string) (*models.Post, error) {
	post := postTable.Postcreation(models.Post{Title: title})
	if post == nil {
		return nil, fmt.Errorf("unable to create post")
	}

	return post, nil
}

func (m *mutationResolver) DeletePost(ctx context.Context, id int) (bool, error) {
	isDeleted := postTable.PostDelete(id)

	if !isDeleted {
		return false, fmt.Errorf("post not found")
	}

	return true, nil
}

func (m *mutationResolver) UpdatePost(ctx context.Context, id int, title string) (*models.Post, error) {
	post := postTable.PostUpdate(id, title)

	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	return post, nil

}

func (q *queryResolver) Post(ctx context.Context, id int) (*models.Post, error) {
	post := postTable.GetPostById(id)

	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	return post, nil
}

func (q *queryResolver) Posts(ctx context.Context) ([]models.Post, error) {
	posts := postTable.GetAllPost()
	return posts, nil
}
