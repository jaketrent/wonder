<script>
import { fly } from 'svelte/transition'
  import { toasts } from './toasts.js'
</script>

<style>
  .toasts {
    position: fixed;
    top: 0;
    right: 0;
  }
  .toast {
    margin: 0.5rem;
    padding: 1rem;
    box-shadow: 0 0.25rem 0.5rem rgba(0, 0, 0, 0.25);
  }
  .toast--success {
    background: #549d3e;
  }
  .toast--error {
    background: #d1342b;
  }
</style>

<ul class="toasts">
  {#each $toasts as toast}
    <li transition:fly="{{ y: -100 }}" class="toast" class:toast--success={toast.status === 'success'}>
      {toast.text}
    </li>
  {/each}
</ul>
