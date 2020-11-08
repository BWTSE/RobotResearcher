<template>
  <v-row
    justify="center"
    align="center"
    class="fill-height"
  >
    <v-card>
      Thank you for participating
    </v-card>
  </v-row>
</template>

<script>
export default {
  name: 'Farewell',
  mounted () {
    this.$auth.logout({
      makeRequest: false,
      redirect: false,
    }) // TODO: logout
  },
}
</script>

<style scoped lang="scss">

</style>
