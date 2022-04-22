<script>
  import {Route, Router} from 'svelte-routing'
  import Layout from './Layout.svelte';
  import Dashboard from './components/Dashboard/Dashboard.svelte';
  import Buttons from './components/Buttons/Buttons.svelte';
  import Forms from './components/Forms/Forms.svelte';
  import Cards from './components/Cards/Cards.svelte';
  import Charts from './components/Charts/Charts.svelte';
  import Modals from './components/Modals/Modals.svelte';
  import Tables from './components/Tables/Tables.svelte';

  import Blank from './components/Pages/Blank.svelte';
  import CreateAccount from './components/Pages/CreateAccount.svelte';
  import ForgotPassword from './components/Pages/ForgotPassword.svelte';
  import Login from './components/Pages/Login.svelte';
  import NotFound from './components/Pages/NotFound.svelte';

  export let url = '';
</script>

<Router url={url}>
  <Route path='/'>
    <Layout><Dashboard/></Layout>
  </Route>

  <Route path='/buttons'>
    <Layout><Buttons/></Layout>
  </Route>

  <Route path='/forms'>
    <Layout><Forms/></Layout>
  </Route>

  <Route path='/cards'>
    <Layout><Cards/></Layout>
  </Route>

  <Route path='/charts'>
    <Layout><Charts/></Layout>
  </Route>

  <Route path='/modals'>
    <Layout><Modals/></Layout>
  </Route>

  <Route path='/tables'>
    <Layout><Tables/></Layout>
  </Route>

  <Route path='/pages/login'>
    <Login/>
  </Route>

  <Route path='/pages/create-account'>
    <CreateAccount/>
  </Route>

  <Route path='/pages/forgot-password'>
    <ForgotPassword/>
  </Route>

  <Route path='/pages/blank'>
    <Layout><Blank/></Layout>
  </Route>

  <Route path='/pages/404'>
    <Layout><NotFound/></Layout>
  </Route>
</Router>

<style global lang='postcss'>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
