<script lang='ts'>
  import { navigate } from 'svelte-routing';
  import { isLoggedIn } from '../store';
  import LoginForm from '../components/LoginForm.svelte';

  function submit(username, password) {
    console.log(`login ${username}:${password}`);
    isLoggedIn.set(true);
    navigate('/');
  }
</script>

<LoginForm submit={submit}/>
