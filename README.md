# Go Link Shortener

A URL shortening service built with Go to understand the fundamentals of Go development. This project provides a simple web interface for creating and managing shortened links.

## Project Goals

This project was created to understand the fundamentals in Go through building a practical application.

The MVP for this project is it must accept a link URL, shorten the URL, store the shortened and full URL, and then redirect a user navigating to the shortened URL to the full URL.

## Features

- URL Shortening: Generate short, unique codes for long URLs  
- URL Validation: Validates target URLs before shortening by checking for valid HTTP responses (200 status)  
- Web Interface: Clean, responsive HTML interface for creating and viewing shortened links  
- Copy to Clipboard: One-click copying of generated short URLs with visual feedback  
- Redirect Handling: Automatic redirection from short codes to original URLs  
- Static File Serving: Serves stylesheets and other static assets  
- JSON-based Storage: Simple file-based data persistence using JSON  

## Getting Started

### Prerequisites

- Go 1.x or higher  
- Make (optional, but recommended)  

### Installation

1. Clone the repository:  
   git clone https://github.com/cbullard-dev/go-link-shortener.git  
   cd go-link-shortener  

2. Bootstrap the project (sets up git hooks):  
   make bootstrap  

3. Build the application:  
   make  

### Running the Application

Start the server:  
make run  

Or run the built binary directly:  
./go-link-shortener  

### Development Commands

Build: Compile the application  
make  

Test: Run all tests  
make test  

Check: Verify the build is working  
make check  

Clean: Remove build artifacts  
make clean  

## Usage

1. Navigate to the application's home page in your browser  
2. Enter a long URL in the input field  
3. Click "Shorten" to generate a shortened link  
4. The shortened URL will be displayed with a "Copy" button  
5. Click "Copy" to copy the shortened URL to your clipboard  
6. Use the short URL to redirect to your original destination  

## Technical Details

### URL Generation

- Generates unique random codes for each URL  
- Validates that generated codes don't conflict with existing entries  
- Checks for code uniqueness before saving  

### URL Validation

- Performs HTTP HEAD requests to validate URLs before shortening  
- Ensures target URLs return a 200 status code at creation time  
- Provides user feedback for invalid URLs  

### Data Storage

- Uses JSON file-based storage for simplicity  
- Includes helper functions for reading and writing data  
- Maintains a clean data structure for URL mappings  

## Continuous Integration

The project includes GitHub Actions workflows that:  

- Run tests on pushes and pull requests to main  
- Verify builds complete successfully  
- Install and validate dependencies  

## Development Notes

### Testing

Some API tests are temporarily disabled (.routes_test.go) due to issues with static file handling in the test environment. These will be reintroduced once the test logic is corrected to properly handle template and static file scenarios.  

### Git Hooks

The project uses git hooks to enforce code formatting. Run make bootstrap after cloning to set up the pre-commit hook that automatically formats code with go fmt.  

## Feedback and Mentoring

Your feedback is very important. Please feel free to leave feedback or report issues by opening an issue on the GitHub repository.

Additionally, I am seeking a mentor to assist me on my software engineering journey. If you are interested in mentoring or providing guidance, please reach out via a GitHub issue or discussion.

## License

[Add your license information here]  

## Acknowledgments

- HTML and CSS implementation developed with AI assistance
