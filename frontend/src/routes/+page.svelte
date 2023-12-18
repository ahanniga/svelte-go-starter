<script>
  import Navbar from "../components/Navbar.svelte";
  import Footer from '../components/Footer.svelte';
</script>

<Navbar />

<div
  class="text-white container mx-auto p-8 overflow-hidden md:rounded-lg md:p-10 lg:p-12 h-fill"
>

  <div class="md:h-32 lg:h-32"></div>

  <p class="font-sans text-4xl font-bold text-gray-200 max-w-5xl lg:text-7xl lg:pr-24 md:text-6xl">
    Index Page
  </p>

  <div class="mb-8"></div>

  <p class="max-w-2xl font-sans text-xl text-gray-400 md:text-2xl">
    Some responsive pages to demonstrate the template
  </p>
  
  <div class="font-sans text-gray-400 h-32 md:h-40 pt-8">
    <p>This is a starter template that provides a Svelte front-end coupled to a Go backend.</p> 
    <p>All the code is packaged into a single, self-contained binary making it ideal for container deployments. </p>
    <p>&nbsp;</p>
    <p>It works by building the frontend using Svelte's <code class="text-green-500">adapter-static</code> plugin. </p>
    <p>The compiled output is embedded in the Go app using the <code class="text-yellow-500">//go:embed</code> directive. </p>
    <p>Single or multiple page apps are possible using simple Svelte <a class="text-green-500" href="https://learn.svelte.dev/tutorial/layouts">routing rules</a>.</p>
  </div>

  <div class="grid gap-8 md:grid-cols-2">
    <div class="flex flex-col justify-center">
    </div>
      <div
        class="-mr-24 rounded-lg md:rounded-l-full bg-gradient-to-br from-gray-900 to-black h-96"
      ></div>
  </div>


</div>  
