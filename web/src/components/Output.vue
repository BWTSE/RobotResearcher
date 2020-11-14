<template>
  <div :id="_uid">
  </div>
</template>

<script>
import Ace from 'ace-builds'
import 'ace-builds/webpack-resolver'

export default {
  name: 'Output',
  data () {
    return {
      editor: null,
      editorOptions: {
        maxLines: 15,
        minLines: 5,
        printMargin: false,
        value: this.code,
        theme: 'ace/theme/monokai',
        fontSize: '12pt',
        tabSize: '2',
        useSoftTabs: false,
        autoScrollEditorIntoView: false,
        scrollPastEnd: false,
        readOnly: true,
      },
    }
  },
  props: {
    code: String,
  },
  watch: {
    code (value) {
      if (this.editor != null && this.editor.getValue() !== value) {
        this.editor.setValue(value)
        this.editor.clearSelection()
      }
    },
  },
  mounted () {
    Ace.require('ace/ext/language_tools')
    this.editor = Ace.edit(String(this._uid), this.editorOptions)
  },
}
</script>

<style scoped lang="scss">
</style>
