<script lang="ts">
	interface Service {
		name: string;
		status: string;
		pid: number;
	}

	let { services }: { services: Service[] } = $props();

	// Filter and search state
	let searchTerm = $state('');
	let statusFilter = $state('all');
	let sortBy = $state('name'); // 'name', 'status', 'pid'
	let sortOrder = $state('asc'); // 'asc', 'desc'

	function getStatusColor(status: string): string {
		switch (status.toLowerCase()) {
			case 'running':
				return 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-300';
			case 'stopped':
				return 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-300';
			case 'pending':
				return 'bg-yellow-100 text-yellow-800 dark:bg-yellow-900 dark:text-yellow-300';
			default:
				return 'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-300';
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

	// Filtered and sorted services
	const filteredServices = $derived(() => {
		if (!services) return [];
		
		return services
			.filter((service: Service) => {
				// Filter by search term
				const matchesSearch = service.name.toLowerCase().includes(searchTerm.toLowerCase());
				// Filter by status
				const matchesStatus = statusFilter === 'all' || service.status.toLowerCase() === statusFilter;
				return matchesSearch && matchesStatus;
			})
			.sort((a: Service, b: Service) => {
				let comparison = 0;
				
				switch (sortBy) {
					case 'name':
						comparison = a.name.localeCompare(b.name);
						break;
					case 'status':
						comparison = a.status.localeCompare(b.status);
						break;
					case 'pid':
						comparison = a.pid - b.pid;
						break;
				}
				
				return sortOrder === 'asc' ? comparison : -comparison;
			});
	});

	// Get unique statuses for filter dropdown
	const uniqueStatuses = $derived(() => {
		if (!services) return [];
		return [...new Set(services.map((s: Service) => s.status.toLowerCase()))];
	});

	function toggleSort(field: string) {
		if (sortBy === field) {
			sortOrder = sortOrder === 'asc' ? 'desc' : 'asc';
		} else {
			sortBy = field;
			sortOrder = 'asc';
		}
	}

	function getSortIcon(field: string): string {
		if (sortBy !== field) return '↕️';
		return sortOrder === 'asc' ? '⬆️' : '⬇️';
	}

	function clearFilters() {
		searchTerm = '';
		statusFilter = 'all';
		sortBy = 'name';
		sortOrder = 'asc';
	}
</script>

<div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
	<div class="px-4 py-5 sm:p-6">
		<div class="flex items-center justify-between">
			<div class="flex-shrink-0">
				<div class="w-8 h-8 bg-green-500 rounded-md flex items-center justify-center">
					<span class="text-white text-sm font-bold">🔧</span>
				</div>
			</div>
			<div class="ml-5 w-0 flex-1">
				<dl>
					<dt class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">
						Services & Processes
					</dt>
					<dd class="text-lg font-medium text-gray-900 dark:text-white">
						{filteredServices().length || 0} of {services?.length || 0} services
					</dd>
				</dl>
			</div>
			<div class="flex-shrink-0">
				<button
					onclick={clearFilters}
					class="text-sm text-indigo-600 hover:text-indigo-500 dark:text-indigo-400 dark:hover:text-indigo-300"
					title="Clear all filters"
				>
					Clear Filters
				</button>
			</div>
		</div>

		<!-- Filter Controls -->
		<div class="mt-4 space-y-3">
			<div class="flex flex-col sm:flex-row gap-3">
				<!-- Search Input -->
				<div class="flex-1">
					<label for="search" class="sr-only">Search services</label>
					<div class="relative">
						<div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
							<span class="text-gray-400 text-sm">🔍</span>
						</div>
						<input
							id="search"
							type="text"
							bind:value={searchTerm}
							placeholder="Search services..."
							class="block w-full pl-10 pr-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md leading-5 bg-white dark:bg-gray-700 text-gray-900 dark:text-white placeholder-gray-500 dark:placeholder-gray-400 focus:outline-none focus:placeholder-gray-400 focus:ring-1 focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
						/>
					</div>
				</div>
				
				<!-- Status Filter -->
				<div class="sm:w-40">
					<label for="status-filter" class="sr-only">Filter by status</label>
					<select
						id="status-filter"
						bind:value={statusFilter}
						class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-white focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
					>
						<option value="all">All Status</option>
						{#each uniqueStatuses() as status}
							<option value={status}>{status.charAt(0).toUpperCase() + status.slice(1)}</option>
						{/each}
					</select>
				</div>
			</div>

			<!-- Sort Controls -->
			<div class="flex items-center space-x-4 text-sm">
				<span class="text-gray-500 dark:text-gray-400">Sort by:</span>
				<button
					onclick={() => toggleSort('name')}
					class="inline-flex items-center space-x-1 text-indigo-600 hover:text-indigo-500 dark:text-indigo-400 dark:hover:text-indigo-300"
				>
					<span>Name</span>
					<span class="text-xs">{getSortIcon('name')}</span>
				</button>
				<button
					onclick={() => toggleSort('status')}
					class="inline-flex items-center space-x-1 text-indigo-600 hover:text-indigo-500 dark:text-indigo-400 dark:hover:text-indigo-300"
				>
					<span>Status</span>
					<span class="text-xs">{getSortIcon('status')}</span>
				</button>
				<button
					onclick={() => toggleSort('pid')}
					class="inline-flex items-center space-x-1 text-indigo-600 hover:text-indigo-500 dark:text-indigo-400 dark:hover:text-indigo-300"
				>
					<span>PID</span>
					<span class="text-xs">{getSortIcon('pid')}</span>
				</button>
			</div>
		</div>

		{#if filteredServices() && filteredServices().length > 0}
			<div class="mt-5">
				<div class="flow-root">
					<ul class="-my-3 divide-y divide-gray-200 dark:divide-gray-700">
						{#each filteredServices() as service}
							<li class="py-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg px-2 transition-colors">
								<div class="flex items-center space-x-4">
									<div class="flex-shrink-0">
										<span class="text-lg">{getStatusIcon(service.status)}</span>
									</div>
									<div class="flex-1 min-w-0">
										<p class="text-sm font-medium text-gray-900 dark:text-white truncate">
											{service.name}
										</p>
										<p class="text-sm text-gray-500 dark:text-gray-400">
											PID: {service.pid || 'N/A'}
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
		{:else if services && services.length > 0}
			<div class="mt-5 text-center py-8">
				<div class="text-4xl mb-2">🔍</div>
				<p class="text-sm text-gray-500 dark:text-gray-400 mb-2">No services match your filters</p>
				<button
					onclick={clearFilters}
					class="text-sm text-indigo-600 hover:text-indigo-500 dark:text-indigo-400 dark:hover:text-indigo-300"
				>
					Clear filters to see all services
				</button>
			</div>
		{:else}
			<div class="mt-5 text-center py-8">
				<div class="text-4xl mb-2">⚙️</div>
				<p class="text-sm text-gray-500 dark:text-gray-400">No services found</p>
			</div>
		{/if}
	</div>
</div>
