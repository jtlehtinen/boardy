<script>
  import CallToAction from '../Misc/CallToAction.svelte';
  import LargerButton from './LargerButton.svelte';
  import LargeButton from './LargeButton.svelte';
  import RegularButton from './RegularButton.svelte';
  import SmallButton from './SmallButton.svelte';
  import HeartIcon from '../Icons/HeartIcon.svelte';
  import EditIcon from '../Icons/EditIcon.svelte';
</script>

<main class='h-full overflow-y-auto'>
  <div class='container grid px-6 mx-auto'>
    <h2 class='my-6 text-2xl font-semibold text-gray-700 dark:text-gray-200'>
      Buttons
    </h2>

    <CallToAction/>

    <h4 class='mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300'>
      Sizes
    </h4>

    <div class='flex flex-col flex-wrap mb-4 space-y-4 md:flex-row md:items-end md:space-x-4'>
      <div><LargerButton text={'Larger'}/></div>
      <div><LargeButton text={'Large'}/></div>
      <div><RegularButton text={'Regular'}/></div>
      <div><RegularButton text={'Regular'} disabled/></div>
      <div><SmallButton text={'Small'}/></div>
    </div>

    <p class='mb-8 text-gray-700 dark:text-gray-400'>
      Apply <code>w-full</code> to any button to create a block level button.
    </p>

    <!-- Buttons with icons -->
    <h4
      class='mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300'
    >
      Icons
    </h4>
    <div
      class='flex flex-col flex-wrap mb-8 space-y-4 md:flex-row md:items-end md:space-x-4'
    >
      <!-- Divs are used just to display the examples. Use only the button. -->
      <div>
        <button
          class='flex items-center justify-between px-4 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple'
        >
          <span>Icon right</span>
          <span class='w-4 h-4 ml-2 -mr-1'><HeartIcon/></span>
        </button>
      </div>

      <!-- Divs are used just to display the examples. Use only the button. -->
      <div>
        <button
          class='flex items-center justify-between px-4 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple'
        >
          <span class='w-4 h-4 mr-3 -ml-1'><HeartIcon/></span>
          <span>Icon left</span>
        </button>
      </div>

      <!-- Divs are used just to display the examples. Use only the button. -->
      <div>
        <button
          class='flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-lg active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple'
          aria-label='Like'
        >
          <span class='w-5 h-5'><HeartIcon/></span>
        </button>
      </div>

      <!-- Divs are used just to display the examples. Use only the button. -->
      <div>
        <button
          class='flex items-center justify-between px-2 py-2 text-sm font-medium leading-5 text-white transition-colors duration-150 bg-purple-600 border border-transparent rounded-full active:bg-purple-600 hover:bg-purple-700 focus:outline-none focus:shadow-outline-purple'
          aria-label='Edit'
        >
        <span class='w-5 h-5'><EditIcon/></span>
        </button>
      </div>
    </div>
  </div>
</main>
