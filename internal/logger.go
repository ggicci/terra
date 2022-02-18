package internal

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ConfigureLogger() error {
	var (
		logger *zap.Logger
		err    error
	)

	if IsDev() {
		logger, err = zap.NewDevelopment(
			zap.WrapCore(func(c zapcore.Core) zapcore.Core {
				return zapcore.RegisterHooks(
					c,
					func(e zapcore.Entry) error {
						// TODO(ggicci): if the error was important, we shall log it
						// to somewhere we can easily be notified, e.g. sentry.
						// Entry: https://pkg.go.dev/go.uber.org/zap@v1.19.1/zapcore#Entry
						return nil
					},
					// add other hooks here...
				)
			}),
			// add other options here...
		)
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)
	return nil
}
