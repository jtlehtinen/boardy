<script>
  import CallToAction from '../Misc/CallToAction.svelte';
  import Table from './Table.svelte';
  import Pagination from '../Dashboard/Pagination.svelte';
  import TableWithActions from './TableWithActions.svelte';

</script>

<main class='h-full pb-16 overflow-y-auto'>
  <div class='container grid px-6 mx-auto'>
    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>
      Tables
    </h2>
    <CallToAction/>

    <!-- With avatar -->
    <Table/>

    <!-- With actions -->
    <TableWithActions/>
  </div>
</main>