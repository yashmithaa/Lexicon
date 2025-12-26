#!/bin/bash

# Sprout Language Quick Start Script

echo "🌱 Building Sprout REPL..."
go build -o sprout cmd/repl/main.go

echo "🌱 Building Sprout Runner..."
go build -o sprun cmd/demo/main.go

if [ $? -eq 0 ]; then
    echo " Build successful!"
    echo ""
    echo "To start the REPL, run:"
    echo "  ./sprout"
    echo ""
    echo "To run a .spr file, use:"
    echo "  ./sprun filename.spr"
    echo ""
    echo "Or run the REPL now? (y/n)"
    read -r response
    if [[ "$response" =~ ^([yY][eE][sS]|[yY])$ ]]; then
        ./sprout
    fi
else
    echo "Build failed!"
    exit 1
fi
