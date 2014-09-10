package feedify

import (
	// Golang packages
	"fmt"
	"strconv"

	// Beego framework packages
	"github.com/astaxie/beego"

	// feedify packages
	"github.com/feedlabs/feedify/config"
	_ "github.com/feedlabs/feedify/stream/adapter/message"
	_ "github.com/feedlabs/feedify/graph/neo4jlang"
)

func Banner() {
	fmt.Printf("Starting app '%s' on port '%s'\n", config.GetConfigKey("appname"), config.GetConfigKey("feedify::port"))
}

func Run() {
	Banner()

	beego.HttpPort, _ = strconv.Atoi(config.GetConfigKey("feedify::port"))
	beego.Run()
}
