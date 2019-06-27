package sale_order_dates

import (
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "sale_order_dates"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})

}
