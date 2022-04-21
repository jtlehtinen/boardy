<script>
  import PaginationButton from './PaginationButton.svelte';
  import LeftArrowHeadIcon from '../Icons/LeftArrowHeadIcon.svelte';
  import RightArrowHeadIcon from '../Icons/RightArrowHeadIcon.svelte';

  let from = 21;
  let to = 30;
  let total = 100;

</script>

<div class='grid px-4 py-3 text-xs font-semibold tracking-wide text-gray-500 uppercase border-t dark:border-gray-700 bg-gray-50 sm:grid-cols-9 dark:text-gray-400 dark:bg-gray-800'>
  <span class='flex items-center col-span-3'>
    Showing {from}-{to} of {total}
  </span>
  <span class='col-span-2'></span>

  <!-- Pagination -->
  <span class='flex col-span-4 mt-2 sm:mt-auto sm:justify-end'>
    <nav aria-label='Table navigation'>
      <ul class='inline-flex items-center'>

        <li>
          <button
            class='px-3 py-1 rounded-md rounded-l-lg focus:outline-none focus:shadow-outline-purple'
            aria-label='Previous'
          >
            <div class='w-4 h-4'><LeftArrowHeadIcon/></div>
          </button>
        </li>

        <li><PaginationButton text={'1'}/></li>
        <li><PaginationButton text={'2'}/></li>
        <li><PaginationButton text={'3'} selected={true}/></li>
        <li><PaginationButton text={'4'}/></li>
        <li><span class='px-3 py-1'>...</span></li>
        <li><PaginationButton text={'8'}/></li>
        <li><PaginationButton text={'9'}/></li>

        <li>
          <button
            class='px-3 py-1 rounded-md rounded-r-lg focus:outline-none focus:shadow-outline-purple'
            aria-label='Next'
          >
            <div class='w-4 h-4'><RightArrowHeadIcon/></div>
          </button>
        </li>
      </ul>
    </nav>
  </span>
</div>
