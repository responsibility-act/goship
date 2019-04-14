package webserver

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

//Instance of webserver
type Instance struct {
	app *iris.Application
}

// URLRequest router method on app
func (objInstance Instance) URLRequest(passedPath string, resHandler context.Handler) {
	objInstance.app.Get(passedPath, resHandler)
}

// StartInstance prepare instance data, before running
func StartInstance(objInstance *Instance) {
	objInstance.app = iris.New()
}

//Run iris instance
func Run(objInstance *Instance) {
	objInstance.app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}
