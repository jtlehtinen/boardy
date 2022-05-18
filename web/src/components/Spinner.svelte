<script>
  import FaSpinner from 'svelte-icons/fa/FaSpinner.svelte'
</script>

<div class='icon spin'>
  <FaSpinner/>
</div>

<style>
  .spin {
    animation: spin 1.2s linear infinite;
  }

  @keyframes spin {
    fron {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
  }
</style>
