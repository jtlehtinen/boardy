<script>
  import NotificationBadge from './NotificationBadge.svelte';
  import BellIcon from '../Icons/BellIcon.svelte';
</script>

<button
  on:click
  on:keydown
  class='relative align-middle rounded-md focus:outline-none focus:shadow-outline-purple'
  aria-label='Notifications'
  aria-haspopup='true'
>
  <div class='w-5 h-5'><BellIcon/></div>
  <NotificationBadge/>
</button>
