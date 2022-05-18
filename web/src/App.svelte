<script lang='ts'>
  import { Router, Route } from 'svelte-routing';
  import Header from './components/Header.svelte';
  import Footer from './components/Footer.svelte';
  import ListingView from './views/ListingView.svelte';
  import DetailsView from './views/DetailsView.svelte';

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
  <div class='content'>
    <Route path='/'>
      <ListingView/>
    </Route>
    <Route path='/game/:id' let:params>
      <DetailsView id={params.id}/>
    </Route>
  </div>
</Router>
<Footer/>
