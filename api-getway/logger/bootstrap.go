package logger

import (
	"fmt"
	"reflect"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var ZapLogger *zap.Logger
var configuration LoggerConfig

type OverloadedEncoder struct { //overloaded encoder
	zapcore.Encoder
	secretKey []string
}

func LoggerInit(conf LoggerConfig) *zap.Logger {
	encoder := zapcore.NewConsoleEncoder(conf.EncoderConfig)
	ZapEncoder := OverloadedEncoder{encoder, conf.SecretValues}
	atomicLevel := zap.AtomicLevel{}
	err := atomicLevel.UnmarshalText([]byte(conf.Level))
	if err != nil {
		fmt.Println("Error in unmarshalling config log level into zap level")
		panic(err.Error())
	}

	fileRotationTime := time.Duration(conf.Rotationtime) * time.Minute
	fileRotation := zapcore.AddSync(NewTimeRotationWriter(conf.Path, conf.Name, fileRotationTime, conf.MaxSize))

	core := zapcore.NewCore(
		ZapEncoder,
		zapcore.NewMultiWriteSyncer(fileRotation),
		atomicLevel,
	)

	ZapLogger := zap.New(core)
	configuration = conf
	return ZapLogger
}

func (m *OverloadedEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {

	filtered := make([]zapcore.Field, 0, len(fields))

	for _, field := range fields {
		res := field
		target := field.Interface
		reflectedType := reflect.TypeOf(target)
		if target != nil {

		}
	}
}
