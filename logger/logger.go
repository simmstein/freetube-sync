package logger

import (
	"context"
	"errors"
	"fmt"
	"time"

	lg "gorm.io/gorm/logger"
)

type LogLevel int

const (
	Silent LogLevel = iota + 1
	Error
	Warn
	Info
)

type logger struct {
	lg.Writer
	lg.Config
	lg.Interface
	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
}

func New(writer lg.Writer, config lg.Config) lg.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		infoStr = lg.Green + "%s\n" + lg.Reset + lg.Green + "[info] " + lg.Reset
		warnStr = lg.BlueBold + "%s\n" + lg.Reset + lg.Magenta + "[warn] " + lg.Reset
		errStr = lg.Magenta + "%s\n" + lg.Reset + lg.Red + "[error] " + lg.Reset
		traceStr = lg.Yellow + "[%.3fms] " + lg.BlueBold + "[rows:%v]" + lg.Reset + " %s"
		traceWarnStr = lg.Green + "%s " + lg.Reset + lg.RedBold + "[%.3fms] " + lg.Yellow + "[rows:%v]" + lg.Magenta + " %s" + lg.Reset
		traceErrStr = lg.RedBold + "%s " + lg.Reset + lg.Yellow + "[%.3fms] " + lg.BlueBold + "[rows:%v]" + lg.Reset + " %s"
	}

	return &logger{
		Writer:       writer,
		Config:       config,
		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

func (l *logger) LogMode(level lg.LogLevel) lg.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l *logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= lg.Info {
		l.Printf(l.infoStr+msg, data...)
	}
}

func (l *logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= lg.Warn {
		l.Printf(l.warnStr+msg, data...)
	}
}

func (l *logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= lg.Error {
		l.Printf(l.errStr+msg, data...)
	}
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= lg.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= lg.Error && (!errors.Is(err, lg.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= lg.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Printf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.LogLevel == lg.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Printf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Printf(l.traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
