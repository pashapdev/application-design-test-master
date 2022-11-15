package usecases

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pashapdev/applicationDesignTest/internal/logger"
	"github.com/pashapdev/applicationDesignTest/internal/utils"
)

func ErrorHandling(w http.ResponseWriter, msg string, err error) {
	logger.Error(msg, err)
	var respondErr error
	if errors.Is(err, ErrNotValid) ||
		errors.Is(err, ErrUserAlreadyExists) {
		respondErr = utils.RespondWith400(w, err.Error())
		if respondErr != nil {
			logger.Error("", respondErr)
		}
		return
	}

	respondErr = utils.RespondWith500(w)
	if respondErr != nil {
		logger.Error("", respondErr)
	}
}

func DecodeBody(w http.ResponseWriter, r *http.Request, request interface{}) error {
	const notValidBodyMessage = "failed to decode request"
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logger.Error(notValidBodyMessage, err)
		respondErr := utils.RespondWith400(w, notValidBodyMessage)
		if respondErr != nil {
			logger.Error("failed to write response", err)
		}
		return err
	}
	return nil
}
