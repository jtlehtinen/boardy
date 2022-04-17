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

  const links = [
    {href: 'forms.html', text: 'Forms', iconComponent: FormIcon},
    {href: 'cards.html', text: 'Cards', iconComponent: CardsIcon},
    {href: 'charts.html', text: 'Charts', iconComponent: ChartsIcon},
    {href: 'buttons.html', text: 'Buttons', iconComponent: ButtonsIcon},
    {href: 'modals.html', text: 'Modals', iconComponent: ModalsIcon},
    {href: 'tables.html', text: 'Tables', iconComponent: TablesIcon},
  ];

  let isPagesMenuOpen = false;
  function togglePagesMenu() {
    isPagesMenuOpen = !isPagesMenuOpen;
  }
</script>

<aside class='z-20 hidden w-64 overflow-y-auto bg-white dark:bg-gray-800 md:block flex-shrink-0'>
  <div class='py-4 text-gray-500 dark:text-gray-400'>
    <a class='ml-6 text-lg font-bold text-gray-800 dark:text-gray-200' href='#'>Boardy</a>

    <ul class='mt-6'>
      <li class='relative px-6 py-3'>
        <span class='absolute inset-y-0 left-0 w-1 bg-purple-600 rounded-tr-lg rounded-br-lg' aria-hidden='true'></span>
        <a href='index.html' class='inline-flex items-center w-full text-sm font-semibold text-gray-800 transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200 dark:text-gray-100'>
          <HomeIcon/>
          <span class='ml-4'>Dashboard</span>
        </a>
      </li>
    </ul>

    <ul>

      {#each links as link}
      <li class='relative px-6 py-3'>
        <SidebarLink href={link.href} text={link.text}>
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
            <PagesIcon/>
            <span class='ml-4'>Pages</span>
          </span>
          <DownArrowIcon/>
        </button>

        {#if isPagesMenuOpen}
        <PagesMenu/>
        {/if}
      </li>
    </ul>

    <div class='px-6 my-6'>
      <button class='flex items-center justify-between w-full px-4 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple'>
        Create account
        <span class='ml-2' aria-hidden='true'>+</span>
      </button>
    </div>
  </div>
</aside>
