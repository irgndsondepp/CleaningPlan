package cleaningplan

import (
	"fmt"
)

func BuildHtml(content string) string {
	return fmt.Sprintf("<html>%v</html>", content)
}
