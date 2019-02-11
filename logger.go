package logger

import (
	"log"

	"github.com/getsentry/raven-go"
	"bitbucket.org/zolotarev-dmitriy/accounts/http/components"
)

func Logger(args ...interface{}) {
	for _, arg := range args {
		if err, ok := arg.(*components.HttpError); ok {
			log.Println("http error:", err.Error(), err.Message())
			continue
		}

		if err, ok := arg.(error); ok {
			raven.CaptureError(err, nil)
			log.Println("error:", arg)
			continue
		}

		log.Println(arg)
	}
}