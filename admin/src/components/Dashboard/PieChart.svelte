<script>
  import { onMount } from 'svelte';

  const config = {
    type: 'doughnut',
    data: {
      datasets: [
        {
          data: [33, 33, 33],
          /**
           * These colors come from Tailwind CSS palette
           * https://tailwindcss.com/docs/customizing-colors/#default-color-palette
           */
          backgroundColor: ['#0694a2', '#1c64f2', '#7e3af2'],
          label: 'Dataset 1',
        },
      ],
      labels: ['Shoes', 'Shirts', 'Bags'],
    },
    options: {
      responsive: true,
      cutoutPercentage: 80,
      /**
       * Default legends are ugly and impossible to style.
       * See examples in charts.html to add your own legends
       */
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

<div class='min-w-0 p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
  <h4 class='mb-4 font-semibold text-gray-800 dark:text-gray-300'>
    Revenue
  </h4>

  <canvas bind:this={canvas}></canvas>

  <div class='flex justify-center mt-4 space-x-3 text-sm text-gray-600 dark:text-gray-400'>
    <!-- Chart legend -->
    <div class='flex items-center'>
      <span class='inline-block w-3 h-3 mr-1 bg-blue-500 rounded-full'></span>
      <span>Shirts</span>
    </div>

    <div class='flex items-center'>
      <span class='inline-block w-3 h-3 mr-1 bg-teal-600 rounded-full'></span>
      <span>Shoes</span>
    </div>

    <div class='flex items-center'>
      <span class='inline-block w-3 h-3 mr-1 bg-purple-600 rounded-full'></span>
      <span>Bags</span>
    </div>
  </div>
</div>
