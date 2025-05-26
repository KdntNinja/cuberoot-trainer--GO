# Cube Root Trainer

A web application that helps you learn and practice calculating cube roots in your head. Built with Go and templUI components.

## Features

- **Practice Mode**: Generate random cube root problems and test your mental calculation skills
- **Learning Resources**: Learn techniques and patterns for calculating cube roots mentally
- **Interactive UI**: Get immediate feedback on your answers

## Installation

This application is built with Go and uses the templUI component library.

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/axzilla/templui-quickstart.git
   cd templui-quickstart
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

## Development

Start the development server with hot reload:

```bash
make dev
```

Your application will be running at [http://localhost:7331](http://localhost:7331)

## Deployment

This application includes a production-ready Dockerfile for easy deployment:

```bash
# Build the image
docker build -t cube-root-trainer .

# Run the container
docker run -p 8090:8090 cube-root-trainer
```

Your application will be available at `http://localhost:8090`

### Docker Image Details

The Docker image is built with all required dependencies:

- Go 1.24 (for building the application)
- templ v0.3.865 or later (for generating Go code from templ files)
- Tailwind CSS v4.0.5 or later (for generating CSS)

The final image is a minimal Alpine Linux image that contains only the compiled application and necessary runtime dependencies.

## Contributing

Issues and pull requests are welcome! Please read our [contributing guidelines](https://github.com/axzilla/templui/blob/main/CONTRIBUTING.md) before submitting a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
