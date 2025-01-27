package preoncall_service

import (
	"context"
	"encoding/json"
	"fmt"
)

// Intercom Canvas Receiver
type IntercomCanvasReceiver struct {
	Content IntercomContent `json:"content"`
}

type IntercomContent struct {
	Components []IntercomComponent `json:"components"`
}

type IntercomComponent struct {
	Type    string   `json:"type"`
	ID      string   `json:"id"`
	Label   string   `json:"label"`
	Options []Option `json:"options"`
	Value   *string  `json:"value,omitempty"`
}

func HandlePreoncallCanvasSubmitAction(ctx context.Context, body string) (CanvasReponse, error) {
	//log. := utils.Get//log.gerWithMethod(ctx, "HandlePreoncallCanvasSubmitAction")
	//log..Infof("HandlePreoncallCanvasSubmitAction request body: %v", body)
	//fmt.Printf("HandlePreoncallCanvasSubmitAction request body: %v \n", body)
	var canvasReq IntercomCanvasRequest
	if err := json.Unmarshal([]byte(body), &canvasReq); err != nil {
		fmt.Println("meet parse err", err)

		return CanvasReponse{}, err
	}

	intercomConversationID := canvasReq.Conversation.ConversationID
	assigneeID := canvasReq.Conversation.AdminAssigneeID
	inputValues := canvasReq.InputValues

	//log..Infof("HandlePreoncallCanvasSubmitAction pretty's request %v \n", larkcore.Prettify(canvasReq))
	//fmt.Printf("HandlePreoncallCanvasSubmitAction pretty's request %v \n", larkcore.Prettify(canvasReq))
	fmt.Printf("HandlePreoncallCanvasSubmitAction pretty's input value %v \n", inputValues)
	fmt.Sprintf("component id %v \n", canvasReq.ComponentID)
	var response CanvasReponse
	// TODO: call intercom pre ocall api to create ticket
	//log..Infof("HandlePreoncallCanvasSubmitAction component id %v", canvasReq.ComponentID)
	fmt.Printf("HandlePreoncallCanvasSubmitAction component id %v \n", canvasReq.ComponentID)
	switch canvasReq.ComponentID {
	case CategorySingleSelectID:
		//log..Infof("HandlePreoncallCanvasSubmitAction single select ")
		if value, ok := canvasReq.InputValues[CategorySingleSelectID]; ok {
			//log..Infof("HandlePreoncallCanvasSubmitAction single select value %v", value)
			fmt.Printf("HandlePreoncallCanvasSubmitAction single select value %v \n", value)
			if value == CreateTicketOptionID {
				response = GetCreateTicketCanvasBody(ctx, inputValues, intercomConversationID, assigneeID, value, canvasReq.CurrentCanvas)
			} else if value == RelatedTicketOptionID {
				response = GetRelatedTicketCanvasBody(ctx, inputValues, intercomConversationID)
			}
		}
	case BizLineSearchButtonID, RegionSearchButtonID, SubmitTicketButtonID, StackSearchButtonID:
		response = GetCreateTicketCanvasBody(ctx, inputValues, intercomConversationID, assigneeID, canvasReq.ComponentID, canvasReq.CurrentCanvas)
		//case RegionSearchButtonID:
		//	response = GetCreateTicketCanvasBody(ctx, inputValues, intercomConversationID, assigneeID, RegionSearchButtonID, canvasReq.CurrentCanvas)
		//case SubmitTicketButtonID:
		//	response = GetCreateTicketCanvasBody(ctx, inputValues, intercomConversationID, assigneeID, SubmitTicketButtonID, canvasReq.CurrentCanvas)
		//case StackSearchButtonID:
		//	response = GetCreateTicketCanvasBody(ctx, inputValues, intercomConversationID, assigneeID, StackSearchButtonID, canvasReq.CurrentCanvas)
	}

	//log..Infof("HandlePreoncallCanvasSubmitAction vanvas response %v", larkcore.Prettify(response))
	//fmt.Printf("------- HandlePreoncallCanvasSubmitAction vanvas response %v \n", larkcore.Prettify(response))
	return response, nil
}

func HandlePreoncallInitializationAction(ctx context.Context) CanvasReponse {
	//log. := utils.Get//log.gerWithMethod(ctx, "HandlePreoncallInitializationAction")
	response := GetInitTicketCanvasBody()

	//log..Infof("HandlePreoncallInitializationAction response %v", larkcore.Prettify(response))

	return response
}
