<script>
  import { onMount } from 'svelte'

  import UserList from './UserList.svelte'
  import PrebakeList from './PrebakeList.svelte'

  export let users = []
  let wonders = []
  let wonderCounts = {}
  let description = ''
  let created = formatInputDate(new Date())

  onMount(async function fetchUserWonders() {
    const res = await fetch('/api/v1/user-wonders')
    const body = await res.json()
    wonders = body.data
    users.forEach(user => wonderCounts[user.id] = countWondersFor(wonders, user.id))
  })

  async function handleWonderCreate(user) {
    function validateNewWonder() {
      return created && description && user
    }

    if (!validateNewWonder()) return

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

  function handlePrebakeClick(desc) {
    description = desc
  }
</script>

<style>
</style>

<label>
  Date
  <input bind:value={created} id="created" type="date" />
</label>

<PrebakeList onClick={handlePrebakeClick} />

<UserList 
  users={users} 
  onClick={handleWonderCreate} 
  counts={wonderCounts} 
/>

