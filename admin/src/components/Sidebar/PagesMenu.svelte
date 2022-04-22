<script>
  import {slide} from 'svelte/transition';
  import {link} from 'svelte-routing';

  const pages = [
    {href: '/pages/login', text: 'Login'},
    {href: '/pages/create-account', text: 'Create account'},
    {href: '/pages/forgot-password', text: 'Forgot password'},
    {href: '/pages/404', text: '404'},
    {href: '/pages/blank', text: 'Blank'},
  ];
</script>

<ul transition:slide class='p-2 mt-2 space-y-2 overflow-hidden text-sm font-medium text-gray-500 rounded-md shadow-inner bg-gray-50 dark:text-gray-400 dark:bg-gray-900' aria-label='submenu'>
  {#each pages as page}
  <li class='px-2 py-1 transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200'>
    <a use:link class='w-full' href={page.href}>{page.text}</a>
  </li>
  {/each}
</ul>
