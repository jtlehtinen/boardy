<script lang='ts'>
  import { navigate } from 'svelte-routing';
  import { Link } from 'svelte-routing';
  import { isLoggedIn } from '../store';

  function logOut() {
    isLoggedIn.set(false);
    navigate('/login');
  }
</script>

<nav>
  <Link to='/'>Home</Link>
  <Link to='/calendar'>Calendar</Link>

  {#if $isLoggedIn}
    <div class='logout' on:click={logOut}>Log Out</div>
  {:else}
    <Link to='/login'>Log In</Link>
  {/if}
</nav>

<style>
  nav {
    overflow: hidden;
    background-color: #333;
    padding-left: 10vw;
    padding-right: 10vw;
  }

  :global(nav a), .logout {
    float: left;
    display: block;
    color: white;
    text-align: center;
    padding: 14px 20px;
    text-decoration: none;
  }

  :global(nav a:hover), .logout:hover {
    background-color: #ddd;
    color: black;
  }

  .logout {
    float: right;
  }
</style>
