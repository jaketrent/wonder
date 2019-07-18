<script>
  export let users = []
  let isOpen = false

  function handleClick() {
    isOpen = !isOpen
  }
</script>
<style>
  button {
    height: 32px;
    width: 32px;
    line-height: 0;
  }
  svg {
    height: 16px;
    width: 16px;
    fill: #fff;
  }
  button:hover,
  button:focus,
  button:active {
    border: 1px solid #FFBF00;
    outline: none;
    background: none;
  }
  button:hover svg,
  button:focus svg,
  button:active svg {
    fill: #FFBF00;
  }

  nav {
    position: fixed;
    top: 0;
    right: 0;
    height: 100vh;
    z-index: 1;
    transition: transform 200ms ease-in-out;
    transform: translateX(100%);
    background: #9E2F96;
    padding: 8px 32px 8px 16px;
  }
  nav.isOpen {
    transform: translateX(0);
    box-shadow: -4px 0 16px rgba(0, 0, 0, 0.3);
  }

  nav h2 {
    font-size: 1rem;
    margin: 1.5rem 0 0.25rem 0;
    color: #bd77b9;
  }
  nav a {
    display: block;
    color: inherit;
    text-decoration: none;
    margin: 0.5rem 0;
  }
  nav a:hover,
  nav a:focus,
  nav a:active {
    outline: none;
    color: #FFBF00;
    text-decoration: underline;
  }
  nav a:before {
    content: ':'
  }
</style>

<button on:click|preventDefault={handleClick}>
  <svg viewBox="0 0 16 16">
    <path d="M2.02 3.572h11.96v.972H2.02zM2.02 7.528h11.96V8.5H2.02zM2.02 11.455h11.96v.972H2.02z"/>
  </svg>
</button>

<nav class:isOpen>
  <ul>
    <li>
      <h2>
        lil' NAV
      </h2>
    </li>
    <li>
      <a href="/">create</a>
    </li>
    <li>
      <h2>
        Profiles
      </h2>
    </li>
    {#each users as user}
      <li>
        <a href="/users/{user.id}">{user.name}</a>
      </li>
    {/each}
  </ul>
</nav>
