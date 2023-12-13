package main

import (
	"fmt"
)

// Component interface for all UI components.
type Component interface {
	Render() string
}

// Action defines the action to be taken when the button is clicked.
type Action struct {
	Type string `json:"type"` // Can be "submit", "URL", or "sheet"
	// Additional fields can be added here based on the action's requirements
}

// NewAction is a constructor for Action
func NewAction(actionType string) Action {
	return Action{Type: actionType}
}

// Button represents a UI button component.
type Button struct {
	Type     string `json:"type"`
	ID       string `json:"id"`
	Label    string `json:"label"`
	Action   Action `json:"action"`
	Style    string `json:"style,omitempty"`    // Primary, Secondary, Link
	Disabled bool   `json:"disabled,omitempty"` // Default is false
}

// NewButton creates a new button with the given parameters.
func NewButton(id, label string, action Action, style string, disabled bool) *Button {
	return &Button{
		Type:     "button",
		ID:       id,
		Label:    label,
		Action:   action,
		Style:    style,
		Disabled: disabled,
	}
}

// Render method for Button
func (b *Button) Render() string {
	return fmt.Sprintf("Button ID: %s, Label: %s, Action: %s, Style: %s, Disabled: %v", b.ID, b.Label, b.Action.Type, b.Style, b.Disabled)
}

// Text component
type Text struct {
	Type  string `json:"type"`
	Text  string
	Style string
}

// NewText is a constructor for Text
func NewText(text, style string) *Text {
	return &Text{Type: "text", Text: text, Style: style}
}

// Render method for Text
func (t *Text) Render() string {
	return fmt.Sprintf("Text: %s, Style: %s", t.Text, t.Style)
}

// Input component
type Input struct {
	Type        string
	ID          string
	Label       string
	Placeholder string
}

// NewInput is a constructor for Input
func NewInput(id, label, placeholder string) *Input {
	return &Input{Type: "input", ID: id, Label: label, Placeholder: placeholder}
}

// TextArea component
type TextArea struct {
	Type        string
	ID          string
	Label       string
	Placeholder string
}

// NewTextArea is a constructor for TextArea
func NewTextArea(id, label, placeholder string) *TextArea {
	return &TextArea{Type: "textarea", ID: id, Label: label, Placeholder: placeholder}
}

// Render method for TextArea
func (ta *TextArea) Render() string {
	return fmt.Sprintf("TextArea ID: %s, Label: %s, Placeholder: %s", ta.ID, ta.Label, ta.Placeholder)
}

// Option for Dropdown and SingleSelect
type Option struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

// NewOption is a constructor for Option
func NewOption(id, text string) *Option {
	return &Option{Type: "option", ID: id, Text: text}
}

// Dropdown component
type Dropdown struct {
	Type    string   `json:"type"`
	ID      string   `json:"id"`
	Label   string   `json:"label"`
	Options []Option `json:"options"`
}

// NewDropdown is a constructor for Dropdown
func NewDropdown(id, label string, options []Option) *Dropdown {
	return &Dropdown{Type: "dropdown", ID: id, Label: label, Options: options}
}

// Render method for Dropdown
func (d *Dropdown) Render() string {
	return fmt.Sprintf("Dropdown ID: %s, Label: %s, Options: %v", d.ID, d.Label, d.Options)
}

// SingleSelect component
type SingleSelect struct {
	Type    string   `json:"type"`
	ID      string   `json:"id"`
	Label   string   `json:"label"`
	Options []Option `json:"options"`
	Action  Action   `json:"action"`
}

func NewSingleSelect(id, selectType, label string, options []Option, action Action) *SingleSelect {
	return &SingleSelect{Type: selectType, ID: id, Label: label, Options: options, Action: action}
}

// Render method for SingleSelect
func (ss *SingleSelect) Render() string {
	return fmt.Sprintf("SingleSelect ID: %s, Label: %s, Options: %v", ss.ID, ss.Label, ss.Options)
}

// Spacer component
type Spacer struct {
	Type string
	Size string
}

func (s Spacer) Render() string {
	return ""
}

// NewSpacer is a constructor for Spacer
func NewSpacer(size string) *Spacer {
	return &Spacer{Type: "spacer", Size: size}
}

// NewCanvas is a constructor for Canvas
func Newcontent(components []Component) *Content {
	return &Content{Components: components}
}

// AddComponent adds a component to the canvas
func (c *Content) AddComponent(component Component) {
	c.Components = append(c.Components, component)
}

// Content represents the content field within canvas.
type Content struct {
	Components []Component `json:"components"`
}

func newContent(components []Component) *Content {
	return &Content{Components: components}
}

// Canvas represents the top-level canvas field in your JSON.
type Canvas struct {
	Content Content `json:"content"`
}

func newCanvas(content Content) *Canvas {
	return &Canvas{Content: content}
}

// Root structure to encapsulate the Canvas
type CanvasReponse struct {
	Canvas Canvas `json:"canvas"`
}

func newCanvasReponse(content Content) *CanvasReponse {
	canvas := newCanvas(content)
	return &CanvasReponse{Canvas: *canvas}
}

//func CreateDemoCanvas() *CanvasReponse {
//	// Creating components using constructor functions
//	text := NewText("*Create a ticket*", "header")
//	input := NewInput("title", "Title", "Enter a title for your issue...")
//	textArea := NewTextArea("description", "Description", "Enter a description of the issue...")
//	option1 := NewOption("bug", "Bug")
//	option2 := NewOption("feedback", "Feedback")
//	dropdown := NewDropdown("label", "Label", []Option{*option1, *option2})
//	option3 := NewOption("low", "Low")
//	option4 := NewOption("medium", "Medium")
//	option5 := NewOption("high", "High")
//	singleSelect := NewSingleSelect("priority", "Priority", []Option{*option3, *option4, *option5})
//	spacer := NewSpacer("s")
//	action := NewAction("submit")
//	button := NewButton("submit", "Submit", action, "primary", false)
//
//	// Creating a canvas and adding components
//	content := newContent([]Component{text, input, textArea, dropdown, singleSelect, spacer, button})
//	canvasResp := newCanvasReponse(*content)
//
//	// Marshalling struct back to JSON
//	marshalledData, err := json.MarshalIndent(canvasResp, "", "  ")
//	if err != nil {
//		fmt.Println(err)
//		return nil
//	}
//
//	fmt.Println("Marshalled Data:")
//	fmt.Println(string(marshalledData))
//	// Optionally, marshal to JSON for demonstration
//	return canvasResp
//}

func InitPreOncallCanvas() CanvasReponse {
	option1 := NewOption(RelatedTicketID, "Related Ticket")
	option2 := NewOption(SubmitTicketID, "Create Ticket")
	action := NewAction("submit")
	singleSelect := NewSingleSelect("pre-oncall-ticket-option", "single-select", "Pre-Oncall Ticket", []Option{*option1, *option2}, action)

	content := newContent([]Component{singleSelect})
	canvasResp := newCanvasReponse(*content)

	return *canvasResp
}
