package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Stores component directory and filepath
//
// Fields:
// - ComponentDir (string): path to directory
// containing components
// - FilePath (string): path to file to compile
//
// Since: 0.1.0
type Compiler struct {
	ComponentDir string
	FilePath     string
}

// Compile the lamb file and components into a parsable
// .html file for standard go parser
//
// Params:
// - filePath (string): path to lamb file for compiling
// - componentDir (string): path to directory of lamb components
//
// Returns:
// - error
//
// Since: 0.1.0
func Compile(filePath string, componentDir string) error {
	compiler := Compiler{
		ComponentDir: componentDir,
		FilePath:     filePath,
	}

	err := compiler.compileLamb()
	if err != nil {
		return err
	}

	return nil
}

// Creates the .cache directory in the
// root of the library
//
// Receiver:
// - c (*Compiler)
//
// Returns:
// - error
//
// Since: 0.1.0
func (c *Compiler) createCache() (string, error) {
	rootDir := getLibraryRoot()

	cacheDir := filepath.Join(rootDir, ".cache")
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		err = os.Mkdir(cacheDir, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create .cache directory: %w", err)
		}
	}

	fmt.Printf("created the cache directory at %s\n", cacheDir)
	return cacheDir, nil
}

// Creates the output file name
//
// Receiver:
// - c (*Compiler)
//
// Returns:
// - string: the compiled filename
//
// Since: 0.1.0
func (c *Compiler) getOutputFileName() string {
	return strings.TrimSuffix(filepath.Base(c.FilePath), ".lamb.html") + ".html"
}

// Creates the output file path
//
// Receiver:
// - c (*Compiler)
//
// Params:
// - fileName (string): the output file name
// - cachePath (string): path to cache directory
//
// Returns:
// - string: the output file path
//
// Since: 0.1.0
func (c *Compiler) getOutputFilePath(fileName string, cachePath string) string {
	return filepath.Join(cachePath, fileName)
}

// Writes the file to the cache
//
// Params:
// - content (string): the parsed content
// - path (string): the output file path
//
// Returns:
// - error
//
// Since: 0.1.0
func writeFileToCache(content string, path string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write compiled file: %w", err)
	}

	fmt.Printf("Compiled file written to %s\n", path)
	return nil
}

// Compiles the specified file to the cache
//
// Receiver:
// - c (*Compiler)
//
// Returns:
// - error
//
// Since: 0.1.0
func (c *Compiler) compileLamb() error {
	// Parse the file to get the content
	parsedContent, err := ParseLamb(c.FilePath, c.ComponentDir)
	if err != nil {
		return err
	}

	cachePath, err := c.createCache()
	if err != nil {
		return err
	}

	// Remove the ".lamb" from the filename and replace it with ".html"
	outputFileName := c.getOutputFileName()

	// Define the path for the compiled .html file
	outputFilePath := c.getOutputFilePath(outputFileName, cachePath)

	err = writeFileToCache(parsedContent, outputFilePath)
	if err != nil {
		return err
	}
	return nil
}
