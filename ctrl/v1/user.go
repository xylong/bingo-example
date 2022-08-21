package v1

import (
	"bingo-example/application/server"
	"bingo-example/domain/entity/user"
	"github.com/xylong/bingo"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	RegisterCtrl(NewUserCtrl())
}

type UserCtrl struct {
	service       *server.UserService `inject:"-"`
	*mongo.Client `inject:"-"`
}

func NewUserCtrl() *UserCtrl {
	return &UserCtrl{}
}

func (c *UserCtrl) Name() string {
	return "UserCtrl"
}

func (c *UserCtrl) Route(group *bingo.Group) {
	group.GET("index", c.index)
	group.POST("users", c.add)
}

func (c *UserCtrl) index(ctx *bingo.Context) interface{} {
	collection := c.Database("test").Collection("users")
	if one, err := collection.InsertOne(ctx, &user.User{
		Name:     "静静",
		Phone:    "13811223344",
		Birthday: "1990-10-01",
		Gender:   0,
	}); err != nil {
		return err.Error()
	} else {
		return one.InsertedID
	}
}

func (c *UserCtrl) add(ctx *bingo.Context) string {
	return c.service.Index()
}
