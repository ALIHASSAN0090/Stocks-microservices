package logger

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type LoggerConfig struct {
	Level         string                `json:"level" bson:"level"`
	Path          string                `json:"path" bson:"path"`
	Name          string                `json:"name" bson:"name"`
	MaxSize       int                   `json:"max_size" bson:"max_size"`
	Rotationtime  int                   `json:"rotation_time" bson:"rotation_time"`
	SecretValues  []string              `json:"secret_values" bson:"secret_values"`
	EncoderConfig zapcore.EncoderConfig `json:"encoder_config" bson:"encoder_config"`
	IsEnabled     bool                  `json:"is_enabled" bson:"is_enabled"`
}

type InfoData struct {
	Date time.Time   `json : "date"`
	Info interface{} `json : "data"`
}
