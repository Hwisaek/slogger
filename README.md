# Slogger - 컨텍스트 기반 Go 로깅 라이브러리

Go의 표준 로깅 패키지 `slog`를 확장한 로깅 라이브러리입니다. 컨텍스트 기반의 로깅과 추가적인 기능을 제공합니다.

## 특징

- 컨텍스트 기반 로깅 - 트레이싱 ID, 스팬 ID 등 컨텍스트 값을 로그에 포함
- 소스 위치 자동 기록 기능
- 타임스탬프 포맷 커스터마이징
- JSON 형식 로그 출력
- 구성 가능한 로그 레벨

## 설치

```shell
go get github.com/Hwisaek/slogger/v2
```

## 기본 사용법

```go
package main

import (
	"github.com/Hwisaek/slogger/v2"
	"log/slog"
)

func main() {
	// 기본 설정으로 초기화
	if err := slogger.Init(nil); err != nil {
		panic(err)
	}
	
	// 기본 로깅
	slog.Info("Hello, world!")
	slog.Debug("Debug message")
	slog.Warn("Warning message")
	slog.Error("Error message")
}
```

## 고급 설정

```go
package main

import (
	"context"
	"github.com/Hwisaek/slogger/v2"
	"log/slog"
)

func main() {
	// 사용자 정의 옵션으로 초기화
	option := slogger.NewOption().
		WithAddSource(true).                           // 소스 위치 기록
		WithTimeFormat("2006-01-02 15:04:05.000").     // 사용자 정의 시간 포맷
		WithLogLevel(slog.LevelDebug).                 // 로그 레벨 설정
		WithSpanIdKey(slogger.ContextKeySpanId).       // 스팬 ID 키 설정
		WithContextKey(slogger.ContextKeyTraceId)      // 컨텍스트 키 추가
	
	if err := slogger.Init(option); err != nil {
		panic(err)
	}
	
	// 컨텍스트와 함께 로깅
	ctx := context.Background()
	ctx = context.WithValue(ctx, slogger.ContextKeyTraceId, "trace-123")
	spanId := 0
	ctx = context.WithValue(ctx, slogger.ContextKeySpanId, &spanId)
	
	slog.InfoContext(ctx, "Request started", "user_id", "user-456")
}
```

## 로그 그룹 사용

로그 그룹을 사용하면 관련 정보를 구조화하여 더 명확하게 로그를 남길 수 있습니다.
컨텍스트 값은 그룹 내부에 포함됩니다.

```go
// 로거에 그룹 추가
logger := slog.Default().WithGroup("request")
logger.InfoContext(ctx, "API 호출", "method", "GET", "path", "/api/users")
```

출력 결과:
```json
{
  "time": "2023-01-01T12:00:00.000+09:00",
  "level": "INFO",
  "msg": "API 호출",
  "request": {
    "method": "GET",
    "path": "/api/users",
    "trace-id": "trace-123",
    "span-id": 1
  }
}
```

## 빌드 옵션

소스 위치를 정확하게 기록하기 위해 빌드 시 `-trimpath` 옵션을 사용하는 것을 권장합니다:

```shell
go build -trimpath
```

## 라이선스

MIT License
