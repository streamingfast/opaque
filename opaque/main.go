package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"

	"github.com/streamingfast/logging"
	"github.com/streamingfast/opaque"
	"github.com/manifoldco/promptui"
)

var zlog = zap.NewNop()

func init() {
	if os.Getenv("DEBUG") != "" || os.Getenv("ZAP_PRETTY") != "" {
		zlog = logging.MustCreateLogger()
	}
}

func main() {
	flag.Parse()

	action := flag.Arg(0)
	args := flag.Args()
	inputs := args[1:len(args)]

	switch strings.ToLower(action) {
	case "decode":
		decode(inputs...)
	case "encode":
		encode(inputs...)
	default:
		printErrorAndExit("Invalid action, must be one of 'decode' or 'encode'", nil)
	}
}

func decode(values ...string) {
	zlog.Debug("Performing decoding")
	if len(values) <= 0 {
		zlog.Debug("Asking user input to decode")
		values = []string{mustAsk("Input ")}
	}

	zlog.Debug("Decoding values", zap.Int("count", len(values)))

	for _, value := range values {
		zlog.Debug("Decoding value", zap.String("value", value))
		decodedOpaque, err := opaque.FromOpaque(value)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Println(decodedOpaque)
	}
}

func encode(values ...string) {
	zlog.Debug("Performing encoding")

	if len(values) <= 0 {
		zlog.Debug("Asking user input to encode")
		values = []string{mustAsk("Input ")}
	}

	zlog.Debug("Encoding values", zap.Int("count", len(values)))

	for _, value := range values {
		zlog.Debug("Encoding value", zap.String("value", value))
		encodedOpaque, err := opaque.ToOpaque(value)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Println(encodedOpaque)
	}
}

func printErrorAndExit(message string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %s", err)
	} else {
		fmt.Fprintln(os.Stderr, message)
		fmt.Fprintln(os.Stderr, "")
		flag.Usage()
		os.Exit(1)
	}
}

func mustAsk(label string) string {
	return mustPrompt(label, nil)
}

func mustPrompt(label string, validator promptui.ValidateFunc) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validator,
	}

	result, err := prompt.Run()
	if err != nil {
		panic(fmt.Sprintf("prompt failed: %v", err))
	}

	return result
}
