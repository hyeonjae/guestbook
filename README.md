# guestbook

go 프로젝트 구조입니다.  

```
.
├── api
│   └── v1
│       └── guestbook.proto
├── cmd
│   └── guestbook
│       ├── cli
│       │   └── flags.go
│       └── main.go
├── config
│   └── local.yaml
├── internal
│   ├── bootstrap
│   │   └── grpc.go
│   ├── config
│   │   └── config.go
│   ├── service
│   │   └── guestbook.go
│   ├── usecase
│   │   └── guestbook.go
...
```

주요 특징은 다음과 같습니다.
 - [uber.fx](https://github.com/uber-go/fx) 를 이용한 의존성 관리
 - bootstrap, gateway, usecase, service 로 이어지는 레이어
 - config 파일을 이용한 설정값 관리

## uber.fx

```go
// cmd/guestbook/main.go
func main() {
	app := fx.New(
		fx.Provide(
			cli.ParseFlags,
			config.NewConfig,
			gateway.NewGuestbook,
			bootstrap.NewGuestbook,
		),
		fx.Invoke(serve),
	)
	// ...
}

func serve(lc fx.Lifecycle, guestbook *bootstrap.Guestbook) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go guestbook.Serve()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
```

```go
// internal/bootstrap/grpc.go
type Guestbook struct {
	config *config.Config
	server *gateway.Guestbook
}

func NewGuestbook(config *config.Config, server *gateway.Guestbook) *Guestbook {
	return &Guestbook{
		config: config,
		server: server,
	}
}
```

## layer dependency

```
main <--- bootstrap 
gateway <--- usecase <--- service <--- adapter 
```

main 은 의존성 주입 관계를 명시하고, fx를 실행시킵니다.  
bootstrap 는 앱이 구동되는 동안 실행되어야 할 서버 (http server, grpc server)나 백그라운드 작업 (tick 등)을 실행시키는 레이어입니다.  
gateway 에는 grpc 나 http server 의 endpoint 를 정의합니다.   
usecase 에는 gateway 의 endpoint 에서 구현이 필요한 부분을 이곳에 구현합니다.
service 에는 비즈니스 로직을 구현합니다.

## config

```go
listen: ":8081"

grpc:
  max_connection_idle: "15s"
  max_connection_age: "30s"
  max_connection_age_grace: "15s"
  time: "15s"
  timeout: "10s"

mysql:
  address: "localhost:3306"
  username: "admin"
  password: "admin"
  database: "guestbook"

redis:
  address: "localhost:6379"
  password: "admin"
```

host 와 port 는 분리하지 않았습니다. 코드에서 host 와 port 를 합치는 코드 (`fmt.Sprintf("%s:%s")`)를 줄일 수 있습니다.  
server 는 listen 으로, client 는 address 로 규칙을 정해보았습니다.  
시간값은 [단위](https://github.com/golang/go/blob/6afa0ae4e54ec049f291050b82c2a770bb3644b1/src/time/format.go#L1389)와 함께 명시한 문자열로 설정합니다.  
