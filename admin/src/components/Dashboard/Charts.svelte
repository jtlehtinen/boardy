<script>
  import BarChart from './BarChart.svelte';
  import LineChart from './LineChart.svelte';
  import PieChart from './PieChart.svelte';
</script>

<h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>
  Charts
</h2>
<div class='grid gap-6 mb-8 md:grid-cols-2'>
  <PieChart/>
  <LineChart/>
</div>
