package preoncall_service

import (
	"context"
	"fmt"
	pre_oncall "github.com/digitalocean/sample-golang/pre_oncall_api"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"strings"
)

// getValuePtr is a helper function to get the pointer of the value
// To fill in the canvas fields with the previous input values
func getValuePtr(key string, selectedValues map[string]string) *string {
	var selectedValue *string
	if val, exist := selectedValues[key]; exist && val != "" {
		selectedValue = &val
	}
	//fmt.Printf("getValuePtr selectedValue %v \n", selectedValue)
	return selectedValue
}

// InitPreOncallCanvas is a constructor for Init CanvasReponse
func InitPreOncallCanvas() CanvasReponse {
	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(CreateTicketOptionID, "Create Ticket")
	action := NewAction("submit")
	singleSelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, nil)

	content := newContent([]Component{singleSelect})
	canvasResp := newCanvasReponse(*content)

	return *canvasResp
}

func InitRelatedTicketCanvas(ctx context.Context, oncallTickets pre_oncall.TickeInfotResponse, inputValues map[string]string) CanvasReponse {
	//log. := utils.Get//log.gerWithMethod(ctx, "InitRelatedTicketCanvas")
	//log..Infof("InitRelatedTicketCanvas oncallTickets %v", larkcore.Prettify(oncallTickets))

	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(CreateTicketOptionID, "Create Ticket")
	action := NewAction("submit")
	singleSelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, getValuePtr(CategorySingleSelectID, inputValues))

	components := []Component{}
	components = append(components, singleSelect)
	// TODO: Replace with the real ticket info

	for index, ticket := range oncallTickets.Data {
		ticketBannerTitle := NewText(fmt.Sprintf("Related Ticket %v", index+1), "header")
		components = append(components, ticketBannerTitle)

		ticketID := NewText(fmt.Sprintf("Ticket id: %v", ticket.TicketId), "paragraph")
		components = append(components, ticketID)

		bizLine := NewText(fmt.Sprintf("Bussiness Line: %v", ticket.BusinessName), "paragraph")
		components = append(components, bizLine)

		TicketTitle := NewText(fmt.Sprintf("Ticket Title: %v", ticket.Title), "paragraph")
		components = append(components, TicketTitle)

		ReportBy := NewText(fmt.Sprintf("Reported by: %v", ticket.Reporter), "paragraph")
		components = append(components, ReportBy)

		Assignee := NewText(fmt.Sprintf("Assignee: %v", ticket.Assignee), "paragraph")
		components = append(components, Assignee)

		CreateTime := NewText(fmt.Sprintf("Create Time: %v", ticket.CreatedAt), "paragraph")
		components = append(components, CreateTime)

		UpdateTime := NewText(fmt.Sprintf("Update Time: %v", ticket.UpdatedAt), "paragraph")
		components = append(components, UpdateTime)

		AdditionalInfo := NewText(fmt.Sprintf("Additional Info: %v", ticket.Remarks), "paragraph")
		components = append(components, AdditionalInfo)

		GroupLink := NewText(fmt.Sprintf("Group Link: %v", ticket.GroupLink), "paragraph")
		components = append(components, GroupLink)

		TicketLink := NewText(fmt.Sprintf("Ticket Link: %v", ticket.TicketLink), "paragraph")
		components = append(components, TicketLink)
	}

	content := newContent(components)
	canvasResp := newCanvasReponse(*content)

	return *canvasResp
}

