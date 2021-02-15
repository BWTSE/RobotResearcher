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
        <v-checkbox v-model="cookie" label="I accept that this site stores a cookie on my computer. The cookie is used to anonymously identify the submission."></v-checkbox>
        <v-btn
          :disabled="processing || code === '' || !cookie"
          :loading="processing"
          @click.prevent.stop="login"
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
      cookie: false,
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
          if (er.response && er.response.status === 401) {
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
  .v-form {
    max-width: 500px;
  }
</style>
