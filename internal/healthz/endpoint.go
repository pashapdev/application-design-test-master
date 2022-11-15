package healthz

import (
	"net/http"

	"github.com/pashapdev/applicationDesignTest/internal/entities"
	"github.com/pashapdev/applicationDesignTest/internal/logger"
	"github.com/pashapdev/applicationDesignTest/internal/utils"
)

func MakeHandler(info *entities.AppInfo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		response := &response{
			Name:         info.Name,
			BuildVersion: info.BuildVersion,
			BuildTime:    info.BuildTime,
			GitTag:       info.GitTag,
			GitHash:      info.GitHash,
		}
		err := utils.SuccessRespondWith200(w, struct{}{})
		if err != nil {
			logger.Error("failed to decode response", response, err)
		}
	}
}
