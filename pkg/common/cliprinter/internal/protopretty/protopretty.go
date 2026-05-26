package protopretty

import (
	"io"

	"google.golang.org/protobuf/proto"
)

func Print(msgs []proto.Message, stdout, _ io.Writer) error { _ = "STUB: not implemented"; return nil }
