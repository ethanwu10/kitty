// License: GPLv3 Copyright: 2022, Kovid Goyal, <kovid at kovidgoyal.net>

package completion

import (
	"fmt"
	"strings"

	"kitty/tools/utils"
)

var _ = fmt.Print

func bash_output_serializer(completions []*Completions, shell_state map[string]string) ([]byte, error) {
	output := strings.Builder{}
	for _, mg := range completions[0].Groups {
		mg.remove_common_prefix()
		for _, m := range mg.Matches {
			w := m.Word
			if !mg.NoTrailingSpace {
				w += " "
			}
			fmt.Fprintln(&output, "COMPREPLY+=("+utils.QuoteStringForSH(w)+")")
		}
	}
	return []byte(output.String()), nil
}

func bash_init_completions(completions *Completions) {
	completions.split_on_equals = true
}

func init() {
	input_parsers["bash"] = shell_input_parser
	output_serializers["bash"] = bash_output_serializer
	init_completions["bash"] = bash_init_completions
}