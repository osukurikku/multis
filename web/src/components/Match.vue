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
          <article class="media" v-for="game in games" :key="game.ID">
            <figure class="media-left" v-if="beatmaps[game.BeatmapID]">
              <img :src="'https://b.ppy.sh/thumb/' + beatmaps[game.BeatmapID].beatmapset_id + 'l.jpg'" style="width: 100px; max-height: 220px">
            </figure>
            <div class="media-content">
              <div class="content">
                <p>
                  <strong>{{ game.Name }}</strong><br>
                  <span v-if="beatmaps[game.BeatmapID]">
                    <span :class="'osu-icon mode-' + game.GameMode"></span>
                    <a :href="'https://osu.ppy.sh/b/' + game.BeatmapID" target="_blank">
                      {{ beatmaps[game.BeatmapID].song_name }}
                    </a>
                    {{ getScoreMods(game.Mods) }}
                  </span>
                </p>
              </div>
              <div class="columns is-multiline is-mobile">
                <div class="column is-half mp-user-column" v-for="(score, uid) in game.Scores" :key="uid">
                  <a :href="'https://ripple.moe/u/' + uid" target="_blank" class="force32">
                    <img :src="'https://a.ripple.moe/' + uid" class="force32">
                  </a>
                  <span style="margin-left: 5px">
                    {{ users[uid] ? users[uid].username : "..." }}
                    &mdash;
                    {{ score.Score.toLocaleString() }}
                    <small>{{ getScoreMods(score.Mods) }}</small>
                  </span>
                </div>
              </div>
            </div>
            <div class="media-right">
              <timeago :since="game.CreatedAt" :title="game.CreatedAt"></timeago>
            </div>
          </article>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import store from '../store'
import api from '../api'
import Hero from './Hero'

const getScoreMods = m => {
  var r = []

  // has nc => remove dt
  if ((m & 512) === 512) {
    m = m & ~64
  }
  // has pf => remove sd
  if ((m & 16384) === 16384) {
    m = m & ~32
  }

  modsString.forEach(function (v, idx) {
    var val = 1 << idx
    if ((m & val) > 0) {
      r.push(v)
    }
  })
  if (r.length > 0) {
    return '+ ' + r.join(', ')
  }
  return ''
}

const modsString = [
  'NF',
  'EZ',
  'NV',
  'HD',
  'HR',
  'SD',
  'DT',
  'RX',
  'HT',
  'NC',
  'FL',
  'AU', // Auto.
  'SO',
  'AP', // Autopilot.
  'PF',
  'K4',
  'K5',
  'K6',
  'K7',
  'K8',
  'K9',
  'RN', // Random
  'LM', // LastMod. Cinema?
  'K9',
  'K0',
  'K1',
  'K3',
  'K2'
]

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
    Hero
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
      this.games.unshift(newGame)
      loadNewEntities(this)
    }
  },
  destroyed () {
    evtSrc && evtSrc.close()
  },
  methods: {
    getScoreMods: getScoreMods
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

<style>
.mp-user-column {
  display: flex;
  align-items: center;
}
.force32 {
  height: 32px;
  width: 32px;
}
</style>
