package main

import (
	"fmt"
)

func getFlagValue(flagName string) (string, error) {
	switch flagName {
	case "Message":
		return flags.Message.GetValue(nil), nil
	case "FontColor":
		return flags.FontColor.GetValue(nil), nil
	case "FontSize":
		return fmt.Sprintf("%d", flags.FontSize.GetValue(nil)), nil
	default:
		return "", fmt.Errorf("flag %s not found", flagName)
	}
}