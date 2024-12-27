
# Go Docker App

This is a simple Go project that demonstrates two basic programs: 

1. **IP Address Formatter** (`main1.go`): Implements the `fmt.Stringer` interface to format IP addresses in a dotted quad representation (e.g., `127.0.0.1`).
2. **ROT13 Reader** (`main2.go`): Implements the `io.Reader` interface to decode a string using the ROT13 substitution cipher.

---

## Project Structure

```plaintext
go-docker-app/
├── cmd/
│   ├── main1.go        # IP Address Formatter
│   ├── main2.go        # ROT13 Reader
├── docker-compose.yml  # Docker Compose configuration
├── Dockerfile          # Dockerfile to build and run the app
├── go.mod              # Go module definition
├── go.sum              # Go dependencies
```

---

## Description of Programs

### 1. **IP Address Formatter (`main1.go`)**
This program formats IP addresses as strings in a dotted quad format using the `fmt.Stringer` interface. It prints a list of IP addresses mapped to their names.

#### Example Output:
```plaintext
loopback: 127.0.0.1
googleDNS: 8.8.8.8
```

---

### 2. **ROT13 Reader (`main2.go`)**
This program implements an `io.Reader` that reads a string and applies the ROT13 substitution cipher to all alphabetical characters. The program demonstrates reading a ROT13-encoded message and decoding it.

#### Example Output:
```plaintext
You cracked the code!
```

---

## Docker Setup

### Prerequisites

- Docker
- Docker Compose

### Building and Running the Project

1. **Build and Run the Docker Container**:
   ```bash
   docker-compose up --build
   ```

2. **View the Output**:
   The output of both programs (`main1.go` and `main2.go`) will be displayed in the terminal.

---

## How It Works

1. The `Dockerfile` builds the Go project and compiles two binaries: `ipaddr` (for `main1.go`) and `rot13` (for `main2.go`).
2. The `docker-compose.yml` file runs the container and executes both binaries in sequence.

---

## Notes

- This project demonstrates basic Go concepts:
  - Implementing interfaces (`fmt.Stringer` and `io.Reader`).
  - Using Docker to build and run Go applications.
- Modify or expand the `main1.go` and `main2.go` files to explore more Go features!
