<template>
  <v-stepper
    :value="step"
  >
    <v-stepper-header>
      <v-stepper-step v-for="s in scenarios" :key="s" :step="s">
        Scenario {{ s }}
      </v-stepper-step>
      <v-stepper-step :step="scenarios + 1">
        Survey
      </v-stepper-step>
    </v-stepper-header>
    <v-stepper-items>
      <Survey v-if="isSurvey"></Survey>
      <Scenario v-else :number="step"></Scenario>
    </v-stepper-items>
  </v-stepper>
</template>

<script>
import Scenario from '../components/Scenario'
import Survey from '../components/Survey'

export default {
  name: 'Experiment',
  components: { Scenario, Survey },
  data () {
    return {
      scenarios: 3,
    }
  },
  computed: {
    step () {
      return Number(this.$route.params.number)
    },
    isSurvey () {
      return this.step === this.scenarios + 1
    },
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
