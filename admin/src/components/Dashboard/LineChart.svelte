<script>
  import { onMount } from 'svelte';

  const config = {
    type: 'line',
    data: {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
      datasets: [
        {
          label: 'Organic',
          /**
           * These colors come from Tailwind CSS palette
           * https://tailwindcss.com/docs/customizing-colors/#default-color-palette
           */
          backgroundColor: '#0694a2',
          borderColor: '#0694a2',
          data: [43, 48, 40, 54, 67, 73, 70],
          fill: false,
        },
        {
          label: 'Paid',
          fill: false,
          /**
           * These colors come from Tailwind CSS palette
           * https://tailwindcss.com/docs/customizing-colors/#default-color-palette
           */
          backgroundColor: '#7e3af2',
          borderColor: '#7e3af2',
          data: [24, 50, 64, 74, 52, 51, 65],
        },
      ],
    },
    options: {
      responsive: true,
      /**
       * Default legends are ugly and impossible to style.
       * See examples in charts.html to add your own legends
       */
      legend: {
        display: false,
      },
      tooltips: {
        mode: 'index',
        intersect: false,
      },
      hover: {
        mode: 'nearest',
        intersect: true,
      },
      scales: {
        x: {
          display: true,
          scaleLabel: {
            display: true,
            labelString: 'Month',
          },
        },
        y: {
          display: true,
          scaleLabel: {
            display: true,
            labelString: 'Value',
          },
        },
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
    Traffic
  </h4>

  <canvas bind:this={canvas}></canvas>

  <div class='flex justify-center mt-4 space-x-3 text-sm text-gray-600 dark:text-gray-400'>
    <!-- Chart legend -->
    <div class='flex items-center'>
      <span class='inline-block w-3 h-3 mr-1 bg-teal-600 rounded-full'></span>
      <span>Organic</span>
    </div>

    <div class='flex items-center'>
      <span class='inline-block w-3 h-3 mr-1 bg-purple-600 rounded-full'></span>
      <span>Paid</span>
    </div>
  </div>
</div>