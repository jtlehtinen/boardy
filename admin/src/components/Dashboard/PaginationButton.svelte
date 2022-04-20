<script>
  export let text = '';
  export let selected = false;
</script>

{#if selected}
<button class='px-3 py-1 text-white transition-colors duration-150 bg-purple-600 border border-r-0 border-purple-600 rounded-md focus:outline-none focus:shadow-outline-purple'>
  {text}
</button>
{:else}
<button class='px-3 py-1 rounded-md focus:outline-none focus:shadow-outline-purple'>
  {text}
</button>
{/if}
