<template>
  <v-app>
    <v-main v-if="$auth.ready()">
      <router-view />
    </v-main>
    <v-main v-else>
      Loading...
    </v-main>
  </v-app>
</template>

<script>

import * as Sentry from '@sentry/browser'
export default {
  name: 'App',
  methods: {
    openDialog () {
      Sentry.showReportDialog({
        eventId,
        title: 'Aw, Snap! You broke the internet!',
        subtitle: 'Send us an email and we will fix it.',
        subtitle2: 'Please refresh your cache to continue (cmd+R / ctrl+R)',
      })
    },
  },
}
</script>

<style scoped lang="scss">

</style>
