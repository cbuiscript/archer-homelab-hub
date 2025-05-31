// Configuration for the frontend application
import { browser } from '$app/environment';

// Get environment variables
const getEnvVar = (name: string, defaultValue: string): string => {
	if (browser) {
		// In the browser, use the values injected at build time
		return (globalThis as any).__API_URL__ || defaultValue;
	} else {
		// On the server, use the environment variables
		return import.meta.env[name] || defaultValue;
	}
};

export const config = {
	// API configuration
	apiUrl: getEnvVar('VITE_API_URL', 'http://localhost:8080'),
	baseUrl: getEnvVar('VITE_BASE_URL', '/'),
	
	// Development mode
	devMode: getEnvVar('VITE_DEV_MODE', 'true') === 'true',
	
	// API endpoints
	endpoints: {
		health: '/api/health',
		system: '/api/system',
		services: '/api/services',
		files: '/api/files',
		network: '/api/network'
	}
};

// Helper function to build full API URLs
export const buildApiUrl = (endpoint: string): string => {
	return `${config.apiUrl}${endpoint}`;
};

export default config;
