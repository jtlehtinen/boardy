<script lang='ts'>
  import { onMount } from 'svelte';

  import Cards from '../components/Cards.svelte';
  import Spinner from '../components/Spinner.svelte';

  let showSpinner = true;

  onMount(() => {
    setTimeout(() => showSpinner = false, 5000)
  });

</script>

<div class='container'>
  {#if showSpinner}
    <Spinner/>
  {:else}
    <Cards/>
  {/if}
</div>

<style>
  .container {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
  }
</style>
