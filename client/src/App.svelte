<script>
  import { Router, Link, Route } from "svelte-routing";

  import WonderCreate from './WonderCreate.svelte'
  import UserDetail from './UserDetail.svelte'

  export let url = "";
</script>

<style>
  h1 {
    text-transform: uppercase;
    margin: 0 0 1rem 0;
    font-size: 1.25rem;
  }
  h1:after {
    content: ':::';
  }
  nav :global(a) {
    padding: 0.25rem 1.25rem;
    border: 1px solid red;
    border-radius: 2px;
    text-decoration: none;
    color: black;
    font-size: 0.85rem;
    text-transform: uppercase;
  }
  nav :global(a):hover,
  nav :global(a):focus {
    background: red;
    color: #fff;
  }
</style>

<div class="app">
  <Router url="{url}">
    <header>
      <h1>Wonder</h1>
    </header>
    <div>
      <Route path="/"><WonderCreate /></Route>
    </div>
  </Router>
</div>
