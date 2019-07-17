<script>
  import { onMount } from 'svelte'

  import UserList from './UserList.svelte'

  let users = []
  let wonders = []
  let wonderCounts = {}
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

  async function fetchUsers() {
    const res = await fetch('/api/v1/users')
    const body = await res.json()
    return body.data
  }

  async function fetchUserWonders() {
    const res = await fetch('/api/v1/user-wonders')
    const body = await res.json()
    return body.data
  }

  onMount(async function loadData() {
    [users, wonders] = await Promise.all([fetchUsers(), fetchUserWonders()])
    users.forEach(user => wonderCounts[user.id] = countWondersFor(wonders, user.id))
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
      wonders = [...wonders, body.data]
      wonderCounts[userId] = wonderCounts[userId] + 1
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

  function countWondersFor(wonders, userId) {
    return (wonders || []).filter(w => w.userId === userId).length
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

<UserList 
  users={users} 
  onClick={handleWonderCreate} 
  counts={wonderCounts} 
/>

