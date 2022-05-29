package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lg *zap.Logger
var cfg *lgconfig

type lgconfig struct {
	filename   string
	level      string
	maxSize    int
	maxBackups int
	maxAge     int
}

func InitZapLogger(v *viper.Viper) (err error) {
	cfg = getLoggerCfg(v)
	writeSyncer := getLogWriter(cfg.filename, cfg.maxSize, cfg.maxBackups, cfg.maxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.level))
	if err != nil {
		return
	}
	core := zapcore.NewCore(encoder, writeSyncer, l)

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

func getEncoder() zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	if cfg.level == "DEBUG" {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,   // 日志切割前单个日志文件的大小，单位MB
		MaxBackups: maxBackup, // 保留旧日志文件的最大个数
		MaxAge:     maxAge,    // 保留旧日志文件的最大天数
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getLoggerCfg(v *viper.Viper) *lgconfig {
	return &lgconfig{
		filename:   v.GetString("filename"),
		level:      v.GetString("level"),
		maxSize:    v.GetInt("maxSize"),
		maxBackups: v.GetInt("maxBackups"),
		maxAge:     v.GetInt("maxAge"),
	}
}

func GinLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		ctx.Next()

		cost := time.Since(start)
		lg.Info(path,
			zap.Int("STATUS", ctx.Writer.Status()),
			zap.String("METHOD", ctx.Request.Method),
			zap.String("PATH", path),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("errors", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			// zap.String("ip", ctx.ClientIP()),
			// zap.String("user-agent", ctx.Request.UserAgent()),
			zap.Duration("COST", cost),
		)

	}
}

func GinRecovery(stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					lg.Error(
						ctx.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					ctx.Error(err.(error))
					ctx.Abort()
					return
				}

				if stack {
					lg.Error(
						"[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
