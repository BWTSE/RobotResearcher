<template>
  <v-row class="fill-height">
    <v-col cols="2" class="list-col fill-height">
      <v-list dense>
        <div class="helper">
          Right click on files to add new, rename and delete. Refreshing the page will reset the whole task.
        </div>
        <v-list-item-group v-model="selected" color="primary" mandatory>
          <v-list-item
            v-for="name in Object.keys(files)"
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
          <v-list-item>
            <v-list-item-title @click="newFileDialog = true">New File</v-list-item-title>
          </v-list-item>
          <v-list-item v-if="contextMenuFile !== 'Main.java'">
            <v-list-item-title @click="pendingRename = contextMenuFile">Rename</v-list-item-title>
          </v-list-item>
          <v-list-item v-if="contextMenuFile !== 'Main.java'">
            <v-list-item-title @click="pendingDelete = contextMenuFile">Delete</v-list-item-title>
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
            >
              <v-text-field
                label="File name"
                suffix=".java"
                v-model="newFileName"
                :error-messages="newFileError"
              ></v-text-field>
              <v-btn
                :disabled="newFileError != null || newFileName === ''"
                @click="addFile()"
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
              align="right"
            >
              <v-text-field
                label="New name"
                suffix=".java"
                v-model="renameFileName"
                :error-messages="renameFileError"
              ></v-text-field>
              <v-btn
                :disabled="renameFileError != null || renameFileName === '' || normalizeFileName(renameFileName) === pendingRename"
                @click="renameFile()"
                color="primary"
              >Rename</v-btn>
            </v-form>
          </v-card-text>
        </v-card>
      </v-dialog>

      <v-dialog
        v-model="deleteDialogOpen"
        max-width="600px"
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
}
.row {
  margin: 0;
}
</style>
