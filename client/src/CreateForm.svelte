<script>
  import { onMount } from 'svelte'

  let description = ''
  let selectedUser

  let users = []

  onMount(async _ => {
    const res = await fetch('/api/v1/users')
    const body = await res.json()
    users = body.data
    selectedUser = users[0]
  })

  async function handleSubmit() {
    const userId = selectedUser.id
    const res = await fetch(`/api/v1/users/${userId}/wonders`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          data: [
            { userId, description }
          ]
        }),
    })

    const body = await res.json()

    console.log('body', body)
  }
</script>

<form on:submit|preventDefault={handleSubmit}>
  Create

  <label for="user">
    Person
    <select bind:value={selectedUser} id="user">
      {#each users as user}
        <option value={user}>{user.name}</option>
      {/each}
    </select>
  </label>

  <label for="description">
    Wonder
    <input bind:value={description} id="description" />
  </label>

  <button>Create</button>
</form>

