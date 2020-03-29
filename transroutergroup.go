// Copyright 2020 trans owner
package trans-gogin

type TRouter interface{
	TRoutes
	Group(string,...HandlerFunc) *RouterGroup
}


