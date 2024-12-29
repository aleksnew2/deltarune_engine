package console

import (
	"deltarune_engine/io/console/message"
	"fmt"
)

type Console struct {
	Settings Settings
}

func (c Console) Log(args ...string) {
	fmt.Println(format(c.Settings, message.Info, args...))
}

func (c Console) Error(error_code int32, args ...string) {
	fmt.Println(formatException(c.Settings, error_code, args...))
}

func (c Console) Warning(error_code int32, args ...string) {
	fmt.Println(format(c.Settings, message.Warning, args...))
}
