module github.com/go-i2p/go-rst

go 1.26.3

require (
	github.com/jung-kurt/gofpdf v1.16.2
	github.com/leonelquinteros/gotext v1.7.2
	github.com/yosssi/gohtml v0.0.0-20201013000340-ee4748c638f4
)

require golang.org/x/net v0.59.0 // indirect

retract (
	v0.1.59999
	v0.1.5999
	v0.1.599
)