// InitCreateOncalTicketCanvas is a constructor for Create Create-Ticket CanvasReponse
func InitCreateOncalTicketCanvas(bizLines []string, regions []string, stackNames []string, selectedValues map[string]string, ticketCreationStatus string) CanvasReponse {
	if selectedValues == nil {
		selectedValues = make(map[string]string)
	}

	// Intercom requires at least 2 options for dropdown to render
	for len(bizLines) < 2 {
		bizLines = append(bizLines, EmptyPlaceHolder)
	}

	for len(regions) < 2 {
		regions = append(regions, EmptyPlaceHolder)
	}

	for len(stackNames) < 2 {
		stackNames = append(stackNames, EmptyPlaceHolder)
	}

	//log..Infof("InitCreateOncalTicketCanvas selectedValues %v", selectedValues)
	//log..Infof("InitCreateOncalTicketCanvas bizLines %v, regions %v, stackNames %v", bizLines, regions, stackNames)
	//fmt.Printf("InitCreateOncalTicketCanvas bizLines %v, regions %v, stackNames %v \n", bizLines, regions, stackNames)
	fmt.Printf("InitCreateOncalTicketCanvas selectedValues %v \n", larkcore.Prettify(selectedValues))
	option1 := NewOption(RelatedTicketOptionID, "Related Ticket")
	option2 := NewOption(CreateTicketOptionID, "Create Ticket")
	action := NewAction("submit")

	categorySelect := NewSingleSelect(CategorySingleSelectID, "single-select", CategorySingleSelectLabel, []Option{*option1, *option2}, &action, getValuePtr(CategorySingleSelectID, selectedValues))

	// bizline
	bizLineSearchText := NewText("Business Line Search", "header")

	var bizLineSearchValue *string
	if val, exist := selectedValues[BizLineSearchInputID]; exist {
		bizLineSearchValue = &val
	}

	bizLineSearchInput := NewInput(BizLineSearchInputID, BizLineSearchLabel, "Enter input here", bizLineSearchValue)
	bizLineSearchBtn := NewButton(BizLineSearchButtonID, BizLineSearchButtonLabel, action, "primary", false)

	bizLineText := NewText("Business Line", "header")

	bizLineDropDownOptions := []Option{}
	for _, bizLine := range bizLines {
		//log..Infof("bizLine %v", bizLine)
		bizLineDropDownOptions = append(bizLineDropDownOptions, *NewOption(bizLine, bizLine))
	}

	bizLineSearchDropDown := NewDropdown(BizLineSearchDropdownID, BizLineSearchDropdownLabel, bizLineDropDownOptions, getValuePtr(BizLineSearchDropdownID, selectedValues))

	// ticket title
	ticketTitleText := NewText("Ticket Title", "header")

	ticketTitleInput := NewInput(TicketTitleInputID, TicketTitleLabel, "Briefly describe the problem", getValuePtr(TicketTitleInputID, selectedValues))

	// region search
	regionSearchText := NewText("Region Search", "header")
	regionSearchInput := NewInput(RegionSearchInputID, RegionSearchLabel, "Enter input here", getValuePtr(RegionSearchInputID, selectedValues))
	regionSearchBtn := NewButton(RegionSearchButtonID, RegionSearchButtonLabel, action, "primary", false)

	regionDropDownOptions := []Option{}
	for _, region := range regions {
		regionDropDownOptions = append(regionDropDownOptions, *NewOption(region, region))
	}

	regionSearchDropDown := NewDropdown(RegionSearchDropdownID, RegionSearchDropdownLabel, regionDropDownOptions, getValuePtr(RegionSearchDropdownID, selectedValues))

	// stack search
	stackSearchText := NewText("Stack", "header")

	stackButton := NewButton(StackSearchButtonID, StackSearchButtonLabel, action, "primary", false)

	stackDropDownOptions := []Option{}
	for _, stackOption := range stackNames {
		stackDropDownOptions = append(stackDropDownOptions, *NewOption(stackOption, stackOption))
	}

	stackSearchDropDown := NewDropdown(StackSearchDropdownID, StackSearchDropdownLabel, stackDropDownOptions, getValuePtr(StackSearchDropdownID, selectedValues))
	//
	// priority
	priorityText := NewText("Priority", "header")
	prioritySingleSelectOptions := []Option{}
	priorityList := []string{P0, P1, P2}
	for _, priority := range priorityList {
		prioritySingleSelectOptions = append(prioritySingleSelectOptions, *NewOption(priority, priority))
	}

	prioritySingleSelect := NewSingleSelect(PrioritySingleSelectID, "single-select", PrioritySingleSelectLabel, prioritySingleSelectOptions, nil, getValuePtr(PrioritySingleSelectID, selectedValues))

	// create group
	createGroupText := NewText("Create Group", "header")
	createGroupSingleSelectOptions := []Option{}
	createGroupList := []string{AutoCreateGroup, AssociateGroup, NotCreateGroup}
	for _, createGroup := range createGroupList {
		createGroupSingleSelectOptions = append(createGroupSingleSelectOptions, *NewOption(createGroup, createGroup))
	}

	createGroupSingleSelect := NewSingleSelect(CreateGroupSingleSelectID, "single-select", CreateGroupSingleSelectLabel, createGroupSingleSelectOptions, nil, getValuePtr(CreateGroupSingleSelectID, selectedValues))

	// user id

	userIDText := NewText("User ID", "header")
	userIDInput := NewInput(userIDInputID, userIDInputLabel, "type in user id", getValuePtr(userIDInputID, selectedValues))

	// tenant id
	tenantIDText := NewText("Tenant ID", "header")
	tenantIDInput := NewInput(tenantIDInputID, tenantIDInputLabel, "type in tenant id", getValuePtr(tenantIDInputID, selectedValues))

	// lark version
	larkVersionText := NewText("Lark Version", "header")
	larkVersionInput := NewInput(LarkVersionInputID, LarkVersionInputLabel, "type in lark version", getValuePtr(LarkVersionInputID, selectedValues))

	// Create button to submit ticket
	submitTicketBtn := NewButton(SubmitTicketButtonID, SubmitTicketLabel, action, "primary", false)

	components := []Component{categorySelect, bizLineSearchText, bizLineSearchInput, bizLineSearchBtn, bizLineText,
		bizLineSearchDropDown, ticketTitleText, ticketTitleInput, regionSearchText, regionSearchInput, regionSearchBtn, regionSearchDropDown,
		stackSearchText, stackButton, stackSearchDropDown, priorityText, prioritySingleSelect,
		createGroupText, createGroupSingleSelect, userIDText, userIDInput,
		tenantIDText, tenantIDInput,
		larkVersionText, larkVersionInput,
		submitTicketBtn}

	switch ticketCreationStatus {
	case WaitingUserInput:
	case CreatTicketFailed:
		failedText := NewText("Create ticket failed, please check your input and try again", "header")
		components = append(components, failedText)
	case CreateTicketSucceed:
		succeedText := NewText("Create ticket succeed, please do not submit again", "header")
		components = append(components, succeedText)
	}

	content := newContent(components)

	//content := newContent([]Component{categorySelect, bizLineText, bizLineSearchInput, bizLineSearchBtn, bizLineSearchDropDown})

	//content := newContent([]Component{singleSelect, bizLineText, bizLineSearchInput, bizLineSearchBtn, bizLineSearchDropDown})
	canvasResp := newCanvasReponse(*content)
	//fmt.Println(" InitCreateOncalTicketCanvas canvasResp %v", larkcore.Prettify(canvasResp))
	return *canvasResp
}

