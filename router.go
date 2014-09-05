package feedify

import (
	"github.com/astaxie/beego"
)

func Router(rootpath string, c beego.ControllerInterface, mappingMethods ...string) *beego.App {
	return beego.Router(rootpath, c, mappingMethods...)
}
