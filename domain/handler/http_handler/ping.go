package http_handler

import (
	"log"
	"net/http"
	"speSkillTest/lib/helper/http_response"
	"speSkillTest/lib/helper/timestamp"

	"github.com/gin-gonic/gin"
)

func (h *handler) Ping(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		log.Println(http_response.SendErrorAborted(c))
		return
	default:
	}

	resp, err := h.usecase.Ping(c.Request.Context())
	if err != nil {
		log.Println(http_response.SendError(err, c))
		return
	}

	response := http_response.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Status:     http_response.STANDARD_200_STATUS,
		Timestamp:  timestamp.GetNow(),
		Data:       resp,
	}

	c.JSON(http.StatusOK, response)
}
