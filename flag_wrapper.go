package main

import (
	"fmt"
)

func getFlagValue(flagName string) (string, error) {
	// This function would typically interact with a feature management service
	// to retrieve the value of the specified flag.
	// For the "ShowMessage" flag, we assume the value is always true.
	switch flagName {
	case "ShowMessage":
		return "true", nil
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