<script>
  import PeopleIcon from '../Icons/PeopleIcon.svelte';
  import AccountBalanceIcon from '../Icons/AccountBalanceIcon.svelte';
  import ShoppingCartIcon from '../Icons/ShoppingCartIcon.svelte';
  import ChatBubbleIcon from '../Icons/ChatBubbleIcon.svelte';
  import Card from './Card.svelte';
</script>

<div class='grid gap-6 mb-8 md:grid-cols-2 xl:grid-cols-4'>
  <Card
    icon={PeopleIcon}
    iconClass={'text-orange-500 bg-orange-100 dark:text-orange-100 dark:bg-orange-500'}
    topText={'Total clients'}
    bottomText={'6389'}
  />

  <Card
    icon={AccountBalanceIcon}
    iconClass={'text-green-500 bg-green-100 dark:text-green-100 dark:bg-green-500'}
    topText={'Account balance'}
    bottomText={'$ 46,760.89'}
  />

  <Card
    icon={ShoppingCartIcon}
    iconClass={'text-blue-500 bg-blue-100  dark:text-blue-100 dark:bg-blue-500'}
    topText={'New sales'}
    bottomText={'376'}
  />

  <Card
    icon={ChatBubbleIcon}
    iconClass={'text-teal-500 bg-teal-100 dark:text-teal-100 dark:bg-teal-500'}
    topText={'Pending contacts'}
    bottomText={'35'}
  />
</div>
