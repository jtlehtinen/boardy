<script>
  import MdWarning from 'svelte-icons/md/MdWarning.svelte'
  import FaEye from 'svelte-icons/fa/FaEye.svelte'
  import FaEyeSlash from 'svelte-icons/fa/FaEyeSlash.svelte'
  import Spinner from './Spinner.svelte';

  export let submit = null;

  let showPassword = false;
  let passwordElement = null;

  function toggleShowPassword() {
    showPassword = !showPassword;
    const type = showPassword ? "text" : "password";
    passwordElement.setAttribute("type", type);
  }

  let errors = {};
  let loggingIn = false;

  function submitInternal(event) {
    const form = event.target;
    const username = form.username.value;
    const password = form.password.value;

    errors = {};
    if (username === '') {
      errors['username'] = 'Please enter your username.';
    }

    if (password === '') {
      errors['password'] = 'Please enter your password.';
    }

    const valid = (Object.keys(errors).length === 0);
    if (valid && submit) {
      loggingIn = true;
      setTimeout(() => {
        submit(username, password);
        // @TODO: The submit function should return promise.
        // Where should we handle the error?
        loggingIn = false;
      }, 1000);
      return true;
    }
    return true;
  }
</script>

<div class='container'>
  <form on:submit|preventDefault={submitInternal} novalidate>
    <input type='text' name='username' placeholder='👤 Username'  spellcheck='false'>
    {#if errors.username}<span class='error'><div class='icon'><MdWarning/></div>{errors.username}</span><br>{/if}

    <div class='input-container'>
      <input bind:this={passwordElement} type='password' name='password' placeholder='🔒 Password' spellcheck='false'>
      <span class='eye' on:click={toggleShowPassword} class:hide={showPassword}><FaEye/></span>
      <span class='eye eye-hide' on:click={toggleShowPassword} class:hide={!showPassword}><FaEyeSlash/></span>
    </div>

    {#if errors.password}<span class='error'><div class='icon'><MdWarning/></div>{errors.password}</span><br>{/if}

    <button type='submit'>
      {#if loggingIn}
        <Spinner/>
      {:else}
        Log In
      {/if}
    </button>
  </form>
</div>

<style>
  .container {
    width: 40%;
  }

  .eye {
    color: #EB984E;
    position: absolute;
    right: 1em;
    height: 1.4em;
    width: auto;
  }

  .eye-hide {
    transform: translateX(0.1em);
  }

  .input-container {
    position: relative;
    align-items: center;
    display: flex;
    width: 100%;
  }

  .hide {
    display: none !important;
  }

  .error {
    color: #ffb400;
    font-size: 0.8em;
    font-style: italic;
  }

  input[type=text], input[type=password] {
    width: 100%;
    padding: 12px;
    margin: 8px 0;
    display: inline-block;
    border: 1px solid #ccc;
    font-size: 1em;
    font-weight: 500;
  }

  button {
    background-color: #EB984E;
    color: white;
    cursor: pointer;
    border: none;
    font-weight: 500;
    font-size: 1em;
    margin: 8px 0;
    padding: 14px 20px;
    width: 100%;
  }

  button:hover {
    opacity: 0.8;
  }
</style>
