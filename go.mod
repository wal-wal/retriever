module retriever

go 1.21.4

require (
	application v0.0.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gofiber/fiber/v2 v2.51.0
	github.com/google/uuid v1.5.0
	persistence v0.0.0
	presentation v0.0.0
)

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.50.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)

replace (
	application v0.0.0 => ./application
	persistence v0.0.0 => ./persistence
	presentation v0.0.0 => ./presentation
)
