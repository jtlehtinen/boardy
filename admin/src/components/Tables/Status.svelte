<script>
  export let status = '';
</script>

{#if status === 'Approved'}
<span
  class='px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-full dark:bg-green-700 dark:text-green-100'
>
  Approved
</span>
{:else if status === 'Pending'}
<span
  class='px-2 py-1 font-semibold leading-tight text-orange-700 bg-orange-100 rounded-full dark:text-white dark:bg-orange-600'
>
  Pending
</span>
{:else if status === 'Denied'}
<span
  class='px-2 py-1 font-semibold leading-tight text-red-700 bg-red-100 rounded-full dark:text-red-100 dark:bg-red-700'
>
  Denied
</span>
{:else if status === 'Expired'}
<span
  class='px-2 py-1 font-semibold leading-tight text-gray-700 bg-gray-100 rounded-full dark:text-gray-100 dark:bg-gray-700'
>
  Expired
</span>
{/if}