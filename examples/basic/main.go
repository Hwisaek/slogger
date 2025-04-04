package main

import (
	"context"
	"log/slog"

	"github.com/Hwisaek/slogger/v2"
)

func main() {
	// 기본 설정으로 초기화
	err := slogger.Init(nil)
	if err != nil {
		panic(err)
	}

	// 기본 로깅
	slog.Info("기본 정보 로그")
	slog.Debug("기본 디버그 로그 - 기본 레벨이 Info이므로 출력되지 않음")
	slog.Warn("경고 로그")
	slog.Error("에러 로그")

	// 구조화된 로깅
	slog.Info("구조화된 로그", "user_id", "user123", "action", "login")

	// 컨텍스트 기반 로깅
	ctx := context.Background()
	slog.InfoContext(ctx, "컨텍스트 로그")

	// 로그 그룹 사용하기
	logger := slog.Default().WithGroup("request")
	logger.Info("로그 그룹 내 메시지", "method", "GET", "path", "/api/users")
}
