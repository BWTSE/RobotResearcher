<template>
  <v-row
    class="root"
  >
    <v-col class="code fill-height">
      <multi-file-editor
        :files="files"
        @update:file="({file, content}) => files[file] = content"
        @delete:file="({file}) => delete files[file]"
        @save="run"
      ></multi-file-editor>
    </v-col>
    <v-col class="desc-toggler" @click="utilVisible = !utilVisible">
      <v-icon v-if="utilVisible">mdi-arrow-expand-right</v-icon>
      <v-icon v-else>mdi-arrow-expand-left</v-icon>
    </v-col>
    <v-col :class="{'hidden': !utilVisible}" class="util fill-height d-flex justify-space-between flex-column">
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
        <v-dialog
          v-model="nextDialog"
          max-width="600px"
          @keydown.enter="next"
        >
          <template v-slot:activator="{on}">
            <v-btn
              color="primary"
              v-on="on"
              width="200px"
              :disabled="loading || nextDialog"
            >
              Submit
            </v-btn>
          </template>
          <v-card>
            <v-card-title>Continue</v-card-title>
            <v-card-text>You can not go back once you have submitted your solution</v-card-text>
            <v-card-actions>
              <v-btn
                @click="nextDialog = false"
                :disabled="loading"
              >
                Go back
              </v-btn>
              <v-spacer></v-spacer>
              <v-btn
                color="primary"
                @click="next"
                :disabled="loading"
              >
                Submit
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
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
      nextDialog: false,
      utilVisible: true,
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
        this.nextDialog = false
        this.$router.push('/experiment/' + (Number(this.$route.params.number) + 1))
      }).catch(() => {
        this.loading = false
        this.nextDialog = false
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
    this.$auth.fetch().then((a) => {
      switch (a.data.stage) {
        case 'agreement':
          this.$router.push('/intro')
          break
        case 'background':
          this.$router.push('/experiment/0')
          break
        case 'experiment':
          if (this.$route.params.number !== a.data.experiment) {
            this.$router.push('/experiment/' + a.data.experiment)
          }
          break
        case 'survey':
          this.$router.push('/experiment/3')
          break
        default:
          this.$router.push('/farewell')
          break
      }
    })
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
    max-width: 700px;
    > * {
      margin-bottom: 30px;
    }
    &.hidden {
      max-width: 0px;
      opacity: 0;
    }
    transition: all ease-in-out .5s;
  }
  .code {
    padding: 0;
  }
}
.desc-toggler {
  max-width: 12px;
  background-color: lightgray;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  &:hover {
    background-color: darkgray;
  }
}
</style>
