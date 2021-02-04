<template>
  <v-row
    justify="center"
    align="center"
    class="fill-height"
  >
    <v-card class="pa-4">
      <v-form
      >
        <span class="sliders" v-for="scenario in scenarios" :key="scenario">
          <span>How would you rate the quality (maintainability) of the preexisting code in the <strong>{{ scenario }}</strong>
            <v-tooltip right>
              <template v-slot:activator="{on, attrs}">
                <v-icon v-bind="attrs" v-on="on" class="infoa" tag="span" small>mdi-information</v-icon>
              </template>
              <span>{{scenarioHelp[scenario]}}</span>
            </v-tooltip>
            task?</span>
          <v-slider
            v-model="surveyAnswers['Scenario' + scenario + 'Quality']"
            :tick-labels="['Very Bad', 'Bad', 'Neutral', 'Good', 'Very Good']"
            ticks
            tick-size="5"
            step="1"
            min="-2"
            max="2"
          ></v-slider>
          <span>The work I did in the <strong>{{ scenario }}</strong>
            <v-tooltip right>
              <template v-slot:activator="{on, attrs}">
                <v-icon v-bind="attrs" v-on="on" class="infoa" tag="span" small>mdi-information</v-icon>
              </template>
              <span>{{scenarioHelp[scenario]}}</span>
            </v-tooltip>
             task made the quality (maintainability) of the system?</span>
          <v-slider
            v-model="surveyAnswers['Scenario' + scenario + 'QualityChange']"
            :tick-labels="['Much Worse', 'Worse', 'Neutral', 'Better', 'Much Better']"
            ticks
            tick-size="5"
            step="1"
            min="-2"
            max="2"
          ></v-slider>
        </span>
        <v-row>
          <v-spacer></v-spacer>
          <v-btn
            color="primary"
            @click="submit"
            :disabled="loading || invalid"
          >Submit</v-btn>
        </v-row>
      </v-form>
    </v-card>
  </v-row>
</template>

<script>
import Vue from 'vue'

export default {
  name: 'Survey',
  data () {
    return {
      loading: false,
      surveyAnswers: { // flat structure to make data analysis easier
        ScenarioTicketsQuality: null,
        ScenarioTicketsQualityChange: null,

        ScenarioBookingQuality: null,
        ScenarioBookingQualityChange: null,
      },
      scenarios: [
        'Shapes',
        'Booking',
        'Shopping',
      ],
      scenarioHelp: {
        Tickets: 'The task where you were asked to extend a buss ticket system',
        Booking: 'The task where you were asked to implement a Cabin class',
      },
    }
  },
  computed: {
    invalid () {
      return Object.values(this.surveyAnswers).some((a) => a === null || a === [])
    },
  },
  mounted () {
    this.$auth.fetch().then((a) => {
      Vue.set(this, 'scenarios', a.data.scenario_order.map((name) => name.charAt(0).toUpperCase() + name.slice(1)))

      switch (a.data.stage) {
        case 'agreement':
          this.$router.push('/intro')
          break
        case 'background':
          this.$router.push('/experiment/0')
          break
        case 'experiment':
          this.$router.push('/experiment/' + a.data.experiment)
          break
        case 'survey':
          break
        default:
          this.$router.push('/farewell')
          break
      }
    })
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
.v-form {
  min-width: 850px;
  padding-bottom: 40px;
  > * {
    padding-bottom: 20px;
  }
  > .sliders  {
    display: block;
    margin-bottom: 30px;
    > span {
      display: block;
      margin-top: 10px;
    }
  }
}
  .infoa {
    position: relative;
    top: -5px;
  }
</style>
