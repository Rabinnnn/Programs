#!/bin/bash

# Count regular files
files_count=$(find . -type f | wc -l)

# Count folders (directories)
folders_count=$(find . -type d | wc -l)

# Print the total count
echo $((files_count + folders_count))
