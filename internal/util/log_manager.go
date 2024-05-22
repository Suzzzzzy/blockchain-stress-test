package util

import (
	"log"
	"os"
)

// LogManager 는 로그를 관리하는 구조체입니다.
type LogManager struct {
	logger *log.Logger
}

// NewLogManager 함수는 새로운 LogManager를 생성합니다.
func NewLogManager() *LogManager {
	return &LogManager{
		logger: log.New(os.Stdout, "[Blockchain]", log.LstdFlags),
	}
}

// Info 함수는 정보를 로그에 기록합니다.
func (lm *LogManager) Info(message string) {
	lm.logger.Printf("[INFO] %s\n", message)
}

// Warning 함수는 경고 메시지를 로그에 기록합니다.
func (lm *LogManager) Warning(message string) {
	lm.logger.Printf("[WARNING] %s\n", message)
}

// Error 함수는 오류 메시지를 로그에 기록합니다.
func (lm *LogManager) Error(message string) {
	lm.logger.Printf("[ERROR] %s\n", message)
}
