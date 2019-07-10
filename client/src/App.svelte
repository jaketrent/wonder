<script>
  import { Router, Link, Route } from "svelte-routing";

  import WonderCreate from './WonderCreate.svelte'
  import UserDetail from './UserDetail.svelte'

  export let url = "";
</script>

<style>
	h1 {
		color: orange;
	}
</style>

<h1>Dad's Independent Summer of Wonder</h1>

<Router url="{url}">
  <nav>
    <Link to="/">Home</Link>
  </nav>
  <div>
    <Route path="/"><WonderCreate /></Route>
    <Route path="/users/:userId" component="{UserDetail}" />
  </div>
</Router>
