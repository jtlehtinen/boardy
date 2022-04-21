<script>
  import HamburgerMenuIcon from '../Icons/HamburgerMenuIcon.svelte';
</script>

<button
  class='p-1 mr-5 -ml-1 rounded-md md:hidden focus:outline-none focus:shadow-outline-purple'
  on:click
  aria-label='Menu'
>
  <span class='w-6 h-6'><HamburgerMenuIcon/></span>
</button>
