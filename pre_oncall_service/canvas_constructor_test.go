package preoncall_service

import (
	"context"
	"fmt"
	pre_oncall "github.com/digitalocean/sample-golang/pre_oncall_api"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetCreateTicketCanvasBody(t *testing.T) {
	inputValues := map[string]string{
		"category_single_select": "create_ticket",
		"business_single_select": "bid_urso2fcl",
		"priority_single_select": "P0",
		"stack_single_select":    "pc",
		"region_single_select":   "India",
		"version_single_select":  "7.4.8",
		"app_single_select":      "Lark",
		"external_single_select": "字节内",
		"source_single_select":   "客服渠道",
		"reporter_single_select": "1110",
		"remarks_text":           "this is 1 intercom ticket",
		"channel_type":           "intercom",
		"biz_ticket_id":          "11111",
	}

	intercomConversationID := "12345"
	assigneeID := 12345
	value := "create_ticket"
	canvas := IntercomCanvasReceiver{
		Content: IntercomContent{
			Components: []IntercomComponent{},
		},
	}

	response := GetCreateTicketCanvasBody(context.Background(), inputValues, intercomConversationID, assigneeID, value, canvas)

	fmt.Println(response)
	//log..Infof("response: %v", response)
}

// TestInitPreOncallCanvas tests the InitPreOncallCanvas function
func TestInitPreOncallCanvas(t *testing.T) {
	// Call the function
	canvasResp := InitPreOncallCanvas()

	// Create the expected response
	expectedOption1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	expectedOption2 := NewOption(CreateTicketOptionID, "Create Ticket")
	expectedAction := NewAction("submit")
	expectedSingleSelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*expectedOption1, *expectedOption2}, &expectedAction, nil)
	expectedContent := newContent([]Component{expectedSingleSelect})
	expectedCanvasResp := newCanvasReponse(*expectedContent)

	// Compare the actual response with the expected response
	if !reflect.DeepEqual(canvasResp, *expectedCanvasResp) {
		t.Errorf("InitPreOncallCanvas() = %v, want %v", canvasResp, *expectedCanvasResp)
	}
}

// Mock data for testing
var mockOncallTickets = pre_oncall.TickeInfotResponse{
	Data: []pre_oncall.Ticket{
		{
			TicketId:     "123",
			BusinessName: "Business A",
			Title:        "Issue with Service A",
			Reporter:     "John Doe",
			CreatedAt:    "2024-01-01 00:00:00",
			UpdatedAt:    "2024-01-01 00:00:00",
			Remarks:      "Urgent issue",
			GroupLink:    "http://example.com/group",
			TicketLink:   "http://example.com/ticket/123",
		},
		// ... Add more mock tickets if needed
	},
}

// TestInitRelatedTicketCanvas tests the InitRelatedTicketCanvas function
func TestInitRelatedTicketCanvas(t *testing.T) {
	ctx := context.Background()
	inputValue := make(map[string]string, 0)
	// Call the function with mock data
	canvasResp := InitRelatedTicketCanvas(ctx, mockOncallTickets, inputValue)

	// Build the expected response
	expectedComponents := []Component{}

	// The initial single select component as in the function
	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(CreateTicketOptionID, "Create Ticket")
	action := NewAction("submit")
	singleSelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, nil)
	expectedComponents = append(expectedComponents, singleSelect)

	// Add the ticket info components
	for _, ticket := range mockOncallTickets.Data {
		expectedComponents = append(expectedComponents, NewText("Related Ticket", "header"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Ticket id: %v", ticket.TicketId), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Bussiness Line: %v", ticket.BusinessName), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Ticket Title: %v", ticket.Title), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Reported by: %v", ticket.Reporter), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Assignee: %v", ticket.Assignee), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Create Time: %v", ticket.CreatedAt), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Update Time: %v", ticket.UpdatedAt), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Additional Info: %v", ticket.Remarks), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Group Link %v", ticket.GroupLink), "paragraph"))
		expectedComponents = append(expectedComponents, NewText(fmt.Sprintf("Ticket Link %v", ticket.TicketLink), "paragraph"))
	}

	expectedContent := newContent(expectedComponents)
	expectedCanvasResp := newCanvasReponse(*expectedContent)

	// Compare the actual response with the expected response
	if !reflect.DeepEqual(canvasResp, *expectedCanvasResp) {
		t.Errorf("InitRelatedTicketCanvas() = %v, want %v", canvasResp, *expectedCanvasResp)
	}
}

