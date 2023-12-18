<script>
  import { showComponent } from "../../store";
  import { onMount, onDestroy } from "svelte";
  import Component_1 from "../../components/Component-1.svelte";
  import Component_2 from "../../components/Component-2.svelte";
  import Component_3 from "../../components/Component-3.svelte";

  let comp = "";
  const unsubscribe = showComponent.subscribe((value) => {
    comp = value;
  });
  onDestroy(unsubscribe);

  let weatherInfo = [];

  onMount(async () => {
    getWeather();
  });

  const getWeather = () => {
    fetch("/api/weather", {
      method: "GET",
      headers: { Accept: "application/json" },
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error(response.error);
        }
        return response.json();
      })
      .then((data) => {
        if (
          data !== null &&
          data.forecast !== null &&
          typeof data.forecast === "object"
        ) {
          if (
            data.forecast.forecastday !== null &&
            data.forecast.forecastday.length > 0
          ) {
            weatherInfo = data.forecast.forecastday[0].hour;
          } else {
            console.error(
              "Unhandled weather data structure (forecast.forecastday)",
            );
          }
        } else {
          console.error("Unhandled weather data structure (forecast)");
        }
      })
      .catch(function (error) {
        console.error("Problem: " + error);
      });
  };
</script>

<div
  class="text-gray-300 container mx-auto p-8 overflow-hidden md:rounded-lg md:p-10 lg:p-12"
>
  <p
    class="font-sans text-4xl font-bold text-gray-200 max-w-5xl lg:text-5xl lg:pr-24 md:text-5xl sm:pb-4 pb-8"
  >
    Single Page App
  </p>

  {#if comp == "comp-1"}
    <Component_1 />
  {:else if comp == "comp-2"}
    <Component_2 />
  {:else if comp == "comp-3"}
    <Component_3 />
  {:else}
    <p>(No component selected)</p>
  {/if}


  <p
  class="font-sans font-bold text-gray-200 text-2xl"
  >
  Backend data fetch
  </p>

  <div class="block w-full mt-5 mb-5 overflow-x-auto">
    <table class="items-center w-full bg-transparent border-collapse">
      <thead>
        <tr>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left text-blueGray-200 border-blueGray-600"
          >
            Time
          </th>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left text-blueGray-200 border-blueGray-600"
          >
            Summary
          </th>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left text-blueGray-200 border-blueGray-600"
          >
            Icon
          </th>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left text-blueGray-200 border-blueGray-600"
          >
            Temp (C)
          </th>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left text-blueGray-200 border-blueGray-600"
          >
            Humidity (%)
          </th>
          <th
            class="px-6 align-middle border border-solid py-3 text-xs uppercase border-l-0 border-r-0 whitespace-nowrap font-semibold text-left text-blueGray-200 border-blueGray-600"
          >
            Wind Dir
          </th>
        </tr>
      </thead>
      <tbody>
        {#each weatherInfo as row}
          <tr>
            <td
              class="border-t-0 px-2 align-middle border-l-0 border-r-0 text-sm whitespace-nowrap p-1"
            >
              <pre>{row.time}</pre>
            </td>
            <td
              class="border-t-0 px-2 align-middle border-l-0 border-r-0 text-sm whitespace-nowrap p-1"
            >
              {row.condition.text}
            </td>
            <td
              class="border-t-0 px-2 align-middle border-l-0 border-r-0 text-sm whitespace-nowrap p-1"
            >
              <img width="36" alt="..." src={row.condition.icon} />
            </td>
            <td
              class="border-t-0 px-2 align-middle border-l-0 border-r-0 text-sm whitespace-nowrap p-1"
            >
              {row.temp_c}
            </td>
            <td
              class="border-t-0 px-2 align-middle border-l-0 border-r-0 text-sm whitespace-nowrap p-1"
            >
              {row.humidity}
            </td>
            <td
              class="border-t-0 px-2 align-middle border-l-0 border-r-0 text-sm whitespace-nowrap p-1"
            >
              {row.wind_dir}
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
  <p class="max-w font-sans text-gray-400 text-right">
    REST test powered by <a
      href="https://www.weatherapi.com"
      class="text-green-600"
      target="_blank">WeatherAPI.com</a
    >
  </p>
</div>
