package utils
import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var LogoutConfig zap.Config

//控制台日志
func init() {
	LogoutConfig = zap.NewDevelopmentConfig()
	LogoutConfig.Level.SetLevel(zap.DebugLevel)
	LogoutConfig.Development = false
	LogoutConfig.Sampling = &zap.SamplingConfig{
		Initial:    100,
		Thereafter: 100,
	}
	LogoutConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	LogoutConfig.EncoderConfig.StacktraceKey = "stack"
	var err error
	Logger, err = LogoutConfig.Build()
	ErrFatal(err)
}

func ErrFatal(err error) {
	if err != nil {
		Logger.Fatal("ErrFatal", zap.NamedError("err", err))
	}
}
