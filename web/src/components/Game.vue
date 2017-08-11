<template>
  <article class="media">
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
            {{ scoreMods(game.Mods) }}
          </span>
        </p>
      </div>
      <div class="columns is-multiline is-mobile">
        <user-score
          v-for="uid in scoresSorted"
          :key="uid"
          :uid="uid"
          :score="game.Scores[uid]"
          :users="users">
        </user-score>
      </div>
    </div>
    <div class="media-right">
      <timeago :since="game.CreatedAt" :title="game.CreatedAt"></timeago>
    </div>
  </article>
</template>

<script>
import scoreMods from '../scoreMods'
import UserScore from './UserScore'

export default {
  components: {
    UserScore
  },
  props: ['game', 'users', 'beatmaps'],
  computed: {
    scoresSorted () {
      return Object.keys(this.game.Scores).sort((a, b) => {
        return this.game.Scores[a].Score < this.game.Scores[b].Score
      })
    }
  },
  methods: {
    scoreMods: scoreMods
  }
}
</script>
