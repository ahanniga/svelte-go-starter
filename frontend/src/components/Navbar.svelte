<!-- Navbar.svelte -->
<script>
  import MenuItems from './MenuItems.svelte';

  import { onMount } from 'svelte';
  import { browser } from '$app/environment';
  import { sidebarVisible } from '../store'

  function toggleSidebar() {
    sidebarVisible.set(!$sidebarVisible) ;
  }

  let collapse = false;
  let location = "";

  // Close the sidebar when the component is mounted (on small screens)
  onMount(() => {
    if (browser) {
      if (window.innerWidth <= 768) {
        sidebarVisible.set(false);
        collapse = true;
      } 
      location = document.location.href;
    }
  });

  // Close the sidebar when the window is resized to a larger size
  if (browser) {
    window.addEventListener('resize', () => {
      if (window.innerWidth > 768) {
        sidebarVisible.set(true);
        collapse = false;
      } else {
        sidebarVisible.set(false);
        collapse = true;
      }
    });
  }
</script>

<nav class="bg-gray-800 p-6 px-8">
  <div class="container flex justify-normal w-fill">
    <button class="{location.indexOf("spa") > 0 ? 'visible' : 'hidden'} text-white" on:click={toggleSidebar}>☰</button>
    <ul class="flex flex-col md:flex-row list-none md:ml-auto md:space-x-6 lg:space-x-10 ml-6">
    <!-- {#if !collapse} -->
    <MenuItems />
    <!-- {/if} -->
    </ul>
  </div>
</nav>

