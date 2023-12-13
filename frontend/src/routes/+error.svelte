<script>
    import { page } from "$app/stores";
</script>

<svelte:head>
    <title>{$page.status}</title>
</svelte:head>

<div class="grid h-screen place-content-center px-4 bg-gray-900">
    <div class="text-center">
      <h1 class="text-8xl text-gray-200 dark:text-gray-700">{$page.status}</h1>
      <p class="mt-4 mb-4 text-4xl text-gray-400">{$page.error.message}</p>
      <a
        href="/"
        class="mt-10 inline-block rounded bg-green-600 px-5 py-3 text-sm font-medium text-white hover:bg-green  -700 focus:outline-none focus:ring"
      >
        Home
      </a>
    </div>
  </div>
  