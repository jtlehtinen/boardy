<script>
  export let game = null;
</script>

{#if game}
<div class='container'>
  <div class='media'>
    <img src={game.images.medium} alt='@TODO'/>
  </div>
  <div class='content'>
    <h2>{game.name}</h2>
    <hr>
    <!--
      @TODO: XSS...
      The description sucks a bit since it contains
      html tags.
    -->
    {@html game.description}

    <table>
      <tr>
        <th>Year Published:</th>
        <td>{game.year_published}</td>
      </tr>
      <tr>
        <th>Playtime:</th>
        <td>{game.min_playtime} - {game.max_playtime} Minutes</td>
      </tr>
      <tr>
        <th>Suggested Age:</th>
        <td>{game.min_age}</td>
      </tr>
      <tr>
        <th>Categories:</th>
        <td>{game.categories.join(", ")}</td>
      </tr>
      <tr>
        <th>Mechanics:</th>
        <td>{game.mechanics.join(", ")}</td>
      </tr>
    </table>
  </div>
</div>
{/if}

<style>
  img {
    max-width: 400px;
    max-height: 400px;
  }

  table {
    padding-top: 2em;
    text-align: left;
    border-collapse: separate;
    border-spacing: 0 1em;
  }

  th {
    white-space: nowrap;
  }

  td {
    padding-left: 2em;
  }

  .container {
    display: flex;
    align-items: start;
    width: 80%;
  }

  .media {
    margin-right: 20px;
    width: 400px;
  }

  .content {
    flex: 1;
  }
</style>
