<script>
  import Status from './Status.svelte';
  import TableHead from './TableHead.svelte';
  import Pagination from '../Dashboard/Pagination.svelte';

  const rows = [
    {
      name: 'Hans Burger',
      profession: '10x Developer',
      amount: 863.45,
      status: 'Approved',
      date: ' 6/10/2020',
      imageUrl: 'https://images.unsplash.com/flagged/photo-1570612861542-284f4c12e75f?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&ixid=eyJhcHBfaWQiOjE3Nzg0fQ',
    },
    {
      name: 'Jolina Angelie',
      profession: 'Unemployed',
      amount: 369.95,
      status: 'Pending',
      date: ' 6/10/2020',
      imageUrl: 'https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&facepad=3&fit=facearea&s=707b9c33066bf8808c934c8ab394dff6',
    },
    {
      name: 'Sarah Curry',
      profession: 'Designer',
      amount: 86.00,
      status: 'Denied',
      date: ' 6/10/2020',
      imageUrl: 'https://images.unsplash.com/photo-1551069613-1904dbdcda11?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&ixid=eyJhcHBfaWQiOjE3Nzg0fQ',
    },
    {
      name: 'Rulia Joberts',
      profession: 'Actress',
      amount: 1276.45,
      status: 'Approved',
      date: ' 6/10/2020',
      imageUrl: 'https://images.unsplash.com/photo-1551006917-3b4c078c47c9?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&ixid=eyJhcHBfaWQiOjE3Nzg0fQ',
    },
    {
      name: 'Wenzel Dashington',
      profession: 'Actor',
      amount: 863.45,
      status: 'Expired',
      date: ' 6/10/2020',
      imageUrl: 'https://images.unsplash.com/photo-1546456073-6712f79251bb?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&ixid=eyJhcHBfaWQiOjE3Nzg0fQ',
    },
    {
      name: 'Dave Li',
      profession: 'Influencer',
      amount: 863.45,
      status: 'Approved',
      date: ' 6/10/2020',
      imageUrl: 'https://images.unsplash.com/photo-1502720705749-871143f0e671?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&s=b8377ca9f985d80264279f277f3a67f5',
    },
    {
      name: 'Maria Ramovic',
      profession: 'Runner',
      amount: 863.45,
      status: 'Approved',
      date: ' 6/10/2020',
      imageUrl: 'https://images.unsplash.com/photo-1531746020798-e6953c6e8e04?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&ixid=eyJhcHBfaWQiOjE3Nzg0fQ',
    },
    {
      name: 'Hitney Wouston',
      profession: 'Singer',
      amount: 863.45,
      status: 'Approved',
      date: ' 6/10/2020',
      imageUrl: 'https://images.unsplash.com/photo-1566411520896-01e7ca4726af?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=200&fit=max&ixid=eyJhcHBfaWQiOjE3Nzg0fQ',
    },
  ];

  const headings = ['Client', 'Amount', 'Status', 'Date'];

</script>

<h4 class='mb-4 text-lg font-semibold text-gray-600 dark:text-gray-300'>
  Table with avatars
</h4>

<div class="w-full mb-8 overflow-hidden rounded-lg shadow-xs">
  <div class='w-full overflow-x-auto'>
    <table class='w-full whitespace-no-wrap'>
      <TableHead {headings}/>
      <tbody class='bg-white divide-y dark:divide-gray-700 dark:bg-gray-800'>

        {#each rows as row}
        <tr class='text-gray-700 dark:text-gray-400'>
          <td class='px-4 py-3'>
            <div class='flex items-center text-sm'>
              <!-- Avatar with inset shadow -->
              <div class='relative hidden w-8 h-8 mr-3 rounded-full md:block'>
                <img
                  class='object-cover w-full h-full rounded-full'
                  src={row.imageUrl}
                  alt=''
                  loading='lazy'
                />
                <div
                  class='absolute inset-0 rounded-full shadow-inner'
                  aria-hidden='true'
                ></div>
              </div>
              <div>
                <p class='font-semibold'>{row.name}</p>
                <p class='text-xs text-gray-600 dark:text-gray-400'>
                  {row.profession}
                </p>
              </div>
            </div>
          </td>

          <td class='px-4 py-3 text-sm'>
            $ {row.amount}
          </td>
          <td class='px-4 py-3 text-xs'>
            <Status status={row.status}/>
          </td>
          <td class='px-4 py-3 text-sm'>
            {row.date}
          </td>
        </tr>
        {/each}
      </tbody>
    </table>
  </div>
  <Pagination/>
</div>