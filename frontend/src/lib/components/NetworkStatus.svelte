<script lang="ts">
	let { networkData } = $props();

	function formatBytes(bytes: number): string {
		if (bytes === 0) return '0 Bytes';
		const k = 1024;
		const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
		const i = Math.floor(Math.log(bytes) / Math.log(k));
		return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
	}

	function getConnectionStatus(count: number): { color: string; text: string } {
		if (count > 100) return { color: 'text-red-600', text: 'High' };
		if (count > 50) return { color: 'text-yellow-600', text: 'Medium' };
		return { color: 'text-green-600', text: 'Normal' };
	}
</script>

<div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
	<div class="px-4 py-5 sm:p-6">
		<div class="flex items-center">
			<div class="flex-shrink-0">
				<div class="w-8 h-8 bg-purple-500 rounded-md flex items-center justify-center">
					<span class="text-white text-sm font-bold">🌐</span>
				</div>
			</div>
			<div class="ml-5 w-0 flex-1">
				<dl>
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">
						Network Status
					</dt>
					<dd class="text-lg font-medium text-gray-900 dark:text-white">
						{networkData?.active_connections || 0} connections
					</dd>
				</dl>
			</div>
		</div>

		{#if networkData}
			<div class="mt-5 grid grid-cols-1 gap-5 sm:grid-cols-2">
				<!-- Data Sent -->
				<div class="bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-lg">
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Data Sent</dt>
					<dd class="mt-1 text-xl font-semibold text-gray-900 dark:text-white">
						{formatBytes(networkData.total_bytes_sent)}
					</dd>
					<div class="mt-2 flex items-center text-sm text-gray-600 dark:text-gray-300">
						<span class="text-green-600">↑</span>
						<span class="ml-1">Outbound</span>
					</div>
				</div>

				<!-- Data Received -->
				<div class="bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-lg">
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Data Received</dt>
					<dd class="mt-1 text-xl font-semibold text-gray-900 dark:text-white">
						{formatBytes(networkData.total_bytes_received)}
					</dd>
					<div class="mt-2 flex items-center text-sm text-gray-600 dark:text-gray-300">
						<span class="text-blue-600">↓</span>
						<span class="ml-1">Inbound</span>
					</div>
				</div>

				<!-- Active Connections -->
				<div class="bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-lg sm:col-span-2">
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Active Connections</dt>
					<dd class="mt-1 flex items-center">
						<span class="text-2xl font-bold text-gray-900 dark:text-white mr-3">
							{networkData.active_connections}
						</span>
						<span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 {getConnectionStatus(networkData.active_connections).color}">
							{getConnectionStatus(networkData.active_connections).text} Activity
						</span>
					</dd>
					<div class="mt-3">
						<div class="w-full bg-gray-200 rounded-full h-2">
							<div
								class="bg-purple-600 h-2 rounded-full"
								style="width: {Math.min((networkData.active_connections / 200) * 100, 100)}%"
							></div>
						</div>
						<p class="mt-1 text-xs text-gray-500 dark:text-gray-400">
							Connection load indicator
						</p>
					</div>
				</div>
			</div>

			<!-- Last Updated -->
			<div class="mt-4 text-xs text-gray-500 dark:text-gray-400 text-center">
				Last updated: {new Date(networkData.timestamp * 1000).toLocaleTimeString()}
			</div>
		{/if}
	</div>
</div>
