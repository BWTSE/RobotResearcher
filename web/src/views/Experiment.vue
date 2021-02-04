<template>
  <v-stepper
    :value="step"
  >
    <v-stepper-header>
      <v-stepper-step :step="0">
        Background info
      </v-stepper-step>
      <v-stepper-step v-for="s in scenarios" :key="s" :step="s">
        {{ scenarioList[s-1] }} Task
      </v-stepper-step>
      <v-stepper-step :step="scenarios + 1">
        Survey
      </v-stepper-step>
    </v-stepper-header>
    <v-stepper-items>
      <Survey v-if="isSurvey"></Survey>
      <Background v-else-if="isBackground">Background questions </Background>
      <Scenario v-else :number="step"></Scenario>
    </v-stepper-items>
  </v-stepper>
</template>

<script>
import Background from '../components/Background'
import Scenario from '../components/Scenario'
import Survey from '../components/Survey'
import Vue from 'vue'

export default {
  name: 'Experiment',
  components: { Scenario, Survey, Background },
  data () {
    return {
      scenarios: 2,
      scenarioList: ['1', '2'],
    }
  },
  computed: {
    step () {
      return Number(this.$route.params.number)
    },
    isSurvey () {
      return this.step === this.scenarios + 1
    },
    isBackground () {
      return this.step === 0
    },
  },
  mounted () {
    this.$auth.fetch().then((a) => {
      Vue.set(this, 'scenarioList', a.data.scenario_order.map((name) => name.charAt(0).toUpperCase() + name.slice(1)))
    })
  },
}
</script>

<style scoped lang="scss">
.v-stepper {
  height: 100%;
  .v-stepper__items {
    height: 100%;
    .v-stepper__content {
      height: 100%;
    }
  }
}
</style>
