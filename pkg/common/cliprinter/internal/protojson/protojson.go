package protojson

import (
	"encoding/json"
	"io"

	"google.golang.org/protobuf/proto"
)

// Print prints one or more protobuf messages formatted as JSON
func Print(msgs []proto.Message, stdout, stderr io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// Unfortunately, we can only marshal one message at a time, so
// we need to build up an array of marshaled messages. We do this
// before printing them to reduce our chances of printing an
// unterminated result

func parseJSONMessages(jms []json.RawMessage) ([]json.RawMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseJSONMessage(jm json.RawMessage) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

func removeNulls(jsonMap map[string]any) { _ = "STUB: not implemented"; return }

func removeNullsFromSlice(slice []any) []any { _ = "STUB: not implemented"; return nil }
