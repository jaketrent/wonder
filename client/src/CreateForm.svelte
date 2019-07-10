<script>
  export let onSubmit
  export let description
  export let created
</script>

<style>
  form {
    background: pink;
  }
</style>

<form on:submit|preventDefault={onSubmit}>
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

