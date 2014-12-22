package feedify

import (
	"github.com/astaxie/beego"
	"github.com/feedlabs/feedify/context"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) GetInput() *context.Input {
	return &context.Input{c.Ctx.Input}
}
