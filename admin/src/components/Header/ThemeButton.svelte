<script>
  import MoonIcon from '../Icons/MoonIcon.svelte';
  import SunIcon from '../Icons/SunIcon.svelte';
  export let darkTheme = false;
</script>

<button
  on:click
  class='rounded-md focus:outline-none focus:shadow-outline-purple'
  aria-label='Toggle color mode'>

  {#if darkTheme}
    <MoonIcon/>
  {:else}
    <SunIcon/>
  {/if}

</button>
