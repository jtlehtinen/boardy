<script>
  import CallToAction from '../Misc/CallToAction.svelte';
  import BarChart from '../Dashboard/BarChart.svelte';
  import LineChart from '../Dashboard/LineChart.svelte';
  import PieChart from '../Dashboard/PieChart.svelte';
</script>

<main class='h-full pb-16 overflow-y-auto'>
  <div class='container px-6 mx-auto grid'>
    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>
      Charts
    </h2>

    <CallToAction/>

    <p class='mb-8 text-gray-600 dark:text-gray-400'>
      Charts are provided by
      <a
        href='https://www.chartjs.org/'
        class='text-purple-600 dark:text-purple-400 hover:underline'
      >Chart.js</a>.
      Note that the default legends are disabled and you should
      provide a description for your charts in HTML. See source code for
      examples.
    </p>

    <div class='grid gap-6 mb-8 md:grid-cols-2'>
      <PieChart/>
      <LineChart/>
      <BarChart/>
    </div>
  </div>
</main>
