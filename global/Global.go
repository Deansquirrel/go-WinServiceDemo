package global

import (
	"context"
	"github.com/Deansquirrel/goWinServiceDemo/config"
)

var SysConfig *config.SysConfig
var Ctx context.Context
var Cancel func()
