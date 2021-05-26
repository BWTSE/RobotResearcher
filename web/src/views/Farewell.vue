<template>
  <v-row
    justify="center"
    align="center"
    class="fill-height pa-16"
  >
    <v-card width="650">
      <v-card-title>Thank you for participating!</v-card-title>
      <v-card-text>
        <p>
          If you want to receive a copy of the completed thesis or want to participate in a short follow up interview you can use the form bellow.
        </p>
      </v-card-text>
      <hr>
      <iframe
        src="https://docs.google.com/forms/d/e/1FAIpQLSeCn-W2hhktZngqNB72XSI5a-d9N1awwuldXkXlm1gZyO7tVQ/viewform?embedded=true&hl=en"
        width="640"
        height="1100"
        seamless
        frameborder="0"
        marginheight="0"
        marginwidth="0"
        v-if="!$demoMode"
      >
        Läser in...
      </iframe><span v-else>No form in demo mode</span>
    </v-card>
  </v-row>
</template>

<script>
export default {
  name: 'Farewell',
  mounted () {
    if (!this.$demoMode) {
      this.$auth.logout({
        makeRequest: true,
        url: 'auth/close',
        redirect: false,
      })
    }
  },
}
</script>

<style scoped lang="scss">
</style>
