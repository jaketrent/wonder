<script>
  export let onClick
  export let selected
  let prebaked = [
    'Make bed',
    'Get dressed',
    'Cleanup breakfast',
    'Participate in scriptures',
    'Ready for a walk',
    'Do chores',
    'School work',
    'Be kind to siblings',
    'Honor your parents',
    'Clean up lunch',
    'Clean up dinner',
    'Clean up toys',
    'Clean room',
    'Get ready for bed',
    'Stay in room',
    'Stay quiet'
  ]
</script>
<style>
  li {
    margin-bottom: 0.5rem;
  }
  button {
    position: relative;
    width: 100%;
    text-align: left;
    border: 1px solid #9E2F96;
    background: none;
    padding: 0.5rem;
    font-weight: 600;
    font-size: 1rem;
  }
  button.active {
    background: #9E2F96;
    color: #FFBF00;
  }
  button:hover,
  button:focus,
  button:active {
    outline: none;
  }
  button:hover:after,
  button:focus:after,
  button:active:after {
    content: '';
    position: absolute;
    top: -0.25rem;
    right: -0.25rem;
    bottom: -0.25rem;
    left: -0.25rem;
    border: 0.1875rem solid #FFBF00;
  }
</style>
<ul>
  {#each prebaked as desc}
    <li>
      <button 
        class:active="{desc === selected}"
        on:click={_ => {
          selected = desc
          onClick(desc)
        }}
      >
        {desc}
      </button>
    </li>
  {/each}
</ul>
