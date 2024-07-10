# Character Count CLI

## Overview

This command-line interface (CLI) tool analyzes the character frequency in files within a specified directory and optionally filters by file type. It generates a sideways bar chart in the terminal to visualize the most common characters.

## Installation

Ensure you have Go installed. Clone the repository and navigate to the project directory:

```bash
git clone github.com/zandacw/charcnt
cd github.com/zandacw/charcnt 
```

Install dependencies:

```bash
go mod tidy
```

Build and install:

```bash
go install
```

## Usage

```bash
charcnt [option] dirpath
```

### Options

- `filetype`: Specify the file type to count characters on (default: * for all files).
- `width`: Maximum width of the bar chart (default: 80)

## Example

```bash
charcnt .
```

```bash
charcnt -filetype=go -width=160 /path/to/directory
```

## Output

The CLI outputs a sideways bar chart where each character is represented by its frequency. Characters are color-coded based on type:
- Red: Digits (0-9)
- Blue: Alphabetic (A-Z, a-z)
- Green: Symbols
- White: Punctuation

## Note

This is a silly tool with no optimisations. Might blow up your machine if run on a massively nested directory :(

## Acknowledgements

- Uses `github.com/fatih/color` for terminal colors

