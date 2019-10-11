package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/PersonalDataRepository/api/presenters"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models"
	"github.com/yuuis/PersonalDataRepository/models/conversation"
	"time"
)

func (r *Registry) GetConversation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := conversation.NewDataStore(r.db)

	cv, err := ds.Get()

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.ConversationView(ctx, *cv)
	}
}

func (r *Registry) FindConversation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := conversation.NewDataStore(r.db)

	tID := c.Param("transactionID")

	cv, err := ds.FindByTransactionId(tID)

	if err != nil {
		switch err {
		case utilities.NotFoundError:
			presenters.ViewNoContent(ctx)
		default:
			presenters.ViewInternalServerError(ctx, err)
		}
	} else {
		presenters.ConversationsView(ctx, cv)
	}
}

func (r *Registry) CreateConversation(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	ds := conversation.NewDataStore(r.db)

	var ipt inputConversation
	if err := c.BindJSON(&ipt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	cv, err := ds.Store(&conversation.Conversation{
		ID:            models.GenerateUUID(),
		TransactionId: ipt.TransactionId,
		RequestText:   ipt.RequestText,
		ResponseText:  ipt.ResponseText,
		CreatedAt:     time.Now(),
	})

	if err != nil {
		presenters.ViewInternalServerError(ctx, err)
	}

	presenters.ConversationView(ctx, *cv)
}

type inputConversation struct {
	TransactionId string `json:"transaction_id"`
	RequestText   string `json:"request_text"`
	ResponseText  string `json:"response_text"`
}
