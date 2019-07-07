<script>
  import { onMount } from 'svelte'

  export let userId
  let description = ''
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
          { userId: parseInt(userId, 10), description }
        ]
      }),
    })
    const body = await res.json()
    wonders.unshift(body)
  }
</script>

<form on:submit|preventDefault={handleSubmit}>
  Create

  <label for="description">
    Wonder
    <input bind:value={description} id="description" />
  </label>

  <button>Create</button>
</form>

<table>
  {#each wonders as wonder}
    <tr>
      <td>{wonder.created}</td>
      <td>{wonder.description}</td>
    </tr>
  {/each}
</table>
