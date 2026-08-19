package zhengtest0817142_rule

import "fmt"

func RunTestProgram(name string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("hello, %s", name)
}
