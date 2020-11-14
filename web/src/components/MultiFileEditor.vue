<template>
  <v-sheet>

    <v-tabs
      show-arrows
      @change="(n) => selected = n"
    >
      <v-tabs-slider color="cyan"></v-tabs-slider>
      <v-tab
        v-for="name in Object.keys(files)"
        :key="name"
      >
        {{ name }}
      </v-tab>
    </v-tabs>
    <editor
      :code="files[selectedKey]"
      @update:code="edited"
    ></editor>
  </v-sheet>
</template>

<script>
import Editor from './Editor'
export default {
  name: 'MultiFileEditor',
  components: { Editor },
  data () {
    return {
      selected: 0,
    }
  },
  props: {
    files: Object,
  },
  computed: {
    selectedKey () {
      return Object.keys(this.files)[this.selected]
    },
  },
  methods: {
    edited (code) {
      this.$emit('update:file', {
        file: this.selectedKey,
        content: code,
      })
    },
  },
}
</script>

<style scoped>

</style>
