module go-example

go 1.15

require (
	echoin v0.0.0
	echoche v0.0.0
	github.com/floostack/transcoder v1.1.1
	github.com/giorgisio/goav v0.1.0
	github.com/golang/protobuf v1.4.3
	github.com/xfrr/goffmpeg v0.0.0-20200825100927-5550d238df5c
	golang.org/x/net v0.0.0-20200822124328-c89045814202
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20201116205149-79184cff4dfe // indirect
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
)

replace (
    echoin => ./echo/echoin
	echoche => ./echo/echoche
)
