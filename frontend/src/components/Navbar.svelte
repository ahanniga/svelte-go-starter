<script>
  import Link from "./Link.svelte";
    let navbarOpen = false;

  function setNavbarOpen() {
    navbarOpen = !navbarOpen;
  }

</script>

<nav
  class="top-0 fixed z-50 w-full flex flex-wrap items-center justify-between px-2 py-3 navbar-expand-lg bg-green-600 shadow"
>
  <div
    class="container px-4 mx-auto flex flex-wrap items-center justify-between"
  >
    <div
      class="w-full relative flex justify-between lg:w-auto lg:static lg:block lg:justify-start"
    >
      <Link
        to="/"
        clazz="text-white  font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap"
      >
        Svelte Go Template
      </Link>
      <button
      class="text-white cursor-pointer text-xl leading-none px-3 py-1 border border-solid border-transparent rounded bg-transparent block lg:hidden outline-none focus:outline-none"
      type="button"
      on:click="{setNavbarOpen}"
    >
      <i class="fa fa-bars"></i>
    </button>
    </div>

    <div
      class="lg:flex flex-grow items-center {navbarOpen ? 'block':'hidden'}"
      id="example-navbar-warning"
    >



    <ul class="flex flex-col lg:flex-row list-none lg:ml-auto">
      <li class="flex items-center">
        <Link
          to="/page-one"
          clazz="text-white sm font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap"
        >
          Page-one
        </Link>
      </li>
      <li class="flex items-center">
        <Link
          to="/page-two"
          clazz="text-white sm font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap"
        >
          Page-two
        </Link>
      </li>
      <li class="flex items-center">
        <Link
          to="/page-three"
          clazz="text-white  font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap"
        >
          Page-three
        </Link>
      </li>
      <li class="flex items-center">
        <Link
          to="/not-found"
          clazz="text-white  font-bold leading-relaxed inline-block mr-4 py-2 whitespace-nowrap"
        >
          404
        </Link>
      </li>
    </ul>
  </div>
</nav>
