package invocation

import (
	"fmt"
)

func EmptyPayloadError(arguments *[]byte, source *[]byte) error {
	return fmt.Errorf("cannot parse payload. Pass at least arguments or source in context.\n no arguments received: %v\n no source received: %v", arguments, source)
}
