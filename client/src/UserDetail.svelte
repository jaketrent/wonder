<script>
  import { onMount } from 'svelte'

  export let userId
  let user = { name: 'el persono' }
  let wonders = []

  onMount(async function fetchUser() {
    const res = await fetch(`/api/v1/users/${userId}`)
    const body = await res.json()
    user = body.data[0]
  })

  onMount(async function fetchUserWonders() {
    const res = await fetch(`/api/v1/users/${userId}/wonders`)
    const body = await res.json()
    wonders = body.data
  })

  function formatDisplayDate(str) {
    const [yyyy, mm, dd] = str.split('T')[0].split('-').map(n => parseInt(n, 10))
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
