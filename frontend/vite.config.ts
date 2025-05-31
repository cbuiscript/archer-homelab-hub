import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig, loadEnv } from 'vite';

export default defineConfig(({ mode }) => {
	// Load environment variables from the parent directory
	const env = loadEnv(mode, '../', '');
	
	return {
		plugins: [tailwindcss(), sveltekit()],
		server: {
			port: parseInt(env.FRONTEND_PORT || '5173'),
			host: env.FRONTEND_HOST || 'localhost',
		},
		envDir: '../', // Look for .env files in the parent directory
		define: {
			// Make some env vars available to the client
			__API_URL__: JSON.stringify(env.VITE_API_URL || 'http://localhost:8080'),
		}
	};
});
