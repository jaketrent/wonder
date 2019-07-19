<script>
  export let onClick
  export let users
  export let counts
</script>

<style>
  ul {
    display: flex;
    justify-content: space-around;
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
  @media (min-width: 48.06rem) {
    ul {
      flex-direction: column;
    }
    li {
      margin-bottom: 0.5rem;
    }
  }
</style>

<ul>
  {#each users as user}
    <li>
      <button 
        on:click|preventDefault={_ => onClick(user)}>{user.name} ({counts[user.id]})
      </button>
    </li>
  {/each}
</ul>
