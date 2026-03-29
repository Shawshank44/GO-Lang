package services

import (
	utilssql "blog_rest_api/internal/repositories/Utils_SQL"
	"blog_rest_api/pkg/utils"
	"time"
)

func StartCleaner() error {
	ticker := time.NewTicker(10 * time.Minute)
	for {
		<-ticker.C
		err := utilssql.UploadCleanUp()
		if err != nil {
			return utils.ErrorHandler(err, "Unable to start the cleaner")
		}

		err = utilssql.SessionsCleanUp()
		if err != nil {
			return utils.ErrorHandler(err, "Unable to cleanup sessions")
		}
	}
}
