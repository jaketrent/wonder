<script>
  import { onMount } from 'svelte'

  let users = []

  onMount(async _ => {
    const res = await fetch('/api/v1/users')
    const body = await res.json()
    users = body.data
  })
</script>

<ul>
  {#each users as user}
    <li><a href="/users/{user.id}">{user.name}</a></li>
  {/each}
</ul>
