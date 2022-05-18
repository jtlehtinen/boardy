<script lang='ts'>
  import { Router, Route } from 'svelte-routing';
  import Header from './components/Header.svelte';
  import ListingView from './views/ListingView.svelte';

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
  </div>
</Router>
