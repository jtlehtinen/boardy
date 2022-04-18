<script>
  import {fade} from 'svelte/transition';
  import { clickAway } from '../../clickAway';

  const links = [
    {href: '#', text: 'Messages', count: 13},
    {href: '#', text: 'Sales', count: 2},
    {href: '#', text: 'Alerts', count: 0},
  ];
</script>

<ul
  out:fade={{duration: 150}}
  use:clickAway
  on:clickAway
  class='absolute right-0 w-56 p-2 mt-2 space-y-2 text-gray-600 bg-white border border-gray-100 rounded-md shadow-md dark:text-gray-300 dark:border-gray-700 dark:bg-gray-700'
>

  {#each links as link}
    <li class='flex'>
      <a
        href={link.href}
        class='inline-flex items-center justify-between w-full px-2 py-1 text-sm font-semibold transition-colors duration-150 rounded-md hover:bg-gray-100 hover:text-gray-800 dark:hover:bg-gray-800 dark:hover:text-gray-200'
      >
        <span>{link.text}</span>
        {#if link.count !== 0}
          <span class='inline-flex items-center justify-center px-2 py-1 text-xs font-bold leading-none text-red-600 bg-red-100 rounded-full dark:text-red-100 dark:bg-red-600'>
            {link.count}
          </span>
        {/if}
      </a>
    </li>
  {/each}
</ul>
