<template>
  <div id="content" class="matches" v-if="match_owner" v-loading="!match_owner">
    <h1 v-if="my_matches">Your matches</h1>
    <h1 v-else-if="all_matches">All matches</h1>
    <h1 v-else>Matches for <router-link :to="'/user/'+match_owner.id">{{match_owner.name}}</router-link></h1>

  <table class="table table-striped">
    <thead>
      <tr>
        <th>Match ID</th>
        <th>Team 1</th>
        <th>Team 2</th>
        <th>Status</th>
        <th v-if="my_matches">Server</th>
        <th v-if="my_matches"></th>
        <th v-else>Owner</th>
      </tr>
    </thead>
    <tbody style="overflow:auto">
      <tr v-for="(match,index) in matches" :key="'m_'+match.id+'_'+index" align="left">
        <td v-if="match" v-loading="loading[match.id]"><router-link :to="'/match/'+match.id">{{match.id}}</router-link></td>
        <td v-if="matchinfo[match.id]">
          <img :src="get_flag_link(matchinfo[match.id].team1)" />
          <router-link :to="'/team/'+match.team1_id">{{matchinfo[match.id].team1.name}}</router-link>
        </td>

        <td v-if="matchinfo[match.id]">
          <img :src="get_flag_link(matchinfo[match.id].team2)"  />
          <router-link :to="'/team/'+match.team2_id">{{matchinfo[match.id].team2.name}}</router-link>
        </td>

        <td v-if="matchinfo[match.id]">
          {{ matchinfo[match.id].status }}
        </td>

        <td v-if="my_matches && matchinfo[match.id].server">{{ matchinfo[match.id].server.display }} </td>
        <td v-if="my_matches && matchinfo[match.id]">
          <a v-if="(match.pending || match.live)" :href="'/match/'+match.id+'cancel'" class="btn btn-danger btn-xs align-right">Cancel</a>
        </td>
        <td v-if="!my_matches && matchinfo[match.id]">
          <router-link :to="'/user/'+matchinfo[match.id].user.id">{{ matchinfo[match.id].user.name }}</router-link>
        </td>
      </tr>
    <el-button type="primary" v-if="!limit" :loading="loadingmore" @click="GetMatches()">Load more...</el-button>
    </tbody>
  </table>

  <!--{{ pagination_buttons(matches) }}-->

  </div>
</template>

<script>
export default {
  name: 'matches',
  data () {
    return {
      user: {},
      loadingmore: false,
      loaded: 0,
      loading: {},
      flag_loading: {},
      limit: false,
      my_matches: false,
      all_matches: true,
      matches: [],
      matchinfo: {},
      match_owner: {},
      teamdatas: {},
      userdatas: {},
      serverdatas: {}
    }
  },
  created () {
    this.Init()
  },
  watch: {
    $route (to, from) {
      this.Init()
    }
  },
  methods: {
    async Init () {
      this.user = {}
      this.loadingmore = false
      this.loaded = 0
      this.loading = {}
      this.flag_loading = {}
      this.limit = false
      this.my_matches = false
      this.all_matches = true
      this.matches = []
      this.matchinfo = {}
      this.match_owner = {}
      this.teamdatas = {}
      this.userdatas = {}
      this.serverdatas = {}
      return new Promise(async (resolve, reject) => {
        this.matches = []
        if (this.$route.params.userid) {
          this.all_matches = false
          let res = await this.axios.get('/api/v1/CheckLoggedIn')
          this.user = res.data
          this.my_matches = this.$route.params.userid === this.user.userid
          this.GetMatches(this.$route.params.userid)
          this.userdatas[this.$route.params.userid] = await this.GetUserData(this.$route.params.userid)
          this.match_owner = this.userdatas[this.$route.params.userid]
          resolve()
        } else {
          const res = await this.axios.get('/api/v1/CheckLoggedIn')
          this.user = res.data
          this.my_matches = this.$route.params.userid === this.user.userid
          await this.GetMatches()
          resolve()
        }
        this.activeIndex = this.$route.name
      })
    },
    async GetMatches (userid) {
      let self = this
      self.loadingmore = true
      return new Promise(async (resolve, reject) => {
        if (userid) {
          const res = await this.axios.get(`/api/v1/GetMatches?userID=${userid}`)
          self.loaded = self.loaded + res.data.length
          for (let i = 0; i < res.data.length; i++) {
            this.matches.push(res.data[i])
            self.$set(self.loading, [res.data[i].id], true)
            this.matchinfo[res.data[i].id] = await this.GetMatchData(res.data[i].id)
            self.$set(self.loading, [res.data[i].id], false)
            if (res.data.length === 0) {
              self.limit = true
            }
            if (i + 1 === res.data.length) {
              self.loadingmore = false
              resolve(res.data)
            }
          }
        } else {
          const res = await this.axios.get(`/api/v1/GetMatches?offset=${this.loaded}`)
          self.loaded = self.loaded + res.data.length
          if (res.data.length === 0) {
            self.limit = true
          }
          for (let i = 0; i < res.data.length; i++) {
            this.matches.push(res.data[i])
            self.$set(self.loading, [res.data[i].id], true)
            this.matchinfo[res.data[i].id] = await this.GetMatchData(res.data[i].id)
            self.$set(self.loading, [res.data[i].id], false)
            if (i + 1 === res.data.length) {
              self.loadingmore = false
              resolve(res.data)
            }
          }
        }
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity .5s
}
.fade-enter, .fade-leave {
  opacity: 0
}
</style>
