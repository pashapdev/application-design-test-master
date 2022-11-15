package getorders

import (
	"errors"
	"net/http"

	appErrors "github.com/pashapdev/applicationDesignTest/internal/errors"
	"github.com/pashapdev/applicationDesignTest/internal/logger"
	"github.com/pashapdev/applicationDesignTest/internal/utils"
)

func MakeHandler(s *Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userEmail := r.URL.Query().Get("email")
		response, err := s.Do(
			&request{
				Email: userEmail,
			})

		if err != nil {
			if errors.Is(err, appErrors.ErrEmailValidation) {
				if err = utils.RespondWith400(w, err.Error()); err != nil {
					logger.Error(err)
				}
				return
			} else {
				if err = utils.RespondWith500(w); err != nil {
					logger.Error(err)
				}
				return
			}
		}
		respondErr := utils.SuccessRespondWith200(w, response)
		if respondErr != nil {
			logger.Error(respondErr)
		}
	}
}
