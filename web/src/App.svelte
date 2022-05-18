<script lang='ts'>
  import { Router, Route, navigate } from 'svelte-routing';
  import Header from './components/Header.svelte';
  import Navbar from './components/Navbar.svelte';
  import Footer from './components/Footer.svelte';
  import CalendarView from './views/CalendarView.svelte';
  import ListingView from './views/ListingView.svelte';
  import LoginView from './views/LoginView.svelte';
  import DetailsView from './views/DetailsView.svelte';
  import { isLoggedIn } from './store';

  if (!$isLoggedIn) {
    navigate('/login');
  }

  export let url = '';

</script>

<style>
  .content {
    display: flex;
    padding-top: 4em;
    padding-bottom: 4em;
    justify-content: center;
    align-items: center;
  }
</style>

<Header/>
<Router url={url}>
  <Navbar/>
  <div class='content'>
    <Route path='/'>
      <ListingView/>
    </Route>
    <Route path='/calendar'>
      <CalendarView/>
    </Route>
    <Route path='/login'>
      <LoginView/>
    </Route>
    <Route path='/game/:id' let:params>
      <DetailsView id={params.id}/>
    </Route>
  </div>
</Router>
<Footer/>