func GetInitTicketCanvasBody() CanvasReponse {
	return InitPreOncallCanvas()
}

func GetRelatedTicketCanvasBody(ctx context.Context, inputValue map[string]string, intercomConversationID string) CanvasReponse {
	//log. := utils.Get//log.gerWithMethod(ctx, "GetRelatedTicketCanvasBody")
	// We use the intercomConversationID to get the tickets external id via pre-oncall api
	//log..Infof("GetRelatedTicketCanvasBody intercomConversation %v", intercomConversationID)

	// TODO we use the intercomConversation to get the tickets via pre-oncall api
	oncallTickets, err := pre_oncall.GetPreOncallTicket(ctx, intercomConversationID, "intercom")
	if err != nil {
		//log..Errorf("GetRelatedTicketCanvasBody GetPreOncallTicket err %v", err)
		return InitPreOncallCanvas()
	}

	return InitRelatedTicketCanvas(ctx, oncallTickets, inputValue)
}

func searchBusinessLine(ctx context.Context, keyword string, bizLines []pre_oncall.Business) []string {
	//log. := utils.Get//log.gerWithMethod(ctx, "searchBusinessLine")
	//log..Infof("searchBusinessLine keyword %v", keyword)
	result := make([]string, 0)
	for _, biz := range bizLines {
		if strings.Contains(strings.ToLower(biz.Name), strings.ToLower(keyword)) || keyword == "" {
			result = append(result, biz.Name)
		}
	}

	//log..Infof("searchBusinessLine result %v", result)
	return result
}

