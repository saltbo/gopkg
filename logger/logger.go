//  Copyright 2019 The Go Authors. All rights reserved.
//  Use of this source code is governed by a BSD-style
//  license that can be found in the LICENSE file.

package logger

import (
	"bytes"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger *zap.Logger
	sugaredLogger *zap.SugaredLogger
	writeSyncer   *WriteSyncer
)

var accessFilename = "access.log"
var atomicLevel = zap.NewAtomicLevel()
var encoderCfg = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "level",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

type Config struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

func init() {
	stdoutCore := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.AddSync(os.Stdout), atomicLevel)
	sugaredLogger = zap.New(stdoutCore, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func Setup(opt Config) error {
	if err := atomicLevel.UnmarshalText([]byte(opt.Level)); err != nil {
		return err
	}
	Infof("[logger level: %s]", atomicLevel.String())

	// access.log
	encoderCfg.TimeKey = ""
	encoderCfg.NameKey = "logger"
	writeSyncer = newAsyncWriter(filepath.Join(opt.Path, accessFilename))
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), writeSyncer, atomicLevel)
	defaultLogger = zap.New(core)
	return nil
}

func Close() {
	_ = writeSyncer.Close()
}

func Debugf(format string, args ...interface{}) {
	sugaredLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	sugaredLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	sugaredLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	sugaredLogger.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	sugaredLogger.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	sugaredLogger.Fatalf(format, args...)
}

func Named(name string) *zap.Logger {
	if defaultLogger == nil {
		return zap.NewNop()
	}

	return defaultLogger.Named(name)
}

// For unit test
func BufLogger(buf *bytes.Buffer) *zap.Logger {
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.AddSync(buf), zap.DebugLevel)
	return zap.New(core)
}
