<script>
  import {Route, Router} from 'svelte-routing'
  import DesktopSidebar from './components/Sidebar/DesktopSidebar.svelte';
  import Header from './components/Header/Header.svelte';
  import Dashboard from './components/Dashboard/Dashboard.svelte';
  import Buttons from './components/Buttons/Buttons.svelte';

  export let url = '';
</script>

<div class='flex h-screen bg-gray-50 dark:bg-gray-900'>
  <DesktopSidebar/>
  <div class="flex flex-col flex-1 w-full">
    <Header/>
    <Router url={url}>
      <Route path='/'>
        <Dashboard/>
      </Route>

      <Route path='/buttons'>
        <Buttons/>
      </Route>
    </Router>
  </div>
</div>

<style global lang='postcss'>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