// Mock data for IntercomCanvasReceiver
func mockIntercomCanvasReceiver() IntercomCanvasReceiver {
	// Assuming Option and Component are defined types in your package
	option1 := Option{Text: "BizLine1"}
	option2 := Option{Text: "BizLine2"}
	component := IntercomComponent{
		ID:      BizLineSearchDropdownID,
		Options: []Option{option1, option2},
	}
	return IntercomCanvasReceiver{
		Content: IntercomContent{
			Components: []IntercomComponent{component},
		},
	}
}

func TestExtractBizlinesFromCurrentCanvas(t *testing.T) {
	ctx := context.Background()
	currentCanvas := mockIntercomCanvasReceiver()

	t.Run("Extract Business Lines", func(t *testing.T) {
		expected := []string{"BizLine1", "BizLine2"}
		got := extractBizlinesFromCurrentCanvas(ctx, currentCanvas)
		assert.Equal(t, expected, got)
	})

	t.Run("No Business Lines", func(t *testing.T) {
		// Modify the currentCanvas to have no business line options
		// and test if the function returns an empty slice
		currentCanvas.Content.Components[0].Options = []Option{}
		expected := []string{}
		got := extractBizlinesFromCurrentCanvas(ctx, currentCanvas)
		assert.Equal(t, expected, got)
	})

}

func TestExtractStackNamesFromCurrentCanvas(t *testing.T) {
	ctx := context.Background()

	t.Run("Extract Stack Names", func(t *testing.T) {
		currentCanvas := IntercomCanvasReceiver{
			Content: IntercomContent{
				Components: []IntercomComponent{
					{ID: StackSearchDropdownID, Options: []Option{{Text: "Stack1"}, {Text: "Stack2"}}},
				},
			},
		}

		expected := []string{"Stack1", "Stack2"}
		got := extractStackNamesFromCurrentCanvas(ctx, currentCanvas)
		assert.Equal(t, expected, got)
	})

	t.Run("No Stack Names", func(t *testing.T) {
		currentCanvas := IntercomCanvasReceiver{
			Content: IntercomContent{
				Components: []IntercomComponent{
					{ID: StackSearchDropdownID, Options: []Option{}},
				},
			},
		}

		expected := []string{}
		got := extractStackNamesFromCurrentCanvas(ctx, currentCanvas)
		assert.Equal(t, expected, got)
	})
}

func TestExtractRegionsFromCurrentCanvas(t *testing.T) {
	ctx := context.Background()

	t.Run("Extract Regions", func(t *testing.T) {
		currentCanvas := IntercomCanvasReceiver{
			Content: IntercomContent{
				Components: []IntercomComponent{
					{ID: RegionSearchDropdownID, Options: []Option{{Text: "Region1"}, {Text: "Region2"}}},
				},
			},
		}

		expected := []string{"Region1", "Region2"}
		got := extractRegionsFromCurrentCanvas(ctx, currentCanvas)
		assert.Equal(t, expected, got)
	})

	t.Run("No Regions", func(t *testing.T) {
		currentCanvas := IntercomCanvasReceiver{
			Content: IntercomContent{
				Components: []IntercomComponent{
					{ID: RegionSearchDropdownID, Options: []Option{}},
				},
			},
		}

		expected := []string{}
		got := extractRegionsFromCurrentCanvas(ctx, currentCanvas)
		assert.Equal(t, expected, got)
	})
}
