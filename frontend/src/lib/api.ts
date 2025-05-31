// API client for the homeserver backend
import { buildApiUrl, config } from './config';

export class ApiClient {
	private baseUrl: string;

	constructor() {
		this.baseUrl = config.apiUrl;
	}

	private async fetchJson<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
		const url = buildApiUrl(endpoint);
		
		const response = await fetch(url, {
			headers: {
				'Content-Type': 'application/json',
				...options.headers,
			},
			...options,
		});

		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}

		return response.json();
	}

	// Health check
	async getHealth() {
		return this.fetchJson(config.endpoints.health);
	}

	// System information
	async getSystemInfo() {
		return this.fetchJson(config.endpoints.system);
	}

	// Services list
	async getServices() {
		return this.fetchJson(config.endpoints.services);
	}

	// Files list
	async getFiles(path?: string) {
		const endpoint = path ? `${config.endpoints.files}?path=${encodeURIComponent(path)}` : config.endpoints.files;
		return this.fetchJson(endpoint);
	}

	// Network status
	async getNetworkStatus() {
		return this.fetchJson(config.endpoints.network);
	}
}

// Export a singleton instance
export const apiClient = new ApiClient();
export default apiClient;