func searchRegion(ctx context.Context, keyword string, regions []pre_oncall.CodeNamePair) []string {
	//log. := utils.Get//log.gerWithMethod(ctx, "searchRegion")
	//log..Infof("searchRegion keyword %v", keyword)
	fmt.Printf("searchRegion keyword %v \n", keyword)
	fmt.Printf("searchRegion regions %v \n", larkcore.Prettify(regions))
	result := make([]string, 0)
	for _, region := range regions {
		if strings.Contains(strings.ToLower(region.Name), strings.ToLower(keyword)) || keyword == "" {
			result = append(result, region.Name)
		}
	}

	//log..Infof("searchRegion result %v", result)
	return result
}

func extractBizlinesFromCurrentCanvas(ctx context.Context, currentCanvas IntercomCanvasReceiver) []string {
	//log. := utils.Get//log.gerWithMethod(ctx, "extractBizlinesFromCurrentCanvas")
	//log..Infof("extractBizlinesFromCurrentCanvas currentCanvas %v", larkcore.Prettify(currentCanvas))
	bizLines := make([]string, 0)
	for _, component := range currentCanvas.Content.Components {
		if component.ID == BizLineSearchDropdownID {
			//log..Infof("extractBizlinesFromCurrentCanvas find bizline dropdown %v", larkcore.Prettify(component))
			existingOptions := component.Options
			for _, option := range existingOptions {
				bizLines = append(bizLines, option.Text)
			}
		}
	}

	//log..Infof("extractBizlinesFromCurrentCanvas bizLines %v", bizLines)

	return bizLines
}

func extractStackNamesFromCurrentCanvas(ctx context.Context, currentCanvas IntercomCanvasReceiver) []string {
	//log. := utils.Get//log.gerWithMethod(ctx, "extractStackNamesFromCurrentCanvas")
	//log..Infof("extractStackNamesFromCurrentCanvas currentCanvas %v", larkcore.Prettify(currentCanvas))
	stackNames := make([]string, 0)
	for _, component := range currentCanvas.Content.Components {
		if component.ID == StackSearchDropdownID {
			//log..Infof("extractStackNamesFromCurrentCanvas find stack dropdown %v", larkcore.Prettify(component))
			existingOptions := component.Options
			for _, option := range existingOptions {
				stackNames = append(stackNames, option.Text)
			}
		}
	}

	//log..Infof("extractStackNamesFromCurrentCanvas stackNames %v", stackNames)

	return stackNames
}

func extractRegionsFromCurrentCanvas(ctx context.Context, currentCanvas IntercomCanvasReceiver) []string {
	//log. := utils.Get//log.gerWithMethod(ctx, "extractRegionsFromCurrentCanvas")
	//log..Infof("extractRegionsFromCurrentCanvas currentCanvas %v", larkcore.Prettify(currentCanvas))
	regions := make([]string, 0)
	for _, component := range currentCanvas.Content.Components {
		if component.ID == RegionSearchDropdownID {
			//log..Infof("extractRegionsFromCurrentCanvas find region dropdown %v", larkcore.Prettify(component))

			existingOptions := component.Options
			for _, option := range existingOptions {
				regions = append(regions, option.Text)
			}
		}
	}

	//log..Infof("extractRegionsFromCurrentCanvas regions %v", regions)
	return regions
}

