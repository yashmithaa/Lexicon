#!/bin/bash

# Sprout Language Quick Start Script

echo "🌱 Building Sprout REPL..."
go build -o sprout cmd/repl/main.go

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo ""
    echo "To start the REPL, run:"
    echo "  ./sprout"
    echo ""
    echo "Or run it now? (y/n)"
    read -r response
    if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
        ./sprout
    fi
else
    echo "❌ Build failed!"
    exit 1
fi
