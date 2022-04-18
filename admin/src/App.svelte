<script>
  import {Route, Router} from 'svelte-routing'
  import DesktopSidebar from './components/Sidebar/DesktopSidebar.svelte';
  import Header from './components/Header/Header.svelte';

  export let url = '';
</script>

<div class='flex h-screen bg-gray-50 dark:bg-gray-900'>
  <!--
  <Router url={url}>
  </Router>
-->
  <DesktopSidebar/>
  <div class="flex flex-col flex-1 w-full">
    <Header/>
  </div>

</div>

<style global lang='postcss'>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