func GetCreateTicketCanvasBody(ctx context.Context, inputValues map[string]string, intercomConversationID string, assigneeID int, buttonClick string, currentCanvas IntercomCanvasReceiver) CanvasReponse {
	//log. := utils.Get//log.gerWithMethod(ctx, "GetCreateTicketCanvasBody")
	//log..Infof("GetCreateTicketCanvasBody buttonClick %v, selectedValue %v, intercom convID %v, assigneeID %v, canvas %v", buttonClick, inputValues, intercomConversationID, assigneeID, larkcore.Prettify(currentCanvas))

	//metaInfoResp, err := pre_oncall.GetPreOncallMetaInfo(ctx, true, true)
	//if err != nil {
	//	//log..Errorf("GetCreateTicketCanvasBody err %v", err)
	//	return InitPreOncallCanvas()
	//}

	ticketStatus := WaitingUserInput

	// extract existing values
	bizLines := extractBizlinesFromCurrentCanvas(ctx, currentCanvas)
	regions := extractRegionsFromCurrentCanvas(ctx, currentCanvas)
	stackNames := extractStackNamesFromCurrentCanvas(ctx, currentCanvas)

	//fmt.Printf("GetCreateTicketCanvasBody buttonClick %v, selectedValue %v, intercom convID %v, assigneeID %v, canvas %v \n", buttonClick, inputValues, intercomConversationID, assigneeID, larkcore.Prettify(currentCanvas))

	//fmt.Printf("GetCreateTicketCanvasBody bizLines %v, regions %v, stackNames %v \n", bizLines, regions, stackNames)

	switch buttonClick {
	case CreateTicketOptionID:
		//log..Infof("GetCreateTicketCanvasBody create ticket option")
		resp, err := pre_oncall.GetPreOncallMetaInfo(ctx, true, true)
		if err != nil {
			//log..Errorf("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v", err)
			return InitPreOncallCanvas()
		}

		bizLines = make([]string, 0)
		regions = make([]string, 0)
		stackNames = make([]string, 0)

		//log..Infof("GetCreateTicketCanvasBody resp %v", larkcore.Prettify(resp))
		//fmt.Printf("GetCreateTicketCanvasBody resp %v \n", larkcore.Prettify(resp))
		bizList := resp.Data.BusinessList
		for _, biz := range bizList {
			bizLines = append(bizLines, biz.Name)
		}
		//log..Infof("GetCreateTicketCanvasBody bizLines %v", bizLines)

		regionList := resp.Data.RegionList
		fmt.Printf("GetCreateTicketCanvasBody regionList %v \n", larkcore.Prettify(regionList))
		for _, region := range regionList {
			fmt.Printf("GetCreateTicketCanvasBody region %v \n", larkcore.Prettify(region))
			regions = append(regions, region.Name)
		}

		//stackNames = append(stackNames, "11")
		//stackNames = append(stackNames, "22")
		//stackNames = append(stackNames, "33")

		//log..Infof("GetCreateTicketCanvasBody regions %v", regions)

		larkVersion := ""
		tenantID := ""
		userID := ""

		//convSrv := service.GetLarkAssistantConversationRecordService()
		//conversation, err := convSrv.GetConversationByIntercomConversationIDWithRetry(ctx, intercomConversationID)
		//
		//if err == nil {
		//	larkConverstionID := conversation.LarkConversationID
		//	//log..Infof("GetCreateTicketCanvasBody find larkConverstionID %v", larkConverstionID)
		//
		//	// get user info related to this conversation
		//	userSrv := service.GetLarkAssistantUserRecordService()
		//	user, userErr := userSrv.GetLarkAssistantUserByConversationIDWithRetry(ctx, larkConverstionID)
		//
		//	if userErr == nil {
		//		//log..Infof("GetCreateTicketCanvasBody find user %v", larkcore.Prettify(user))
		//		// Prefill the user id and tenant id
		//		tenantID = user.TenantID
		//		userID = user.UserID
		//		userOpenID := user.UserOpenID
		//
		//		// Begin to fetch the lark version
		//		redisCli := cache.GetClient()
		//		segmentID, cacheErr := intercom.GetSegmentID(ctx, redisCli, userOpenID)
		//
		//		if cacheErr == nil {
		//			segSrv := service.GetLarkAssistantSegmentInfoRecordService()
		//			//log..Infof("GetCreateTicketCanvasBody find segment info by segmentID %v and userOpenID %v", segmentID)
		//			segmentInfo, segErr := segSrv.GetBySegmentIDOrUserID(ctx, segmentID, userOpenID)
		//			if segmentInfo != nil && segErr == nil {
		//				// Prefill the lark version
		//				larkVersion = segmentInfo.ExtraInfo.Version
		//				//log..Infof("GetCreateTicketCanvasBody find lark version %v", larkVersion)
		//			}
		//		}
		//
		//	}
		//
		//}
		inputValues[userIDInputID] = userID
		inputValues[tenantIDInputID] = tenantID
		inputValues[LarkVersionInputID] = larkVersion
	case BizLineSearchButtonID:
		var bizLineSearchKeyword string
		if v, ok := inputValues[BizLineSearchInputID]; ok {
			bizLineSearchKeyword = v
		}

		resp, err := pre_oncall.GetPreOncallMetaInfo(ctx, true, true)
		if err != nil {
			//log..Errorf("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v", err)
			return InitPreOncallCanvas()
		}
		bussinessList := resp.Data.BusinessList
		bizLines = searchBusinessLine(ctx, bizLineSearchKeyword, bussinessList)
		inputValues[BizLineSearchDropdownID] = ""
	case RegionSearchButtonID:
		var regionSearchKeyword string
		if v, ok := inputValues[RegionSearchInputID]; ok {
			regionSearchKeyword = v
		}

		resp, err := pre_oncall.GetPreOncallMetaInfo(ctx, true, true)
		if err != nil {
			//log..Errorf("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v", err)
			return InitPreOncallCanvas()
		}
		regionList := resp.Data.RegionList
		regions = searchRegion(ctx, regionSearchKeyword, regionList)
		inputValues[RegionSearchDropdownID] = ""
	case SubmitTicketButtonID:
		ticket, isValid := validSubmitForm(ctx, inputValues, intercomConversationID)
		fmt.Printf("GetCreateTicketCanvasBody validSubmitForm ticket %v \n", larkcore.Prettify(ticket))
		if !isValid {
			fmt.Printf("GetCreateTicketCanvasBody validSubmitForm failed \n")
			ticketStatus = CreatTicketFailed
		} else {
			fmt.Printf("GetCreateTicketCanvasBody validSubmitForm success \n")
			resp, err := pre_oncall.SubmitPreOncallTicket(ctx, ticket)
			fmt.Printf("GetCreateTicketCanvasBody SubmitPreOncallTicket resp %v \n", larkcore.Prettify(resp))
			if err != nil {
				//log..Errorf("GetCreateTicketCanvasBody SubmitPreOncallTicket err %v", err)
				fmt.Printf("GetCreateTicketCanvasBody SubmitPreOncallTicket err %v \n", err)
				ticketStatus = CreatTicketFailed
			} else {
				if resp.Data.TicketId != "" {
					ticketStatus = CreateTicketSucceed
				} else {
					ticketStatus = CreatTicketFailed
				}
			}
		}
	case StackSearchButtonID:
		//log..Infof("GetCreateTicketCanvasBody stack search button")
		fmt.Printf("GetCreateTicketCanvasBody stack search button \n")
		stackNames = make([]string, 0)
		inputValues[StackSearchDropdownID] = ""
		if value, ok := inputValues[BizLineSearchDropdownID]; ok {
			resp, err := pre_oncall.GetPreOncallMetaInfo(ctx, true, true)
			if err != nil {
				//log..Errorf("GetCreateTicketCanvasBody GetPreOncallMetaInfo err %v", err)
				return InitPreOncallCanvas()
			}

			bizName := value
			fmt.Printf("Get bizName %v when doing stack search \n", bizName)

			bussinessList := resp.Data.BusinessList
			for _, biz := range bussinessList {
				if biz.Name == bizName {
					stackNames = biz.Stacks
				}
			}

			fmt.Printf("Search results stackNames is %v \n", stackNames)
		}

		fmt.Printf("GetCreateTicketCanvasBody stackNames %v \n", stackNames)
	}

	return InitCreateOncalTicketCanvas(bizLines, regions, stackNames, inputValues, ticketStatus)
}

