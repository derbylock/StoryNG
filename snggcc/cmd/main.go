package main

import (
	"flag"
	"fmt"
	"github.com/derbylock/snggcc/pkg/model/github.com/derbylock/snggcc"
	"github.com/derbylock/snggcc/pkg/repository/mem"
	"github.com/goccy/go-yaml"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//go:generate protoc -I=../proto --go_out=../pkg/model ../proto/sng-story.proto

// Config holds the configuration from command line flags
type Config struct {
	rootDir    string
	outPath    string
	verbose    bool
	extensions string
	maxSize    int64
}

var config Config

var rep = mem.NewStoryRepository()

// ProcessSngFile is the function that gets called for each .sng file
// filePath: the full path to the .sng file
// content: the file's content as bytes
func ProcessSngFile(filePath string, content []byte, rep *mem.StoryRepository) error {
	var pt snggcc.ScriptPart

	if err := yaml.Unmarshal(content, &pt); err != nil {
		return err
	}

	for _, character := range pt.Characters {
		if err := rep.ConsumeCharacter(character); err != nil {
			return fmt.Errorf("processing character %s: %s", character.Id, err)
		}
	}

	for _, location := range pt.Locations {
		if err := rep.ConsumeLocation(location); err != nil {
			return fmt.Errorf("processing location %s: %s", location.Id, err)
		}
	}

	for _, action := range pt.Actions {
		if err := rep.ConsumeAction(action); err != nil {
			return fmt.Errorf("processing action %s: %s", action.Id, err)
		}
	}

	for _, environment := range pt.Environments {
		if err := rep.ConsumeEnvironment(environment); err != nil {
			return fmt.Errorf("processing environment %s: %s", environment.Id, err)
		}
	}

	for _, dialogue := range pt.Dialogues {
		if err := rep.ConsumeDialogue(dialogue); err != nil {
			return fmt.Errorf("processing dialogue %s: %s", dialogue.Id, err)
		}
	}

	for _, passage := range pt.Passages {
		if err := rep.ConsumePassage(passage); err != nil {
			return fmt.Errorf("processing dialogue %s: %s", passage.Id, err)
		}
	}

	for _, story := range pt.Stories {
		if err := rep.ConsumeStory(story); err != nil {
			return fmt.Errorf("processing story %s: %s", story.Id, err)
		}
	}

	return nil
}

// processDirectory recursively traverses directories and processes .sng files
func processDirectory(rootPath string) error {
	// Get the list of extensions to process
	extList := strings.Split(config.extensions, ",")
	for i, ext := range extList {
		extList[i] = strings.ToLower(strings.TrimSpace(ext))
		if !strings.HasPrefix(extList[i], ".") {
			extList[i] = "." + extList[i]
		}
	}

	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return nil // Continue walking despite errors
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check file size limit
		if config.maxSize > 0 && info.Size() > config.maxSize {
			if config.verbose {
				fmt.Printf("Skipping large file: %s (%d bytes)\n", path, info.Size())
			}
			return nil
		}

		// Check if file has one of the specified extensions
		fileExt := strings.ToLower(filepath.Ext(path))
		for _, ext := range extList {
			if fileExt == ext {
				// Read file content
				content, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Printf("Error reading file %s: %v\n", path, err)
					return nil // Continue with next file
				}

				// Process the file
				if err := ProcessSngFile(path, content, rep); err != nil {
					fmt.Printf("Error processing file %s: %v\n", path, err)
				}
				break
			}
		}

		return nil
	})
}

func init() {
	// Set up command line flags
	flag.StringVar(&config.rootDir, "dir", ".", "Root directory to scan recursively")
	flag.StringVar(&config.outPath, "out", "bundle.json", "Output filename")
	flag.StringVar(&config.extensions, "ext", ".sng", "Comma-separated list of file extensions to process (e.g., .sng,.txt,.log)")
	flag.BoolVar(&config.verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&config.verbose, "v", false, "Enable verbose output (shorthand)")
	flag.Int64Var(&config.maxSize, "max-size", 0, "Maximum file size to process in bytes (0 = no limit)")
}

func main() {
	// Parse command line flags
	flag.Parse()

	// Use the first positional argument as directory if provided
	if flag.NArg() > 0 {
		config.rootDir = flag.Arg(0)
	}

	// Verify the directory exists
	if _, err := os.Stat(config.rootDir); os.IsNotExist(err) {
		fmt.Printf("Error: Directory does not exist: %s\n", config.rootDir)
		os.Exit(1)
	}

	// Display configuration
	if config.verbose {
		fmt.Printf("Scanning directory: %s\n", config.rootDir)
		fmt.Printf("File extensions: %s\n", config.extensions)
		if config.maxSize > 0 {
			fmt.Printf("Maximum file size: %d bytes\n", config.maxSize)
		}
		fmt.Println("---")
	}

	if err := processDirectory(config.rootDir); err != nil {
		fmt.Printf("Error scanning directory: %v\n", err)
		os.Exit(1)
	}

	jsonBundle, err := rep.ExportToJSON()
	if err != nil {
		fmt.Printf("Error exporting JSON: %v\n", err)
	}

	if err := os.WriteFile(config.outPath, jsonBundle, 0644); err != nil {
		fmt.Printf("Error writing file: %v\n", err)
	}

	fmt.Println("Compilation completed!")
}
