<template>
  <v-row
    justify="center"
    align="center"
    class="fill-height"
  >
    <v-card class="pa-4">
      <v-form
        align="right"
      >
        <v-select
          :items="[{text: 'yes', value: true}, {text: 'no', value: false}]"
          v-model="surveyAnswers.LikesChocolate"
          label="I like chocolate!*"
        ></v-select>
        <v-btn
          color="primary"
          @click="submit"
          :disabled="loading || invalid"
        >Submit</v-btn>
      </v-form>
    </v-card>
  </v-row>
</template>

<script>
export default {
  name: 'Survey',
  data () {
    return {
      loading: false,
      surveyAnswers: {
        LikesChocolate: null,
      },
    }
  },
  computed: {
    invalid () {
      return Object.values(this.surveyAnswers).some((a) => a === null)
    },
  },
  methods: {
    submit () {
      if (this.loading) return
      this.loading = true
      this.axios.post('survey', this.surveyAnswers).then(() => {
        this.loading = false
        this.$router.push('/farewell')
      }).catch(() => {
        this.loading = false
      })
    },
  },
}
</script>

<style scoped lang="scss">

</style>
