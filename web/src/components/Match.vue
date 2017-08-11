<template>
  <div>
    <hero
      :title="match === 404 ? 'Not found' : (match === null ? 'Loading' : match.Name)"
      class="is-info"></hero>
    <section class="section">
      <div class="container">
        <router-link class="button" to="/" style="margin-bottom: 1.5rem">
          <span class="icon is-small">
            <i class="fa fa-arrow-left"></i>
          </span>
          <span>Go back</span>
        </router-link>
        <div class="box" v-if="games.length > 0">
          <p>Latest played games at the top!</p>
          <br>
          <game
            v-for="game in games"
            :key="game.ID"
            :game="game"
            :users="users"
            :beatmaps="beatmaps">
          </game>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import store from '../store'
import api from '../api'
import Hero from './Hero'
import Game from './Game'

// loads new users and beatmaps
const loadNewEntities = component => {
  let users = new Set()
  let beatmaps = new Set()
  component.games.forEach(game => {
    if (typeof component.beatmaps[game.BeatmapID] === 'undefined') {
      beatmaps.add(game.BeatmapID)
    }
    Object.keys(game.Scores).forEach(val => {
      if (typeof component.users[val] === 'undefined') {
        users.add(val)
      }
    })
  })
  api.loadEntities(component, [...users], [...beatmaps])
}

const loadGames = component => {
  api.games(component.match.ID, games => {
    component.games = games
    loadNewEntities(component)
  })
}

let evtSrc

export default {
  components: {
    Hero,
    Game
  },
  data () {
    return {
      match: store.state.clickedMatch,
      games: [],
      beatmaps: {},
      users: {}
    }
  },
  mounted () {
    if (this.match === null) {
      api.match(this.$route.params.id, matches => {
        this.match = matches.length > 0 ? matches[0] : 404
        loadGames(this)
      })
    } else {
      loadGames(this)
    }
    evtSrc = new EventSource('/api/new_games')
    evtSrc.onmessage = e => {
      if (e.data === 'h') {
        // heartbeat
        return
      }
      let newGame = JSON.parse(e.data)
      if (newGame.MatchID === this.$route.params.id) {
        this.games.unshift(newGame)
        loadNewEntities(this)
      }
    }
  },
  destroyed () {
    evtSrc && evtSrc.close()
  }
}
</script>

<style>
@font-face {
    font-family: "OsuFont";
    src: url("../assets/osu-font.woff") format('woff');
}

.osu-icon {
  display: inline-block;
  font: normal normal normal 14px/1 OsuFont;
  font-size: inherit;
  text-rendering: auto;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
.mode-0:before { content: '\e800' }
.mode-1:before { content: '\e803' }
.mode-2:before { content: '\e801' }
.mode-3:before { content: '\e802' }
.osu-logo:before { content: '\e805' }
</style>
