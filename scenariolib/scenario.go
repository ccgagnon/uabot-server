package scenariolib

// Scenario Represents one visit to the search
// Name A name given to the scenario for easier logging
// Weight A Weight for randomizing scenarios
// Events An array of actions the user will take
type Scenario struct {
	Name   string  `json:"name"`
	Weight int     `json:"weight"`
	Events []Event `json:"events"`
}

// Event An action taken by the user such as a search, a click, a SearchAndClick, etc.
// Type A string describing the type of event
// Arguments An array of the arguments to the event, specific to the type of event.
type Event struct {
	Type      string                 `json:"type"`
	Arguments map[string]interface{} `json:"arguments"`
}