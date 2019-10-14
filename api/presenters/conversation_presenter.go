package presenters

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"github.com/yuuis/PersonalDataRepository/models/conversation"
	"net/http"
)

func ConversationView(ctx context.Context, cv conversation.Conversation) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, cv)
}

func ConversationsView(ctx context.Context, cv *[]conversation.Conversation) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, cv)
}
