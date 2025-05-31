# Homeserver

A modern homeserver application with Go backend and SvelteKit frontend.

## Project Structure

```
├── backend/               # Go API server
│   ├── main.go           # Main server file
│   ├── go.mod            # Go dependencies
│   ├── go.sum            # Go dependency checksums
│   └── Dockerfile        # Backend container
├── frontend/             # SvelteKit web application
│   ├── src/              # Source files
│   │   ├── routes/       # SvelteKit routes
│   │   ├── lib/          # Components and utilities
│   │   │   └── components/ # Svelte components
│   │   └── app.html      # HTML template
│   ├── package.json      # Node dependencies
│   ├── pnpm-lock.yaml    # Package lock file
│   ├── svelte.config.js  # SvelteKit configuration
│   ├── tsconfig.json     # TypeScript configuration
│   └── vite.config.ts    # Vite configuration
├── .env.example          # Environment variables template
├── .gitignore            # Git ignore rules
├── .dockerignore         # Docker ignore rules
├── CONTRIBUTING.md       # Development guidelines
├── LICENSE               # MIT License
├── Makefile              # Development automation
├── docker-compose.yml    # Docker services
└── README.md             # This file
```

## Features

- 🏠 **System Monitoring Dashboard** - Real-time system information
- 📊 **Resource Usage Statistics** - CPU, memory, and disk usage
- 🔧 **Service Management** - View running services and their status
- 📁 **File Browser** - Navigate your server's file system
- 🌐 **Network Monitoring** - Track network usage and connections
- 🔒 **Secure API** - CORS-enabled REST API
- 📱 **Responsive Design** - Works on desktop and mobile
- 🌙 **Dark Mode Support** - Automatic dark/light theme switching

## Quick Start

### Option 1: Using Make (Recommended)

1. **Setup the project**:
   ```bash
   make setup
   ```

2. **Start development servers**:
   ```bash
   make dev
   ```

3. **Access the application**:
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080

### Option 2: Manual Setup

1. **Install dependencies**:
   ```bash
   # Backend
   cd backend && go mod tidy
   
   # Frontend
   cd frontend && pnpm install  # or npm install
   ```

2. **Start the Backend**:
   ```bash
   cd backend
   go run main.go
   ```

3. **Start the Frontend** (in a new terminal):
   ```bash
   cd frontend
   pnpm dev  # or npm run dev
   ```

### Option 3: Using Docker

```bash
make docker-up
# or manually:
# docker-compose up
```

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/health` | GET | Health check and server status |
| `/api/system` | GET | Complete system information |
| `/api/services` | GET | List of running services |
| `/api/files` | GET | File browser (supports ?path= parameter) |
| `/api/network` | GET | Network usage statistics |

### Example API Response

```json
// GET /api/system
{
  "hostname": "homeserver",
  "platform": "linux",
  "os": "Ubuntu 22.04",
  "architecture": "amd64",
  "cpu_count": 4,
  "memory": {
    "total": 8589934592,
    "used": 2147483648,
    "used_percent": 25.0
  },
  "cpu": {
    "usage_percent": [15.2, 18.7, 12.1, 20.3],
    "cores": 4
  }
}
```

## Development

### Available Make Commands

```bash
make help          # Show all available commands
make dev           # Start both backend and frontend
make install       # Install all dependencies
make build         # Build for production
make test          # Run tests
make lint          # Run linters
make clean         # Clean build artifacts
make docker-up     # Start with Docker
make docker-down   # Stop Docker services
```

### Frontend Development
- Built with SvelteKit 5 and TypeScript
- Styled with Tailwind CSS 4
- Real-time data updates every 5 seconds
- Responsive design for all screen sizes

### Backend Development
- Written in Go 1.21 with Gin framework
- Uses gopsutil for system monitoring
- CORS enabled for development
- RESTful API design

### Adding New Features

1. **Backend**: Add new endpoints in `backend/main.go`
2. **Frontend**: Create components in `frontend/src/lib/components/`
3. **Styling**: Use Tailwind CSS classes for consistent design

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed development guidelines.

## Production Deployment

### Quick Build
```bash
make build
```

### Manual Build
1. Build the frontend:
   ```bash
   cd frontend && pnpm build
   ```

2. Build the backend:
   ```bash
   cd backend && go build -o homeserver main.go
   ```

### Docker Deployment
```bash
# Install Node.js adapter for Docker builds
cd frontend && pnpm add -D @sveltejs/adapter-node

# Build and run with Docker
make docker-build
make docker-up
# or manually:
# docker-compose up --build
```

## Configuration

### Environment Variables

Copy `.env.example` to `.env` and configure:

- `PORT`: Backend server port (default: 8080)
- `GIN_MODE`: Gin framework mode (debug/release)
- `CORS_ORIGINS`: Allowed CORS origins
- `VITE_API_URL`: Frontend API endpoint
- `ENABLE_FILE_BROWSER`: Enable/disable file browser

See [`.env.example`](.env.example) for all available options.

### CORS Configuration

The backend is configured to allow requests from:
- `http://localhost:5173` (Vite dev server)
- `http://localhost:3000` (Alternative dev port)

## Troubleshooting

### Common Issues

1. **Backend won't start**: Make sure Go 1.21+ is installed and run `go mod tidy`
2. **Frontend won't start**: Make sure Node.js 18+ is installed and run `pnpm install`
3. **API connection issues**: Check that backend is running on port 8080
4. **Permission errors**: File browser requires appropriate file system permissions
5. **Make command not found**: Install make or use manual commands from [CONTRIBUTING.md](CONTRIBUTING.md)

### Logs

- Backend logs are output to console
- Frontend dev server logs are in the terminal
- Production logs can be configured in docker-compose.yml

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines, setup instructions, and coding standards.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
