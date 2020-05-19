package log

/************************* Dummy Logger *************************/

type dummyLogger struct { }

func NewDummyLogger() *dummyLogger {
	return &dummyLogger{ }
}

func (dlog *dummyLogger) Trace(v ...interface{}) { }

func (dlog *dummyLogger) Tracef(format string, v ...interface{}) { }

func (dlog *dummyLogger) Debug(v ...interface{}) { }

func (dlog *dummyLogger) Debugf(format string, v ...interface{}) { }

func (dlog *dummyLogger) Info(v ...interface{}) { }

func (dlog *dummyLogger) Infof(format string, v ...interface{}) { }

func (dlog *dummyLogger) Warn(v ...interface{}) { }

func (dlog *dummyLogger) Warnf(format string, v ...interface{}) { }

func (dlog *dummyLogger) Error(v ...interface{}) { }

func (dlog *dummyLogger) Errorf(format string, v ...interface{}) { }


/************************* Dummy Logger Builder *************************/

type dummyLoggerBuilder struct { }

func NewDummyLoggerBuilder() *dummyLoggerBuilder {
	return &dummyLoggerBuilder{ }
}

func (builder *dummyLoggerBuilder) Build() Logger {
	return &dummyLogger{ }
}
