# API Key Scanner v2.0
**Created by HeJo-1**

---

# English — Quick Repo Guide

## What this repo is
A simple, aesthetic terminal-based API key scanner written in Go.  
It scans provided URLs for *potential* exposed API keys using regex patterns (Google, AWS, GitHub, Slack). This repository includes a `main.go` scanner and a `test/` folder with safe test pages you can use locally.

> 🔒 **Important:** Do **not** use this tool to scan third-party websites without explicit permission. Unauthorized scanning can be illegal or unethical. Use only on sites you own or on local test files.

---

## Repo layout (suggested)
```
.
├── main.go          # Scanner program (HeJo-1)
├── found_keys.txt   # (generated) scan results
├── README.md
└── test/
    ├── test.html    # example test page with fake API keys
    └── ...          # other test files (optional)
```

---

## Requirements
- Go 1.18+ (modules enabled)
- Internet access only to download dependencies during setup (optional afterwards)
- Permission to scan the target resources (for legal/ethical reasons)

---

## Setup (one-time)

Open a terminal in the project root and run:

```bash
# initialize go module (choose a module name if you like)
go mod init apikeyscanner

# fetch the colored-output dependency
go get github.com/fatih/color@latest
```

You should see a `go.mod` and `go.sum` created.

---

## Local test (safe demo)

A safe way to test the scanner is to host the `test/test.html` locally and scan it.

1. Create `test/test.html` (example content — contains **fake** API key formats):

```html
<!DOCTYPE html>
<html>
  <body>
    <h3>Local Test Page</h3>
    <p>Fake Google API key example:</p>
    <code>AIzaSyD9P9bQzBfVxkP8JrL6E6DExAMPLEKEY12345678</code>
  </body>
</html>
```

2. Start a simple HTTP server (from repo root):
```bash
python3 -m http.server 8080
```

3. Run the scanner against the local test page:

```bash
# English output
go run main.go --lang en http://localhost:8080/test/test.html

# or Turkish output
go run main.go --lang tr http://localhost:8080/test/test.html
```

Expected (example) console output:
```
🚀 API Key Scanner v2.0 - Created by HeJo-1
==============================================

🔍 Scanning: http://localhost:8080/test/test.html
✅ Potential Google Key Found: http://localhost:8080/test/test.html
   ➜ AIzaSyD9P9bQzBfVxkP8JrL6E6DExAMPLEKEY12345678

🎯 Scan complete. Results saved to 'found_keys.txt'.
```

`found_keys.txt` will contain entries like:
```
URL: http://localhost:8080/test/test.html
Type: Google
Key: AIzaSyD9P9bQzBfVxkP8JrL6E6DExAMPLEKEY12345678
```

---

## Usage examples

Scan multiple URLs in parallel:
```bash
go run main.go --lang en https://example.com http://localhost:8080/test/test.html
```

Show Turkish messages:
```bash
go run main.go --lang tr http://localhost:8080/test/test.html
```

If you prefer to build a binary:
```bash
go build -o apiscanner main.go
./apiscanner --lang en http://localhost:8080/test/test.html
```

---

## Customization
- Add more regexes in `apiRegex` (in `main.go`) for more providers.
- Modify the messages map to change text/emoji/localization.
- Replace `github.com/fatih/color` usage with another console library if desired.

---

## Security & Ethics
- Only scan resources you own or have explicit permission to scan.
- The tool reports *potential* keys by regex; it may produce false positives.
- Do **not** use discovered keys — responsibly inform the owners and follow disclosure guidelines.

---

## License
This repo is provided as-is. You can add a license file (e.g., `LICENSE`), such as MIT, if you want to open-source it.

```
MIT License

Copyright (c) 2025 HeJo-1

Permission is hereby granted...
```

---

## Author
**HeJo-1**
