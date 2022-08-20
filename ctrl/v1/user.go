package v1

import (
	"bingo-example/domain/model/user"
	"github.com/xylong/bingo"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	RegisterCtrl(NewUserCtrl())
}

type UserCtrl struct {
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
