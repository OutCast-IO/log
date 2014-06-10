// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

package tracelog

import (
	"fmt"
)

//** STARTED AND COMPLETED

// Started uses the Serialize destination and adds a Started tag to the log line
func Started(title string, functionName string) {
	logger.Serialize.Lock()
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Started\n", title, functionName))
	logger.Serialize.Unlock()
}

// Startedf uses the Serialize destination and writes a Started tag to the log line
func Startedf(title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Started : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

// Completed uses the Serialize destination and writes a Completed tag to the log line
func Completed(title string, functionName string) {
	logger.Serialize.Lock()
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Completed\n", title, functionName))
	logger.Serialize.Unlock()
}

// COMPLETEDf uses the Serialize destination and writes a Completed tag to the log line
func Completedf(title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Completed : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

// CompletedError uses the Error destination and writes a Completed tag to the log line
func CompletedError(err error, title string, functionName string) {
	logger.Serialize.Lock()
	logger.Error.Output(2, fmt.Sprintf("%s : %s : Completed : ERROR : %s\n", title, functionName, err))
	logger.Serialize.Unlock()
}

// CompletedErrorf uses the Error destination and writes a Completed tag to the log line
func CompletedErrorf(err error, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Error.Output(2, fmt.Sprintf("%s : %s : Completed : ERROR : %s : %s\n", title, functionName, fmt.Sprintf(format, a...), err))
	logger.Serialize.Unlock()
}

//** TRACE

// Trace writes to the Trace destination
func Trace(title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Trace.Output(2, fmt.Sprintf("%s : %s : Info : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

//** INFO

// Info writes to the Info destination
func Info(title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Info.Output(2, fmt.Sprintf("%s : %s : Info : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

//** WARNING

// Warning writes to the Warning destination
func Warning(title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Warning.Output(2, fmt.Sprintf("%s : %s : Info : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

//** ERROR

// Error writes to the Error destination and accepts an err
func Error(err error, title string, functionName string) {
	logger.Serialize.Lock()
	logger.Error.Output(2, fmt.Sprintf("%s : %s : ERROR : %s\n", title, functionName, err))
	logger.Serialize.Unlock()
}

// Errorf writes to the Error destination and accepts an err
func Errorf(err error, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Error.Output(2, fmt.Sprintf("%s : %s : ERROR : %s : %s\n", title, functionName, fmt.Sprintf(format, a...), err))
	logger.Serialize.Unlock()
}

//** ALERT

// Alert write to the Error destination and sends email alert
func Alert(subject string, title string, functionName string, format string, a ...interface{}) {
	message := fmt.Sprintf("%s : %s : ALERT : %s\n", title, functionName, fmt.Sprintf(format, a...))

	logger.Serialize.Lock()
	logger.Error.Output(2, message)
	logger.Serialize.Unlock()

	SendEmailException(subject, message)
}

// CompletedAlert write to the Error destination, writes a Completed tag to the log line and sends email alert
func CompletedAlert(subject string, title string, functionName string, format string, a ...interface{}) {
	message := fmt.Sprintf("%s : %s : Completed : ALERT : %s\n", title, functionName, fmt.Sprintf(format, a...))

	logger.Serialize.Lock()
	logger.Error.Output(2, message)
	logger.Serialize.Unlock()

	SendEmailException(subject, message)
}
