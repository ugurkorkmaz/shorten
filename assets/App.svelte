<svelte:options tag="shorten-app" />
<script>
    let url;
    let short;
    
    function shorten() {
        fetch(`/shorten?url=${url}`)
            .then(res => res.json())
            .then(data => {
                console.log(data);
                short = data.uid;
                short = location.href + short;
            });
    }
</script>
<div class="shorten">
    {#if short != null}
    <span><a href="{short}">{short}</a></span>
    {/if}
  <input class="url" contenteditable="true" placeholder="https://github.com/ugurkorkmaz/shorten" bind:value={url}/>
  <button on:click="{shorten}">SHORTEN</button>
</div>

<style>
  .shorten {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 0;
    margin: 0;
    overflow: hidden;
    width: 90vw;
    height: 90vh;
  }
  .shorten span {
    background-color: #ccc;
    margin-bottom: 12px;
    padding: 12px;
  }
  .shorten .url {
    border: 1px solid #ccc;
    border-bottom: none;
    width: 60%;
    height: 36px;
    font-size: 2em;
    margin: 0;
    overflow: hidden;
  }
  .shorten .url:focus {
    outline: none;
  }
  .shorten button {
    width: 60%;
    height: 36px;
    background-color: #ccc;
    border: none;
    font-size: 2em;
    color: #fff;
    cursor: pointer;
  }
  .shorten button:hover {
    background-color: #aaa;
  }
</style>
