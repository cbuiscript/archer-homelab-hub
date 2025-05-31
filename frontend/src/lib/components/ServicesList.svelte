<script lang="ts">
	let { services } = $props();

	function getStatusColor(status: string): string {
		switch (status.toLowerCase()) {
			case 'running':
				return 'bg-green-100 text-green-800';
			case 'stopped':
				return 'bg-red-100 text-red-800';
			case 'pending':
				return 'bg-yellow-100 text-yellow-800';
			default:
				return 'bg-gray-100 text-gray-800';
		}
	}

	function getStatusIcon(status: string): string {
		switch (status.toLowerCase()) {
			case 'running':
				return '🟢';
			case 'stopped':
				return '🔴';
			case 'pending':
				return '🟡';
			default:
				return '⚪';
		}
	}
</script>

<div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
	<div class="px-4 py-5 sm:p-6">
		<div class="flex items-center">
			<div class="flex-shrink-0">
				<div class="w-8 h-8 bg-green-500 rounded-md flex items-center justify-center">
					<span class="text-white text-sm font-bold">🔧</span>
				</div>
			</div>
			<div class="ml-5 w-0 flex-1">
				<dl>
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">
						Running Services
					</dt>
					<dd class="text-lg font-medium text-gray-900 dark:text-white">
						{services?.length || 0} services
					</dd>
				</dl>
			</div>
		</div>

		{#if services && services.length > 0}
			<div class="mt-5">
				<div class="flow-root">
					<ul class="-my-3 divide-y divide-gray-200 dark:divide-gray-700">
						{#each services as service}
							<li class="py-3">
								<div class="flex items-center space-x-4">
									<div class="flex-shrink-0">
										<span class="text-lg">{getStatusIcon(service.status)}</span>
									</div>
									<div class="flex-1 min-w-0">
										<p class="text-sm font-medium text-gray-900 dark:text-white truncate">
											{service.name}
										</p>
										<p class="text-sm text-gray-500 dark:text-gray-400">
											PID: {service.pid}
										</p>
									</div>
									<div class="inline-flex items-center">
										<span
											class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium {getStatusColor(service.status)}"
										>
											{service.status}
										</span>
									</div>
								</div>
							</li>
						{/each}
					</ul>
				</div>
			</div>
		{:else}
			<div class="mt-5 text-center">
				<p class="text-sm text-gray-500 dark:text-gray-400">No services found</p>
			</div>
		{/if}
	</div>
</div>
