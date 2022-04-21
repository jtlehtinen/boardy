<script>
  export let text = '@TODO';
  export let disabled = false;

  const disabledClasses = 'opacity-50 cursor-not-allowed';
  const enabledClasses = 'active:bg-purple-600 hover:bg-purple-700 focus:shadow-outline-purple';
</script>

<button
  on:click
  class='{disabled ? disabledClasses : enabledClasses} px-5 py-3 font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg focus:outline-none'
>
  {text}
</button>
