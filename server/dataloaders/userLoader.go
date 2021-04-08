package dataloaders

import (
	"context"
	"lireddit/db"
	"lireddit/models"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const userLoaderKey = "userLoaderKey"

var psql *gorm.DB

func init() {
	psql, _ = db.Connect()
}

func DataLoaderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userLoader := UserLoader{
			maxBatch: 51,
			wait:     1 * time.Millisecond,
			fetch: func(keys []string) ([]*models.User, []error) {
				var users []*models.User

				err := psql.Model(&users).Where("id IN ?", keys).Find(&users).Error
				if err != nil {
					return nil, []error{err}
				}
				return users, nil
			},
		}
		ctx := context.WithValue(c.Request().Context(), userLoaderKey, &userLoader)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
