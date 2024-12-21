package build

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Build() error {
	fmt.Println("Building the WebAssembly module...")

	err := sh.RunWith(map[string]string{
		"GOOS":   "js",
		"GOARCH": "wasm",
	}, "go", "build", "-o", "web/js/main.wasm", "internal/wasm/wasm.go")

	if err != nil {
		return fmt.Errorf("failed to build: %w", err)
	}

	fmt.Println("Build successful!")
	return nil
}

func Clean() error {
	fmt.Println("Cleaning up...")
	if err := sh.Rm("web/js/main.wasm"); err != nil {
		return fmt.Errorf("failed to clean: %w", err)
	}
	fmt.Println("Clean successful!")
	return nil
}
