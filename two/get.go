package two

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
)

func Get() string {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger.Log("two", "Get")
	return fmt.Sprintf("%d", 2)
}
