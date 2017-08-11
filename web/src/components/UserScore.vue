<template>
  <div class="column is-half mp-user-column">
    <div :class="'team-indicator team-' + score.Team"></div>
    <a :href="'https://ripple.moe/u/' + uid" target="_blank" class="force32">
      <img :src="'https://a.ripple.moe/' + uid" class="force32">
    </a>
    <span style="margin-left: 5px">
      <span>
        {{ users[uid] ? users[uid].username : "..." }}
        &mdash;
        {{ score.Score.toLocaleString() }}
        <small>{{ scoreMods(score.Mods) }}</small>
      </span>
      <span class="icon">
        <i class="fa fa-times" v-if="score.Failed" title="Failed (including at the end)"></i>
        <i class="fa fa-wheelchair-alt" v-else-if="!score.Pass" title="Not passed (hp restored before the end), also displayed on older games"></i>
      </span>
    </span>
  </div>
</template>

<script>
import scoreMods from '../scoreMods'

export default {
  props: ['score', 'uid', 'users'],
  methods: {
    scoreMods: scoreMods
  }
}
</script>

<style scoped>
.mp-user-column {
  display: flex;
  align-items: center;
}
.force32 {
  height: 32px;
  width: 32px;
}

.team-indicator {
  width: 6px;
  height: 32px;
  margin-right: 5px;
}
.team-indicator.team-0 { background: #e5e5e5 }
.team-indicator.team-1 { background: #0074db }
.team-indicator.team-2 { background: #dd300d }
</style>
