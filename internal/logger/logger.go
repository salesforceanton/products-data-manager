package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func LogServerIssue(err error) {
	logrus.WithFields(logrus.Fields{
		"point":   "server",
		"problem": fmt.Sprintf("Server running error: %s", err.Error()),
	}).Error(err)
}
