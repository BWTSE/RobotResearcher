<template>
  <div class="editor" @keydown.ctrl.83.prevent.stop="$emit('save')">
    <div :id="_uid">
    </div>
  </div>
</template>

<script>
import Ace from 'ace-builds'
import 'ace-builds/webpack-resolver'

export default {
  name: 'Editor',
  data () {
    return {
      editor: null,
      editorOptions: {
        maxLines: 1000,
        minLines: 50,
        printMargin: false,
        value: this.code,
        mode: 'ace/mode/java',
        theme: 'ace/theme/monokai',
        fontSize: '12pt',
        tabSize: '4',
        useSoftTabs: false,
        autoScrollEditorIntoView: false,
        scrollPastEnd: false,
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

    this.editor.on('change', () => {
      this.$emit('update:code', this.editor.getValue())
    })
  },
}
</script>

<style scoped lang="scss">
  .editor {
    position: relative;
    overflow-y: auto;
    height: calc(100vh - 72px);
    background-color: #272822;
    >div {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
    }
  }
</style>
