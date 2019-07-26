<script>
  import { onMount } from 'svelte'

  import * as toasts from './toasts.js'

  export let userId
  let user = { name: 'el persono' }
  let wonders = []

  onMount(async function fetchUser() {
    const res = await fetch(`/api/v1/users/${userId}`)
    const body = await res.json()
    if (res.ok) {
      user = body.data[0]
    } else {
      toasts.add({ text: 'User fetch failed', status: 'error' })
    }
  })

  onMount(async function fetchUserWonders() {
    const res = await fetch(`/api/v1/users/${userId}/wonders`)
    const body = await res.json()
    if (res.ok) {
      wonders = body.data
    } else {
      toasts.add({ text: 'User wonders fetch failed', status: 'error' })
    }
  })

  function formatDisplayDate(str) {
    const [yyyy, mm, dd] = str
      .split('T')[0]
      .split('-')
      .map(n => parseInt(n, 10))
    const date = new Date(yyyy, mm - 1, dd)
    const month = date.toLocaleString('en-us', { month: 'short' })
    return dd + ' ' + month
  }
</script>

<h2>{user.name} ({wonders.length})</h2>
<table>
  {#each wonders as wonder}
    <tr>
      <td>{formatDisplayDate(wonder.created)}</td>
      <td>{wonder.description}</td>
    </tr>
  {/each}
</table>
