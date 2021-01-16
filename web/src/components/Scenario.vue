<template>
  <v-row
    class="root"
  >
    <v-col cols="8" class="code fill-height">
      <multi-file-editor
        :files="files"
        @update:file="({file, content}) => files[file] = content"
      ></multi-file-editor>
    </v-col>
    <v-col class="util fill-height d-flex justify-space-between flex-column">
      <div
        v-html="htmlinfo"
      >
      </div>
      <div>
        <v-btn
          color="primary"
          @click="run"
          width="100%"
          :disabled="loading"
        >
          Run Code
        </v-btn>
        <Output :code="runOutput + 'Exit code: ' + runCode"></Output>
      </div>
      <div class="d-flex justify-space-between">
        <v-spacer></v-spacer>
        <v-btn
          color="primary"
          @click="next"
          width="200px"
          :disabled="loading"
        >
          Submit
        </v-btn>
      </div>
    </v-col>
  </v-row>
</template>

<script>
import moment from 'moment'
import marked from 'marked'
import MultiFileEditor from './MultiFileEditor'
import Output from './Output'

export default {
  name: 'Scenario',
  data: () => {
    return {
      scenario: null,
      files: {},
      info: '',
      startedAt: null,
      loading: false,
      runOutput: '',
      runCode: null,
    }
  },
  props: {
    number: Number,
  },
  components: { Output, MultiFileEditor },
  computed: {
    htmlinfo () {
      return marked(this.info)
    },
  },
  methods: {
    next () {
      if (this.loading) return
      this.loading = true
      this.axios.post('scenario/' + (Number(this.$route.params.number) - 1), { submission: this.prepareSubmission(this.files) }).then(() => {
        this.loading = false
        this.$router.push('/experiment/' + (Number(this.$route.params.number) + 1))
      }).catch(() => {
        this.loading = false
      })
    },
    run () {
      if (this.loading) return
      this.loading = true
      this.axios.post('code/run', { submission: this.prepareSubmission(this.files) }).then((resp) => {
        this.runOutput = resp.data.out
        this.runCode = resp.data.code
        this.loading = false
      }).catch((err) => {
        if (err.response) {
          this.runOutput = err.response.data.out
          this.runCode = err.response.data.code
          this.loading = false
        }
      })
    },
    prepareSubmission (files) {
      const submission = {}
      Object.entries(files).forEach(([name, content]) => {
        submission[name] = btoa(content)
      })
      return submission
    },
    loadScenario () {
      this.loading = true
      this.axios.get('scenario/' + (Number(this.$route.params.number) - 1)).then((resp) => {
        this.scenario = resp.data
        this.info = this.scenario.info
        this.files = {}
        Object.entries(this.scenario.given).forEach(([name, encodedContent]) => {
          this.files[name] = atob(encodedContent)
        })
        this.startedAt = moment(this.scenario.started_at)
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
  },
  mounted () {
    this.loadScenario()
  },
  watch: {
    number (n, o) {
      if (n !== o) {
        this.loadScenario()
      }
    },
  },
}
</script>

<style scoped lang="scss">
.root {
  height: calc(100vh - 72px);
  .util {
    padding-right: 24px;
    overflow-y: auto;
    > * {
      margin-bottom: 30px;
    }
  }
  .code {
    padding: 0;
  }
}
</style>
