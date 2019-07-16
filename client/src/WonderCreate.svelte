<script>
  import { onMount } from 'svelte'

  let users = []
  let description = ''
  let created = formatInputDate(new Date())
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
  })

  async function handleWonderCreate(user) {
    const userId = user.id
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
  }
</style>

<h2>A creation arises</h2>

<label>
  Date
  <input bind:value={created} id="created" type="date" />
</label>

<ul>
  {#each prebaked as desc}
    <li>
      <button 
        class:active="{desc === description}"
        on:click={_ => description = desc}
      >
        {desc}
      </button>
    </li>
  {/each}
</ul>

<ul>
  {#each users as user}
    <li>
      <button 
        on:click|preventDefault={_ => handleWonderCreate(user)}>{user.name} (XX)</button></li>
  {/each}
</ul>
