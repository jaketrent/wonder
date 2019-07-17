<script>
  export let onClick
  export let users
  export let counts
</script>
<style>
</style>
<ul>
  {#each users as user}
    <li>
      <button 
        on:click|preventDefault={_ => onClick(user)}>{user.name} ({counts[user.id]})</button></li>
  {/each}
</ul>