func validSubmitForm(ctx context.Context, inputValues map[string]string, intercomConversationID string) (pre_oncall.TicketSubmitRequest, bool) {
	//log. := utils.Get//log.gerWithMethod(ctx, "validSubmitForm")
	//log..Infof("validSubmitForm inputValues %v", inputValues)
	ticket := pre_oncall.TicketSubmitRequest{
		Title:         "",
		Business:      "",
		Priority:      "",
		Stack:         "",
		Region:        "",
		UserId:        "",
		Version:       "",
		CreateChatWay: "",
		// prefill
		Type:        "Pre-Oncall",
		App:         "Lark",
		External:    "字节外",
		Source:      "客服渠道",
		Reporter:    "test@larksuite.com",
		Remarks:     " ",
		ChannelType: "intercom",
		BizTicketId: intercomConversationID,
	}

	if inputValues == nil {
		return ticket, false
	}

	// Check Biz Line
	if val, ok := inputValues[BizLineSearchDropdownID]; !ok || val == "" {
		return ticket, false
	}

	fmt.Printf("validSubmitForm BizLineSearchDropdownID %v \n", inputValues[BizLineSearchDropdownID])

	resp, err := pre_oncall.GetPreOncallMetaInfo(ctx, true, true)
	if err != nil {
		//log..Errorf("validSubmitForm GetPreOncallMetaInfo err %v", err)
		return ticket, false
	}

	bizList := resp.Data.BusinessList
	for _, biz := range bizList {
		if biz.Name == inputValues[BizLineSearchDropdownID] {
			ticket.Business = biz.Bid
			break
		}
	}

	// Check Ticket Title
	if val, ok := inputValues[TicketTitleInputID]; !ok || val == "" {
		return ticket, false
	}
	ticket.Title = inputValues[TicketTitleInputID]

	// Check Region
	if val, ok := inputValues[RegionSearchDropdownID]; !ok || val == "" {
		return ticket, false
	}
	ticket.Region = inputValues[RegionSearchDropdownID]

	// Check Stack
	if val, ok := inputValues[StackSearchDropdownID]; !ok || val == "" {
		return ticket, false
	}
	ticket.Stack = chooseFormValue(inputValues[StackSearchDropdownID] == EmptyPlaceHolder, "", inputValues[StackSearchDropdownID])

	// Check Priority
	if val, ok := inputValues[PrioritySingleSelectID]; !ok || val == "" {
		return ticket, false
	}
	ticket.Priority = inputValues[PrioritySingleSelectID]

	// Check Create Group
	if val, ok := inputValues[CreateGroupSingleSelectID]; !ok || val == "" {
		return ticket, false
	}
	ticket.CreateChatWay = inputValues[CreateGroupSingleSelectID]

	// Check User ID
	if val, ok := inputValues[userIDInputID]; !ok || val == "" {
		return ticket, false
	}
	ticket.UserId = inputValues[userIDInputID]

	// Check Tenant ID
	if val, ok := inputValues[tenantIDInputID]; !ok || val == "" {
		return ticket, false
	}

	// Check Lark Version
	if val, ok := inputValues[LarkVersionInputID]; !ok || val == "" {
		return ticket, false
	}
	ticket.Version = inputValues[LarkVersionInputID]

	fmt.Printf("validSubmitForm ticket ========== %v \n", larkcore.Prettify(ticket))
	return ticket, true
}

// Ternary-like function
func chooseFormValue(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}
