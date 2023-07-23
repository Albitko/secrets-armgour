package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap *zap.SugaredLogger

func Init() {
	file, err := os.Create("armgour-server.log")
	if err != nil {
		panic(fmt.Errorf("error open file for logs: %w", err))
	}

	pe := zap.NewDevelopmentEncoderConfig()
	fileEncoder := zapcore.NewJSONEncoder(pe)
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	level := zap.DebugLevel

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	Zap = zap.New(core).Sugar()
}
