<template>
  <v-app>
    <v-main v-if="$demoMode || $auth.ready()">
      <router-view />
    </v-main>
    <v-main v-else>
      Loading...
    </v-main>
  </v-app>
</template>

<script>

export default {
  name: 'App',
}
</script>

<style scoped lang="scss">

</style>
