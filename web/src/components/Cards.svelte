<script lang='ts'>
  import Card from './Card.svelte';
  import { games } from './../../db.json'
</script>

<div class='cards'>
  {#each games as game}
    <Card title={game.name} img={game.images.medium} id={game.id}/>
  {/each}
</div>

<style>
  .cards {
    display: flex;
    flex-wrap: wrap;
    width: 80%;

    margin-left: -0.5em;
    margin-right: -0.5em;
  }
</style>
