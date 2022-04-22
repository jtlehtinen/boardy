<script>
  import { onMount } from 'svelte';

  const config = {
    type: 'bar',
    data: {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
      datasets: [
        {
          label: 'Shoes',
          backgroundColor: '#0694a2',
          // borderColor: window.chartColors.red,
          borderWidth: 1,
          data: [-3, 14, 52, 74, 33, 90, 70],
        },
        {
          label: 'Bags',
          backgroundColor: '#7e3af2',
          // borderColor: window.chartColors.blue,
          borderWidth: 1,
          data: [66, 33, 43, 12, 54, 62, 84],
        },
      ],
    },
    options: {
      responsive: true,
      legend: {
        display: false,
      },
    },
  };

  let canvas;
  let chart;

  onMount(() => {
    chart = new Chart(canvas, config);
  });

</script>

<div
  class='min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'
>
  <h4 class='mb-4 font-semibold text-gray-800 dark:text-gray-300'>
    Bars
  </h4>
  <canvas bind:this={canvas}></canvas>
  <div
    class='flex justify-center mt-4 space-x-3 text-sm text-gray-600 dark:text-gray-400'
  >
    <!-- Chart legend -->
    <div class='flex items-center'>
      <span
        class='inline-block w-3 h-3 mr-1 bg-teal-500 rounded-full'
      ></span>
      <span>Shoes</span>
    </div>
    <div class='flex items-center'>
      <span
        class='inline-block w-3 h-3 mr-1 bg-purple-600 rounded-full'
      ></span>
      <span>Bags</span>
    </div>
  </div>
</div>