// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

package tracelog

import (
	"fmt"
)

//** STARTED AND COMPLETED

// Startedcd uses the Trace destination and adds a Started tag to the log line
func Startedcd(callDepth int, title string, functionName string) {
	logger.Serialize.Lock()
	logger.Trace.Output(callDepth, fmt.Sprintf("%s : %s : Started\n", title, functionName))
	logger.Serialize.Unlock()
}

// Startedfcd uses the Trace destination and writes a Started tag to the log line
func Startedfcd(callDepth int, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Trace.Output(callDepth, fmt.Sprintf("%s : %s : Started : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

// Completedcd uses the Trace destination and writes a Completed tag to the log line
func Completedcd(callDepth int, title string, functionName string) {
	logger.Serialize.Lock()
	logger.Trace.Output(callDepth, fmt.Sprintf("%s : %s : Completed\n", title, functionName))
	logger.Serialize.Unlock()
}

// Completedfcd uses the Trace destination and writes a Completed tag to the log line
func Completedfcd(callDepth int, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Trace.Output(callDepth, fmt.Sprintf("%s : %s : Completed : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

// CompletedErrorcd uses the Error destination and writes a Completed tag to the log line
func CompletedErrorcd(callDepth int, err error, title string, functionName string) {
	logger.Serialize.Lock()
	logger.Error.Output(callDepth, fmt.Sprintf("%s : %s : Completed : ERROR : %s\n", title, functionName, err))
	logger.Serialize.Unlock()
}

// CompletedErrorfcd uses the Error destination and writes a Completed tag to the log line
func CompletedErrorfcd(callDepth int, err error, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Error.Output(callDepth, fmt.Sprintf("%s : %s : Completed : ERROR : %s : %s\n", title, functionName, fmt.Sprintf(format, a...), err))
	logger.Serialize.Unlock()
}

//** TRACE

// Tracecd writes to the Trace destination
func Tracecd(callDepth int, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Trace.Output(callDepth, fmt.Sprintf("%s : %s : Info : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

//** INFO

// Infocd writes to the Info destination
func Infocd(callDepth int, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Info.Output(callDepth, fmt.Sprintf("%s : %s : Info : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

//** WARNING

// Warningcd writes to the Warning destination
func Warningcd(callDepth int, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Warning.Output(callDepth, fmt.Sprintf("%s : %s : Info : %s\n", title, functionName, fmt.Sprintf(format, a...)))
	logger.Serialize.Unlock()
}

//** ERROR

// Errorcd writes to the Error destination and accepts an err
func Errorcd(callDepth int, err error, title string, functionName string) {
	logger.Serialize.Lock()
	logger.Error.Output(callDepth, fmt.Sprintf("%s : %s : ERROR : %s\n", title, functionName, err))
	logger.Serialize.Unlock()
}

// Errorfcd writes to the Error destination and accepts an err
func Errorfcd(callDepth int, err error, title string, functionName string, format string, a ...interface{}) {
	logger.Serialize.Lock()
	logger.Error.Output(callDepth, fmt.Sprintf("%s : %s : ERROR : %s : %s\n", title, functionName, fmt.Sprintf(format, a...), err))
	logger.Serialize.Unlock()
}

//** ALERT

// Alertcd write to the Error destination and sends email alert
func Alertcd(callDepth int, subject string, title string, functionName string, format string, a ...interface{}) {
	message := fmt.Sprintf("%s : %s : ALERT : %s\n", title, functionName, fmt.Sprintf(format, a...))

	logger.Serialize.Lock()
	logger.Error.Output(callDepth, message)
	logger.Serialize.Unlock()

	SendEmailException(subject, message)
}

// CompletedAlertcd write to the Error destination, writes a Completed tag to the log line and sends email alert
func CompletedAlertcd(callDepth int, subject string, title string, functionName string, format string, a ...interface{}) {
	message := fmt.Sprintf("%s : %s : Completed : ALERT : %s\n", title, functionName, fmt.Sprintf(format, a...))

	logger.Serialize.Lock()
	logger.Error.Output(callDepth, message)
	logger.Serialize.Unlock()

	SendEmailException(subject, message)
}