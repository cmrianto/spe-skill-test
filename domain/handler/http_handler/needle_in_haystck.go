package http_handler

import (
	"log"
	"net/http"
	"speSkillTest/domain/request"
	"speSkillTest/lib/helper/http_response"
	"speSkillTest/lib/helper/timestamp"

	"github.com/gin-gonic/gin"
)

func (h *handler) NeedleInHaystack(c *gin.Context) {
	select {
	case <-c.Request.Context().Done():
		log.Println(http_response.SendErrorAborted(c))
		return
	default:
	}

	in := request.NeedleInHaystackRequest{}

	err := c.ShouldBindJSON(&in)
	if err != nil {
		log.Println(http_response.SendError(err, c))
		return
	}

	resp, err := h.usecase.NeedleInHaystack(c.Request.Context(), &in)
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
