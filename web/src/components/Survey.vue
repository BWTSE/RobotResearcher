<template>
  <v-row
    justify="center"
    align="center"
    class="fill-height"
  >
    <v-card>
      <v-form
        align="right"
      >
        Some  questions
        <v-btn color="primary" @click="submit">Submit</v-btn>
      </v-form>
    </v-card>
  </v-row>
</template>

<script>
export default {
  name: 'Survey',
  methods: {
    submit () {
      this.$router.push('/farewell')
    },
  },
}
</script>

<style scoped lang="scss">

</style>
