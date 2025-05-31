<script lang="ts">
	import { onMount } from 'svelte';
	import SystemInfo from '$lib/components/SystemInfo.svelte';
	import ServicesList from '$lib/components/ServicesList.svelte';
	import NetworkStatus from '$lib/components/NetworkStatus.svelte';
	import FileBrowser from '$lib/components/FileBrowser.svelte';

	let systemData: any = null;
	let services: any[] = [];
	let networkData: any = null;
	let loading = true;
	let error = '';

	const API_BASE = 'http://localhost:8080/api';

	async function fetchData() {
		try {
			loading = true;
			error = '';

			// Fetch all data in parallel
			const [systemRes, servicesRes, networkRes] = await Promise.all([
				fetch(`${API_BASE}/system`),
				fetch(`${API_BASE}/services`),
				fetch(`${API_BASE}/network`)
			]);

			if (!systemRes.ok || !servicesRes.ok || !networkRes.ok) {
				throw new Error('Failed to fetch data from server');
			}

			systemData = await systemRes.json();
			services = await servicesRes.json();
			networkData = await networkRes.json();
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error occurred';
			console.error('Error fetching data:', err);
		} finally {
			loading = false;
		}
	}

	onMount(() => {
		fetchData();
		// Refresh data every 5 seconds
		const interval = setInterval(fetchData, 5000);
		return () => clearInterval(interval);
	});
</script>

<svelte:head>
	<title>Homeserver Dashboard</title>
</svelte:head>

<div class="min-h-screen bg-gray-50 dark:bg-gray-900">
	<!-- Header -->
	<header class="bg-white dark:bg-gray-800 shadow">
		<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
			<div class="flex justify-between items-center py-6">
				<div class="flex items-center">
					<h1 class="text-3xl font-bold text-gray-900 dark:text-white">
						🏠 Homeserver Dashboard
					</h1>
				</div>
				<div class="flex items-center space-x-4">
					<div class="flex items-center space-x-2">
						<div class="w-3 h-3 bg-green-500 rounded-full animate-pulse"></div>
						<span class="text-sm text-gray-600 dark:text-gray-300">Live</span>
					</div>
					<button
						on:click={fetchData}
						class="inline-flex items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
					>
						Refresh
					</button>
				</div>
			</div>
		</div>
	</header>

	<!-- Main Content -->
	<main class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
		{#if loading && !systemData}
			<div class="flex justify-center items-center h-64">
				<div class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
			</div>
		{:else if error}
			<div class="bg-red-50 border border-red-200 rounded-md p-4">
				<div class="flex">
					<div class="ml-3">
						<h3 class="text-sm font-medium text-red-800">Error</h3>
						<div class="mt-2 text-sm text-red-700">
							<p>{error}</p>
						</div>
					</div>
				</div>
			</div>
		{:else}
			<div class="px-4 py-6 sm:px-0">
				<div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-6">
					<!-- System Information -->
					<SystemInfo {systemData} />
					
					<!-- Services List -->
					<ServicesList {services} />
				</div>

				<div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
					<!-- Network Status -->
					<NetworkStatus {networkData} />
					
					<!-- File Browser -->
					<FileBrowser />
				</div>
			</div>
		{/if}
	</main>
</div>
