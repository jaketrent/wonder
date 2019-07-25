<script>
  import { onMount } from 'svelte'

  import UserList from './UserList.svelte'
  import PrebakeList from './PrebakeList.svelte'
  import * as toasts from './toasts.js'

  export let users = []
  let wonders = []
  let wonderCounts = {}
  let description = ''
  let created = formatInputDate(new Date())

  onMount(async function fetchUserWonders() {
    const res = await fetch('/api/v1/user-wonders')
    const body = await res.json()
    wonders = body.data
    users.forEach(
      user => (wonderCounts[user.id] = countWondersFor(wonders, user.id))
    )
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
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        data: [
          {
            userId: parseInt(userId, 10),
            description,
            created: formatServerDate(created)
          }
        ]
      })
    })
    const body = await res.json()
    if (res.ok && Array.isArray(body.data)) {
      wonders = [...wonders, body.data]
      wonderCounts[userId] = wonderCounts[userId] + 1
      toasts.add('Toast created')
    }
  }

  function formatServerDate(str) {
    const [yyyy, mm, dd] = str.split('-')
    const date = new Date(yyyy, parseInt(mm, 10) - 1, dd)
    return date
  }

  function formatDisplayDate(str) {
    const [yyyy, mm, dd] = str
      .split('T')[0]
      .split('-')
      .map(n => parseInt(n, 10))
    const date = new Date(yyyy, mm - 1, dd)
    const month = date.toLocaleString('en-us', { month: 'short' })
    return dd + ' ' + month
  }

  function formatInputDate(date) {
    return (
      date.getFullYear() +
      '-' +
      pad(date.getMonth() + 1) +
      '-' +
      pad(date.getDate())
    )
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
  .page {
    display: grid;
    gap: 1rem;
    grid-template-rows: auto 1fr auto;
    grid-template-areas:
      'form'
      'prebake'
      'users';
  }
  label {
    grid-area: form;
  }
  input {
    font-size: 1rem;
  }
  .prebake {
    grid-area: prebake;
    padding-bottom: 5rem;
  }
  .users {
    position: fixed;
    left: 0;
    bottom: 0;
    width: 100%;
    grid-area: users;
    background: #9e2f96;
    padding: 0.875rem 0 1.25rem 0;
  }
  @media (min-width: 48.06rem) {
    .page {
      grid-template-rows: auto;
      grid-template-columns: 1fr minmax(12.5rem, auto);
      grid-template-areas:
        'form    .'
        'prebake users';
    }
    .users {
      position: static;
      padding: 0;
      background: none;
    }
  }
</style>

<div class="page">
  <label>
    Date
    <input bind:value={created} id="created" type="date" />
  </label>

  <div class="prebake">
    <PrebakeList onClick={handlePrebakeClick} />
  </div>

  <div class="users">
    <UserList {users} onClick={handleWonderCreate} counts={wonderCounts} />
  </div>
</div>
