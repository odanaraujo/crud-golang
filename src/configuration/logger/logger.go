package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	log         *zap.Logger
	LOG_OUTPUT  = "LOG_OUTPUT"
	LOG_LEVEL   = "LOG_LEVEL"
	LEVEL_KEY   = "LEVEL"
	TIME_KEY    = "time"
	MESSAGE_KEY = "message"
	ENCODING    = "json"
)

func init() {
	//zap.config faz com que possamos configurar o logger da nossa forma
	logConfig := zap.Config{
		OutputPaths: []string{getOutPutLogs()},            //onde jogar esse logger? stdout, por exemplo, joga no terminal
		Level:       zap.NewAtomicLevelAt(getLevelLogs()), //Qual o level do log? Degub, info, error...a ideia é colocar uma variável de ambiente
		Encoding:    ENCODING,                             //qual o tipo de return irei utilizar? json ou txt?
		EncoderConfig: zapcore.EncoderConfig{ //podemos configurar os nomes dos campos
			LevelKey:     LEVEL_KEY,                     // campo levelkey seja level dentro do json
			TimeKey:      TIME_KEY,                      //campo timekey seja time dentro do json
			MessageKey:   MESSAGE_KEY,                   //campo message seja message no json
			EncodeTime:   zapcore.ISO8601TimeEncoder,    //qual horário
			EncodeLevel:  zapcore.LowercaseLevelEncoder, //vai padronizar os logs de alguma forma - nesse caso, escolhi string e nada maiúsculo
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfig.Build() // builda o log
}

func Info(message string, tags ...zap.Field) {
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}

func getOutPutLogs() string {
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}

	return output
}

func getLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))) {
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}
