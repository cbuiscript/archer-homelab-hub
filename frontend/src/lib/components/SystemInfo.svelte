<script lang="ts">
	let { systemData } = $props();

	function formatBytes(bytes: number): string {
		if (bytes === 0) return '0 Bytes';
		const k = 1024;
		const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
		const i = Math.floor(Math.log(bytes) / Math.log(k));
		return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
	}

	function formatUptime(seconds: number): string {
		const days = Math.floor(seconds / 86400);
		const hours = Math.floor((seconds % 86400) / 3600);
		const minutes = Math.floor((seconds % 3600) / 60);
		return `${days}d ${hours}h ${minutes}m`;
	}
</script>

<div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
	<div class="px-4 py-5 sm:p-6">
		<div class="flex items-center">
			<div class="flex-shrink-0">
				<div class="w-8 h-8 bg-blue-500 rounded-md flex items-center justify-center">
					<span class="text-white text-sm font-bold">💻</span>
				</div>
			</div>
			<div class="ml-5 w-0 flex-1">
				<dl>
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">
						System Information
					</dt>
					<dd class="text-lg font-medium text-gray-900 dark:text-white">
						{systemData?.hostname || 'Loading...'}
					</dd>
				</dl>
			</div>
		</div>

		{#if systemData}
			<div class="mt-5 grid grid-cols-1 gap-5 sm:grid-cols-2">
				<!-- OS Information -->
				<div class="bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-lg">
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Operating System</dt>
					<dd class="mt-1 text-sm text-gray-900 dark:text-white">
						{systemData.os} ({systemData.platform})
					</dd>
					<dd class="text-xs text-gray-500 dark:text-gray-400">
						{systemData.architecture} • {systemData.cpu_count} cores
					</dd>
				</div>

				<!-- Uptime -->
				<div class="bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-lg">
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Uptime</dt>
					<dd class="mt-1 text-sm text-gray-900 dark:text-white">
						{formatUptime(systemData.uptime)}
					</dd>
				</div>

				<!-- Memory Usage -->
				<div class="bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-lg">
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Memory Usage</dt>
					<dd class="mt-1 text-sm text-gray-900 dark:text-white">
						{formatBytes(systemData.memory.used)} / {formatBytes(systemData.memory.total)}
					</dd>
					<div class="mt-2 w-full bg-gray-200 rounded-full h-2">
						<div
							class="bg-blue-600 h-2 rounded-full"
							style="width: {systemData.memory.used_percent}%"
						></div>
					</div>
					<dd class="text-xs text-gray-500 dark:text-gray-400 mt-1">
						{systemData.memory.used_percent.toFixed(1)}% used
					</dd>
				</div>

				<!-- CPU Usage -->
				<div class="bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-lg">
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400">CPU Usage</dt>
					{#if systemData.cpu.usage_percent && systemData.cpu.usage_percent.length > 0}
						<dd class="mt-1 text-sm text-gray-900 dark:text-white">
							{(systemData.cpu.usage_percent.reduce((a: number, b: number) => a + b, 0) / systemData.cpu.usage_percent.length).toFixed(1)}%
						</dd>
						<div class="mt-2 w-full bg-gray-200 rounded-full h-2">
							<div
								class="bg-green-600 h-2 rounded-full"
								style="width: {(systemData.cpu.usage_percent.reduce((a: number, b: number) => a + b, 0) / systemData.cpu.usage_percent.length)}%"
							></div>
						</div>
					{:else}
						<dd class="mt-1 text-sm text-gray-900 dark:text-white">N/A</dd>
					{/if}
				</div>

				<!-- Disk Usage -->
				{#if systemData.disks && systemData.disks.length > 0}
					<div class="bg-gray-50 dark:bg-gray-700 px-4 py-5 rounded-lg sm:col-span-2">
						<dt class="text-sm font-medium text-gray-500 dark:text-gray-400">Disk Usage</dt>
						{#each systemData.disks as disk}
							<div class="mt-2">
								<div class="flex justify-between text-sm">
									<span class="text-gray-900 dark:text-white">{disk.mountpoint}</span>
									<span class="text-gray-500 dark:text-gray-400">
										{formatBytes(disk.used)} / {formatBytes(disk.total)}
									</span>
								</div>
								<div class="mt-1 w-full bg-gray-200 rounded-full h-2">
									<div
										class="bg-yellow-600 h-2 rounded-full"
										style="width: {disk.used_percent}%"
									></div>
								</div>
								<div class="text-xs text-gray-500 dark:text-gray-400 mt-1">
									{disk.used_percent.toFixed(1)}% used • {disk.fstype}
								</div>
							</div>
						{/each}
					</div>
				{/if}
			</div>
		{/if}
	</div>
</div>
