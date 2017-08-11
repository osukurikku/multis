<template>
  <div>
    <hero title="Vinse" subtitle="A directory of multiplayer games on Ripple" class="is-info"></hero>
    <section class="section">
      <div class="container box content">
        <p>
          Welcome to Vinse. Vinse is a website from which you can see all the
          multiplayer matches that are happening right now on Ripple. On this
          page, you can see all the multiplayer matches, sorted by time of
          creation. Please note, though, that in order for a match to pop up
          here, at least one score must have been submitted.
        </p>
      </div>
      <div class="container box" v-if="matches === null">
        Loading matches...
      </div>
      <div class="container box" v-else>
        <article class="media" v-for="match in matches" :key="match.ID">
          <div class="media-content">
            <div class="content">
              <p>
                <strong>{{ match.Name }}</strong>
                <timeago auto-update="30" :since="match.CreatedAt" :title="match.CreatedAt"></timeago>
              </p>
            </div>
          </div>
          <div class="media-right">
            <a class="button" :href="'/match/' + match.ID" @click.prevent="goToMatch(match)">
              <span class="icon is-small"><i class="fa fa-eye"></i></span>
            </a>
          </div>
        </article>
        <paginator v-model="page" :right-disabled="matches.length < 100"></paginator>
      </div>
    </section>
  </div>
</template>

<script>
import Hero from './Hero'
import Paginator from './Paginator'
import api from '../api'
import store from '../store'

export default {
  components: {
    Hero,
    Paginator
  },
  data () {
    return {
      matches: null,
      page: 0
    }
  },
  mounted () {
    this.loadMatches()
  },
  methods: {
    loadMatches () {
      api.matches(this.page, results => {
        this.matches = results
      })
    },
    goToMatch (match) {
      store.commit('setClickedMatch', match)
      this.$router.push('/match/' + match.ID)
    }
  }
}
</script>
