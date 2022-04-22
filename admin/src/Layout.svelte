<script>
  import DesktopSidebar from './components/Sidebar/DesktopSidebar.svelte';
  import Header from './components/Header/Header.svelte';
</script>

<div class='flex h-screen bg-gray-50 dark:bg-gray-900'>
  <DesktopSidebar/>
  <div class='flex flex-col flex-1 w-full'>
    <Header/>
    <slot></slot>
  </div>
</div>