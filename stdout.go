package gotils

import (
	"encoding/json"
	"os"
)

// PrettyPrint prints the given value to stdout in a pretty format. Intended to be used with structs and maps.
func PrettyPrint(v interface{}) {
	prettyPrint(v)
}

func prettyPrint(v interface{}) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	encoder.Encode(v)
}
