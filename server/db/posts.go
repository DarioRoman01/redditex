package db

import (
	"fmt"
	"lireddit/models"

	"gorm.io/gorm"
)

type PostTable struct {
	Table gorm.DB
}

func (p *PostTable) Postcreation(post models.Post) *models.Post {
	result := p.Table.Create(&post)
	if result.RowsAffected == 0 || result.Error != nil {
		return nil
	}

	return &post
}

// delete post from the db and validate that the requesting user is owner of the post
func (p *PostTable) PostDelete(id int, userId int) error {
	result := p.Table.Where("id = ? and creator_id = ?", id, userId).Delete(&models.Post{})
	if result.RowsAffected == 0 || result.Error != nil {
		return fmt.Errorf("post does not exist or you are not the owner of the post")
	}

	return nil
}

// update post data and validate that requesting user is owner of the post
func (p *PostTable) PostUpdate(id int, userId int, options models.PostInput) (*models.Post, error) {
	var post models.Post

	p.Table.First(&post, id)
	if post.ID == 0 {
		return nil, fmt.Errorf("post not found")
	}

	if post.CreatorId != userId {
		return nil, fmt.Errorf("you do not have permissions tp perform this action")
	}

	p.Table.Model(&post).Updates(&models.Post{
		Title: options.Title,
		Text:  options.Text,
	})

	return &post, nil
}

func (p *PostTable) GetPostById(id int) *models.Post {
	var post models.Post
	p.Table.
		Table("posts").
		Where("id = ?", id).
		Preload("Creator").
		Find(&post)

	if post.ID == 0 {
		return nil
	}
	return &post
}

// get all posts and order the post by createdAt
// if a cursor is recibed only retrieves the posts created before that post
func (p *PostTable) GetAllPost(limit int, userId int, cursor *string) ([]models.Post, bool) {
	var posts []models.Post
	if limit > 50 {
		limit = 50
	}
	limit++

	if cursor != nil {
		p.Table.Raw(`
			SELECT p.*,
			( SELECT "value" from "updoots" 
			WHERE "user_id" = ? and "post_id" = p.id) as "StateValue"
			FROM posts p
			WHERE p.created_at < ?
			ORDER BY p.created_at DESC
			LIMIT ?
		`, userId, *cursor, limit).
			Preload("Creator").Find(&posts)

	} else {
		p.Table.Raw(`
			SELECT p.*,
			( SELECT "value" from "updoots" 
			WHERE "user_id" = ? and "post_id" = p.id) as "StateValue"
			FROM posts p
			ORDER BY p.created_at DESC
			LIMIT ?
		`, userId, limit).
			Preload("Creator").Find(&posts)
	}

	if len(posts) == 0 {
		return nil, false
	}

	if len(posts) == limit {
		return posts[0 : limit-1], true
	}

	return posts[0 : len(posts)-1], false
}

// set vote in updoots join table and update post points
func (p *PostTable) SetVote(postId, userId, value int) bool {
	var updoot models.Updoot
	isUpdoot := value != -1
	var realValue int

	if isUpdoot {
		realValue = 1
	} else {
		realValue = -1
	}

	p.Table.Table("updoots").Where("user_id = ? and post_id = ?", userId, postId).Find(&updoot)

	// user is vote the post before and
	// they are changing their vote
	if updoot.PostID != 0 && updoot.Value != realValue {
		query := fmt.Sprintf(`
			START TRANSACTION;

			UPDATE "updoots"
			SET value = %d
			WHERE post_id = %d AND user_id = %d; 

			UPDATE "posts"
			SET points = points + %d
			WHERE posts.id = %d;

			COMMIT;
		`, realValue, postId, userId, 2*realValue, postId)

		if err := p.Table.Exec(query).Error; err != nil {
			return false
		}

	} else if updoot.PostID == 0 {
		// user has never voted before
		query := fmt.Sprintf(`
			START TRANSACTION;

			INSERT INTO "updoots" ("user_id", "post_id", "value")
			values(%d, %d, %d);

			UPDATE "posts"
			SET points = points + %d
			WHERE posts.id = %d;

			COMMIT;
		`, userId, postId, realValue, realValue, postId)

		if err := p.Table.Exec(query).Error; err != nil {
			return false
		}
	}

	return true
}
