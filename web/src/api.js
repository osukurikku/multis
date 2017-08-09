import axios from 'axios'

const errorHandler = err => {
  console.error(err)
}

const urlRippleAPI = 'https://api.ripple.moe/api/v1'

export default {
  matches (page, callback) {
    axios.get('/api/matches?p=' + page)
      .then(resp => callback(resp.data))
      .catch(errorHandler)
  },
  match (id, callback) {
    axios.get('/api/matches?id=' + id)
      .then(resp => callback(resp.data))
      .catch(errorHandler)
  },
  games (matchID, callback) {
    axios.get('/api/games?match_id=' + matchID)
      .then(resp => callback(resp.data))
      .catch(errorHandler)
  },
  loadEntities (component, users, beatmaps) {
    let usersURL = urlRippleAPI + '/users?l=100&'
    users.forEach(u => { usersURL += 'ids=' + u + '&' })
    axios.get(usersURL)
      .then(resp => {
        if (!resp.data.users) {
          return
        }
        resp.data.users.forEach(user => {
          component.$set(component.users, user.id, user)
        })
      }).catch(errorHandler)

    let beatmapURL = urlRippleAPI + '/beatmaps?l=100&'
    beatmaps.forEach(s => { beatmapURL += 'bb=' + s + '&' })
    axios.get(beatmapURL)
      .then(resp => {
        if (!resp.data.beatmaps) {
          return
        }
        resp.data.beatmaps.forEach(beatmap => {
          component.$set(component.beatmaps, beatmap.beatmap_id, beatmap)
        })
      }).catch(errorHandler)
  }
}
