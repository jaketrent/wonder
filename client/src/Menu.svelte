<script>
  import { Link } from 'svelte-routing'

  import CloseIcon from './CloseIcon.svelte'

  export let users = []
  let isOpen = false

  function toggleOpen() {
    isOpen = !isOpen
  }
</script>

<style>
  .button {
    position: relative;
    z-index: 1;
    height: 3rem;
    width: 3rem;
    line-height: 0;
    background: none;
    border: none;
  }
  .button :global(svg) {
    height: 1.5rem;
    width: 1.5rem;
    fill: #fff;
  }
  .button:hover,
  .button:focus,
  .button:active {
    border: 1px solid #9e2f96;
    outline: none;
    background: none;
  }
  .button:hover :global(svg),
  .button:focus :global(svg),
  .button:active :global(svg) {
    fill: #9e2f96;
  }
  .button.isOpen:hover :global(svg),
  .button.isOpen:focus :global(svg),
  .button.isOpen:active :global(svg) {
    fill: #ffbf00;
  }

  nav {
    position: fixed;
    top: 0;
    right: 0;
    z-index: 1;
    height: 100vh;
    transition: transform 200ms ease-in-out;
    transform: translateX(100%);
    background: #9e2f96;
    padding: 0.5rem 2rem 0.5rem 1rem;
  }
  nav.isOpen {
    transform: translateX(0);
    box-shadow: -0.25rem 0 1rem rgba(0, 0, 0, 0.3);
  }

  nav ul {
    margin-top: -0.5rem;
  }
  nav h2 {
    font-size: 1rem;
    margin: 1.5rem 0 0.25rem 0;
    color: #bd77b9;
  }
  nav :global(a) {
    display: block;
    color: inherit;
    text-decoration: none;
    margin: 0.5rem 0;
  }
  nav :global(a):hover,
  nav :global(a):focus,
  nav :global(a):active {
    outline: none;
    color: #ffbf00;
    text-decoration: underline;
  }
  nav :global(a):before {
    content: ':';
  }
</style>

<nav class:isOpen>
  <ul>
    <li>
      <h2>NAV</h2>
    </li>
    <li>
      <Link to="/" on:click={toggleOpen}>create</Link>
    </li>
    <li>
      <h2>Profiles</h2>
    </li>
    {#each users as user}
      <li>
        <Link to="/users/{user.id}" on:click={toggleOpen}>{user.name}</Link>
      </li>
    {/each}
    <li>
      <h2>Reports</h2>
    </li>
    <li>
      <Link to="/reports/dip" on:click={toggleOpen}>dip</Link>
    </li>
  </ul>
</nav>

{#if !isOpen}
  <button class="button" on:click|preventDefault={toggleOpen}>
    <svg viewBox="0 0 16 16">
      <path
        d="M2.02 3.572h11.96v.972H2.02zM2.02 7.528h11.96V8.5H2.02zM2.02
        11.455h11.96v.972H2.02z" />
    </svg>
  </button>
{/if}

{#if isOpen}
  <button class="button" class:isOpen on:click|preventDefault={toggleOpen}>
    <CloseIcon />
  </button>
{/if}
