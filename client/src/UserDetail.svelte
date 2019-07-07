<script>
  import { onMount } from 'svelte'

  export let userId
  let description = ''
  let created = formatInputDate(new Date())
  let wonders = []

  onMount(async _ => {
    const res = await fetch(`/api/v1/users/${userId}/wonders`)
    const body = await res.json()
    wonders = body.data
  })

  async function handleSubmit() {
    const res = await fetch(`/api/v1/users/${userId}/wonders`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        data: [
          { userId: parseInt(userId, 10), description, created: formatServerDate(created) }
        ]
      }),
    })
    const body = await res.json()
    if (res.ok && Array.isArray(body.data)) {
    wonders = body.data.concat(wonders)
    }
  }

  function formatServerDate(str) {
    const [yyyy, mm, dd] = str.split('-')
    const date = new Date(yyyy, parseInt(mm, 10) - 1, dd)
    return date
  }

  function formatDisplayDate(str) {
    const [yyyy, mm, dd] = str.split('T')[0].split('-').map(n => parseInt(n, 10))
    const date = new Date(yyyy, mm - 1, dd)
    const month = date.toLocaleString('en-us', { month: 'short' })
    return dd + ' ' + month
  }

  function formatInputDate(date) {
    return date.getFullYear() + '-' + pad(date.getMonth() + 1) + '-' + pad(date.getDate())
  }

  function pad(n) {
    return ('0' + n).slice(-2)
  }
</script>

<form on:submit|preventDefault={handleSubmit}>
  Create

  <label for="description">
    Wonder
    <input bind:value={description} id="description" />
  </label>

  <label for="created">
    Date
    <input bind:value={created} id="created" type="date" />
  </label>

  <button>Create</button>
</form>

<table>
  {#each wonders as wonder}
    <tr>
      <td>{formatDisplayDate(wonder.created)}</td>
      <td>{wonder.description}</td>
    </tr>
  {/each}
</table>
