# GPT-4o-mini-CLI


# Real-time GPT-4o CLI with WebSockets

## Overview

This project is a real-time command-line interface (CLI) that interacts with OpenAI's GPT-4o-mini using WebSockets. It supports streaming responses and function calling for an efficient and responsive experience.

## Project Structure

```
├── main.go                # Entry point of the application
├── handlers
│   ├── websocket.go       # WebSocket connection handler
│   ├── handlers.go        # Handles incoming messages
├── functions
│   ├── addition.go        # Adds numbers
│   ├── multiplication.go  # Multiplies numbers
├── models
│   ├── function.go        # Defines request/response structures
├── config
│   ├── settings.go        # Configuration settings
├── README.md              # Project documentation
```

## Installation

### Prerequisites

- Go 1.22.1
- github.com/gorilla/websocket

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/Yair0001/GPT-4o-mini-CLI
   cd GPT-4o-mini-CLI
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Run the application:
   ```sh
   go run main.go
   ```



# RAG X

## Project Structure
    src
        embedding.py
        extraction.py
        main.py
        utils.py
        requirements.txt
## Setup
    pip install -r requirements.txt
    python3 main.py <FOLDER PATH1> <FOLDER PATH2>
