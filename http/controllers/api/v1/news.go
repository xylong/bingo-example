package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/xylong/bingo"
)

func init() {
	RegisterCtrl(NewNewController())
}

type NewController struct {
	Es *elastic.Client `inject:"-"`
}

func NewNewController() *NewController {
	return &NewController{}
}

func (c *NewController) index(ctx *gin.Context) (int, string, interface{}) {
	m, err := c.Es.GetMapping().Index("news").Do(ctx)
	if err != nil {
		return 500, err.Error(), m
	}

	return 0, "", m
}

func (c *NewController) Name() string {
	return "news"
}

func (c *NewController) Route(group *bingo.Group) {
	group.GET("news", c.index)
}
