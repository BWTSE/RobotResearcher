<template>
  <v-container
    class="fill-height"
  >
    <v-row
      justify="center"
      align="center"
    >
      <v-form
        align="right"
        v-on:keydown.enter="login"
      >
        <v-text-field
          prepend-icon="mdi-console-line"
          label="Code"
          v-model="code"
          :error-messages="errorMessage"
        ></v-text-field>
        <v-btn
          :disabled="processing || code === ''"
          :loading="processing"
          @click="login"
          type="submit"
        >Enter</v-btn>
      </v-form>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Join',
  data () {
    return {
      code: '',
      rememberMe: true,
      processing: false,
      errorMessage: '',
    }
  },
  methods: {
    login: function () {
      if (!(this.processing || this.code === '')) {
        this.processing = true
        this.$auth.login({
          data: { code: this.code },
          rememberMe: true,
          redirect: '/intro',
          fetchUser: true,
        }).then((res) => {
          this.processing = false
        }).catch((er) => {
          this.processing = false
          if (er.response && er.response.status === 400) {
            this.errorMessage = 'Invalid Code'
            this.code = ''
          } else {
            this.errorMessage = 'Something went wrong... Contact organizer.'
            this.code = ''
          }
        })
      }
    },
  },
  created () {
    if (this.$auth.check()) {
      this.$router.push('/intro')
    }
  },
}
</script>

<style scoped lang="scss">
</style>
