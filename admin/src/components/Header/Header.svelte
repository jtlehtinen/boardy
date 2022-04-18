<script>
  import Search from './Search.svelte';
  import HamburgerButton from './HamburgerButton.svelte';
  import ThemeButton from './ThemeButton.svelte';
  import NotificationsButton from './NotificationsButton.svelte';
  import NotificationsMenu from './NotificationsMenu.svelte';
  import ProfileButton from './ProfileButton.svelte';
  import ProfileMenu from './ProfileMenu.svelte';

  // @TODO: mouse and key events properly

  let showSideMenu = false;
  function toggleSideMenu() {
    // @TODO: handle toggling the side menu...
    showSideMenu = !showSideMenu;
  }

  let darkTheme = false; // @TODO: get theme from local storage
  function toggleTheme() {
    darkTheme = !darkTheme;
    // @TODO: set theme to local storage
  }

  let showNotificationsMenu = false;
  function toggleNotificationsMenu() {
    showNotificationsMenu = !showNotificationsMenu;
  }
  function closeNotificationsMenu() {
    showNotificationsMenu = false;
  }

  let showProfileMenu = false;
  function toggleProfileMenu() {
    showProfileMenu = !showProfileMenu;
  }
  function closeProfileMenu() {
    showProfileMenu = false;
  }
</script>

<header class='z-10 py-4 bg-white shadow-md dark:bg-gray-800'>
  <div
    class='container flex items-center justify-between h-full px-6 mx-auto text-purple-600 dark:text-purple-300'
  >
    <HamburgerButton on:click={toggleSideMenu}/>
    <Search/>

    <ul class='flex items-center flex-shrink-0 space-x-6'>
      <li class='flex'>
        <ThemeButton on:click={toggleTheme} {darkTheme}/>
      </li>

      <li class='relative'>
        <NotificationsButton
          on:click={toggleNotificationsMenu}
          on:keydown={e => e.key === 'Escape' && closeNotificationsMenu()}
        />
        {#if showNotificationsMenu}
          <NotificationsMenu
            on:clickAway={closeNotificationsMenu}
            on:keydown={e => e.key === 'Escape' && closeNotificationsMenu()}
          />
        {/if}
      </li>

      <li class='relative'>
        <ProfileButton
          on:click={toggleProfileMenu}
          on:keydown={e => e.key === 'Escape' && closeProfileMenu()}
        />
        {#if showProfileMenu}
          <ProfileMenu
            on:clickAway={closeProfileMenu}
            on:keydown={e => e.key === 'Escape' && closeNotificationsMenu()}
          />
        {/if}
      </li>
    </ul>
  </div>
</header>
