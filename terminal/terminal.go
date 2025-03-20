package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func PrintEntities(entities []any) {
	if len(entities) == 0 {
		fmt.Println("No entities found")
		return
	}

	for _, entity := range entities {
		parsed := entity.(map[string]any)

		for key, value := range parsed {
			if strings.HasPrefix(key, "@") {
				continue
			}

			switch v := value.(type) {
			case string:
				println(key, v)
			case float64:
				println(key, v)
			case bool:
				println(key, v)
			default:
				println(key, v)
			}
		}

		println("\n")
	}
}

func Clear() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}
