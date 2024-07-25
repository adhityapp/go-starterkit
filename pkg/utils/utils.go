package utils

import (
	"github.com/sirupsen/logrus"
)

func LogError(loc string, funs string, err error) {
	logrus.WithFields(logrus.Fields{
		"loc":  loc,
		"func": funs,
	}).Error(err.Error())
}
