<template>
<div id="app">
    <div id="header">
        <nav class="navbar navbar-default">
            <div class="container-fluid">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#myNavbar">
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <transition name="fade">
                        <router-link v-if="LogoTransition" to="/" class="navbar-brand" >Get5 Web Panel</router-link>
                    </transition>
                </div>
                <div class="collapse navbar-collapse" id="myNavbar">
                    <el-menu
                        class="nav navbar-nav"
                        :default-active="activeIndex"
                        mode="horizontal"
                        router
                    >
                        <el-menu-item index="Matches" id="matches" :route="{ path:'/matches' }" v-if="user.isLoggedIn">All Matches</el-menu-item>
                        <el-menu-item index="My matches" id="mymatches" :route="{ path:'/mymatches' }" v-if="user.isLoggedIn">My Matches</el-menu-item>
                        <el-menu-item index="match_create" id="match_create" :route="{ path:'/match/create' }" v-if="user.isLoggedIn">Create a Match</el-menu-item>
                        <el-menu-item index="myteams" id="myteams" :route="{ path:'/myteams' }" v-if="user.isLoggedIn">My Teams</el-menu-item>
                        <el-menu-item index="team_create" id="matches" :route="{ path:'/team/create' }" v-if="user.isLoggedIn">Create a Team</el-menu-item>
                        <el-menu-item index="myservers" id="myservers" :route="{ path:'/myservers' }" v-if="user.isLoggedIn">My Servers</el-menu-item>
                        <el-menu-item index="server_create" id="server_create" :route="{ path:'/server/create' }" v-if="user.isLoggedIn">Add a Server</el-menu-item>
                        <el-menu-item index="logout" id="logout" v-if="user.isLoggedIn"> <a href='/api/v1/logout' v-if="user.isLoggedIn">Logout</a> </el-menu-item>
                        <el-menu-item index="login" id="login" v-if="!user.isLoggedIn"> <a href='/api/v1/login' v-if="!user.isLoggedIn"> <img src="/img/login_small.png" height="18" /></a> </el-menu-item>
                    </el-menu>
                </div>
            </div>
        </nav>
    </div>
    <router-view />
    <div class="panel-footer text-muted">
        <div>
            Powered by <a href="http://steampowered.com">Steam</a> -
            <router-link id="metrics" to="/metrics" >Stats</router-link>
            <div v-if="version">- Version <a href="https://github.com/FlowingSPDG/get5-web-go">{{ version }}</a></div>
        </div>
    </div>
</div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'App',
  data () {
    return {
      LogoTransition: false,
      version: '',
      activeIndex: '',
      user: {
        isLoggedIn: false,
        adminaccess: false,
        steamid: '',
        userid: ''
      } // should be object from JSON response
    }
  },
  async mounted () {
    this.LogoTransition = true
    this.activeIndex = this.$route.name
    let LoggedIn = await axios.get('/api/v1/CheckLoggedIn')
    this.user = LoggedIn.data
    let Version = await axios.get('/api/v1/GetVersion')
    this.version = Version.data.version
  },
  methods: {

  }
}
</script>

<style>
#app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

.container {
  padding-right: 15px;
  padding-left: 15px;
  margin-right: auto;
  margin-left: auto;
}
</style>
