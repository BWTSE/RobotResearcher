<template>
  <v-row class="fill-height">
    <v-col cols="2" class="list-col fill-height">
      <v-list dense>
        <v-list-item-group v-model="selected" color="primary">
          <v-list-item :key="name" v-for="name in Object.keys(files)" :value="name">{{ name }}</v-list-item>
        </v-list-item-group>
      </v-list>
    </v-col>
    <v-col>
      <editor
        :code="files[selected]"
        @update:code="edited"
      ></editor>
    </v-col>
  </v-row>
</template>

<script>
import Editor from './Editor'

const defaultFile = 'Main.java'

export default {
  name: 'MultiFileEditor',
  components: { Editor },
  data () {
    return {
      selected: defaultFile,
    }
  },
  props: {
    files: Object,
  },
  watch: {
    files () {
      if (this.files[this.selected] === undefined) {
        this.selected = defaultFile
      }
    },
  },
  methods: {
    edited (code) {
      this.$emit('update:file', {
        file: this.selected,
        content: code,
      })
    },
  },
}
</script>

<style scoped>
.col {
  padding: 0;
}
.list-col {
  overflow-y: auto;
  min-width: 200px;
}
.row {
  margin: 0;
}
</style>
