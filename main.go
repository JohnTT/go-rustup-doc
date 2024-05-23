package main

import (
	"flag"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		log.Fatalf("unsupported platform")
		return
	}

	log.Printf("Running command '%s'", cmd)
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}

func runCommand(command string, args ...string) error {
	// Create the command
	cmd := exec.Command(command, args...)
	log.Printf("Running command '%s'", cmd)

	// Get the output of the command
	output, err := cmd.CombinedOutput()
	stdout := strings.TrimRight(string(output), "\r\n")
	log.Printf("Output '%s'", stdout)
	if err != nil {
		return err
	}

	// Create an HTTP handler function
	index := strings.TrimRight(string(output), "\r\n")

	// Define the directory to be served
	dir := filepath.Dir(index)

	// Create a file server handler
	fs := http.FileServer(http.Dir(dir))

	// Register the file server handler
	http.Handle("/", fs)

	// Define the port to listen on
	port := ":8000"

	// Start the HTTP server in a separate goroutine
	go func() {
		log.Printf("Starting server on %s...\n", port)
		err = http.ListenAndServe(port, nil)
		if err != nil {
			panic(err)
		}
	}()

	openBrowser("http://localhost" + port)
	return nil
}

func main() {
	// Define command-line flags
	toolchain := flag.String("toolchain", "", "Toolchain name, such as 'stable', 'nightly', or '1.8.0'")
	alloc := flag.Bool("alloc", false, "The Rust core allocation and collections library")
	book := flag.Bool("book", false, "The Rust Programming Language book")
	cargo := flag.Bool("cargo", false, "The Cargo Book")
	core := flag.Bool("core", false, "The Rust Core Library")
	editionGuide := flag.Bool("edition-guide", false, "The Rust Edition Guide")
	nomicon := flag.Bool("nomicon", false, "The Dark Arts of Advanced and Unsafe Rust Programming")
	procMacro := flag.Bool("proc_macro", false, "A support library for macro authors when defining new macros")
	reference := flag.Bool("reference", false, "The Rust Reference")
	rustByExample := flag.Bool("rust-by-example", false, "A collection of runnable examples that illustrate various Rust concepts and standard libraries")
	rustc := flag.Bool("rustc", false, "The compiler for the Rust programming language")
	rustdoc := flag.Bool("rustdoc", false, "Documentation generator for Rust projects")
	std := flag.Bool("std", false, "Standard library API documentation")
	test := flag.Bool("test", false, "Support code for rustc's built in unit-test and micro-benchmarking framework")
	unstableBook := flag.Bool("unstable-book", false, "The Unstable Book")
	embeddedBook := flag.Bool("embedded-book", false, "The Embedded Rust Book")

	// Parse the flags
	flag.Parse()

	// Build the rustup doc command
	args := []string{"doc"}

	if *toolchain != "" {
		args = append(args, "--toolchain", *toolchain)
	}
	if *alloc {
		args = append(args, "--alloc")
	}
	if *book {
		args = append(args, "--book")
	}
	if *cargo {
		args = append(args, "--cargo")
	}
	if *core {
		args = append(args, "--core")
	}
	if *editionGuide {
		args = append(args, "--edition-guide")
	}
	if *nomicon {
		args = append(args, "--nomicon")
	}
	if *procMacro {
		args = append(args, "--proc_macro")
	}
	if *reference {
		args = append(args, "--reference")
	}
	if *rustByExample {
		args = append(args, "--rust-by-example")
	}
	if *rustc {
		args = append(args, "--rustc")
	}
	if *rustdoc {
		args = append(args, "--rustdoc")
	}
	if *std {
		args = append(args, "--std")
	}
	if *test {
		args = append(args, "--test")
	}
	if *unstableBook {
		args = append(args, "--unstable-book")
	}
	if *embeddedBook {
		args = append(args, "--embedded-book")
	}
	args = append(args, "--path")

	// Run the rustup doc command
	err := runCommand("rustup", args...)
	if err != nil {
		panic(err)
	}
	select {}
}
