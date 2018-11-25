//go:generate gorunpkg github.com/99designs/gqlgen

package iotendancebe

import (
	context "context"
	"fmt"
	"math/rand"

	user "github.com/luizcavalieri/IoTendance-be/service/user"
)

type Resolver struct{
	users []user.User
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input *NewUser) (user.User, error) {
	user := user.User{
		UserID: fmt.Sprintf("T%d", rand.Int()),
		Username: input.Username,
		UserFname: input.UserFname,
		UserLname: input.UserLname,
		Password: input.Password,
	}
	r.users = append(r.users, user)
	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]user.User, error) {
	return r.users, nil
}
