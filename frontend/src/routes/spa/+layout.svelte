<script>
  import Navbar from '../../components/Navbar.svelte';
  import Sidebar from '../../components/Sidebar.svelte';
  import Footer from '../../components/Footer.svelte';
</script>

<style>
</style>

<div class="flex h-screen ">
  <Sidebar />
  <div class="flex-1 flex flex-col overflow-hidden">
    <Navbar />
    <main class="flex-1 overflow-x-hidden overflow-y-auto ">
        <slot />
    </main>
    <Footer />
  </div>
</div>
