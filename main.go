package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"log"
	"net/http"
	"go-study/app"
	"go-study/lib/httprouter"
	"go-study/lib/negroni"
)

func main() {

	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcerSafe("./src/casbin-examples/roles/auth_model.conf", "./src/casbin-examples/roles/policy.csv")
	if err != nil {
		log.Fatal("casbin:",err)
	}

	n := negroni.Classic()
	n.Use(negroni.Wrap(checkRole(authEnforcer)(NewRouter())))
	//n.UseHandler()
	n.Run(":3000")
	log.Print("Server started on localhost:3000")
}

//中间件
func ErrPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "进入了错误页面")
}

func checkRole(ae *casbin.Enforcer)func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		sub := "member"
		obj := "/home"
		act := "GET"

		fn := func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("这里可以写权限验证", r.URL)

			obj = r.URL.String()
			if ae.Enforce(sub, obj, act) == true {
				// permit alice to read data1
				fmt.Println("路由成功？")
				next.ServeHTTP(w, r)
			} else {
				fmt.Println("路由失败？")
				ErrPage(w, r)
			}


		}

		return http.HandlerFunc(fn)
	}
}

func NewRouter() *httprouter.Router {

	router := httprouter.New()
	routes := app.GetRoutes()
	for _, route := range routes {
		router.Handle(route.Method, route.Pattern, route.Handle)
	}
	return router
}
