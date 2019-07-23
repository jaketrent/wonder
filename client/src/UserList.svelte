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
  li {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  span {
    margin: 0 0 0.5rem 0;
    color: #151515;
    font-size: 0.875rem;
  }
  button {
    color: #ffbf00;
    border: 1px solid #ffbf00;
    background: none;
    font-size: 1rem;
    font-weight: 600;
  }
  button:hover,
  button:focus,
  button:active {
    background: #ffbf00;
    color: #9e2f96;
    outline: none;
  }
  @media (min-width: 34rem) {
    button {
      font-size: 1.5rem;
    }
  }
  @media (min-width: 48.06rem) {
    ul {
      flex-direction: column;
    }
    li {
      flex-direction: row-reverse;
      justify-content: flex-end;
      margin-bottom: 0.5rem;
    }
    span {
      margin: 0 0 0 0.5rem;
    }
  }
</style>

<ul>
  {#each users as user}
    <li>
      <span>({counts[user.id]})</span>
      <button on:click|preventDefault={_ => onClick(user)}>{user.name}</button>
    </li>
  {/each}
</ul>
