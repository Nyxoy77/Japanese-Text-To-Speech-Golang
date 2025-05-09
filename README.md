# Japanese Text-to-Speech CLI Tool

This project is a command-line interface (CLI) tool that converts Japanese text into speech using Google Cloud Text-to-Speech API. It provides multiple voice options, adjustable speech rate, and pitch, allowing users to generate high-quality audio output.

## Features
- Supports **Google Cloud Text-to-Speech API**
- Multiple **Japanese voices** (Standard & WaveNet)
- Adjustable **speech rate** (0.25x - 4.0x) and **pitch** (-20.0 to 20.0)
- Outputs audio in **MP3 or WAV format**
- CLI-based tool for quick usage
- **Error handling** for input validation and API requests

## Prerequisites
Before using this tool, make sure you have the following:
- **Google Cloud account** with **Text-to-Speech API enabled**
- **Google Cloud SDK (gcloud CLI) installed**
- A valid **service account JSON key** for authentication
- **Go installed** on your machine

## Setup
1. **Clone the repository:**
   ```sh
   git clone https://github.com/nyxoy77/go-text-to-speech.git
   cd go-text-to-speech
   ```
2. **Set up authentication:**
   ```sh
   export GOOGLE_APPLICATION_CREDENTIALS="/path/to/service-account-key.json"
   ```
   *(For Windows PowerShell:)*
   ```sh
   $env:GOOGLE_APPLICATION_CREDENTIALS="C:\path\to\service-account-key.json"
   ```
3. **Install dependencies:**
   ```sh
   go mod tidy
   ```
4. **Build and run the project:**
   ```sh
   go run main.go --text "こんにちは" --voice "wave-a" -O output.mp3
   ```

## Usage
Run the CLI tool with customizable parameters:
```sh
./go-text-to-speech --text "Your Japanese text" --voice "wave-a" --rate 1.0 --pitch 0.0 -O output.mp3
```

### Command-line Options
| Option         | Description                                     |
|---------------|---------------------------------|
| `--text`      | The Japanese text to convert to speech |
| `--voice`     | Select voice (`stand-a`, `wave-a`, etc.)|
| `--rate`      | Speech rate (0.25 - 4.0) |
| `--pitch`     | Adjusts pitch (-20.0 to 20.0) |
| `-O`          | Output file name with extension (`.mp3`, `.wav`) |

## Learning Experience
Through this project, I learned:
- How to **integrate Google Cloud APIs** into a Go project
- Handling **service account authentication** in Go
- Working with **Google's Text-to-Speech API**
- Building a **CLI application** with Go
- Managing **audio encoding formats** (MP3, WAV)
- Using **flags for CLI input handling**

## Future Improvements
- Add **SSML support** for richer speech synthesis
- Extend **language support** beyond Japanese
- Implement **a web-based UI** for ease of use

## Contributions
Feel free to contribute by submitting pull requests or reporting issues!

## License
This project is open-source and available under the MIT License.

