<script>
  export let onClick
  export let users
  export let counts
</script>

<style>
  ul {
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    display: flex;
    justify-content: space-around;
    background: #9E2F96;
    margin: 0;
    padding: 12px 0;
  }
  button {
    color: #FFBF00;
    border: 1px solid #FFBF00;
    background: none;
    font-weight: 600;
  }
  button:hover,
  button:focus,
  button:active {
    background: #FFBF00;
    color: #9E2F96;
    outline: none;
  }
</style>

<ul>
  {#each users as user}
    <li>
      <button 
        on:click|preventDefault={_ => onClick(user)}>{user.name} ({counts[user.id]})</button></li>
  {/each}
</ul>
