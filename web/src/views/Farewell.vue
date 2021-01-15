<template>
  <v-row
    justify="center"
    align="center"
    class="fill-height"
  >
    <v-card>
      <v-card-title>Thank you for participating!</v-card-title>
      <v-card-text>If you want to receive a copy of the thesis after its completion you may provide your email at: <a target="_blank" href="https://forms.gle/WJ3kXBfBYjCV35E89">https://forms.gle/WJ3kXBfBYjCV35E89</a></v-card-text>
    </v-card>
  </v-row>
</template>

<script>
export default {
  name: 'Farewell',
  mounted () {
    this.$auth.logout({
      makeRequest: true,
      url: 'auth/close',
      redirect: false,
    })
  },
}
</script>

<style scoped lang="scss">
</style>
