package migrations

import (
	"context"
	"fmt"

	"github.com/calloc134/golang-ddd-test/cmd/migrate/schemas"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")

		models := []interface{}{
			(*schemas.UserTable)(nil),
			(*schemas.UserDetailTable)(nil),
			(*schemas.PostTable)(nil),
			(*schemas.PostDetailTable)(nil),
			(*schemas.CommentTable)(nil),
			(*schemas.CommentDetailTable)(nil),
		}

		for _, model := range models {

			q := db.NewCreateTable().
				IfNotExists().
				Model(model)

			fmt.Println(q.String())

			_, err := q.Exec(ctx)

			if err != nil {
				return err
			}
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")

		q := db.NewDropTable().
			IfExists().
			Model((*schemas.UserTable)(nil))

		_, err := q.Exec(ctx)

		if err != nil {
			return err
		}
		return nil
	})
}
