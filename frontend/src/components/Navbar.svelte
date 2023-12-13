<script>
  import Link from "./Link.svelte";
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
    </div>
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
