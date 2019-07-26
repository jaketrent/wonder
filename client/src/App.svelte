<script>
  import { onMount } from 'svelte'
  import { Router, Link, Route } from 'svelte-routing'

  import CreatePage from './CreatePage.svelte'
  import Menu from './Menu.svelte'
  import DipReport from './DipReport.svelte'
  import * as toasts from './toasts.js'
  import Toasts from './Toasts.svelte'
  import UserDetail from './UserDetail.svelte'
  import WonderLogo from './WonderLogo.svelte'

  export let url = ''
  let users = []

  onMount(async function fetchUsers() {
    const res = await fetch('/api/v1/users')
    const body = await res.json()
    if (res.ok) {
      users = body.data
    } else {
      toasts.add({ text: 'Users fetch failed', status: 'error' })
    }
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
    display: flex;
    align-items: center;
    text-transform: uppercase;
    margin: 0 0 0 -1rem;
    font-size: 1.25rem;
    color: #000;
    line-height: 0;
  }
  h1 img {
    transform: translate(0.5rem, 0.75rem);
  }
  h1 :global(a) {
    margin-left: 0.25rem;
    color: inherit;
    text-decoration: none;
  }
  h1 :global(a:hover),
  h1 :global(a:focus) {
    color: #9e2f96;
    border-bottom: 1px solid #ffbf00;
    outline: none;
  }
  h1 :global(svg) {
    height: 5rem;
    margin: -1rem 0;
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
        <Link to="/">
        <WonderLogo />
        </Link>
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
