# Hevy MCP Server

A Model Context Protocol (MCP) server that provides access to workout data from [Hevy](https://www.hevyapp.com/), a popular weightlifting and workout tracking application.

## Overview

This MCP server allows AI assistants like Claude to query and retrieve workout data from your Hevy account, enabling natural language interactions with your fitness history.

## Features

### Tools

- **get_workout_count**: Get the total number of workouts in your Hevy account
- **get_workouts**: Retrieve your workouts from newest to oldest with pagination support
  - Default limit: 5 workouts
  - Maximum limit: 10 workouts per request
  - Supports offset for pagination

## Prerequisites

- A Hevy account with API access
- Hevy API key

## Installation

### Option 1: Download Pre-built Binary (Recommended)

Download the latest release for your platform from the [releases page](https://github.com/swrm-io/hevy-mcp/releases):

- **Linux**: `hevy-mcp_Linux_x86_64.tar.gz`
- **macOS (Intel)**: `hevy-mcp_Darwin_x86_64.tar.gz`
- **macOS (Apple Silicon)**: `hevy-mcp_Darwin_arm64.tar.gz`
- **Windows**: `hevy-mcp_Windows_x86_64.zip`

Extract the archive and you're ready to use the `hevy-mcp` binary.

### Option 2: Build from Source

Requirements: Go 1.25.1 or higher

1. Clone the repository:
```bash
git clone https://github.com/swrm-io/hevy-mcp.git
cd hevy-mcp
```

2. Build the server:
```bash
go build -o hevy-mcp
```

## Configuration

### Environment Variables

- `HEVY_API_KEY` (required): Your Hevy API key

### Setting up with Claude Desktop

Add this configuration to your Claude Desktop config file:

**MacOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "hevy": {
      "command": "/path/to/hevy-mcp",
      "env": {
        "HEVY_API_KEY": "your-api-key-here"
      }
    }
  }
}
```

## Usage

### Running the Server

```bash
export HEVY_API_KEY="your-api-key-here"
./hevy-mcp
```

### Example Queries (via Claude)

Once configured with Claude Desktop, you can ask questions like:

- "How many workouts do I have in Hevy?"
- "Show me my last 5 workouts"
- "What exercises did I do in my recent workouts?"

## Development

The server is built using:
- [go-sdk](https://github.com/modelcontextprotocol/go-sdk) - Official MCP SDK for Go
- [go-hevy](https://github.com/swrm-io/go-hevy) - Hevy API client library

### Project Structure

- `main.go` - Server initialization and tool registration
- `types.go` - Type definitions for requests and service
- `workouts.go` - Workout-related tool implementations

## License

See LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
