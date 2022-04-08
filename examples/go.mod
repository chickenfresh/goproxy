module github.com/chickenfresh/goproxy/examples/goproxy-transparent

go 1.18

require (
	github.com/chickenfresh/goproxy v0.0.0-20181111060418-2ce16c963a8a
	github.com/inconshreveable/go-vhost v0.0.0-20160627193104-06d84117953b
)

require github.com/rogpeppe/go-charset v0.0.0-20190617161244-0dc95cdf6f31 // indirect

replace github.com/chickenfresh/goproxy => ../
