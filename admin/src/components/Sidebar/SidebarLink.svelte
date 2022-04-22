<script>
  import {link} from 'svelte-routing';

  export let text = '';
  export let href = '';
  export let selected = false;
</script>

<a use:link href={href}  class:text-gray-800={selected} class='inline-flex items-center w-full text-sm font-semibold transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200'>
  <span class='w-5 h-5'><slot name='icon'></slot></span>
  <span class='ml-4'>{text}</span>
</a>
