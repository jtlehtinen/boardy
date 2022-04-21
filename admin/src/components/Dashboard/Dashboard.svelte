<script>
  import CallToAction from '../Misc/CallToAction.svelte';
  import Cards from './Cards.svelte';
  import Table from './Table.svelte';
  import Pagination from './Pagination.svelte';
  import Charts from './Charts.svelte';
</script>

<main class='h-full overflow-y-auto'>
  <div class='container px-6 mx-auto grid'>
    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>
      Dashboard
    </h2>

    <CallToAction/>
    <Cards/>

    <div class='w-full overflow-hidden rounded-lg shadow-xs'>
      <Table/>
      <Pagination/>
    </div>
    <Charts/>
  </div>
</main>
