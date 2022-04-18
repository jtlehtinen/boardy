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
  <BellIcon/>
  <NotificationBadge/>
</button>
