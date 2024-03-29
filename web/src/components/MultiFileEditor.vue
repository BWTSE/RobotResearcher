<template>
  <v-row class="fill-height">
    <v-col class="list-col fill-height">
      <v-list dense>
        <div class="helper">
          Right click on files to add new, rename and delete. Refreshing the page will reset the whole task.
        </div>
        <v-list-item-group v-model="selected" color="primary" mandatory>
          <v-list-item
            v-for="name in Object.keys(files).sort()"
            :key="name"
            :value="name"
            @contextmenu="(e) => openContextMenu(name, e)"
          >{{ name }}</v-list-item>
        </v-list-item-group>
      </v-list>
      <v-menu
        absolute
        :position-x="contextMenuX"
        :position-y="contextMenuY"
        v-model="contextMenuOpen"
      >
        <v-list class="ctx-menu">
          <v-list-item @click="newFileDialog = true">
            <v-list-item-title>New File</v-list-item-title>
          </v-list-item>
          <v-list-item v-if="contextMenuFile !== 'Main.java'" @click="pendingRename = contextMenuFile">
            <v-list-item-title>Rename</v-list-item-title>
          </v-list-item>
          <v-list-item v-if="contextMenuFile !== 'Main.java'" @click="pendingDelete = contextMenuFile">
            <v-list-item-title>Delete</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
      <v-dialog
        v-model="newFileDialog"
        max-width="600px"
      >
        <v-card>
          <v-card-title>New File</v-card-title>
          <v-card-text>
            Create a new file
            <v-form
              align="right"
              @submit.prevent.stop=""
            >
              <v-text-field
                label="File name"
                suffix=".java"
                v-model="newFileName"
                ref="newFileInput"
                @keyup.enter.prevent.stop="addFile()"
                :error-messages="newFileError"
              ></v-text-field>
              <v-btn
                :disabled="newFileError != null || newFileName === ''"
                @click.prevent.stop="addFile()"
                color="primary"
              >Create</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-dialog>

      <v-dialog
        v-model="renameDialogOpen"
        max-width="600px"
      >
        <v-card>
          <v-card-title>Rename File</v-card-title>
          <v-card-text>
            Rename {{pendingRename}}
            <v-form
              @submit.prevent.stop=""
              align="right"
            >
              <v-text-field
                label="New name"
                suffix=".java"
                ref="renameFileInput"
                v-model="renameFileName"
                @keyup.enter.prevent.stop="renameFile()"
                autofocus
                :error-messages="renameFileError"
              ></v-text-field>
              <v-btn
                :disabled="renameFileError != null || renameFileName === '' || normalizeFileName(renameFileName) === pendingRename"
                @click.prevent.stop="renameFile()"
                color="primary"
              >Rename</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-dialog>

      <v-dialog
        v-model="deleteDialogOpen"
        max-width="600px"
        @keydown.enter="deleteFile"
      >
        <v-card>
          <v-card-title>Delete file</v-card-title>
          <v-card-text>
            Do you want to delete {{ pendingDelete}}?
            <v-card-actions>
              <v-btn
                @click="deleteDialogOpen = false"
              >Cancel</v-btn>
              <v-spacer></v-spacer>
              <v-btn
                @click="deleteFile"
                color="error"
              >Delete</v-btn>
            </v-card-actions>
          </v-card-text>
        </v-card>
      </v-dialog>
    </v-col>
    <v-col>
      <editor
        :file="{name: selected, code: files[selected]}"
        @update:code="edited"
        @save="$emit('save')"
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
      contextMenuFile: null,
      contextMenuX: 0,
      contextMenuY: 0,
      newFileDialog: false,
      newFileName: '',
      renameFileName: '',
      pendingDelete: null,
      pendingRename: null,
    }
  },
  props: {
    files: Object,
  },
  computed: {
    newFileError () {
      if (this.newFileName === '') {
        return null
      }
      return this.fileNameErrors(this.normalizeFileName(this.newFileName))
    },
    renameFileError () {
      if (this.renameFileName === '') {
        return null
      }
      if (this.normalizeFileName(this.renameFileName) === this.pendingRename) {
        return null
      }
      return this.fileNameErrors(this.normalizeFileName(this.renameFileName))
    },
    deleteDialogOpen: {
      get () {
        return this.pendingDelete !== null
      },
      set (x) {
        if (x === false) {
          this.pendingDelete = null
        }
      },
    },
    renameDialogOpen: {
      get () {
        return this.pendingRename !== null
      },
      set (x) {
        if (x === false) {
          this.pendingRename = null
        }
      },
    },
    contextMenuOpen: {
      get () {
        return this.contextMenuFile !== null
      },
      set (x) {
        if (x === false) {
          this.contextMenuFile = null
        }
      },
    },
  },
  watch: {
    files () {
      if (this.files[this.selected] === undefined) {
        this.selected = defaultFile
      }
    },
    pendingRename (val) {
      if (val != null) {
        this.$nextTick(() => {
          setTimeout(() => {
            this.$refs.renameFileInput.focus()
          }, 0)
        })
      }
    },
    newFileDialog (val) {
      if (val) {
        this.$nextTick(() => {
          setTimeout(() => {
            this.$refs.newFileInput.focus()
          }, 0)
        })
      }
    },
  },
  methods: {
    renameFile () {
      this.$emit('update:file', {
        file: this.normalizeFileName(this.renameFileName),
        content: this.files[this.pendingRename],
      })
      this.$emit('delete:file', {
        file: this.pendingRename,
      })
      this.selected = this.normalizeFileName(this.renameFileName)
      this.renameFileName = ''
      this.pendingRename = null
    },
    deleteFile () {
      this.$emit('delete:file', {
        file: this.pendingDelete,
      })
      this.pendingDelete = null
      this.selected = defaultFile
    },
    normalizeFileName (name) {
      if (name.indexOf('.java', name.length - '.java'.length) === -1) {
        name += '.java'
      }
      return name[0].toUpperCase() + name.substring(1)
    },
    fileNameErrors (name) {
      if (name.indexOf('.java', name.length - '.java'.length) === -1) {
        return 'only java files allowed'
      }
      if (name.indexOf('.') !== name.length - '.java'.length) {
        return 'dots are not allowed in file name'
      }
      if (Object.keys(this.files).some((f) => f === name)) {
        return 'that file already exists'
      }
      return null
    },
    addFile () {
      this.newFileDialog = false
      this.$emit('update:file', {
        file: this.normalizeFileName(this.newFileName),
        content: '\n',
      })
      this.selected = this.normalizeFileName(this.newFileName)
      this.newFileName = ''
    },
    edited (code) {
      this.$emit('update:file', {
        file: this.selected,
        content: code,
      })
    },
    openContextMenu (file, event) {
      event.preventDefault()
      this.contextMenuX = event.clientX
      this.contextMenuY = event.clientY
      this.contextMenuFile = file
    },
  },
}
</script>

<style scoped>

.helper {
  background-color: #cafac1;
  padding: 10px 10px 10px 20px;
}

.ctx-menu > *:hover {
  background-color: rgba(0, 0, 0, 0.1);
  cursor: pointer;
}
.col {
  padding: 0;
}
.list-col {
  overflow-y: auto;
  min-width: 200px;
  max-width: 300px;
}
.row {
  margin: 0;
}
</style>
