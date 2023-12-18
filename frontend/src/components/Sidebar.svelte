<!-- Sidebar.svelte -->
<script>
  import { sidebarVisible, showComponent } from "../store";
  import { onDestroy } from "svelte";
  import { browser } from "$app/environment";
    import MenuItems from "./MenuItems.svelte";

  let visible = false;

  const unsubscribe1 = sidebarVisible.subscribe((value) => {
    visible = value;
  });
  onDestroy(unsubscribe1);

  function toggleSidebar() {
    sidebarVisible.set(!$sidebarVisible);
  }

</script>

<div
  class="{browser && window.innerWidth <= 768
    ? 'absolute top-0 bottom-0'
    : ''} {visible ? 'visible' : 'hidden'} bg-green-600 w-64"
>
  <!-- Sidebar content goes here -->
  <div class="p-4 px-6">
    <div class="flex justify-start mb-4">
      <img src="/assets/img/svelte.png" width="40" alt="..." />
      <h2 class="text-white text-xl font-semibold px-4 pt-2"><a href="/">Go Starter</a></h2>
      <button
        class="visible md:hidden lg:hidden text-white -mt-2"
        on:click={toggleSidebar}>☰</button
      >
    </div>

    <hr class="h-8 mt-2">

    <ul
      class="md:flex-col md:min-w-full flex flex-col list-none md:mb-4 space-y-3"
    >
      <li class="items-center">
        <a class="text-white hover:text-green-300  font-bold block" href="##" on:click={()=>{showComponent.set("comp-1")}}>
          <i class="fa fa-caret-right text-white mr-2"></i>
          Component 1
        </a>
      </li>
      <li class="items-center">
        <a class="text-white hover:text-green-300  font-bold block" href="##" on:click={()=>{showComponent.set("comp-2")}} >
          <i class="fa fa-caret-right text-white mr-2"></i>
          Component 2
        </a>
      </li>

      <li class="items-center">
        <a class="text-white hover:text-green-300  font-bold block" href="##" on:click={()=>{showComponent.set("comp-3")}} >
          <i class="fa fa-caret-right text-white mr-2"></i>
          Component 3
        </a>
      </li>
    </ul>
    <hr class="h-8 mt-8">
    <ul class="flex flex-col list-none space-y-3">
      <MenuItems />
    </ul>
  </div>
</div>
