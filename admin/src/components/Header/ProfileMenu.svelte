<script>
  import ProfileIcon from '../Icons/ProfileIcon.svelte';
  import GearIcon from '../Icons/GearIcon.svelte';
  import LogOutIcon from '../Icons/LogOutIcon.svelte';
  import {fade} from 'svelte/transition';
  import { clickAway } from '../../clickAway';

  const links = [
    {href:'#', iconComponent: ProfileIcon, text: 'Profile'},
    {href:'#', iconComponent: GearIcon, text: 'Settings'},
    {href:'#', iconComponent: LogOutIcon, text: 'Log Out'},
  ];
</script>

<ul
  out:fade={{duration: 150}}
  use:clickAway
  on:clickAway
  class='absolute right-0 w-56 p-2 mt-2 space-y-2 text-gray-600 bg-white border border-gray-100 rounded-md shadow-md dark:border-gray-700 dark:text-gray-300 dark:bg-gray-700'
  aria-label='submenu'
>
  {#each links as link}
    <li class='flex'>
      <a
        href={link.href}
        class='inline-flex items-center w-full px-2 py-1 text-sm font-semibold transition-colors duration-150 rounded-md hover:bg-gray-100 hover:text-gray-800 dark:hover:bg-gray-800 dark:hover:text-gray-200'
      >
        <span class='w-4 h-4 mr-3'>
          <svelte:component this={link.iconComponent}/>
        </span>
        <span>{link.text}</span>
      </a>
    </li>
  {/each}
</ul>
