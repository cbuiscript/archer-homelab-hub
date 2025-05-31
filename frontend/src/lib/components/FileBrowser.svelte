<script lang="ts">
	import { onMount } from 'svelte';

	let files: any[] = [];
	let currentPath = '/';
	let loading = false;
	let error = '';

	const API_BASE = 'http://localhost:8080/api';

	async function loadFiles(path: string = '/') {
		try {
			loading = true;
			error = '';
			const response = await fetch(`${API_BASE}/files?path=${encodeURIComponent(path)}`);
			
			if (!response.ok) {
				throw new Error('Failed to load files');
			}

			const data = await response.json();
			files = data.files || [];
			currentPath = data.path || path;
		} catch (err) {
			error = err instanceof Error ? err.message : 'Failed to load files';
		} finally {
			loading = false;
		}
	}

	function formatFileSize(bytes: number): string {
		if (bytes === 0) return '0 Bytes';
		const k = 1024;
		const sizes = ['Bytes', 'KB', 'MB', 'GB'];
		const i = Math.floor(Math.log(bytes) / Math.log(k));
		return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
	}

	function formatDate(dateString: string): string {
		return new Date(dateString).toLocaleDateString();
	}

	function getFileIcon(fileName: string, isDir: boolean): string {
		if (isDir) return '📁';
		
		const ext = fileName.split('.').pop()?.toLowerCase();
		switch (ext) {
			case 'txt': case 'md': case 'readme': return '📄';
			case 'js': case 'ts': case 'json': return '📜';
			case 'html': case 'css': return '🌐';
			case 'png': case 'jpg': case 'jpeg': case 'gif': return '🖼️';
			case 'pdf': return '📑';
			case 'zip': case 'tar': case 'gz': return '📦';
			default: return '📄';
		}
	}

	function navigateToParent() {
		const parentPath = currentPath.split('/').slice(0, -1).join('/') || '/';
		loadFiles(parentPath);
	}

	function navigateToFile(file: any) {
		if (file.is_dir) {
			const newPath = currentPath === '/' ? `/${file.name}` : `${currentPath}/${file.name}`;
			loadFiles(newPath);
		}
	}

	onMount(() => {
		loadFiles();
	});
</script>

<div class="bg-white dark:bg-gray-800 overflow-hidden shadow rounded-lg">
	<div class="px-4 py-5 sm:p-6">
		<div class="flex items-center justify-between">
			<div class="flex items-center">
				<div class="flex-shrink-0">
					<div class="w-8 h-8 bg-yellow-500 rounded-md flex items-center justify-center">
						<span class="text-white text-sm font-bold">📁</span>
					</div>
				</div>
				<div class="ml-5">
					<h3 class="text-lg font-medium text-gray-900 dark:text-white">File Browser</h3>
					<p class="text-sm text-gray-500 dark:text-gray-400">Current: {currentPath}</p>
				</div>
			</div>
			<button
				on:click={() => loadFiles(currentPath)}
				class="inline-flex items-center px-3 py-2 border border-gray-300 dark:border-gray-600 shadow-sm text-sm leading-4 font-medium rounded-md text-gray-700 dark:text-gray-200 bg-white dark:bg-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
			>
				Refresh
			</button>
		</div>

		{#if loading}
			<div class="mt-5 flex justify-center">
				<div class="animate-spin rounded-full h-8 w-8 border-b-2 border-indigo-600"></div>
			</div>
		{:else if error}
			<div class="mt-5 bg-red-50 border border-red-200 rounded-md p-4">
				<p class="text-sm text-red-700">{error}</p>
			</div>
		{:else}
			<div class="mt-5">
				<!-- Navigation -->
				{#if currentPath !== '/'}
					<div class="mb-3">
						<button
							on:click={navigateToParent}
							class="inline-flex items-center text-sm text-indigo-600 hover:text-indigo-500"
						>
							← Back to parent directory
						</button>
					</div>
				{/if}

				<!-- Files List -->
				{#if files && files.length > 0}
					<div class="bg-gray-50 dark:bg-gray-700 rounded-lg overflow-hidden">
						<div class="divide-y divide-gray-200 dark:divide-gray-600">
							{#each files as file}
								<div
									class="p-3 hover:bg-gray-100 dark:hover:bg-gray-600 cursor-pointer transition-colors"
									class:cursor-pointer={file.is_dir}
									on:click={() => navigateToFile(file)}
									on:keydown={(e) => e.key === 'Enter' && navigateToFile(file)}
									role="button"
									tabindex="0"
								>
									<div class="flex items-center justify-between">
										<div class="flex items-center space-x-3">
											<span class="text-lg">{getFileIcon(file.name, file.is_dir)}</span>
											<div>
												<p class="text-sm font-medium text-gray-900 dark:text-white">
													{file.name}
												</p>
												<p class="text-xs text-gray-500 dark:text-gray-400">
													{#if file.is_dir}
														Directory
													{:else}
														{formatFileSize(file.size)} • {formatDate(file.modified)}
													{/if}
												</p>
											</div>
										</div>
										{#if file.is_dir}
											<span class="text-gray-400">→</span>
										{/if}
									</div>
								</div>
							{/each}
						</div>
					</div>
				{:else}
					<div class="text-center py-8">
						<span class="text-4xl">📂</span>
						<p class="mt-2 text-sm text-gray-500 dark:text-gray-400">
							This directory is empty
						</p>
					</div>
				{/if}
			</div>
		{/if}
	</div>
</div>
