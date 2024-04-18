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

			q := db.NewCreateTable().
				IfNotExists().
				Model((*schemas.UserTable)(nil))

			fmt.Println(q.String())
			
			_, err := q.Exec(ctx)

			if err != nil {
				return err
			}
			return nil
	}, func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [down migration] ")

		q := db.NewDropTable().
			IfExists().
			Model((*schemas.UserTable)(nil))

		_, err  := q.Exec(ctx)

		if err != nil {
			return err
		}
		return nil
	})
}
