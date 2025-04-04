package main

import (
	"context"
	"log/slog"

	"github.com/Hwisaek/slogger/v2"
)

func main() {
	// 고급 설정으로 초기화
	option := slogger.NewOption().
		WithAddSource(true).                       // 소스 위치 추가
		WithTimeFormat("2006-01-02 15:04:05.000"). // 타임스탬프 포맷 설정
		WithLogLevel(slog.LevelDebug).             // 로그 레벨을 Debug로 설정
		WithSpanIdKey(slogger.ContextKeySpanId).   // 스팬 ID 키 설정
		WithContextKey(slogger.ContextKeyTraceId)  // 트레이스 ID 키 추가

	if err := slogger.Init(option); err != nil {
		panic(err)
	}

	// 컨텍스트 생성 및 값 설정
	ctx := context.WithValue(context.Background(), slogger.ContextKeyTraceId, "trace-123")
	spanId := 0
	ctx = context.WithValue(ctx, slogger.ContextKeySpanId, &spanId)

	// 다양한 레벨의 로그 출력
	slog.DebugContext(ctx, "디버그 메시지", "step", 1)
	slog.InfoContext(ctx, "정보 메시지", "user", "hwisaek")

	// 로그 그룹 사용
	logger := slog.Default().WithGroup("process")
	logger.WarnContext(ctx, "경고 메시지", "status", "slow")
	logger.ErrorContext(ctx, "에러 메시지", "error", "connection refused")

	// spanId 자동 증가 확인
	newCtx := context.WithValue(context.Background(), slogger.ContextKeyTraceId, "trace-1234")
	newSpanId := 0
	newCtx = context.WithValue(newCtx, slogger.ContextKeySpanId, &newSpanId)
	slog.InfoContext(newCtx, "첫 번째 작업 완료")
	slog.InfoContext(newCtx, "두 번째 작업 완료")
	slog.InfoContext(newCtx, "세 번째 작업 완료")
}
