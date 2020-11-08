<template>
  <v-container
    class="fill-height"
  >
    <v-row
      justify="center"
      align="center"
    >
      <v-card
        max-width="700px"
      >
        <v-card-title>Intro</v-card-title>
        <v-card-text>
          <p>Hello, today you are to ......</p>
          <v-divider />
          <v-subheader>Confidentiality agreement</v-subheader>
          <p>My data will, be used to....</p>
          <v-checkbox
            label="I accept the agreement"
            v-model="accepted"
          ></v-checkbox>
          <v-card-actions>
            <v-spacer />
            <v-btn
              color="primary"
              :disabled="!accepted"
              @click="submit"
            >Start</v-btn>
          </v-card-actions>
        </v-card-text>
      </v-card>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Intro',
  data () {
    return {
      accepted: false,
    }
  },
  methods: {
    submit () {
      if (this.accepted === true) {
        this.axios.post('agreement/accept').then(() => {
          this.$router.push('/experiment/1')
        }) // TODO handle error
      }
    },
  },
}
</script>

<style scoped lang="scss">

</style>
