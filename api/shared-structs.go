package main

// Effect struct for Equipment
type Effect struct {
	Text  string `json:"text,omitempty" bson:"text,omitempty"`
	Value int    `json:"value,omitempty" bson:"value,omitempty"`
}
