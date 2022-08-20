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
          <h3>
            Welcome and thanks for participating in our experiment.
          </h3>
          <p>
            The experiment aims to investigate the evolution of software projects by asking the participants to complete tasks in preexisting projects. There will be two (2) tasks to be performed followed by a short survey. You will also be asked to answer some background questions before you start. Your participation is voluntary and you are free to drop out of the experiment at any time.
            <strong>We do greatly appreciate if you complete the whole experiment</strong> so that we may acquire some useful results.
          </p>
          <p v-if="!$demoMode">
            This experiment is part of a master thesis by William Leven (<a target="_blank" href="mailto:william@leven.id">william@leven.id</a>) and Hampus Broman (<a target="_blank" href="mailto:bromanh@student.chalmers.se">bromanh@student.chalmers.se</a>) studying at the Software Engineering master program at Chalmers University. The supervisors are Terese Besker and Richard Torkar.
          </p>
          <p v-if="!$demoMode"><!-- Also in Farewell.vue -->
            If you want to receive a copy of the thesis after its completion you may provide your email at: <a v-if="!$demoMode" target="_blank" href="https://forms.gle/WJ3kXBfBYjCV35E89">https://forms.gle/WJ3kXBfBYjCV35E89</a><span v-else>DEMO MODE</span>
          </p>
          <v-divider />
          <h3>Confidentiality agreement</h3>
          <p>
            During the experiment we will collect your solutions to the tasks, your completion time, and survey answers for further analysis to investigate the evolution of software projects. Responsible for the collection are William Leven (<a target="_blank" href="mailto:william@leven.id">william@leven.id</a>) and Hampus Broman (<a target="_blank" href="mailto:bromanh@student.chalmers.se">bromanh@student.chalmers.se</a>).
          </p>
          <p>
            The data will be continuously collected for the whole duration of the experiment and will be stored even if the experiment isn't completed. The data we collect from you will also be enriched by data about the place of your employment. After the completion of the thesis, the collected (anonymized) data will be made publicly available.
          </p>
          <p>
            None of the collected data can, by its nature, uniquely identify you as an individual.
            In those cases where the submitted data contains any uniquely identifying information, the dataset will be anonymized before further processing or publication to ensure confidentiality.
          </p>

          <p v-if="$demoMode">
            Demo mode! No data will be stored.
          </p>

          <v-checkbox
            label="I accept that my data is collected and processed in accordance to the agreement above."
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
        if (this.$demoMode) {
          return this.$router.push('/experiment/0')
        }
        this.axios.post('agreement/accept').then(() => {
          this.$router.push('/experiment/0')
        })
      }
    },
  },
  mounted () {
    if (!this.$demoMode) {
      this.$auth.fetch().then((a) => {
        switch (a.data.stage) {
          case 'agreement':
            // this.$router.push('/intro')
            break
          case 'background':
            this.$router.push('/experiment/0')
            break
          case 'experiment':
            this.$router.push('/experiment/' + a.data.experiment)
            break
          case 'survey':
            this.$router.push('/experiment/3')
            break
          default:
            this.$router.push('/farewell')
            break
        }
      })
    }
  },
}
</script>

<style scoped lang="scss">

</style>
