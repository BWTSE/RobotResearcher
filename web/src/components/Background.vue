<template>
  <v-row
    justify="center"
    align="center"
    class="fill-height"
  >
    <v-card class="pa-4">
      <v-form
      >
        <v-select
          :items="edLevels"
          v-model="surveyAnswers.EducationLevel"
          label="What is the highest degree or level of education you have completed?"
        ></v-select>
        <v-text-field
          v-model="surveyAnswers.EducationField"
          label="In what field did you attain that level of education?"
          placeholder="e.g. computer science, electrical engineering, software engineering..."
        ></v-text-field>
        <v-text-field
          v-model="surveyAnswers.ProgrammingExperience"
          label="How many years of professional programming experience do you have?"
          placeholder="e.g. 1, 2, 3..."
          suffix="years"
        ></v-text-field>
        <v-text-field
          v-model="surveyAnswers.JavaExperience"
          label="How many years of professional experience do you have working with java?"
          placeholder="e.g. 1, 2, 3..."
          suffix="years"
        ></v-text-field>
        <v-text-field
          v-model="surveyAnswers.WorkDomain"
          label="Within what domain did you attain the most work-experience?"
          placeholder="e.g. telecom, game development, automotive..."
        ></v-text-field>
        <v-divider></v-divider>
        <span>My most recent place of employment...</span>
        <v-select
          :items="[{text: 'Yes', value: true}, {text: 'No', value: false}]"
          label="used peer code reviews."
          v-model="surveyAnswers.CompanyUsesCodeReviews"
        ></v-select>
        <v-select
          :items="[{text: 'Yes', value: true}, {text: 'No', value: false}]"
          label="used pair programing."
          v-model="surveyAnswers.CompanyUsesPairProgramming"
        ></v-select>
        <v-select
          :items="[{text: 'Yes', value: true}, {text: 'No', value: false}]"
          label="tracked technical debt."
          v-model="surveyAnswers.CompanyTracksTechnicalDebt"
        ></v-select>
        <v-select
          :items="[{text: 'Yes', value: true}, {text: 'No', value: false}]"
          label="had established coding standards."
          v-model="surveyAnswers.CompanyHasCodingStandards"
        ></v-select>
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

export default {
  name: 'Survey',
  data () {
    return {
      loading: false,
      surveyAnswers: { // flat structure to make data analysis easier
        EducationLevel: null,
        EducationField: null,
        ProgrammingExperience: null,
        JavaExperience: null,
        WorkDomain: null,

        CompanyUsesCodeReviews: null,
        CompanyUsesPairProgramming: null,
        CompanyTracksTechnicalDebt: null,
        CompanyHasCodingStandards: null,
      },
      edLevels: [
        'None',
        'Some bachelor studies',
        'Bachelor degree',
        'Some master studies',
        'Master degree',
        'Some Ph. D. studies',
        'Ph. D.',
      ],
    }
  },
  computed: {
    invalid () {
      return Object.values(this.surveyAnswers).some((a) => a === null || a === [])
    },
  },
  methods: {
    submit () {
      if (this.loading) return
      this.loading = true
      this.axios.post('background', this.surveyAnswers).then(() => {
        this.loading = false
        this.$router.push('/experiment/1')
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
</style>
