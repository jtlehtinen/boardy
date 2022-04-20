<script>
  import PeopleIcon from '../Icons/PeopleIcon.svelte';

  export let icon;
  export let topText;
  export let bottomText;
  export let iconClass;

</script>

<div class='flex items-center p-4 bg-white rounded-lg shadow-xs dark:bg-gray-800'>
  <div class='p-3 mr-4 rounded-full {iconClass}'>
    <svelte:component this={icon}/>
  </div>
  <div>
    <p class='mb-2 text-sm font-medium text-gray-600 dark:text-gray-400'>
      {topText}
    </p>
    <p class='text-lg font-semibold text-gray-700 dark:text-gray-200'>
      {bottomText}
    </p>
  </div>
</div>
