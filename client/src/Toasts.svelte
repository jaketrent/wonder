<script>
  import { fly } from 'svelte/transition'

  import CloseIcon from './CloseIcon.svelte'
  import { remove, toasts } from './toasts.js'
</script>

<style>
  .toasts {
    position: fixed;
    z-index: 1;
    top: 0;
    right: 0;
  }
  .toast {
    display: flex;
    align-items: center;
    margin: 0.5rem;
    padding: 1rem;
    box-shadow: 0 0.25rem 0.5rem rgba(0, 0, 0, 0.25);
    max-width: 30rem;
  }
  .toast--success {
    background: #549d3e;
  }
  .toast--error {
    background: #d1342b;
  }

  .close {
    height: 1.5rem;
    width: 1.5rem;
    margin-left: 0.25rem;
    line-height: 0;
    background: none;
    border: none;
  }
  .close :global(svg) {
    height: 1rem;
    width: 1rem;
    fill: #fff;
  }
  .close:hover,
  .close:focus,
  .close:active {
    outline: none;
  }
  .close:hover :global(svg),
  .close:focus :global(svg),
  .close:active :global(svg) {
    fill: #ffbf00;
  }
</style>

<ul class="toasts">
  {#each $toasts as toast}
    <li
      transition:fly={{ y: -100 }}
      class="toast"
      class:toast--success={toast.status === 'success'}>
      {toast.text}
      <button class="close" on:click|preventDefault={_ => remove(toast.id)}>
        <CloseIcon />
      </button>
    </li>
  {/each}
</ul>
