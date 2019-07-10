<script>
  import { onMount } from 'svelte'

  let users = []
  let selectedUser
  let description = ''
  let created = formatInputDate(new Date())
  let wonders = []
  let prebaked = [
    'Make bed',
    'Get dressed',
    'Cleanup breakfast',
    'Participate in scriptures',
    'Ready for a walk',
    'Do chores',
    'School work',
    'Be kind to siblings',
    'Honor your parents',
    'Clean up lunch',
    'Clean up dinner',
    'Clean up toys',
    'Clean room',
    'Get ready for bed',
    'Stay in room',
    'Stay quiet'
  ]

  onMount(async _ => {
    const res = await fetch('/api/v1/users')
    const body = await res.json()
    users = body.data
    handleUserSelect(users[0])
  })

  async function handleUserSelect(user) {
    selectedUser = user
    const res = await fetch(`/api/v1/users/${selectedUser.id}/wonders`)
    const body = await res.json()
    wonders = body.data
  }

  async function handleWonderCreate() {
    const userId = selectedUser.id
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

<style>
  .active {
    background: red;
    color: white;
  }
</style>

<ul>
  {#each users as user}
    <li>
      <button 
        class:active="{user.id === selectedUser.id}"
        on:click|preventDefault={_ => handleUserSelect(user)}>{user.name}</button></li>
  {/each}
</ul>

<form on:submit|preventDefault={handleWonderCreate}>
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

<h2>Prebaked</h2>
<ul>
  {#each prebaked as desc}
    <li>
      <button on:click={_ => description = desc}>
      {desc}
      </button>
    </li>
  {/each}
</ul>

<h2>Wonders ({wonders.length})</h2>
<table>
  {#each wonders as wonder}
    <tr>
      <td>{formatDisplayDate(wonder.created)}</td>
      <td>{wonder.description}</td>
    </tr>
  {/each}
</table>
