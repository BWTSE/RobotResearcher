<template>
  <v-row
    class="fill-height"
  >
    <v-col cols="8" class="fill-height">
      <Editor class="fill-height"></Editor>
    </v-col>
    <v-col class="fill-height">
      Instructions for scenario {{ number }}
      <v-spacer></v-spacer>
      <v-btn
        color="primary"
        @click="next"
      >
        Continue
      </v-btn>
    </v-col>
  </v-row>
</template>

<script>
import Editor from './Editor'

export default {
  name: 'Scenario',
  props: {
    number: Number,
  },
  components: { Editor },
  methods: {
    next () {
      this.$router.push('/experiment/' + (Number(this.$route.params.number) + 1))
    },
  },
}
</script>

<style scoped lang="scss">

</style>
