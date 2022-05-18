<script>
  import { Link } from 'svelte-routing';
</script>

<nav>
  <Link to='/'>Home</Link>
  <Link to='/calendar'>Calendar</Link>
</nav>

<style>
  nav {
    overflow: hidden;
    background-color: #333;
    padding-left: 10vw;
  }

  :global(nav a) {
    float: left;
    display: block;
    color: white;
    text-align: center;
    padding: 14px 20px;
    text-decoration: none;
  }
</style>
