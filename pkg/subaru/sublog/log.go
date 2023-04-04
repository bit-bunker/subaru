package sublog

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func SetDeveplomentLogger() {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	Logger = zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	)).Sugar().WithOptions(zap.AddCaller())
}

func DiscordGoBindLog() {
	discordgo.Logger = func(level, caller int, format string, args ...interface{}) {
		log := Logger.WithOptions(zap.AddCallerSkip(caller + 1))
		switch level {
		case discordgo.LogError:
			log.Errorf(format, args...)
		case discordgo.LogWarning:
			log.Warnf(format, args...)
		case discordgo.LogInformational:
			log.Infof(format, args...)
		case discordgo.LogDebug:
			log.Debugf(format, args...)
		}
	}
}
