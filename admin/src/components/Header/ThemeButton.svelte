<script>
  import MoonIcon from '../Icons/MoonIcon.svelte';
  import SunIcon from '../Icons/SunIcon.svelte';
  export let darkTheme = false;
</script>

<button
  on:click
  class='rounded-md focus:outline-none focus:shadow-outline-purple'
  aria-label='Toggle color mode'>

  <div class='w-5 h-5'>
  {#if darkTheme}
  <SunIcon/>
  {:else}
  <MoonIcon/>
  {/if}
  </div>

</button>
