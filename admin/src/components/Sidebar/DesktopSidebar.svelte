<script>
  import ButtonsIcon from './../Icons/ButtonsIcon.svelte';
  import CardsIcon from './../Icons/CardsIcon.svelte';
  import ChartsIcon from './../Icons/ChartsIcon.svelte';
  import DownArrowIcon from './../Icons/DownArrowIcon.svelte';
  import FormIcon from './../Icons/FormIcon.svelte';
  import HomeIcon from './../Icons/HomeIcon.svelte';
  import ModalsIcon from './../Icons/ModalsIcon.svelte';
  import PagesIcon from './../Icons/PagesIcon.svelte';
  import TablesIcon from './../Icons/TablesIcon.svelte';
  import PagesMenu from './PagesMenu.svelte';
  import SidebarLink from './SidebarLink.svelte';
  import CreateAccountButton from './CreateAccountButton.svelte';

  const links = [
    {href: '/forms', text: 'Forms', iconComponent: FormIcon},
    {href: '/cards', text: 'Cards', iconComponent: CardsIcon},
    {href: '/charts', text: 'Charts', iconComponent: ChartsIcon},
    {href: '/buttons', text: 'Buttons', iconComponent: ButtonsIcon},
    {href: '/modals', text: 'Modals', iconComponent: ModalsIcon},
    {href: '/tables', text: 'Tables', iconComponent: TablesIcon},
  ];

  let showPagesMenu = false;
  function togglePagesMenu() {
    showPagesMenu = !showPagesMenu;
  }

  const pathname = window.location.pathname;
</script>

<aside class='z-20 hidden w-64 overflow-y-auto bg-white dark:bg-gray-800 md:block flex-shrink-0'>
  <div class='py-4 text-gray-500 dark:text-gray-400'>
    <a class='ml-6 text-lg font-bold text-gray-800 dark:text-gray-200' href='/'>Boardy</a>

    <ul class='mt-6'>
      <li class='relative px-6 py-3'>
        {#if pathname === '/'}
        <span class='absolute inset-y-0 left-0 w-1 bg-purple-600 rounded-tr-lg rounded-br-lg' aria-hidden='true'></span>
        {/if}
        <a href='/' class:text-gray-800={pathname === '/'} class='inline-flex items-center w-full text-sm font-semibold text-gray-800 transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200 dark:text-gray-100'>
          <span class='w-5 h-5'><HomeIcon/></span>
          <span class='ml-4'>Dashboard</span>
        </a>
      </li>
    </ul>

    <ul>

      {#each links as link}
      <li class='relative px-6 py-3'>
        {#if pathname === link.href}
          <span class='absolute inset-y-0 left-0 w-1 bg-purple-600 rounded-tr-lg rounded-br-lg' aria-hidden='true'></span>
        {/if}
        <SidebarLink href={link.href} text={link.text} selected={pathname === link.href}>
          <svelte:component this={link.iconComponent} slot='icon'/>
        </SidebarLink>
      </li>
      {/each}

      <li class='relative px-6 py-3'>
        <button class='inline-flex items-center justify-between w-full text-sm font-semibold transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200'
          on:click={togglePagesMenu}
          aria-haspopup='true'
        >
          <span class='inline-flex items-center'>
            <span class='w-5 h-5'><PagesIcon/></span>
            <span class='ml-4'>Pages</span>
          </span>
          <span class='w-4 h-4'><DownArrowIcon/></span>
        </button>

        {#if showPagesMenu}
        <PagesMenu/>
        {/if}
      </li>
    </ul>

    <CreateAccountButton/>
  </div>
</aside>
