<script>
  import { onMount } from 'svelte'
  import { Router, Link, Route } from 'svelte-routing'

  import Menu from './Menu.svelte'
  import DipReport from './DipReport.svelte'
  import Toasts from './Toasts.svelte'
  import UserDetail from './UserDetail.svelte'
  import CreatePage from './CreatePage.svelte'

  export let url = ''
  let users = []

  onMount(async function fetchUsers() {
    const res = await fetch('/api/v1/users')
    const body = await res.json()
    users = body.data
  })
</script>

<style>
  header {
    display: flex;
    align-items: center;
    margin-bottom: 1rem;
  }
  .menu {
    margin-left: auto;
  }
  h1 {
    text-transform: uppercase;
    margin: 0;
    font-size: 1.25rem;
    color: #000;
  }
  h1 :global(a) {
    color: inherit;
    text-decoration: none;
  }
  h1 :global(a:hover),
  h1 :global(a:focus) {
    color: #9e2f96;
    border-bottom: 1px solid #ffbf00;
    outline: none;
  }
  h1:after {
    content: ':::';
    color: #fff;
  }
  .logo {
    position: relative;
    top: 0.25rem;
    right: -0.25rem;
    height: 1.5rem;
  }
  nav :global(a) {
    padding: 0.25rem 1.25rem;
    border: 1px solid red;
    border-radius: 0.125rem;
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
  <Router {url}>
    <header>
      <h1>
        <img class="logo" alt="" src="/palm.png" />
        <Link to="/">Wonder</Link>
      </h1>
      <div class="menu">
        <Menu {users} />
      </div>
    </header>
    <div>
      <Route path="/">
        <CreatePage {users} />
      </Route>
      <Route path="/users/:userId" component={UserDetail} />
      <Route path="/reports/dip" component={DipReport} {users} />
    </div>
    <Toasts />
  </Router>
</div>
