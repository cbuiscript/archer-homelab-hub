# Homeserver

A modern homeserver application with Go backend and SvelteKit frontend.

## Project Structure

```
├── backend/          # Go API server
│   ├── main.go       # Main server file
│   ├── go.mod        # Go dependencies
│   └── Dockerfile    # Backend container
├── frontend/         # SvelteKit web application
│   ├── src/          # Source files
│   │   ├── routes/   # SvelteKit routes
│   │   └── lib/      # Components and utilities
│   ├── package.json  # Node dependencies
│   └── tailwind.config.js # Tailwind configuration
├── docker-compose.yml
└── README.md
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

### Option 1: Manual Setup (Recommended for Development)

1. **Start the Backend**:
   ```bash
   cd backend
   go mod tidy
   go run main.go
   ```

2. **Start the Frontend** (in a new terminal):
   ```bash
   cd frontend
   npm install  # or pnpm install
   npm run dev  # or pnpm dev
   ```

3. **Access the Application**:
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080

### Option 2: Using Docker (Production-like)

```bash
docker-compose up
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

### Frontend Development
- Built with SvelteKit 5 and TypeScript
- Styled with Tailwind CSS 4
- Real-time data updates every 5 seconds
- Responsive design for all screen sizes

### Backend Development
- Written in Go with Gin framework
- Uses gopsutil for system monitoring
- CORS enabled for development
- RESTful API design

### Adding New Features

1. **Backend**: Add new endpoints in `backend/main.go`
2. **Frontend**: Create components in `frontend/src/lib/components/`
3. **Styling**: Use Tailwind CSS classes for consistent design

## Production Deployment

1. Build the frontend:
   ```bash
   cd frontend && npm run build
   ```

2. Build the backend:
   ```bash
   cd backend && go build -o homeserver main.go
   ```

3. Or use Docker:
   ```bash
   docker-compose up --build
   ```

## Configuration

### Environment Variables

- `PORT`: Backend server port (default: 8080)
- `NODE_ENV`: Frontend environment (development/production)

### CORS Configuration

The backend is configured to allow requests from:
- `http://localhost:5173` (Vite dev server)
- `http://localhost:3000` (Alternative dev port)

## Troubleshooting

### Common Issues

1. **Backend won't start**: Make sure Go 1.21+ is installed and run `go mod tidy`
2. **Frontend won't start**: Make sure Node.js 18+ is installed and run `npm install`
3. **API connection issues**: Check that backend is running on port 8080
4. **Permission errors**: File browser requires appropriate file system permissions

### Logs

- Backend logs are output to console
- Frontend dev server logs are in the terminal
- Production logs can be configured in docker-compose.yml
