package server

import (
	v1 "kx-boutique/api/shop/v1"
	"kx-boutique/app/shop/internal/conf"
	"kx-boutique/app/shop/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, healthCheck *service.HealthService, products *service.ProductsService, carts *service.CartsService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterHealthHTTPServer(srv, healthCheck)
	v1.RegisterProductsHTTPServer(srv, products)
	v1.RegisterCartsHTTPServer(srv, carts)
	return srv
}
