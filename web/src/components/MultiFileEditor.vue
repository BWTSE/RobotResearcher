<template>
  <v-row class="fill-height">
    <v-col cols="2" class="list-col fill-height">
      <v-list dense>
        <v-list-item-group v-model="selected" color="primary">
          <v-list-item :key="name" v-for="name in Object.keys(files)">{{ name }}</v-list-item>
        </v-list-item-group>
      </v-list>
    </v-col>
    <v-col>
      <editor
        :code="files[selectedKey]"
        @update:code="edited"
      ></editor>
    </v-col>
  </v-row>
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
