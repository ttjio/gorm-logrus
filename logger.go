package gorm_logrus

import (
"context"
"errors"
"time"

"github.com/sirupsen/logrus"
"gorm.io/gorm"
gormLogger "gorm.io/gorm/logger"
"gorm.io/gorm/utils"
)

type logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	log *logrus.Logger
}

func New(log *logrus.Logger) *logger {
	return &logger{
		SkipErrRecordNotFound: true,
		log: log,
	}
}

func (l *logger) LogMode(gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
	l.log.WithContext(ctx).Infof(s, args)
}

func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.log.WithContext(ctx).Warnf(s, args)
}

func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
	l.log.WithContext(ctx).Errorf(s, args)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		l.log.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.log.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	l.log.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
