<script>
  export let headings = [];
</script>

<thead>
  <tr class='text-xs font-semibold tracking-wide text-left text-gray-500 uppercase border-b dark:border-gray-700 bg-gray-50 dark:text-gray-400 dark:bg-gray-800'>
    {#each headings as heading}
    <th class='px-4 py-3'>{heading}</th>
    {/each}
  </tr>
</thead>
