<script>
  export let text = '';
  export let href = '';
</script>

<a  href={href} class='inline-flex items-center w-full text-sm font-semibold transition-colors duration-150 hover:text-gray-800 dark:hover:text-gray-200'>
  <slot name='icon'></slot>
  <span class='ml-4'>{text}</span>
</a>
