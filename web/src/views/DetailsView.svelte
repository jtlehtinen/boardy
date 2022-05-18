<script>
  import GameDetails from '../components/GameDetails.svelte';
  import {games} from '../../db.json'

  export let id = '';

  const game = games.find(game => game.id === id);
</script>

<GameDetails game={game}/>
