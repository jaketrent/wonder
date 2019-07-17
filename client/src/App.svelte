<script>
  import { Router, Link, Route } from "svelte-routing";

  import Menu from './Menu.svelte'
  import UserDetail from './UserDetail.svelte'
  import WonderCreate from './WonderCreate.svelte'

  export let url = "";
</script>

<style>
  header {
    display: flex;
    align-items: center;
    margin-bottom: 16px;
  }
  .menu {
    margin-left: auto;
  }
  h1 {
    text-transform: uppercase;
    margin: 0;
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
      <div class="menu">
        <Menu />
      </div>
    </header>
    <div>
      <Route path="/"><WonderCreate /></Route>
    </div>
  </Router>
</div>
