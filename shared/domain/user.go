package domain

type (
	User struct {
		Base
		Name string
	}

	// UserUsecase interface {
	// 	Fetch(ctx context.Context, cursor string, num int64) ([]User, string, error)
	// 	GetByID(ctx context.Context, id int64) (User, error)
	// 	Update(ctx context.Context, u *User) error
	// 	Store(ctx context.Context, u *User) error
	// 	Delete(ctx context.Context, id int64) error
	// }

	// UserRepository interface {
	// 	Fetch(ctx context.Context, cursor string, num int64) ([]User, string, error)
	// 	GetByID(ctx context.Context, id int64) (User, error)
	// 	Update(ctx context.Context, u *User) error
	// 	Store(ctx context.Context, u *User) error
	// 	Delete(ctx context.Context, id int64) error
	// }
)
