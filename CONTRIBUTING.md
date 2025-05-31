# Contributing to Homeserver

Thank you for your interest in contributing to the Homeserver project! This document provides guidelines and information for contributors.

## Development Setup

### Prerequisites

- **Go 1.21+** for backend development
- **Node.js 18+** for frontend development
- **pnpm** for package management (recommended)
- **Docker** (optional, for containerized development)

### Quick Start

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd homeserver
   ```

2. **Install dependencies**:
   ```bash
   make install
   # Or manually:
   cd backend && go mod tidy
   cd ../frontend && pnpm install
   ```

3. **Set up environment**:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Start development servers**:
   ```bash
   make dev
   # Or manually:
   # Terminal 1: cd backend && go run main.go
   # Terminal 2: cd frontend && pnpm dev
   ```

## Project Structure

```
├── backend/              # Go API server
│   ├── main.go          # Main server file
│   ├── go.mod           # Go dependencies
│   └── Dockerfile       # Backend container
├── frontend/            # SvelteKit web application
│   ├── src/            # Source files
│   │   ├── routes/     # SvelteKit routes
│   │   └── lib/        # Components and utilities
│   └── package.json    # Node dependencies
├── .env.example        # Environment variables template
├── .gitignore          # Git ignore rules
├── docker-compose.yml  # Docker services
└── Makefile           # Development commands
```

## Development Guidelines

### Backend (Go)

- Follow standard Go conventions and formatting
- Use `go fmt` and `go vet` before committing
- Add tests for new functionality
- Keep endpoints RESTful and well-documented

### Frontend (SvelteKit)

- Use TypeScript for type safety
- Follow SvelteKit 5 patterns and conventions
- Use Tailwind CSS for styling
- Ensure responsive design
- Test components when possible

### Code Style

- Use descriptive variable and function names
- Add comments for complex logic
- Keep functions small and focused
- Follow the existing code patterns

## Making Changes

### Workflow

1. **Fork the repository** (if external contributor)
2. **Create a feature branch**:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. **Make your changes**
4. **Test your changes**:
   ```bash
   make test
   make lint
   ```
5. **Commit your changes**:
   ```bash
   git commit -m "feat: add your feature description"
   ```
6. **Push and create a pull request**

### Commit Messages

Use conventional commit messages:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `style:` for formatting changes
- `refactor:` for code refactoring
- `test:` for test additions
- `chore:` for maintenance tasks

## Adding New Features

### Backend Features

1. Add new endpoints in `backend/main.go`
2. Follow existing patterns for error handling
3. Update API documentation in README.md
4. Test the endpoint manually and with curl/Postman

### Frontend Features

1. Create components in `frontend/src/lib/components/`
2. Add routes in `frontend/src/routes/`
3. Use TypeScript interfaces for API data
4. Ensure responsive design with Tailwind CSS
5. Update the main page to include new features

## Testing

### Backend Testing

```bash
cd backend
go test ./...
```

### Frontend Testing

```bash
cd frontend
pnpm test  # When tests are configured
pnpm check # Type checking
```

## Building and Deployment

### Local Build

```bash
make build
```

### Docker Build

```bash
make docker-build
make docker-up
```

### Production Deployment

1. Build the frontend: `cd frontend && pnpm build`
2. Build the backend: `cd backend && go build -o homeserver main.go`
3. Or use Docker: `docker-compose up --build`

## Environment Variables

See `.env.example` for all available environment variables. Key variables:

- `PORT`: Backend server port (default: 8080)
- `GIN_MODE`: Gin framework mode (debug/release)
- `VITE_API_URL`: Frontend API endpoint

## Troubleshooting

### Common Issues

1. **Port conflicts**: Make sure ports 8080 and 5173 are available
2. **Permission errors**: File browser requires appropriate file system permissions
3. **CORS errors**: Check that frontend URL is in CORS_ORIGINS
4. **Module not found**: Run `go mod tidy` in backend directory

### Getting Help

- Check existing issues in the repository
- Review the README.md for setup instructions
- Ask questions in discussions or create an issue

## Security Considerations

- Be careful with file browser functionality
- Validate all user inputs
- Follow Go security best practices
- Don't commit sensitive information (use .env files)

Thank you for contributing!
