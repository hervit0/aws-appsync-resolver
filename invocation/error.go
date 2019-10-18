package invocation

import (
	"encoding/json"
	"fmt"
)

func EmptyPayloadError(arguments *json.RawMessage, source *json.RawMessage) error {
	return fmt.Errorf("cannot parse payload. Pass at least arguments or source in context.\n no arguments received: %v\n no source received: %v", arguments, source)
}
